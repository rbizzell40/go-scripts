package main

import (
  "fmt"
)

func main() {
	const lightSpeed = 299792 //km/s
	distance := 5600000   // km
	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
}