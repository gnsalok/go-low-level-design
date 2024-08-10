package main

// Link : https://refactoring.guru/design-patterns/observer/go/example#example-0

// Define the interface for the observable type
type Observable interface {
	register(obs Observer)
	unregister(obs Observer)
	notifyAll()
}
