package main

var plugins = make([]Plugin, 0)

func InitPlugin() {
	for _, p := range []Plugin{
		new(pluginRootSecret),
		new(pluginCSVSecret),
	} {
		if p.Init() != nil {
			plugins = append(plugins, p)
		}
	}
}

func GetPlugins() []Plugin {
	return plugins
}

type Plugin interface {
	Authorization() map[string][]Authorization
	Init() Plugin
}
