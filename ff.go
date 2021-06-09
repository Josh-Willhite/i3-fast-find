package main

import (
	"flag"
	"fmt"

	"go.i3wm.org/i3/v4"
)

type windowNames struct {
	names []string
}

func (w *windowNames) getWindows(n *i3.Node) {
	if n.Type == "con" && n.Name != "content" && n.Name != "" && n.WindowProperties.Class != "" {
		w.names = append(w.names, n.WindowProperties.Class)
	}
	for _, c := range n.Nodes {
		w.getWindows(c)
	}
}

func main() {
	var focus string
	flag.StringVar(&focus, "focus", "", "focus window")
	flag.Parse()

	if focus != "" {
		_, err := i3.RunCommand(fmt.Sprintf("[class=\"%s\"] focus", focus))
		if err != nil {
			fmt.Printf("Failed with: %s", err)
		}
		return
	}

	n := windowNames{names: []string{}}

	tree, _ := i3.GetTree()
	n.getWindows(tree.Root)

	for _, name := range n.names {
		fmt.Println(name)
	}
}
