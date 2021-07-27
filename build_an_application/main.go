package main

import (
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

	db, err := os.OpenFile(dir+"/"+currentFolder+dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}
	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
