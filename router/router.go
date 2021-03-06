package router

import (
	"net/http"

	"github.com/AnthonyCapirchio/t-mux/tree"
)

type Router struct {
	Tree *tree.TreeNode
}

type RouteHandler func(w http.ResponseWriter, r *http.Request, params map[string]string)

type RouteHandlers map[string]interface{}

func NewRouter() *Router {
	return &Router{
		Tree: tree.NewTree(),
	}
}

func (r *Router) Get(path string, handler tree.Handler) {
	r.Tree.AddNode(path, "GET", handler)
}

func (r *Router) Post(path string, handler tree.Handler) {
	r.Tree.AddNode(path, "POST", handler)
}

func (r *Router) Put(path string, handler tree.Handler) {
	r.Tree.AddNode(path, "PUT", handler)
}

func (r *Router) Delete(path string, handler tree.Handler) {
	r.Tree.AddNode(path, "DELETE", handler)
}

func (r *Router) GetHandler(path, method string) (tree.Handler, map[string]string) {
	handler, params := r.Tree.GetNode(path, method)

	return handler, params
}
