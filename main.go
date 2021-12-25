package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var (
	APIAddr     string
	TTL         int
	Identity    string
	IdentityURL string
	RootSecret  string
)

func main() {
	flag.StringVar(&APIAddr, "addr", ":1325", "api port default :1325")
	flag.StringVar(&Identity, "identity", "zhimiaox-nsq-auth", "identity default zhimiaox-nsq-auth")
	flag.StringVar(&IdentityURL, "auth-url", "http://localhost:1325", "auth-url")
	flag.IntVar(&TTL, "ttl", 60, "auth expire duration unit s, default 60")

	go StartAPI()
	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
