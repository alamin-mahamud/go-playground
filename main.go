package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testing"), 10)
	fmt.Println(string(hashedPassword))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
