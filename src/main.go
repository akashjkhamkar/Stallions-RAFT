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

	peer_ids := []string{"a", "b", "c", "d"}
	applyMsg := make(chan raft.ApplyMsg)
	raft.Make(
		peer_ids,
		*serverid,
		applyMsg,
		*port,
	)
}
