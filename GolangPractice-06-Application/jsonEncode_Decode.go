package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	const jsonData = `
    {"Name": "Alice", "Age": 25}
    {"Name": "Bob", "Age": 22}
`
	reader := strings.NewReader(jsonData)
	writer := os.Stdout

	dec := json.NewDecoder(reader)
	enc := json.NewEncoder(writer)

	for {
		// Read one JSON object and store it in a map.
		var m map[string]interface{}
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// Remove all key-value pairs with key == "Age" from the map.
		for k := range m {
			if k == "Age" {
				delete(m, k)
			}
		}

		// Write the map as a JSON object.
		if err := enc.Encode(&m); err != nil {
			log.Println(err)
		}
	}
}
