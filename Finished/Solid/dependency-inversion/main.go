package main

/*
This principle states that high-level modules should, but rather both should depend on abstractions.
This helps to reduce the coupling between components and make the code more flexible and maintainable.


How to implement :

Assume :

You can two type :
1. Worker
2. Supervisor and they implement some methods


Now let's assume you have type Department which is high level module. Now it's is anti-pattern if you implement
like below :

type Department struct {
  Workers []Worker
  Supervisors []Supervisor
}


How to fix that ? Well Using Interface:

1. Make Employee as an interface

type Employee interface {
  GetID() int
  GetName() string
}

2. Under department use Employees

type Department struct {
  Employees []Employee
}

*/

type Employee interface {
	GetID() int
	GetName() string
}

type Department struct {
	Employees []Employee
}

func main() {

}
