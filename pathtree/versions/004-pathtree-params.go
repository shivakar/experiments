package main

import (
	"fmt"
	"strings"
	"time"
)

// Handler ...
type Handler func() []string

type node struct {
	// Cannot use a map when using path params like :name
	// For example, for route /user/:name and an input of
	// /user/foo, we cannot check for foo in the map since
	// it is an instance for the param :name
	// and there is no way to realize to search for :name
	//children  map[string]*node
	children  []*node
	component string
	path      string
	isParam   bool
	handler   Handler
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
	components := strings.Split(path, "/")[1:]

	if len(components) == 1 && n.component == components[0] {
		return
	}

	pn := n
	for i, c := range components {
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
			// Insert new child
			child := &node{component: c,
				handler: func() []string { return components[:i+1] }}
			if c[0] == ':' {
				child.isParam = true
			}
			pn.children = append(pn.children, child)
			pn = child
		}
	}
}

func (n *node) getHandler(path string) Handler {
	components := strings.Split(path, "/")[1:]
	count := len(components)

	pn := n
	for count > 0 {
		componentFound := false
		for _, child := range pn.children {
			if child.component == components[0] || child.isParam {
				pn = child
				components = components[1:]
				count--
				componentFound = true
				break
			}
		}
		if !componentFound {
			return nil
		}
	}

	return pn.handler
}

func main() {
	// the pathtree starts with a root of "/".
	// Here we are removing slashes from all paths therefore name=""
	tree := &node{component: ""}

	/*
		data, err := ioutil.ReadFile("golang_src_paths.txt")
		if err != nil {
			panic(err)
		}
		paths := strings.Split(string(data), "\n")
	*/

	paths := []string{
		"/faq",
		"/user/new",
		"/user/:name",
	}

	st := time.Now()
	nPaths := 0
	for _, path := range paths {
		if len(path) == 0 || path[len(path)-3:] == ".go" {
			continue
		}
		tree.add(path)
		nPaths++
	}
	dur := time.Since(st)
	// Print the tree
	fmt.Println("\n\nPathtree after inserts: ")
	tree.print("")

	// Print total time
	fmt.Printf("\n\nTotal time for inserting %d paths:  %v\n", nPaths, dur)

	st = time.Now()
	nPaths = 0
	paths = []string{
		"/faq",
		"/faqs",
		"/user/foo",
		"/usr/foo",
		"/users/foo",
		"/user/foo/hello",
		"/user/new",
	}
	for _, path := range paths {
		if len(path) == 0 || path[len(path)-3:] == ".go" {
			continue
		}
		handler := tree.getHandler(path)
		if handler != nil {
			fmt.Println(path, handler())
		} else {
			fmt.Println(path, handler)
		}
		_ = tree.getHandler(path)
		nPaths++
	}
	dur = time.Since(st)
	fmt.Printf("\n\nTotal time for retrieving  %d paths:  %v\n", nPaths, dur)
}
