package main

import (
	"encoding/json"
	"fmt"
)

// define a struct for car
type Car struct {
	RegNum string
	Colour string
}

func main() {
	// initialize struct car
	c1 := Car{
		RegNum: "MH-12-4575",
		Colour: "White",
	}

	c2 := Car{
		RegNum: "KA-11-8596",
		Colour: "Grey",
	}
	c3 := Car{
		RegNum: "MH-43-6785",
		Colour: "Black",
	}

	cars := []Car{c1, c2, c3}
	fmt.Println(cars)

	bs, err := json.Marshal(cars)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
