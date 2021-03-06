package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable

	// #3
	// Key        : Last name
	// Element    : Phone numbers

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity

	var phone map[string]string
	var productId map[int]bool
	var phones map[string][]string
	var customer map[int]map[int]int

	fmt.Printf("phones     : %#v\n", phone)
	fmt.Printf("products   : %#v\n", productId)
	fmt.Printf("multiPhones: %#v\n", phones)
	fmt.Printf("basket     : %#v\n", customer)
}
