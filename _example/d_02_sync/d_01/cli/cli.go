package main

import (
	"context"
	"flag"
	"log"
	"rpcx_source/_example/d_01_simple/d_01/model"
	"rpcx_source/client"
	"rpcx_source/protocol"
	"rpcx_source/share"
	"time"
)

var (
	addr2 = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	var op = client.Option{
		Retries:        10,
		RPCPath:        share.DefaultRPCPath,
		ConnectTimeout: 10 * time.Second,
		SerializeType:  protocol.MsgPack,
		CompressType:   protocol.None,
		BackupLatency:  10 * time.Millisecond,
	}

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr2, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, op)
	defer xclient.Close()

	args := &model.Args{
		A: 10,
		B: 20,
	}

	reply := &model.Reply{}
	// Go 代替 call
	call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
	}

}
