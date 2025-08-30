package main

import (
	"fmt"
	"log"
	"os"

	cleaner "github.com/adrian-borovnik/node-modules-cleaner/internal"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory")
	}

	c := cleaner.NewCleaner(workingDir)

	fmt.Println("Searching for node_modules in the", workingDir)
	nodeModules := c.Search()

	for _, path := range nodeModules {
		fmt.Printf(" %s\n", path)
	}

	fmt.Println("Cleaning node_modules...")
	c.Clean()
}
