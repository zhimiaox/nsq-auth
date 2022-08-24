package main

import (
	"flag"
)

type Options struct {
	APIAddr     string `flag:"address"`
	Identity    string `flag:"identity"`
	IdentityURL string `flag:"auth-url"`
	TTL         int    `flag:"ttl"`
	Secret      string `flag:"secret"`
	CSV         string `flag:"csv"`
}

type config map[string]interface{}

func NewOptions() *Options {
	return &Options{
		APIAddr:     ":1325",
		Identity:    "zhimiaox-nsq-auth",
		IdentityURL: "http://localhost:1325",
		TTL:         60,
		Secret:      "",
		CSV:         "",
	}
}

func FlagSet(opts *Options) *flag.FlagSet {
	flagSet := flag.NewFlagSet("nsqauth", flag.ExitOnError)

	flagSet.String("config", "", `path to config file`)

	flagSet.String("address", opts.APIAddr, `api port, default: ":1325"`)
	flagSet.String("identity", opts.Identity, `identity default: "zhimiaox-nsq-auth"`)
	flagSet.String("auth-url", opts.IdentityURL, `auth-url`)
	flagSet.Int("ttl", opts.TTL, `auth expire duration unit s, default: 60`)
	flagSet.String("secret", "", `root secret allow all push and sub topic and channel`)
	flagSet.String("csv", "", `csv secret file path`)

	return flagSet
}
