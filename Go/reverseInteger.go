package main

import (
	"log"
	"strconv"
)

func reverseInt(x int) int {
	stringInt := strconv.Itoa(x)
	reversedString := ""
	for i := len(stringInt) - 1; i >= 0; i-- {

		if stringInt[i] == '-' {
			reversedString = "-" + reversedString
			continue
		}

		reversedString += string(stringInt[i])
	}
	reversedInt, _ := strconv.Atoi(reversedString)

	if reversedInt > 2147483647 || reversedInt < -2147483648 {
		return 0
	}

	return reversedInt
}

func main() {
	res := reverseInt(123)

	log.Println(res)
}
