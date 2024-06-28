package main

import "fmt"

// STEP-1 :: Component Interface

// Notifier defines the interface for sending notifications.
type Notifier interface {
	Send(message string) string
}

// Step 2 : Concrete Component

// EmailNotifier is a concrete component that sends email notifications.
type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) string {
	return fmt.Sprintf("Sending email: %s", message)
}

// Step 3 : DECORATOR

// NotifierDecorator is a base decorator that implements the Notifier interface.
type NotifierDecorator struct {
	Notifier Notifier
}

func (d *NotifierDecorator) Send(message string) string {
	return d.Notifier.Send(message)
}

// Step 4 : ADD CONCRETE DECORATOR

// SMSDecorator is a concrete decorator that adds SMS notification functionality.
type SMSDecorator struct {
	NotifierDecorator
}

func (s *SMSDecorator) Send(message string) string {
	smsMessage := fmt.Sprintf("Sending SMS: %s", message)
	return fmt.Sprintf("%s\n%s", smsMessage, s.NotifierDecorator.Send(message))
}

// FacebookDecorator is a concrete decorator that adds Facebook notification functionality.
type FacebookDecorator struct {
	NotifierDecorator
}

func (f *FacebookDecorator) Send(message string) string {
	facebookMessage := fmt.Sprintf("Sending Facebook message: %s", message)
	return fmt.Sprintf("%s\n%s", facebookMessage, f.NotifierDecorator.Send(message))
}

func main() {
	// Create an EmailNotifier
	emailNotifier := &EmailNotifier{}

	// Decorate the EmailNotifier with SMSDecorator
	smsDecorator := &SMSDecorator{
		NotifierDecorator: NotifierDecorator{Notifier: emailNotifier},
	}

	// Further decorate with FacebookDecorator
	facebookDecorator := &FacebookDecorator{
		NotifierDecorator: NotifierDecorator{Notifier: smsDecorator},
	}

	// Send notifications
	message := "Hello, World!"
	result := facebookDecorator.Send(message)
	fmt.Println(result)

}
