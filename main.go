package main

import (
	"fmt"

	locations_provider "example.com/hello/providers/location_provider"
)

func main() {

	country, err := locations_provider.GetCountry("AR")

	fmt.Println(err)
	fmt.Println(country)

}
