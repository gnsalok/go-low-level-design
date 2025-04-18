package main

/*
Factory Pattern
- Factory pattern allows **loose coupling** between client and implementation.
- Adding a new notifier (e.g., SlackNotifier) only requires implementing `Notifier` and updating the factory.
- The client code remains unchanged.
- This pattern is useful when the exact type of object to be created is not known until runtime.
- It encapsulates the object creation process, allowing for better separation of concerns.
*/

import "fmt"

// 1. Define Interface
type Notifier interface {
	Send(message string) string
}

// 2. Define Concrete Implementations

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) string {
	return "Email sent with message: " + message
}

type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) string {
	return "SMS sent with message: " + message
}

// 3. Define Factory Function
func NewNotifier(notifierType string) Notifier {
	switch notifierType {
	case "email":
		return &EmailNotifier{}
	case "sms":
		return &SMSNotifier{}
	default:
		return nil
	}
}

func main() {
	var notifier Notifier

	notifier = NewNotifier("email")
	fmt.Println(notifier.Send("Hello via Email"))

	notifier = NewNotifier("sms")
	fmt.Println(notifier.Send("Hello via SMS"))
}
