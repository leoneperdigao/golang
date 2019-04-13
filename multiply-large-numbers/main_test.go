package main

import (
	"math/big"
	"testing"
)

var EXPECTED, _ = new(big.Int).SetString("8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184", 10)

func Test(t *testing.T) {

	var prime1, _ = new(big.Int).SetString("3141592653589793238462643383279502884197169399375105820974944592", 10)
	var prime2, _ = new(big.Int).SetString("2718281828459045235360287471352662497757247093699959574966967627", 10)
	var product = multiply(prime1, prime2)

	if product.Cmp(EXPECTED) != 0 {
		t.Errorf("The product expected for the multiplication of %v and %v was %v but is %v", prime1, prime2, EXPECTED, product)
	}
}
