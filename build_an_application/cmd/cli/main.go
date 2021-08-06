package main

import (
	"../.."
	"fmt"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFileFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFileFunc()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
