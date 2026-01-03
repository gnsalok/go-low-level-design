package main

import "errors"

// Core entities
type User struct {
	Name string
}

type Driver struct {
	Name string
}

// Concrete type
type Trip struct {
	ID              string
	Rider           *User
	Driver          *Driver
	PickupLocation  *Location
	DropoffLocation *Location
	State           TripState
	Fare            float64
}

func NewTrip(id string, rider *User, pickup, dropoff *Location) *Trip {
	return &Trip{
		ID:              id,
		Rider:           rider,
		PickupLocation:  pickup,
		DropoffLocation: dropoff,
		State:           Requested,
		Fare:            0.0,
	}
}

func (t *Trip) AssignDriver(driver *Driver) error {
	if t.State != Requested {
		return errors.New("cannot assign driver in current state")
	}
	t.Driver = driver
	t.State = DriverAssigned
	return nil
}

func (t *Trip) StartTrip() error {
	if t.State != DriverAssigned {
		return errors.New("cannot start trip in current state")
	}
	t.State = InProgress
	return nil
}

func (t *Trip) CompleteTrip(calculator PriceCalculator) error {
	if t.State != InProgress {
		return errors.New("cannot complete trip in current state")
	}
	t.Fare = calculator.CalculateFare(t)
	t.State = Completed
	return nil
}

func (t *Trip) CancelTrip() error {
	if t.State == Completed {
		return errors.New("cannot cancel completed trip")
	}
	t.State = Cancelled
	return nil
}
