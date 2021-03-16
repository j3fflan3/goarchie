package main

import (
	"fmt"
	"service"
	"service/storage/mongo"
	"service/storage/postgres"
)

func main() {
	fmt.Println("Hi!")
	dbm := mongo.Db{}
	dbp := postgres.Db{}

	p1 := service.Person{First: "Jenny"}
	p2 := service.Person{First: "James"}
	ps := service.NewPersonService(dbm)
	service.Put(dbm, 1, p1)
	service.Put(dbm, 2, p2)
	fmt.Println(service.Get(dbm, 1))
	fmt.Println(service.Get(dbm, 2))

	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))
	service.Put(dbp, 1, p1)
	service.Put(dbp, 2, p2)
	fmt.Println(service.Get(dbp, 1))
	fmt.Println(service.Get(dbp, 2))

}
