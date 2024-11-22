package main

import (
	_ "datahub_exporter/config" // 初始化配置文件
	"datahub_exporter/router"
	"flag"
)

func main() {

	var addr string
	flag.StringVar(&addr, "web.listen-address", ":48080", "The listen-address of web")

	flag.Parse()

	router := router.Router()
	router.Run(addr)

}
