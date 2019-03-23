package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/pascaloseko/Events/events/rest"

	"github.com/pascaloseko/Events/lib/config"
)

func main() {
	confPath := flag.String("conf", `../lib/config/config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	// extract config details
	conf, _ := config.ExtractConfiguration(*confPath)
	fmt.Println("Server running on " + conf.RestfulEndpoint)
	log.Println(rest.ServeAPI(conf.RestfulEndpoint, nil))
}