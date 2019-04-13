package main

import (
	"fmt"
	"math/big"
)

func main() {
	var prime1, _ = new(big.Int).SetString("3141592653589793238462643383279502884197169399375105820974944592", 10)
	var prime2, _ = new(big.Int).SetString("2718281828459045235360287471352662497757247093699959574966967627", 10)

	fmt.Printf("Prime1: %v\n", prime1)
	fmt.Printf("Prime2: %v\n", prime2)
	fmt.Printf("Product: %v\n", multiply(prime1, prime2))
}

func multiply(prime1 *big.Int, prime2 *big.Int) *big.Int {
	var product = new(big.Int)
	return product.Mul(prime1, prime2)
}
