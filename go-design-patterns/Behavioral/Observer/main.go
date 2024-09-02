package main

/*

// Link : https://refactoring.guru/design-patterns/observer/go/example#example-0

- Subject, the instance which publishes an event when anything happens.
- Observer, which subscribes to the subject events and gets notified when they happen.

* Subject ; Concrete Subject -> Item
* Observer ; Concrete Observer -> Customer
* main.go -> client


*/

func main() {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
