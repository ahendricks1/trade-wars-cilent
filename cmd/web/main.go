package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/grid", grid)
	mux.HandleFunc("/players", playersHandler)

	godotenv.Load()

	PORT := os.Getenv("PORT")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server...")
	err := http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}
