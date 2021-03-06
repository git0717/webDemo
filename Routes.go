package main

import (
	"net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{Name:"Index", Method:"Get",Pattern: "/",HandlerFunc:Index},
	Route{Name:"TodoIndex",Method:"GET", Pattern:"/todos", HandlerFunc:TodoIndex},
	Route{Name: "TodoShow",Method: "GET",Pattern: "/todos/{todoId}",HandlerFunc:TodoShow},
	Route{Name: "TodoPostForm",Method: "POST",Pattern: "/TodoPostForm",HandlerFunc:TodoPostForm},
	Route{Name: "TodoPostJson",Method: "POST",Pattern: "/TodoPostJson",HandlerFunc:TodoPostJson},
}