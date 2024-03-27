package main

import "fmt"

func main() {
	website := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(website)
	fmt.Println(website["Amazon Web Services"])
	website["LinkedIn"] = "https://linkedin.com"

	delete(website,"Google")
	fmt.Println(website)
}

// struct can't delete use for tell type
// map can remove use for collection