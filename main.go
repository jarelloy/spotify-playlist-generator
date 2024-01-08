package main

import (
	"fmt"
	"net/http"
	"spotify-go/handler"
	"spotify-go/utils"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	handler.SetupRoutes(r)

	go func() {
		client := <-handler.Ch
		utils.SearchByGenre("", client)
	}()

	http.Handle("/", r)

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
