package main

type pluginRootSecret struct {
	Secret string
}

func (p *pluginRootSecret) Init() Plugin {
	if Opts.Secret != "" {
		p.Secret = Opts.Secret
		return p
	}
	return nil
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
