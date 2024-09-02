package main

// Define the interface for the observable type
type Observable interface {
	register(obs Observer)
	unregister(obs Observer)
	notifyAll()
}
