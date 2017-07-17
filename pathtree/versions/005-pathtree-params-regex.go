package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Handler ...
type Handler func() []string

// node ...
type node struct {
	children  []*node
	component string
	isParam   bool
	regex     *regexp.Regexp
	handler   Handler
}

// newNode returns a new node
func newNode() *node {
	return &node{
		children:  nil,
		component: "",
		isParam:   false,
		regex:     nil,
		handler:   nil,
	}
}

// print prints the tree
func (n *node) print(indent string) {
	fmt.Printf("%s/%s\n", indent, n.component)
	for _, c := range n.children {
		// Add two spaces to the indent for the children
		c.print(indent + "  ")
	}
}

func newUserHandler() []string {
	return []string{"new", "users"}
}

func existingUserHandler() []string {
	return []string{"existing", "users"}
}

// add adds a new path to the pathtree
// returns an error if the path could not be added
func (n *node) add(path string) error {
	if path[0] != '/' {
		return errors.New("add requires an absolute path")
	}
	// Split path into components
	components := strings.Split(path, "/")

	// Start at the current node and then descent into children
	pn := n

	// If the current node's component is not equal to
	// components[0] we are probably checking at the wrong node
	if pn.component != components[0] {
		return errors.New("pn.component != components[0] in add")
	}
	for i, c := range components {
		// c == "" if multiple slashes are specified in the path
		if i == 0 || c == "" {
			continue
		}

		// Handle param and regex
		params := []string{}
		if c[0] == ':' {
			// The format should be :var:regexp
			// Using SplitN instead of Split so that : can be included
			// in the regex as it is a valid character in URLs
			// Therefore, there will only be two elements in params
			params = strings.SplitN(c, ":", 3)[1:]
			c = ":" + params[0]
		}

		// Check if this component is already inserted
		inserted := false
		for _, child := range pn.children {
			if child.component == c {
				pn = child
				inserted = true
				break
			}
		}

		//fmt.Println("here with i: ", i, " path: ", path, " and component: ", c,
		//" and inserted: ", inserted)
		// if not already inserted, insert it
		if !inserted {
			child := newNode()
			child.component = c
			//fmt.Println("Setting up component: ", c, " handler: ", components[1:i+1])

			// Handle param and regex
			if c[0] == ':' {
				child.isParam = true
				// Check if it is a regex
				if len(params) > 1 {
					regex, err := regexp.Compile(params[1])
					if err != nil {
						return err
					}
					child.component = c
					child.regex = regex
				}
			}
			pn.children = append(pn.children, child)
			pn = child
		}

		// Add or Update handler
		if i == len(components)-1 {
			pn.handler = func() []string { return components[1 : i+1] }
		}
	}

	return nil
}

// getHandler returns the Handler registered for the given path if any
// returns nil if no handler could be found
// returns an optional error when things go wrong
func (n *node) getHandler(path string) (Handler, error) {
	components := strings.Split(path, "/")

	pn := n
	if pn.component != components[0] {
		return nil, errors.New("n.component != components[0] in getHandler")
	}
	for i, c := range components {
		if i == 0 || c == "" {
			continue
		}
		componentFound := false
		for _, child := range pn.children {
			if child.component == c || child.isParam {
				if child.regex != nil {
					if !child.regex.Match([]byte(c)) {
						return nil, errors.New("Regex match failed in getHandler")
					}
				}
				pn = child
				componentFound = true
				break
			}
		}
		if !componentFound {
			return nil, nil
		}
	}
	return pn.handler, nil
}

// main entry point to the application
func main() {
	tree := newNode()
	paths := []string{
		"/faq",
		"/user/new",
		"/user/:name",
		"/source/new",
		`/source/:name/:id:[0-9]+`,
		`/source/:name:\w+`,
		`/foo/:name:[a-z:0-9]+`, // : is part of the regex
		`/source/:name/:id:[0-9]+`,
	}

	st := time.Now()
	nPaths := 0
	for _, path := range paths {
		err := tree.add(path)
		if err != nil {
			panic(err)
		}
		nPaths++
	}
	dur := time.Since(st)
	// Print the tree
	fmt.Println("\n\nPathtree after inserts: ")
	tree.print("")

	// Print total time
	fmt.Printf("\nTotal time for inserting %d paths: %v\n\n\n", nPaths, dur)

	// Test the paths
	nPaths = 0
	paths = []string{
		"/faq",                // match /faq
		"/faqs",               // nil handler
		"/user/foo",           // match /user/:name
		"/user/new",           // match /user/new
		"/user",               // nil
		"/usr/foo",            // nil
		"/users/foo",          // nil
		"/user/foo/hello",     // nil
		"/source/bar",         // match /source/:name:\w+
		"/source/123",         // match /source/:name:\w+
		"/source/123abc",      // match /source/:name:\w+
		"/source/abc123",      // match /source/:name:\w+
		"/source/hello world", // match /source/:name:\w+
		"/source/abc123",      // match /source/:name:\w+
		"/source/bar/123",     // match /source/:name/:id[0-9]+
		"/source/bar/baz",     // nil /source/:name/:id[0-9]+ needs digits

	}

	st = time.Now()
	for _, path := range paths {
		handler, err := tree.getHandler(path)
		if handler == nil {
			fmt.Println(path, handler, err)
		} else {
			fmt.Println(path, handler(), err)
		}
		nPaths++
	}
	dur = time.Since(st)
	fmt.Printf("\n\nTotal time for retrieving %d paths: %v\n", nPaths, dur)
}
