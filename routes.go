package main

import (
	"github.com/ant0ine/go-urlrouter"
)

//order matters so most general should go towards the bottom
var router urlrouter.Router = urlrouter.Router{
	Routes: []urlrouter.Route{
		urlrouter.Route{
			PathExp: "/string",
			Dest: map[string]interface{}{
				"POST": ReturnString,
			},
		},
	},
}
