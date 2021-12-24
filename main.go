package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	APIAddr             string
	ExpiredAuthDuration time.Duration
	Identity            string
	IdentityURL         string
)

func main() {
	flag.StringVar(&APIAddr, "addr", ":1325", "api port default :1325")
	flag.StringVar(&Identity, "identity", "zhimiaox-nsq-auth", "identity default zhimiaox-nsq-auth")
	flag.StringVar(&IdentityURL, "auth-url", "http://localhost:1325", "auth-url")
	flag.DurationVar(&ExpiredAuthDuration, "expire", time.Minute, "auth expire duration default 1m")

	go StartAPI()
	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
