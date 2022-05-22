package main

import (
	"fmt"
)

func main() {

	customers := GetCustomers()

	for _, cucustomer := range customers {
		fmt.Println(cucustomer)
	}
}
