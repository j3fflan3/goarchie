package postgres

import "service"

type Db map[int]service.Person

func (pg Db) Save(n int, p service.Person) {
	pg[n] = p
}

func (pg Db) Retrieve(n int) service.Person {
	return pg[n]
}
