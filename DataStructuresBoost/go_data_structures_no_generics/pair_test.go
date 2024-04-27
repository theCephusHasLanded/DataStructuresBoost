package go_data_structures_no_generics

import (
    "testing"
	"fmt"


)

func TestPair(t *testing.T) {
	// Create an instance of Pair from the go_data_structures_no_generics package
	pair := Pair{
		Key:      "exampleKey",
		Priority: "high",
		Value:    "This is the value.",
	}

	// Print out the Pair
	fmt.Println("Pair contents:", pair)
}
