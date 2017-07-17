package main

import (
	"fmt"
	"strings"
)

type node struct {
	children  []*node
	component string
}

func (n *node) print(indent string) {
	fmt.Printf("%s/%s\n", indent, n.component)
	for _, c := range n.children {
		// Add two spaces to the indent of this node's children
		c.print(indent + "  ")
	}
}

// add the path to the path tree
func (n *node) add(path string) {
	fmt.Println("Inserting path:", path)
	components := strings.Split(path, "/")[1:]

	if len(components) == 1 && n.component == components[0] {
		return
	}

	pn := n
	for _, c := range components {
		if c == "" {
			continue
		}
		inserted := false
		for _, child := range pn.children {
			if child.component == c {
				pn = child
				inserted = true
				break
			}
		}
		if !inserted {
			child := node{component: c}
			pn.children = append(pn.children, &child)
			pn = &child
		}
	}
}

func main() {
	// the pathtree starts with a root of "/".
	// Here we are removing slashes from all paths therefore name=""
	tree := &node{component: ""}

	// Add a new path to the tree
	tree.add("/hello")

	// Adding "/" should do nothing
	tree.add("/")

	// Adding some more paths
	tree.add("/users")
	tree.add("/hello/world")
	tree.add("users/world")
	tree.add("/how/are/you/doing")
	tree.add("/abc//def//ghi")

	// Print the tree
	fmt.Println("\n\nPathtree after inserts: ")
	tree.print("")
}
