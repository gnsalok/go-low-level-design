package main

/*
Prototype Pattern
- The Prototype pattern is a creational design pattern that allows you to create new objects by copying an existing object, known as the prototype.
- It is useful when the cost of creating a new object is more expensive than copying an existing one.
- The Prototype pattern is particularly useful when you want to create a large number of similar objects or when the object creation process is complex.
*/

import "fmt"

// Prototype interface
type Cloner interface {
	Clone() Cloner
	Print()
}

// Concrete prototype
type Document struct {
	Title   string
	Content string
}

// Clone method (returns a copy of Document)
func (d *Document) Clone() Cloner {
	// shallow copy is enough here, deep copy if needed
	clone := *d
	return &clone
}

func (d *Document) Print() {
	fmt.Printf("Document Title: %s\nContent: %s\n", d.Title, d.Content)
}

func main() {
	// Original document
	doc1 := Document{
		Title:   "Prototype Pattern",
		Content: "This is the original content.",
	}

	doc1.Print()

	// Cloning the document
	doc2 := doc1.Clone().(*Document)
	doc2.Title = "Cloned Document"
	doc2.Content = "This is the cloned version."

	doc2.Print()

	// Original remains unchanged
	doc1.Print()
}
