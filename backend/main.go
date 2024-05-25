package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	"chainlink-git-tasks/githubauth"
)

func main() {
	router := httprouter.New()

	router.POST("/addrepo", githubauth.AddRepo)

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":3001", handler))
}