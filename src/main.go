package main

import (
	"flag"
	"fmt"
	"log"
	"stallionraft/src/kvserver"
	"stallionraft/src/raft"
	"strconv"

	"github.com/caarlos0/env"
)

var (
	port     = flag.Int("port", 50051, "The server port")
	serverid = flag.Int("id", 0, "The server id")
)

type config struct {
	Pod_name string `env:"POD_NAME"`
}

func main() {
	flag.Parse()

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)

	// Make sure pod name ends with a number !
	pod_name := cfg.Pod_name

	if pod_name == "" {
		log.Fatalf("Pod name not found, please pass the pod name through the env variables")
		return
	}

	pod_id_str := string(pod_name[len(pod_name)-1])
	pod_id, _ := strconv.ParseInt(pod_id_str, 0, 64)

	peer_ids := []string{
		"stallion-raft-0.stallion-raft-service.default.svc.cluster.local:50051",
		"stallion-raft-1.stallion-raft-service.default.svc.cluster.local:50051",
		"stallion-raft-2.stallion-raft-service.default.svc.cluster.local:50051",
	}

	applyMsg := make(chan raft.ApplyMsg)
	raft.Make(
		peer_ids,
		int(pod_id),
		applyMsg,
		*port,
	)

	kvserver.StartKVserver()

}
