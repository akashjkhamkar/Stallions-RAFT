package main

import (
	"flag"
	"stallionraft/src/raft"
)

var (
	port = flag.Int("port", 50051, "The server port")
	serverid = flag.Int("id", 0, "The server id")
)

func main() {
	flag.Parse()

	peer_ids := []string{"localhost:50051", "localhost:50052", "localhost:50053"}
	applyMsg := make(chan raft.ApplyMsg)
	raft.Make(
		peer_ids,
		*serverid,
		applyMsg,
		*port,
	)
}
