package main

import (
	"log"
	"net/http"
	_ "github.com/robatipoor/very-simple-panel-admin/config"
	"github.com/robatipoor/very-simple-panel-admin/router"
	"github.com/spf13/viper"
)

var addr string

func main() {
	routers := router.GetRouters()
	err := http.ListenAndServe(addr, routers)
	if err != nil {
		log.Fatalln(err)
	}
	println("Start Server ...")
}

func init() {
	a := viper.GetString("server.addr")
	port := viper.GetString("server.port")
	addr = a + ":" + port
}
