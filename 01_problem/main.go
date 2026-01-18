package main

import (
	"errors"
	"fmt"
)

func calRemaAndMod(denominator, numarator int) (int, int, error) {
	if numarator == 0 {
		return 0, 0, errors.New("it is are try to divided by zero")
	}
	return (denominator / numarator), (denominator % numarator), nil
}
func main() {
	a, b, c := calRemaAndMod(10, 0)

	if c != nil {
		fmt.Println("cought and error", c)
		return
	}
	fmt.Println("remander is ", a, " modulus of ", b)
}
