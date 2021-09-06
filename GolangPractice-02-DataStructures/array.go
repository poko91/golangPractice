package main

import "fmt"

func main() {
	// Declare an Array Variable with a Specific Size
	var names [5]string
	var ids [5]int

	//Assign a Value to a Specific Element in an Array
	names[0] = "Ryan"
	names[1] = "Jason"
	names[2] = "John"
	names[3] = "Anthony"
	names[4] = "Luis"

	ids[0] = 101
	ids[1] = 102
	ids[2] = 103
	ids[3] = 104
	ids[4] = 105

	fmt.Println("1. String Array : ")
	fmt.Println("Names of employees:", names)

	fmt.Println("\n2. Int Array : ")
	fmt.Println("Employee IDs:", ids)

	//Initialize and Assign values to Array at the Same time
	ar := [5]int{1, 2, 3, 4, 5}
	fmt.Println("\n3. Declare and Initialize Array at the same time : ")
	fmt.Println("Array of int:", ar)

	//print length of array
	fmt.Println("\n4. Length of array:", len(ar))

	//default value of an array
	ar1 := [3]int{}    //default value of int is 0
	ar2 := [2]string{} //default value is an empty array of strings
	ar3 := [2]bool{}   //default value of bool is false

	fmt.Println("\n5. Default Values of Array:")
	fmt.Println("Int:", ar1)
	fmt.Println("String:", ar2)
	fmt.Println("Bool", ar3)

	// range through array elements
	fmt.Println("\n6. Loop through Array using For and Range : ")
	for index, value := range names {
		fmt.Println(index, " = ", value)
	}

	// Multi dimensional array
	fmt.Println("\n7. Multi dimensional array : ")

	count := 1
	var multi [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			multi[i][j] = count
			count++
		}
	}
	fmt.Println("Array 5 x 6 : ", multi)

}
