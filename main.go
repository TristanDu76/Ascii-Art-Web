package main

import (
	"ascii-art-web/handlers"
	"ascii-art-web/utils"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", routeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/download", handlers.HandleDownload)
	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlers.HandleHome(w, r)
	case "/download":
		handlers.HandleDownload(w, r)
	default:
		utils.Serve404(w, r)
	}
}
