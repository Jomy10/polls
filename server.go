package main

import (
	"fmt"
	"github.com/jomy10/poll/api"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting server at port ")
	http.HandleFunc("/api/vote", api.VoteHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server at %s\n", port)

	// Serve
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
