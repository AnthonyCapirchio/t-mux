package router

import (
	"octopus-project.t-mux/tree"
)

type Router struct {
	Tree *tree.TreeNode
}

type RouteHandler interface{}

type RouteHandlers map[string]interface{}

func NewRouter() *Router {
	return &Router{
		Tree: tree.NewTree(),
	}
}

func (r *Router) Get(path string, handler interface{}) {
	r.Tree.AddNode(path, "GET", handler)
	//r.Routes["GET"].AddNode(path, RouteHandlers{})
}

func (r *Router) GetHandler(path, method string) interface{} {
	return r.Tree.GetNode(path, method)
}
