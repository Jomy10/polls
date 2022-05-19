package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jomy10/polls/api"
	"github.com/jomy10/polls/middleware"
)

func main() {
	fmt.Println("Starting server at port ")
	authenticated := middleware.NewAuth(os.Getenv("SECRET_KEY"))

	http.Handle("/api/vote", authenticated(http.HandlerFunc(api.VoteHandler)))
	http.HandleFunc("/api/info", api.PollInfoHandler)

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
