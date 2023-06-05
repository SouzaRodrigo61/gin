package main

import (
	"log"
	"net/http"
	"strings"
)

type httpMethod string
type urlPattern string

type routeRules struct {
	methods map[httpMethod]http.Handler
}

type router struct {
	routes map[urlPattern]routeRules
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	foundRoute, exists := r.routes[urlPattern(req.URL.Path)]
	if !exists {
		http.NotFound(w, req)
		return
	}
	handler, exists := foundRoute.methods[httpMethod(req.Method)]
	if !exists {
		notAllowed(w, req, foundRoute)
		return
	}
	handler.ServeHTTP(w, req)
}

func (r *router) HandleFunc(method httpMethod, pattern urlPattern, f func(w http.ResponseWriter, req *http.Request)) {
	rules, exists := r.routes[pattern]
	if !exists {
		rules = routeRules{methods: make(map[httpMethod]http.Handler)}
		r.routes[pattern] = rules
	}
	rules.methods[method] = http.HandlerFunc(f)
}

func notAllowed(w http.ResponseWriter, req *http.Request, r routeRules) {
	methods := make([]string, 1)
	for k := range r.methods {
		methods = append(methods, string(k))
	}
	w.Header().Set("Allow", strings.Join(methods, " "))
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func New() *router {
	return &router{routes: make(map[urlPattern]routeRules)}
}

func handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	r := New()
	r.HandleFunc(http.MethodPost, "/test", handler)
	log.Print("Listening... ")
	http.ListenAndServe(":8000", r)
}
