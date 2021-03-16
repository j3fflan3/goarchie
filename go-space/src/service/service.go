package service

import "fmt"

// Person is a struct containing information for a person.
type Person struct {
	First string
}

// Accessor is an interface for saving and retrieving a persons information.
type Accessor interface {
	Save(int, Person)
	Retrieve(int) Person
}

// PersonService is a struct that contains the Accessor
type PersonService struct {
	a Accessor
}

// NewPersonService is a constructor to assign an implementation of an Accessor to a PersonService object.
func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

// Get is a method PersonService receiver method to retrieve a person of id n. If n does not
// exist, an error is returned.
func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("no person with n of %d", n)
	}
	return p, nil
}

// Put saves a person with id of n to the underlying db via Accessor
func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}
func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}
