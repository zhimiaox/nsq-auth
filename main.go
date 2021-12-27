package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jessevdk/go-flags"
)

var SystemOpts struct {
	APIAddr     string `short:"a" long:"address" description:"api port default :1325" default:":1325"`
	Identity    string `short:"i" long:"identity" description:"identity default zhimiaox-nsq-auth" default:"zhimiaox-nsq-auth"` //nolint:lll
	IdentityURL string `short:"u" long:"auth-url" description:"auth-url" default:"http://localhost:1325"`
	TTL         int    `short:"t" long:"ttl" description:"auth expire duration unit s, default 60" default:"60"`
	Secret      string `short:"s" long:"secret" description:"root secret allow all push and sub topic and channel" default:""` //nolint:lll
	CSV         string `short:"f" long:"csv" description:"csv secret file path" default:""`
}

func main() {
	_, err := flags.NewParser(&SystemOpts, flags.Default).Parse()
	if err != nil {
		log.Fatalln(err.Error())
	}
	InitPlugin()
	// 初始化权限数据
	GetStorage().Refresh()
	go func() {
		if err := APIRoute().Run(SystemOpts.APIAddr); err != nil {
			panic(err)
		}
	}()
	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
