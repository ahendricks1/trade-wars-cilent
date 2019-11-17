package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
	}
	
	w.Write([]byte("This is a test run!"))
}

func players(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet && r.Method != http.MethodPost{
		w.Header().Set("Allow", http.MethodGet)
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Display a player snippet."))
}

func showSnippet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
        w.WriteHeader(405)
        w.Write([]byte("Method Not Allowed"))
        return
    }
	w.Write([]byte("Create a new snippet."))
}
 

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
	
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
