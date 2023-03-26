package raft

//
// this is an outline of the API that raft must expose to
// the service (or tester). see comments below for
// each of these functions for more details.
//
// rf = Make(...)
//   create a new Raft server.
// rf.Start(command interface{}) (index, term, isleader)
//   start agreement on a new log entry
// rf.GetState() (term, isLeader)
//   ask a Raft for its current term, and whether it thinks it is leader
// ApplyMsg
//   each time a new entry is committed to the log, each Raft peer
//   should send an ApplyMsg to the service (or tester)
//   in the same server.
//

import (
	//	"bytes"

	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	pb "stallionraft/raftrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make(). set
// CommandValid to true to indicate that the ApplyMsg contains a newly
// committed log entry.
//
// in part 2D you'll want to send other kinds of messages (e.g.,
// snapshots) on the applyCh, but set CommandValid to false for these
// other uses.
//
type ApplyMsg struct {
	CommandValid bool
	Command      interface{}
	CommandIndex int

	// For 2D:
	SnapshotValid bool
	Snapshot      []byte
	SnapshotTerm  int
	SnapshotIndex int
}


type LogEntry struct {
	Command interface{}
	Term int
}
//
// A Go object implementing a single Raft peer.
//
type Raft struct {
	mu        sync.Mutex          // Lock to protect shared access to this peer's state
	peer_ids     [] string // RPC end points of all peers
	peers     [] pb.RaftRpcClient // RPC end points of all peers
	raft_rpc_server *server
	majority int
	me        int                 // this peer's index into peers[]
	dead      int32               // set by Kill()

	// Your data here (2A, 2B, 2C).
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.
	election_timeout int
	last_reset_time int64
	random_sleep_time_range int
	base_sleep_time int

	candidate int32
	leader int32

	term int
	voted int

	commitIndex int
	lastApplied int

	nextIndex[] int
	matchIndex[] int

	log[] * pb.LogEntry
	apply_channel chan ApplyMsg
}

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	return rf.term, rf.is_leader()
}

//
// save Raft's persistent state to stable storage,
// where it can later be retrieved after a crash and restart.
// see paper's Figure 2 for a description of what should be persistent.
//
func (rf *Raft) persist() {
	// Your code here (2C).
	// Example:
	// w := new(bytes.Buffer)
	// e := labgob.NewEncoder(w)
	// e.Encode(rf.xxx)
	// e.Encode(rf.yyy)
	// data := w.Bytes()
	// rf.persister.SaveRaftState(data)
}


//
// restore previously persisted state.
//
func (rf *Raft) readPersist(data []byte) {
	if data == nil || len(data) < 1 { // bootstrap without any state?
		return
	}
	// Your code here (2C).
	// Example:
	// r := bytes.NewBuffer(data)
	// d := labgob.NewDecoder(r)
	// var xxx
	// var yyy
	// if d.Decode(&xxx) != nil ||
	//    d.Decode(&yyy) != nil {
	//   error...
	// } else {
	//   rf.xxx = xxx
	//   rf.yyy = yyy
	// }
}


//
// A service wants to switch to snapshot.  Only do so if Raft hasn't
// have more recent info since it communicate the snapshot on applyCh.
//
func (rf *Raft) CondInstallSnapshot(lastIncludedTerm int, lastIncludedIndex int, snapshot []byte) bool {

	// Your code here (2D).

	return true
}

// the service says it has created a snapshot that has
// all info up to and including index. this means the
// service no longer needs the log through (and including)
// that index. Raft should now trim its log as much as possible.
func (rf *Raft) Snapshot(index int, snapshot []byte) {
	// Your code here (2D).

}

// the service using Raft (e.g. a k/v server) wants to start
// agreement on the next command to be appended to Raft's log. if this
// server isn't the leader, returns false. otherwise start the
// agreement and return immediately. there is no guarantee that this
// command will ever be committed to the Raft log, since the leader
// may fail or lose an election. even if the Raft instance has been killed,
// this function should return gracefully.
//
// the first return value is the index that the command will appear at
// if it's ever committed. the second return value is the current
// term. the third return value is true if this server believes it is
// the leader.
//
func (rf *Raft) Start(command string) (int, int, bool) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	index := len(rf.log) + 1
	term := rf.term
	isLeader := rf.is_leader()

	if isLeader {
		// append the entry
		entry := &pb.LogEntry{
			Command: command,
			Term: int32(term),
		}
		
		rf.log = append(rf.log, entry)
		rf.Debug(dMake, "Adding entry : ", command)
	}

	return index, term, isLeader
}

//
// the tester doesn't halt goroutines created by Raft after each test,
// but it does call the Kill() method. your code can use killed() to
// check whether Kill() has been called. the use of atomic avoids the
// need for a lock.
//
// the issue is that long-running goroutines use memory and may chew
// up CPU time, perhaps causing later tests to fail and generating
// confusing debug output. any goroutine with a long-running loop
// should call killed() to check whether it should stop.
//
func (rf *Raft) Kill() {
	atomic.StoreInt32(&rf.dead, 1)
	// Your code here, if desired.
}

func (rf *Raft) killed() bool {
	z := atomic.LoadInt32(&rf.dead)
	return z == 1
}

// execute the commands upto the commit point
func (rf *Raft) executer() {
	index := rf.lastApplied + 1

	for rf.commitIndex >= index {
		msg := ApplyMsg{
			CommandValid: true,
			Command: rf.log[index - 1].Command,
			CommandIndex: index,
		}

		rf.Debug(dExecuter, "Executing cmd : ", rf.log[index - 1])
		rf.apply_channel <- msg
		index++
	}

	rf.lastApplied = rf.commitIndex
}

// The ticker go routine starts a new election if this peer hasn't received
// heartsbeats recently.
func (rf *Raft) ticker() {
	for !rf.killed() {
		// Sleep atleast for 100ms to avoid excessive locking
		time.Sleep(100*time.Millisecond)
		rf.random_sleep_2(100)

		rf.mu.Lock()

		if !rf.is_election_timeout() || rf.is_leader() {
			rf.mu.Unlock()
			continue
		} else if rf.is_candidate() {
			rf.Debug(dTicker, "Stopping election and then sleeping for random time")
			rf.set_candidate(false)
			rf.voted = -1
			rf.mu.Unlock()
			rf.random_sleep_2(200)
			continue
		} else {
			rf.Debug(dTicker, "Starting election")
			rf.set_candidate(true)
			rf.election()
			rf.reset_election_timeout()
		}

		rf.mu.Unlock()
	}
}

func (rf *Raft) connect_peers() {
	for i, addr := range rf.peer_ids {
		if i == rf.me{
			rf.peers = append(rf.peers, nil)
			continue
		}

		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			rf.Debug("did not connect address %s : %v", addr, err)
		}

		raft_rpc_client := pb.NewRaftRpcClient(conn)
		rf.peers = append(rf.peers, raft_rpc_client)
	}
}

type server struct {
	pb.UnimplementedRaftRpcServer
}

var Raft_instance *Raft;

func (rf *Raft) startGrpcServer(grpc_server grpc.Server, listener net.Listener) {
	if err := grpc_server.Serve(listener); err != nil {
		rf.Debug(dInit, "failed to serve: %v", err)
	}
}

func Make(peers [] string, me int, applyCh chan ApplyMsg, port int) *Raft {
	rf := &Raft{}
	rf.majority = len(peers)/2 + 1
	rf.me = me
	rf.peer_ids = peers
	rf.voted = -1
	rf.apply_channel = applyCh
	rf.random_sleep_time_range = 300
	rf.base_sleep_time = 300
	rf.election_timeout = 300

	// connect to of all the peers
	rf.connect_peers()

	// Starting grpc server setup
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		rf.Debug(dInit, "failed to listen")
	}

	grpc_server := grpc.NewServer()
	rf.raft_rpc_server = &server{}
	Raft_instance = rf

	pb.RegisterRaftRpcServer(grpc_server, rf.raft_rpc_server)

	rf.Debug(dInit, "server listening at %v", lis.Addr())

	// Listening for grpc calls
	rf.Debug(dInit, "Starting the grpc server ...")
	go rf.startGrpcServer(*grpc_server, lis)

	// Starting the ticker
	rf.Debug(dInit, "Starting the ticker ...")
	rf.ticker()

	return rf
}
