package main

import (
	"log"
	"net/http"

	poker "github.com/Brett-Tanner/go_with_tests/firstapp"
)

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(poker.DbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatalf("could not listen on port 3000 %v", err)
	}
}
