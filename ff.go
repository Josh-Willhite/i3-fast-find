package main

import (
	"flag"
	"log"
	"strings"

	"go.i3wm.org/i3/v4"
)

func main() {
	tree, _ := i3.GetTree()
	match := tree.Root.FindChild(func(n *i3.Node) bool {
		return strings.Contains(n.Name, "Emacs")
	},
	)

	if match != nil {
		log.Println("Found match! ", match.Name)
		return
	}

	log.Fatal("Failed to find match")
}
