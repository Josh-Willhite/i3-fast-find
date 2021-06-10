package main

import (
	"flag"
	"fmt"
	"strings"

	"go.i3wm.org/i3/v4"
)

type windowNames struct {
	names map[string]string
}

func (w *windowNames) getWindows(n *i3.Node) {
	title := n.WindowProperties.Title

	if title != "" && !strings.Contains(title, "i3bar") {
		w.names[strings.ToLower(title)] = title
	}
	for _, c := range n.Nodes {
		w.getWindows(c)
	}
}

func main() {
	var focus string
	flag.StringVar(&focus, "focus", "", "focus window")
	flag.Parse()

	n := windowNames{names: map[string]string{}}

	tree, _ := i3.GetTree()
	n.getWindows(tree.Root)

	if focus != "" {
		_, err := i3.RunCommand(fmt.Sprintf("[title=\"%s\"] focus", n.names[focus]))
		if err != nil {
			fmt.Printf("Failed with: %s", err)
		}
		return
	}

	for k, _ := range n.names {
		fmt.Println(k)
	}
}
