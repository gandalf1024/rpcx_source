package main

import (
	"flag"
	"time"

	example "github.com/rpcxio/rpcx-examples"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	a := serverplugin.NewAliasPlugin()
	//Mul 别名 Times
	//Arith 别名 a.b.c.D
	a.Alias("a.b.c.D", "Times<-Mul", "Arith", "Mul")

	option := server.WithReadTimeout(time.Second * 3)
	option2 := server.WithWriteTimeout(time.Second * 3)
	s := server.NewServer(option, option2)
	s.Plugins.Add(a)
	s.RegisterName("Arith", new(example.Arith), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
