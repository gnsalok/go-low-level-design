package main

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true

	fmt.Println("Checkup has completed!!")
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}
