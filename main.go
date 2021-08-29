package main

import (
	"fmt"
	"math"
)

func main() {
	//x := 10.0
	//n := 5.0
	//lim := 2.5
	//pow(x, n, lim)

	//input := int64(10)
	//getSwitch(input)

	checkValue := int32(0)
	getIfSwitch(checkValue)
}

func getIfSwitch(checkValue int32) {
	fmt.Println(checkValue)
	switch {
	case checkValue <= 0:
		checkValue = 10
		getIfSwitch(checkValue)
	case checkValue <= 10:
		checkValue--
		if checkValue != 0 {
			getIfSwitch(checkValue)
		}
	}
}

// math.pow = Возведение в число
func pow(x, n, lim float64) float64 {
	if result := math.Pow(x, n); result < lim {
		fmt.Println(result)
	}
	fmt.Println(lim)
	return lim
}
func getSwitch(input int64) int64 {
	//var checkValue uint32
	//fmt.Scan(&checkValue)
	switch input {
	case 0:
		fmt.Println("Zero")
		return 1
	case 1:
		fmt.Println("One")
		return 1
	case 2:
		fmt.Println("Two")
		return 2
	default:
		fmt.Println("Default")
		return 0
	}
}
