// Package main starts the simple server on port and serves HTML,
// CSS, and JavaScript to clients.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Interceptor struct {
	origWriter http.ResponseWriter
	overridden bool
}

// templates parses the specified templates and caches the parsed results
// to help speed up response times.
var templates = template.Must(template.ParseFiles("./templates/base.html", "./templates/searchBody.html"))

func (i *Interceptor) WriteHeader(rc int) {
	switch rc {
	case 500:
		http.Error(i.origWriter, "Custom 500 message / content", 500)
	case 404:
		http.Error(i.origWriter, "Custom 404 message", 404)
	case 403:
		i.origWriter.WriteHeader(403)
		fmt.Fprintln(i.origWriter, "Custom 403 message")
	default:
		i.origWriter.WriteHeader(rc)
		return
	}
	// if the default case didn't execute (and return) we must have overridden the output
	i.overridden = true
	log.Println(i.overridden)
}

func (i *Interceptor) Write(b []byte) (int, error) {
	if !i.overridden {
		return i.origWriter.Write(b)
	}

	// Return nothing if we've overriden the response.
	return 0, nil
}

func (i *Interceptor) Header() http.Header {
	return i.origWriter.Header()
}

func ErrorHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w = &Interceptor{origWriter: w}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// logging is middleware for wrapping any handler we want to track response
// times for and to see what resources are requested.
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)
		next.ServeHTTP(w, r)
		log.Println(req, "completed in", time.Now().Sub(start))
	})
}

// search is the handler responsible for rending the search results page for the site.
func search() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("search results: couldn't parse template: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

// public serves static assets such as CSS and JavaScript to clients.
func public() http.Handler {
	return http.StripPrefix("/pub/", http.FileServer(http.Dir("./pub")))
}

// public serves static assets such as CSS and JavaScript to clients.
func publicRedir() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "./pub", 301)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/pub/", logging(public()))
	mux.Handle("/pub/search", logging(search()))
	mux.Handle("/", logging(publicRedir()))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
/*
	addr := fmt.Sprintf(":%s", port)
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Println("Dore server on port", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start Dore server: %v\n", err)
	}
*/
	log.Println("Dore server on port", port)
	log.Fatal(http.ListenAndServe("localhost:8080", ErrorHandler(mux)))
}
