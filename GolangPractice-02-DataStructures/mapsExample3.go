package main

import "fmt"

func main() {
	employeeID := map[int]map[string]string{
		101: map[string]string{
			"First Name": "Aiden",
			"Last Name":  "Smith",
		},
		102: map[string]string{
			"First Name": "John",
			"Last Name":  "Ray",
		},
		103: map[string]string{
			"First Name": "Anthony",
			"Last Name":  "Reed",
		},
		104: map[string]string{
			"First Name": "Jacob",
			"Last Name":  "Steal",
		},
	}

	for key, value := range employeeID {
		fmt.Println(value["First Name"], value["Last Name"], "-- Employee ID:", key)
	}

	if eID, ok := employeeID[103]; ok {
		fmt.Println(eID["First Name"], eID["Last Name"])
	}

}
