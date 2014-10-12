package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	bindAddr   = flag.String("bind", "127.0.0.1:8080", "bind address")
	dockerAddr = flag.String("dockerapi", "unix:///var/run/docker.sock", "docker address")
)

func main() {
	flag.Parse()

	proxy := &Proxy{}
	log.Println("HTTP Proxy running on", *bindAddr)
	err := http.ListenAndServe(*bindAddr, proxy)
	if err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func checkFatalError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
