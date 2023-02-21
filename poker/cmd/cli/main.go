package main

import (
	"fmt"
	"log"
	"os"

	"poker"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	poker.NewCli(os.Stdin, os.Stdout, game).PlayPoker()
}
