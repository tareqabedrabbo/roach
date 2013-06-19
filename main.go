package main

import (
	"log"
	"net/http"
	"roach/hashtable"
)

func main() {
	db := new(hashtable.Hashtable)
	log.Printf("created %+v\n", db)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
