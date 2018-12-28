package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name, Phone string
}

func main() {

	// init mgo session
	session, err := mgo.Dial("localhost:27017")
	checkErr(err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	err = c.Insert(
		&Person{"Alex", "+880 168 60 434"},
		&Person{"Ryan", "+880 168 70 444"},
	)
	checkErr(err)

	result := Person{}
	err = c.Find(
		bson.M{"name": "Alex"},
	).One(&result)
	checkErr(err)

	fmt.Println("Phone:", result.Phone)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
