package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Wizard Printer
//
//   In this exercise, your goal is to display a few famous scientists
//   in a pretty table.
//
//   1. Create a multi-dimensional array
//   2. In each inner array, store the scientist's name, lastname and his/her
//      nickname
//   3. Print their information in a pretty table using a loop.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

func main() {
	scientist_names := [...][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	fmt.Printf("%-15s", "First Name")
	fmt.Printf("%-15s", "Last Name")
	fmt.Printf("%-15s", "Last Name")
	fmt.Println()
	fmt.Println("=============================================")

	for _, Col := range scientist_names {
		for _, elem := range Col {
			fmt.Printf("%-15s", elem)

		}
		fmt.Println()
	}
}
