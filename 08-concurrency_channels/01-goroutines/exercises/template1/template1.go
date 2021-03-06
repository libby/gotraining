// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/TZFvBbg3OJ

// Create a program that declares two anonymous functions. Once that counts up to
// 100 from 0 and one that counts down to 0 from 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

import "runtime"

// Add imports.

// main is the entry point for all Go programs.
func main() {
	// Declare a wait group and set the count to two.

	// Allocate one logical processor.
	runtime.GOMAXPROCS(1)

	// Declare an anonymous function and create a goroutine.
	{
		// Declare a loop that counts down from 100 to 0 and
		// display each value.

		// Decrements the count of the wait group.
	}

	// Declare an anonymous function and create a goroutine.
	{
		// Declare a loop that counts up from 0 to 100 and
		// display each value.

		// Decrements the count of the wait group.
	}

	// Wait for the goroutines to finish.
}
