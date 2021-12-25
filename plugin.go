package main

var plugins []Plugin = make([]Plugin, 0)

func AddPlugin(p Plugin) {
	plugins = append(plugins, p)
}

func GetPlugins() []Plugin {
	return plugins
}

type Plugin interface {
	Authorization() map[string][]Authorization
}
