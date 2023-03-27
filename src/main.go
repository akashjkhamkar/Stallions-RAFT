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

	peer_ids := []string{
		"stallion-raft-0.stallion-raft-service.default.svc.cluster.local:50051",
		"stallion-raft-1.stallion-raft-service.default.svc.cluster.local:50051",
		"stallion-raft-2.stallion-raft-service.default.svc.cluster.local:50051",
	}
	applyMsg := make(chan raft.ApplyMsg)
	raft.Make(
		peer_ids,
		*serverid,
		applyMsg,
		*port,
	)
}
