package main

import "fmt"

// define a struct for car
type car struct {
	regNum string
	colour string
}

// define a struct for parking lot
type parkingLot struct {
	capacity int
	slots    []car
}

func main() {

	// initialize struct car
	c1 := car{
		regNum: "MH-12-4575",
		colour: "White",
	}

	c2 := car{
		regNum: "KA-11-8596",
		colour: "Grey",
	}
	c3 := car{
		regNum: "MH-43-6785",
		colour: "Black",
	}

	// initialize struct parkingLot
	p1 := parkingLot{
		capacity: 10,
		slots:    []car{c1, c2, c3},
	}

	assignParkingSlot(p1)
	capacity(p1)
	checkFreeSpace(p1)
}

// create function that iterates over all the cars, assign slot no. and print details of each car
func assignParkingSlot(p parkingLot) {
	cnt := 0
	for _, v := range p.slots {
		cnt++
		fmt.Println("Parking slot -", cnt, "-- Vehicle Details:", v)
	}
}

// create function to check capacity of parking
func capacity(p parkingLot) {
	fmt.Println("\nTotal number of parking spaces:", p.capacity)
}

// create function to check free space available in the parking
func checkFreeSpace(p parkingLot) {

	if len(p.slots) < p.capacity {
		occ := len(p.slots)
		free := p.capacity - len(p.slots)
		fmt.Println("Occupied spaces:", occ)
		fmt.Println("Free spaces:", free)
	}
}
