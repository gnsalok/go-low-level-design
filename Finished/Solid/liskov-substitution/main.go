package main

/* Liskov Substitution:
This principle states that objects of a superclass should be replaceable with objects of a subclass.
*/
import "fmt"

type AnimalBehaviour interface {
	MakeSound()
}

type Animal struct {
	Name string
}

func (a Animal) MakeSound() {
	fmt.Println("Animal sound")
}

type Bird struct {
	Animal
}

func (b Bird) MakeSound() {
	fmt.Println("Chirp chirp")
}

func MakeSound(ab AnimalBehaviour) {
	ab.MakeSound()
}

func main() {
	a := Animal{}
	b := Bird{}
	MakeSound(a)
	MakeSound(b)
}
