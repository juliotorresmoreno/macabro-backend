package main

import (
	"github.com/juliotorresmoreno/macabro/bootstrap"
	"github.com/juliotorresmoreno/macabro/server"
	_ "github.com/lib/pq"
)

func main() {
	bootstrap.Init()
	svr := server.NewFastServerHTTP()
	svr.Listen()
}
