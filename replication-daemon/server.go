package simplepb

//
// This is a outline of primary-backup replication based on a simplifed version of Viewstamp replication.
//
//
//

import (
	"log"
	"net"
	"sync"

	"github.com/adamsanghera/blurber-protobufs/dist/replication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// the 3 possible server status
const (
	NORMAL = iota
	VIEWCHANGE
	RECOVERING
)

type callbackArg struct {
	callback chan bool
	args     *replication.PrepareArgs
	handled  bool
}

// PBServer defines the state of a replica server (either primary or backup)
type PBServer struct {
	mu             *sync.Mutex                     // Lock to protect shared access to this peer's state
	peers          []replication.ReplicationClient // RPC end points of all peers
	peerAddresses  []string                        // Addresses of all peers
	me             int32                           // this peer's index into peers[]
	currentView    int32                           // what this peer believes to be the current active view
	status         int32                           // the server's current status (NORMAL, VIEWCHANGE or RECOVERING)
	lastNormalView int32                           // the latest view which had a NORMAL status

	log         []*replication.Command // the log of "commands"
	commitIndex int32                  // all log entries <= commitIndex are considered to have been committed.
	commitChan  chan *replication.Command

	prepChan chan *callbackArg // Channel used by prep calls to communicate with the central prep-processor
}

// GetPrimary is an auxilary function that returns the server index of the
// primary server given the view number (and the total number of replica servers)
func GetPrimary(view int32, nservers int32) int32 {
	return view % nservers
}

// IsCommitted is called by tester to check whether an index position
// has been considered committed by this server
func (srv *PBServer) IsCommitted(index int32) (committed bool) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	if srv.commitIndex >= index {
		return true
	}
	return false
}

func (srv *PBServer) connectPeer(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("ReplicationD: Failed to connect to the Replication Daemon at %s", addr)
		return err
	}
	self := replication.NewReplicationClient(conn)
	log.Printf("ReplicationD: Successfully connected to peer at %s", addr)

	srv.peers = append(srv.peers, self)
	srv.peerAddresses = append(srv.peerAddresses, addr)
	return nil
}

// ViewStatus is called by tester to find out the current view of this server
// and whether this view has a status of NORMAL.
func (srv *PBServer) ViewStatus() (currentView int32, statusIsNormal bool) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	return srv.currentView, srv.status == NORMAL
}

// GetEntryAtIndex is called by tester to return the command replicated at
// a specific log index. If the server's log is shorter than "index", then
// ok = false, otherwise, ok = true
func (srv *PBServer) GetEntryAtIndex(index int) (ok bool, command interface{}) {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	if len(srv.log) > index {
		return true, srv.log[index]
	}
	return false, command
}

func runGRPCServer(thisAddress string, srv *PBServer) {
	log.Printf("ReplicationD: Registering to listen on (%s)", thisAddress)
	lis, err := net.Listen("tcp", thisAddress)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	replication.RegisterReplicationServer(s, srv)
	reflection.Register(s)

	log.Printf("ReplicationD: Registered successfully (%s)", thisAddress)

	if err = s.Serve(lis); err != nil {
		panic("ReplicationD: Failed to start gRPC server")
	}
}

// NewReplicationDaemon spawns a new replication daemon.
// If thisAddress and leaderAddress are not the same,
// then the new daemon will start in recovery mode.
func NewReplicationDaemon(thisAddress string, leaderAddress string) *PBServer {
	srv := &PBServer{
		mu:             &sync.Mutex{},
		peers:          make([]replication.ReplicationClient, 0),
		peerAddresses:  make([]string, 0),
		me:             0,
		currentView:    0,
		status:         NORMAL,
		lastNormalView: 0,
		log:            make([]*replication.Command, 0),
		commitIndex:    0,
		commitChan:     make(chan *replication.Command, 1000),
		prepChan:       make(chan *callbackArg),
	}

	srv.mu.Lock()
	defer srv.mu.Unlock()

	// Init log
	srv.log = append(srv.log, &replication.Command{})

	// Initting daemon connections
	if thisAddress != leaderAddress {
		log.Printf("ReplicationD: Spawning as follower")
		err := srv.connectPeer(leaderAddress)
		if err != nil {
			panic(err)
		}
		err = srv.connectPeer(thisAddress)
		if err != nil {
			panic(err)
		}
		srv.me = 1
	} else {
		log.Printf("ReplicationD: Spawning as leader")
		err := srv.connectPeer(thisAddress)
		if err != nil {
			panic(err)
		}
	}

	log.Printf("Spwaning replication processor")
	go srv.prepareProcessor()

	// Registering server
	go runGRPCServer(thisAddress, srv)

	if leaderAddress != thisAddress {
		log.Printf("ReplicationD: Follower entering recovery mode")
		go srv.sendRecovery()
	}

	return srv
}