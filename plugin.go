package main

var plugins []Plugin = make([]Plugin, 0)

func InitPlugin() {
	if p := new(pluginRootSecret).Init(); p != nil {
		plugins = append(plugins, p)
	}
}

func GetPlugins() []Plugin {
	return plugins
}

type Plugin interface {
	Authorization() map[string][]Authorization
	Init() Plugin
}
