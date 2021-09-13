package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	RegNum string `json:"RegNum"`
	Colour string `json:"Colour"`
}

func main() {
	s := `[{"RegNum":"MH-12-4575","Colour":"White"},{"RegNum":"KA-11-8596","Colour":"Grey"},{"RegNum":"MH-43-6785","Colour":"Black"}]`
	bs := []byte(s)

	var cars []Car

	err := json.Unmarshal(bs, &cars)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range cars {
		fmt.Println("\nCar Number:", i+1)
		fmt.Println("Registration Number:", v.RegNum, "\nColour:", v.Colour)
	}
}
