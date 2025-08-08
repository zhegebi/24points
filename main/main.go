package main

import (
	"24points/twentyfourPoints"
	"fmt"
)

func main() {
	var n1, n2, n3, n4 int
	fmt.Println("Enter four numbers for calculating 24 points:")
	fmt.Scan(&n1, &n2, &n3, &n4)
	if !twentyfourPoints.Calculate24(n1, n2, n3, n4) {
		fmt.Println("No solution!")
	}
}
