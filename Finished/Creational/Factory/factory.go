package main

import "fmt"

/*
- Using factory we delegate the creation of object to our factory.
- With this approach, we can add more types of object as long as they implement the interface.

Benefits : Make application loosely coupled
*/

func newPublication(pubType string, name string, pg int, pub string) (iPublication, error) {
	// Create the right kind of publication based on the given type
	if pubType == "newspaper" {
		return createNewspaper(name, pg, pub), nil
	}
	if pubType == "magazine" {
		return createMagazine(name, pg, pub), nil
	}
	return nil, fmt.Errorf("No such publication type")
}
