package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/mreiferson/go-options"
)

var Opts *Options

func main() {
	Opts = NewOptions()

	flagSet := FlagSet(Opts)
	flagSet.Parse(os.Args[1:])

	var cfg config
	configFile := flagSet.Lookup("config").Value.String()
	if configFile != "" {
		_, err := toml.DecodeFile(configFile, &cfg)
		if err != nil {
			log.Fatalf("failed to load config file %s - %s", configFile, err)
		}
	}

	options.Resolve(Opts, flagSet, cfg)

	InitPlugin()
	// 初始化权限数据
	GetStorage().Refresh()
	go func() {
		if err := APIRoute().Run(Opts.APIAddr); err != nil {
			panic(err)
		}
	}()
	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
