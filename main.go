package main

import (
	"fmt"
)

func main() {
	result, err := toInt("431212")
	if err != nil {
		fmt.Println("deu erro moral")
	}

	fmt.Println(result)
}

func toInt(num string) (int, error) {
	byteArray := []byte(num)

	fmt.Println(byteArray)

	result, err := byteArrayToInt(byteArray)

	if err != nil {
		return 0, err
	}

	return result, nil
}

func byteArrayToInt(byteArray []byte) (int, error) {
	var result int
	for _, charDigit := range byteArray {
		var zeroInAscii byte = 48
		decimalNumberValue := charDigit - zeroInAscii
		result = result*10 + int(decimalNumberValue)
	}
	return result, nil
}
