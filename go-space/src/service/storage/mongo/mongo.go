package mongo

import "service"

type Db map[int]service.Person

func (m Db) Save(n int, p service.Person) {
	m[n] = p
}

func (m Db) Retrieve(n int) service.Person {
	return m[n]
}
