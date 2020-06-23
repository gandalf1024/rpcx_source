package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/soheilhy/cmux"
	"net"
	"net/http"
	"time"
)

func main() {

	ln, err := net.Listen("tcp", "0.0.0.0:9988")
	if err != nil {
		panic(err)
	}

	m := cmux.New(ln)

	jsonrpc2Ln := m.Match(cmux.HTTP1HeaderField("X-JSONRPC-2.0", "true"))
	go startJSONRPC2(jsonrpc2Ln)

	httpLn := m.Match(cmux.HTTP1Fast())
	go startHTTP1APIGateway(httpLn)

	time.Sleep(time.Minute * 5)
}

type contextKey struct {
	name string
}

func startJSONRPC2(ln net.Listener) {
	HttpConnContextKey := &contextKey{"http-conn"}
	newServer := http.NewServeMux()
	newServer.HandleFunc("/json", jsonrpcHandler)

	srv := http.Server{ConnContext: func(ctx context.Context, c net.Conn) context.Context {
		return context.WithValue(ctx, HttpConnContextKey, c)
	}}

	srv.Handler = newServer
	go srv.Serve(ln)

}

func jsonrpcHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------jsonrpcHandler--")
}

func startHTTP1APIGateway(ln net.Listener) {
	router := httprouter.New()
	//router.POST("/*servicePath", handleGatewayRequest)
	router.GET("/api", handleGatewayRequest)
	//router.PUT("/*servicePath", handleGatewayRequest)

	hrou := &http.Server{Handler: router}
	go hrou.Serve(ln)
}

func handleGatewayRequest(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("--------------handleGatewayRequest--")
}
