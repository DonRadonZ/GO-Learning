package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 10)
	// userNames := []string{}

    userNames[0] = "Erron"

	userNames = append(userNames, "Don")
	userNames = append(userNames, "Sinkaseam")

	fmt.Println(userNames)

	coursesRating := make(floatMap, 3)

	coursesRating["go"] = 4.7
	coursesRating["react"] = 4.8
	coursesRating["angular"] = 4.8

	coursesRating.output()

	// fmt.Println(coursesRating)

	for index, value := range userNames {
		// ...
		fmt.Println("Index:",index)
		fmt.Println("Value:",value)
	}

	for key, value := range coursesRating {
		fmt.Println("Key:",key)
		fmt.Println("Value:",value)
	}
}