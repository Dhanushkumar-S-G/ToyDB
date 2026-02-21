package main

import (
	"fmt"
	"github.com/Dhanushkumar-S-G/ToyDB/toy"
	"github.com/Dhanushkumar-S-G/ToyDB/server"
)


func main() {
	fmt.Println("Welcome to the ToyDB project!!")
	toyDB := toy.New()

	server.Start("8080", toyDB)
}