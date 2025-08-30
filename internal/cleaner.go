package cleaner

import (
	"fmt"
	"os"

	queue "github.com/adrian-borovnik/node-modules-cleaner/pkg"
)

type Cleaner struct {
	WorkingDir string
	Paths      []string
}

func NewCleaner(workingDir string) *Cleaner {
	return &Cleaner{
		WorkingDir: workingDir,
		Paths:      []string{},
	}
}

func (c *Cleaner) Search() []string {
	q := queue.NewQueue[string]()
	q.Enqueue(c.WorkingDir)

	var nodeModules []string
	for curr, err := q.Dequeue(); err == nil; curr, err = q.Dequeue() {
		c, err := os.ReadDir(curr)
		if err != nil {
			// fmt.Println("Failed to read directory:", err)
			continue
		}

		for _, entity := range c {
			if !entity.IsDir() {
				continue
			}

			if entity.Name() == "node_modules" {
				nodeModules = append(nodeModules, curr+"/"+entity.Name())
				continue
			}

			q.Enqueue(curr + "/" + entity.Name())
		}
	}

	c.Paths = nodeModules
	return nodeModules
}

func (c *Cleaner) Clean() {
	for _, path := range c.Paths {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
		}
	}
}
