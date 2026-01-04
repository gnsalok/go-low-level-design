package main

import "fmt"

/*
Builder Pattern
- TA builder is a helper that lets you create a complex object step by step without worrying about the order or messy construction details.
- It's used when an object has many optional parts or configuration choices.
- It separates the construction of a complex object from its representation, allowing the same construction process to create different representations.
- The Builder pattern is particularly useful when an object needs to be created with many optional parameters or when the construction process involves multiple steps.
- In this example, we will create a NotificationBuilder that allows us to build a Notification object step by step.
*/

func main() {
	var bldr = newNotificationBuilder()
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetTitle("This is a subtitle")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(5)
	bldr.SetMessage("This is a basic notification")
	bldr.SetType("alert")

	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating the notification:", err)
	} else {
		fmt.Printf("Notification: %+v\n", notif)
	}
}
