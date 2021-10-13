package main

import (
	"fmt"
	"goDemo/demo"
)

func main() {
	printDemo()
	printPerson()
	printRider()
	printArray()
	demo.PrintSlice()
}

func printDemo() {
	fmt.Println("hello world")
	demo.Demo1()
	demo.Demo2()
	fmt.Println("------------------------")
}

func printPerson() {
	person := demo.NewPerson("tony", 20, 10)
	person.Add()
	fmt.Println(person.Age)
	fmt.Println("------------------------")

	demo.Add(person)
	fmt.Println(person.Age)
	fmt.Println("------------------------")
}

func printRider() {
	rider := &demo.Rider{
		Bike: &demo.Bike{"Merida"},
		Name: "Tony",
	}
	rider.Ride()
	fmt.Println("------------------------")
}

func printArray() {
	arrayList := [4]int{100, 200, 300, 400}
	for _, value := range arrayList {
		fmt.Println(value)
	}
}
