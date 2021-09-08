package main

import "fmt"

func main() {
	// Declare empty slice
	var strSlice []string
	var intSlice []int

	fmt.Println("1. Empty Slice of String:", strSlice)
	fmt.Println("\n2. Empty Slice of Int:", intSlice)

	// Declare Slice using Make
	var strSlice1 = make([]string, 10, 20) // when length and capacity is different
	var intSlice1 = make([]int, 10)        // when length and capacity is same

	fmt.Println("\n3. Length and Capcity of Slice:")
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice1), cap(strSlice1))
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))

	// Assign a Value to Slice using a Slice literal
	var strSlice2 = []string{"India", "Italy", "Japan", "Netherlands", "Germany"}
	var intSlice2 = []int{1, 2, 3, 4, 5}

	fmt.Println("\n4. Slice Values:")
	fmt.Println(strSlice2)
	fmt.Println(intSlice2)

	// Initialize Slice using make
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("\n5. Slice A:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

	// Add more items to Slice
	a = append(a, 30, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println("\n6. Slice A afer appending:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))

	// Access items in a Slice
	fmt.Println("\n7. Accessing specific items in a slice:")
	fmt.Println("Value at index 2:", a[2])
	fmt.Println("Value at from index 3-7:", a[3:7])

	// Change Value
	names := []string{"Janet", "Jones", "Jamie", "John", "Jacob"}
	fmt.Println("\n8. Changing values of Array:")
	fmt.Println("Array before changing value:", names)

	names[1] = "Jared"
	fmt.Println("Array after changing value:", names)

	// Access parts of data
	fmt.Println("\n9. Accessing parts of Data:")
	fmt.Println("Index 0-3", names[:3], "\nIndex 2-last", names[2:])

	// Append a slice to an existing slice
	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9, 10}
	x = append(x, y...)

	fmt.Println("\n10. Append a slice to an existing slice:", x)

	// Append a slice
	ab := make([]int, 0, 3)
	fmt.Println(ab, len(ab), cap(ab))
	ab = append(ab, 2, 2, 3, 4, 1, 2)
	fmt.Println(ab, len(ab), cap(ab))

}
