package main

import (
	r "github.com/unprofession-al/routing"
)

func (s Server) routes() r.Route {
	return r.Route{
		R: r.Routes{
			"hooks/{hook}": {
				H: r.Handlers{"POST": r.Handler{F: s.HookHandler}},
			},
		},
	}
}
