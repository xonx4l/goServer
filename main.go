package main

import "fmt"
       "os"

func main() {
	fmt.Println("hello world")

	portString := os.Getenv("PORT")
	if portString = "" {
		log.fatal("PORT is not found in the environment")

	}

	fmt.println("Port:", portString)
}
