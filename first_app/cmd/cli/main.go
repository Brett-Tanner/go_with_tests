package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/Brett-Tanner/go_with_tests/firstapp"
)

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(poker.DbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println(("Let's play poker!"))
	fmt.Println(("Type '{name} wins' to record a win"))

	poker.NewCLI(store, os.Stdin).PlayPoker()
}
