package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	cleaner "github.com/adrian-borovnik/node-modules-cleaner/internal"
)

func main() {

	var workingDir string
	var err error
	if len(os.Args) < 2 {
		workingDir, err = os.Getwd()
		if err != nil {
			log.Fatal("Failed to get working directory")
		}
	} else {
		workingDir = os.Args[1]
		workingDir = strings.TrimRight(workingDir, "/")
	}

	c := cleaner.NewCleaner(workingDir)

	fmt.Println("Searching for node_modules in the", workingDir)
	nodeModules := c.Search()
	count := len(nodeModules)
	fmt.Println()

	if count == 0 {
		fmt.Println("No node_modules found")
		return
	}

	for _, path := range nodeModules {
		fmt.Println(path)
	}
	fmt.Printf("Found %d node_modules\n\n", count)

	var answer string
	fmt.Printf("Do you want to clean all %d node_modules? (y/n): ", count)
	fmt.Scanln(&answer)
	answer = strings.ToLower(strings.TrimSpace(answer))

	switch answer {
	case "y", "yes":
		c.Clean()
		fmt.Println("Cleanup complete")
	case "n", "no":
		fmt.Println("Skipping cleanup")
	default:
		fmt.Println("Invalid input")
	}
}
