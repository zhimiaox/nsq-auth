package main

import (
	"flag"
)

type pluginRootSecret struct {
	Secret string
}

func init() {
	p := &pluginRootSecret{}
	flag.StringVar(&p.Secret, "secret", "", "root secret allow all push and sub topic and channel")
	if p.Secret != "" {
		AddPlugin(p)
	}
}

func (p *pluginRootSecret) Authorization() map[string][]Authorization {
	data := make(map[string][]Authorization)
	data[p.Secret] = []Authorization{
		{
			Topic:       ".*",
			Permissions: []string{Subscribe, Publish},
			Channels:    []string{".*"},
		},
	}
	return data
}
