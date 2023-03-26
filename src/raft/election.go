package raft

import (
	"context"
	pb "stallionraft/raftrpc"
	"time"
)

type RequestVoteArgs struct {
	Term int
	Server int
	LastLogIndex int
	LastLogTerm int
}

type RequestVoteReply struct {
	Term int
	Vote bool
}

func (rf *Raft) is_log_upto_date(candidate_last_entry_index, candidate_last_entry_term int) bool {
	last_index := len(rf.log)

	if last_index == 0 {
		return true
	}

	current_last_entry := rf.log[last_index - 1]
	
	if int(current_last_entry.Term) > candidate_last_entry_term {
		return false
	} else if int(current_last_entry.Term) < candidate_last_entry_term {
		return true
	} else if candidate_last_entry_index >= last_index {
		return true
	}

	return false
}

func (s *server) RequestVoteHandler(ctx context.Context, args *pb.RequestVoteArgs) (*pb.RequestVoteReply, error) {
	rf := Raft_instance
	rf.mu.Lock()
	defer rf.mu.Unlock()

	reply := &pb.RequestVoteReply{}
	reply.Term = int32(rf.term)

	is_log_upto_date := rf.is_log_upto_date(int(args.Lastlogindex), int(args.Lastlogterm))

	if int(args.Term) < rf.term {
		rf.Debug(dTicker, "No vote for S%d because old term (%d).", args.Server, args.Term)
		reply.Vote = false
		return reply, nil
	}
	
	if int(args.Term) > rf.term{
		rf.term = int(args.Term)
		rf.become_follower()

		if is_log_upto_date {
			rf.Debug(dTicker, "Voting to a higher term candidate S%d", args.Server)
			rf.voted = -1
		}
	}

	if (rf.voted == -1 || int32(rf.voted) == args.Server) && is_log_upto_date {
		// Grant the vote and convert to follower
		rf.Debug(dTicker, "Voted to S%d", args.Server)
		rf.voted = int(args.Server)
		reply.Term = int32(rf.term)
		reply.Vote = true
		rf.reset_election_timeout()
		return reply, nil
	}

	rf.Debug(dTicker, "No vote for S%d beacause already voted S%d", args.Server, rf.voted)
	reply.Vote = false
	return reply, nil
}

func (rf *Raft) sendRequestVote(server int, args RequestVoteArgs, reply *RequestVoteReply) bool {
	rpc_req_data := pb.RequestVoteArgs{
		Term: int32(args.Term),
		Server: int32(args.Server),
		Lastlogindex: int32(args.LastLogIndex),
		Lastlogterm: int32(args.LastLogIndex),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := rf.peers[server].RequestVoteHandler(ctx, &rpc_req_data)
	if err != nil {
		rf.Debug(dElection, "could not send vote req: %v", err)
		return false
	}

	reply.Term = int(res.Term)
	reply.Vote = res.Vote
	
	return true
}

func (rf *Raft) request_vote(term int, server int, vote_result chan RequestVoteReply) {
	LastIndex := len(rf.log)
	LastTerm := 0
	
	if LastIndex != 0 {
		LastTerm = int(rf.log[LastIndex - 1].Term)
	}

	args := RequestVoteArgs{
		Term: term,
		Server: rf.me,
		LastLogIndex: LastIndex,
		LastLogTerm: LastTerm,
	}
	
	reply := &RequestVoteReply{}

	for rf.is_candidate() && !rf.killed() && rf.get_current_term() == term {
		ok := rf.sendRequestVote(server, args, reply)

		if !ok {
			continue
		}

		break
	}

	vote_result <- *reply
}

func (rf *Raft) send_vote_requests(term int) {
	result_ch := make(chan RequestVoteReply, len(rf.peers) - 1)
	
	// Spawn the workers to get the votes 
	for i := 0; i < len(rf.peers); i++ {
		if i == rf.me {
			continue
		}

		go rf.request_vote(term, i, result_ch)
	}

	// Receive the votes
	votes := 1
	for i := 0; i < len(rf.peers) - 1; i++ {
		result := <- result_ch

		if result.Vote {
			votes++
		}

		rf.mu.Lock()

		if !rf.is_candidate() || rf.term != term {
			rf.mu.Unlock()
			break
		}
		
		if result.Term > term {
			rf.term = result.Term
			rf.mu.Unlock()
			break
		}

		if votes == rf.majority {
			// Become leader and end the election
			rf.Debug(dElection, "Won the election")
			rf.set_leader(true)
			rf.voted = -1
			go rf.heartbeats(rf.term)
						
			rf.mu.Unlock()
			return
		}

		rf.mu.Unlock()
	}

	rf.Debug(dElection, "Lost the election", term)

	rf.mu.Lock()
	defer rf.mu.Unlock()
	rf.become_follower()
	rf.voted = -1
}

func (rf *Raft) election() {
	// Get the term
	// Start the main routine
	
	rf.term++
	rf.voted = rf.me
	go rf.send_vote_requests(rf.term)
}