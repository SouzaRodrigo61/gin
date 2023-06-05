package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vardius/gorouter/v4"
	"github.com/vardius/gorouter/v4/context"
)

func logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func example(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// do smth
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request) {
	params, _ := context.Parameters(r.Context())
	fmt.Fprintf(w, "hello, %s!\n", params.Value("name"))
}

func main() {
	// apply middleware to all routes
	// can pass as many as you want
	router := gorouter.New(logger, example)

	router.GET("/", http.HandlerFunc(index))
	router.GET("/hello/{name}", http.HandlerFunc(hello))
	port := os.Getenv("PORT")
	log.Printf("port: " + port)
	if port == "" {
		port = "8080"
	}
	formatedPort := ":" + port
	log.Printf("formatedPort: " + formatedPort)


	log.Fatal(http.ListenAndServe(formatedPort, router))
}
