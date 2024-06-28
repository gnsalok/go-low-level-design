package main

type Employee struct {
	Name    string
	Salary  float64
	Address string
}

// ------- Can be broken into -- Look below

type EmployeeInfo struct {
	Name   string
	Salary float64
}

type EmployeeAddress struct {
	Address string
}

func main() {

}
