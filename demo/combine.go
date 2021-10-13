package demo

import "fmt"

type Bike struct {
	Name string
}

func (bike *Bike) BikeRide() string {
	return "骑" + bike.Name
}

type Rider struct {
	*Bike
	Name string
}

func (rider *Rider) Ride() {
	fmt.Println(rider.Name + rider.BikeRide())
}
