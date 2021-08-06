package main

import (
	"../.."
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"
const currentFolder = "build_an_application/"

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	store, closeFileFunc, err := poker.FileSystemPlayerStoreFromFile(dir + "/" + dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFileFunc()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
