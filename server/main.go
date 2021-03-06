// Package main starts the simple server on port and serves HTML,
// CSS, and JavaScript to clients.
package main

import (
	"os"
	"time"
	"fmt"
    "strings"
    "strconv"
    "io/ioutil"
	"log"
	"net/http"
	"html/template"
    "encoding/json"
)

type InvertedIndexEntry struct {
	Term            string
	Frequency       int
	DocumentListing []int
}

type InvertedIndex struct {
	DocList []string
	TopicName map[string]string
	HashMap map[string]*InvertedIndexEntry
	Items   []*InvertedIndexEntry
}

func termNotFound() {
    if r := recover(); r!= nil {
        log.Println("Term not found")
    }
}

type Interceptor struct {
	origWriter http.ResponseWriter
	overridden bool
}

func (i *Interceptor) WriteHeader(rc int) {
	log.Printf("%d\n", rc)
	switch rc {
	case 500:
		http.Error(i.origWriter, "Error:  500 Internal server error.", 500)
	case 404:
		http.Error(i.origWriter, "Error:  404 Requested page does not exist.\n\tReturn to /pub/index.html", 404)
	default:
		i.origWriter.WriteHeader(rc)
		return
	}
	// if the default case didn't execute and return, must have overridden the output
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

// page request error handler
func ErrorHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w = &Interceptor{origWriter: w}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// track response  times and see what resources are requested
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)
		next.ServeHTTP(w, r)
		log.Println(req, "completed in", time.Now().Sub(start))
	})
}

var queryString string

// template function: populate search results
func renderResults() template.HTML {

	var invertedIndex InvertedIndex

	defer termNotFound()

	log.Printf("QUERYSTRING: %s\n", queryString)

	term := strings.ToLower(queryString)

	//
	// get inverted index
	//

	jsonFile, err := os.Open("./pub/searchIndex.json")
	if err != nil {
	    log.Println(err)
	}
	log.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &invertedIndex)

//	fmt.Printf("%d topics match the search term:\n", len(invertedIndex.HashMap[term].DocumentListing))

	//
	// populate search results page
	//

    newhtml := ""
//	newhtml += `<h2>` + strconv.Itoa(invertedIndex.HashMap[term].Frequency) + ` occurrences of "` + queryString + `" in `+ strconv.Itoa(len(invertedIndex.HashMap[term].DocumentListing)) + ` topics:</h2>`
	newhtml += `<h2>"` + queryString + `" found in `+ strconv.Itoa(len(invertedIndex.HashMap[term].DocumentListing)) + ` topics:</h2>`

	for _, docID := range invertedIndex.HashMap[term].DocumentListing {
//		fmt.Println(invertedIndex.TopicName[invertedIndex.DocList[docID]])
//		fmt.Println(invertedIndex.DocList[docID])

		// target topic title

	 	newhtml += `<div class="searchTopicTitle"><a class="reference internal" href="` + invertedIndex.DocList[docID] + `.html?highlight=` + term + `">` + invertedIndex.TopicName[invertedIndex.DocList[docID]] + `</a>`
	 	newhtml += `</div>`
	}
	return template.HTML(newhtml)
}

// search is the handler responsible for rending the search results page for the site.
func search() http.Handler {
	templates, err := template.New("").Funcs(template.FuncMap{
	    "renderResults": renderResults,
	}).ParseFiles("./pub/static/templates/base.html", "./pub/static/templates/searchBody.html")
//	templates, err := template.ParseFiles("./pub/static/templates/base.html", "./pub/static/templates/searchBody.html")
    if err != nil {
        log.Println(err.Error())
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        	http.Error(w, fmt.Sprintf("Error: Search template %v", err), http.StatusInternalServerError)
        })
    }
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queryString = r.FormValue("q")
		err := templates.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: Parsing search template %v", err), http.StatusInternalServerError)
			return
		}
//		w.WriteHeader(http.StatusOK)
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
