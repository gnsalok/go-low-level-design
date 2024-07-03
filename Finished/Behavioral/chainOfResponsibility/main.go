package main

func main() {
	medical := &Medical{}

	//Set next for doctor department
	doctor := &Doctor{}
	medical.setNext(doctor)

	patient := &Patient{name: "alex"}

	//Patient visiting
	medical.execute(patient)

}
