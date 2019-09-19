package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Welcome Krish..!</b>")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", router))
	//fmt.Println("Another test from Krish using command line ..!!")
}
