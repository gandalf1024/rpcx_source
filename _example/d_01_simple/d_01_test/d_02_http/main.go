package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:9988")
	if err != nil {
		panic(err)
	}
	//m := cmux.New(ln)
	//jsonrpc2Ln := m.Match(cmux.HTTP1HeaderField("X-JSONRPC-2.0", "true"))

	newServer := http.NewServeMux()
	newServer.HandleFunc("/json", jsonrpcHandler)

	srv := http.Server{ConnContext: func(ctx context.Context, c net.Conn) context.Context {
		return context.WithValue(ctx, "value--0", c)
	}}

	srv.Handler = newServer
	go srv.Serve(ln)

	time.Sleep(time.Minute * 5)
}

func jsonrpcHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------")
}
