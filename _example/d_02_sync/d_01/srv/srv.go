package main

import (
	"flag"
	"rpcx_source/_example/d_01_simple/d_01/model"
	"rpcx_source/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()
	// 启动服务
	// 初始化操作
	// 参数非必填 参数为配置信息
	s := server.NewServer()
	//s.Plugins = 自定义插件必须在Register前
	s.Register(new(model.Arith), "")
	s.Serve("tcp", *addr)
}
