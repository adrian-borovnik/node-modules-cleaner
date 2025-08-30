package main

import (
	"fmt"
	"log"
	"os"

	queue "github.com/adrian-borovnik/node-modules-cleaner/pkg"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory")
	}

	fmt.Println("Searching for node_modules in the", workingDir)

	nodeModules := findNodeModules(workingDir)

	for _, path := range nodeModules {
		fmt.Printf(" %s\n", path)
	}

	deleteNodeModules(nodeModules)
}

func findNodeModules(path string) []string {
	q := queue.NewQueue[string]()
	q.Enqueue(path)

	var nodeModules []string
	for curr, err := q.Dequeue(); err == nil; curr, err = q.Dequeue() {
		c, err := os.ReadDir(curr)
		if err != nil {
			log.Fatal("Failed to read directory")
		}

		for _, entity := range c {
			if entity.IsDir() {
				if entity.Name() == "node_modules" {
					nodeModules = append(nodeModules, curr+"/"+entity.Name())
					continue
				}

				// fmt.Printf("Found directory: %s\n", curr+"/"+entity.Name())
				q.Enqueue(curr + "/" + entity.Name())
			}
		}
	}

	return nodeModules
}

func deleteNodeModules(paths []string) {
	for _, path := range paths {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
		}
	}

}
