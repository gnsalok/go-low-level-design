package main

import "fmt"

// Define the television interface
type television interface {
	turnOn()
	turnOff()
	volumeUp()
	volumeDown()
	channelUp()
	channelDown()
	goToChannel(channel int)
}

// SammysangTV represents a TV with a different interface
type SammysangTV struct {
	currentChan   int
	currentVolume int
	tvOn          bool
}

// SohneeTV implements the television interface directly
type SohneeTV struct {
	vol     int
	channel int
	isOn    bool
}

// Adapter for SammysangTV to conform to the television interface
type sammysangAdapter struct {
	// IMP : you can assume this is a third-party library which you cannot modify
	sstv *SammysangTV
}

func (s *sammysangAdapter) turnOn() {
	// Adapted method to turn on the TV
	fmt.Println("Turning on SammysangTV")
	s.sstv.tvOn = true
}

func (s *sammysangAdapter) turnOff() {
	// Adapted method to turn off the TV
	fmt.Println("Turning off SammysangTV")
	s.sstv.tvOn = false
}

func (s *sammysangAdapter) volumeUp() {
	// Adapted method to increase volume
	fmt.Println("Increasing SammysangTV volume")
	s.sstv.currentVolume++
}

func (s *sammysangAdapter) volumeDown() {
	// Adapted method to decrease volume
	fmt.Println("Decreasing SammysangTV volume")
	s.sstv.currentVolume--
}

func (s *sammysangAdapter) channelUp() {
	// Adapted method to increase channel
	fmt.Println("Increasing SammysangTV channel")
	s.sstv.currentChan++
}

func (s *sammysangAdapter) channelDown() {
	// Adapted method to decrease channel
	fmt.Println("Decreasing SammysangTV channel")
	s.sstv.currentChan--
}

func (s *sammysangAdapter) goToChannel(channel int) {
	// Adapted method to go to a specific channel
	fmt.Printf("Changing SammysangTV to channel %d\n", channel)
	s.sstv.currentChan = channel
}

// SohneeTV methods implementing the television interface
func (s *SohneeTV) turnOn() {
	fmt.Println("Turning on SohneeTV")
	s.isOn = true
}

func (s *SohneeTV) turnOff() {
	fmt.Println("Turning off SohneeTV")
	s.isOn = false
}

func (s *SohneeTV) volumeUp() {
	fmt.Println("Increasing SohneeTV volume")
	s.vol++
}

func (s *SohneeTV) volumeDown() {
	fmt.Println("Decreasing SohneeTV volume")
	s.vol--
}

func (s *SohneeTV) channelUp() {
	fmt.Println("Increasing SohneeTV channel")
	s.channel++
}

func (s *SohneeTV) channelDown() {
	fmt.Println("Decreasing SohneeTV channel")
	s.channel--
}

func (s *SohneeTV) goToChannel(channel int) {
	fmt.Printf("Changing SohneeTV to channel %d\n", channel)
	s.channel = channel
}

func main() {
	// Create instances of the two TV types with some default values
	tv1 := &SammysangTV{
		currentChan:   13,
		currentVolume: 35,
		tvOn:          true,
	}
	tv2 := &SohneeTV{
		vol:     20,
		channel: 9,
		isOn:    true,
	}

	// SohneeTV directly implements the television interface
	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.channelUp()
	tv2.channelDown()
	tv2.goToChannel(68)
	tv2.turnOff()

	fmt.Println("--------------------")

	// Use the adapter for SammysangTV
	ssAdapt := &sammysangAdapter{
		sstv: tv1,
	}
	ssAdapt.turnOn()
	ssAdapt.volumeUp()
	ssAdapt.volumeDown()
	ssAdapt.channelUp()
	ssAdapt.channelDown()
	ssAdapt.goToChannel(68)
	ssAdapt.turnOff()
}
