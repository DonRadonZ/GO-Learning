package main

import "fmt"

type Product struct {
	id string
	title string
	price float64
} 

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.

	hobby := [3]string{"watch anime", "play game", "write code"}
	fmt.Println(hobby)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobby[0])
	fmt.Println(hobby[1:3])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)

	mainHobbies := hobby[:2]

	fmt.Println(mainHobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.

	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Understan Go","API"}
	fmt.Print(courseGoals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Make User Login"
	courseGoals = append(courseGoals, "Learn to make Authentication")
	fmt.Print(courseGoals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{
		{"first","milk", 25.00},
		{"second","tea", 45.00},
		
	}

	newProduct := Product{
		"third","youghurt", 85.00,
	}
	products = append(products,newProduct)

	fmt.Println(products)

}
