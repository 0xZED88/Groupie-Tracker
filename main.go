package main

import (
	"fmt"
	"net/http"

	"groupie/handlers"
)

func main() {
	////serve style
	http.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./styles"))))

	///server
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artists/", handlers.GetArtist)
	handlers.GetArtistsData()
	const PORT = "8000"
	fmt.Printf("running on %v...", PORT)
	http.ListenAndServe("localhost:"+PORT, nil)

	///Template
}
