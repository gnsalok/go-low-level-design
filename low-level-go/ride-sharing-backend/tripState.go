package main

type TripState string

const (
	Requested      TripState = "REQUESTED"
	DriverAssigned TripState = "DRIVER_ASSIGNED"
	InProgress     TripState = "IN_PROGRESS"
	Completed      TripState = "COMPLETED"
	Cancelled      TripState = "CANCELLED"
)
