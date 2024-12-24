package main

import (
	"fmt"
	"os"
)

var (
	// Initialize the variables
	num1 int = 0
	num2 int = 0
	num3 int = 0

	multAnswer int = 0
	addAnswer  int = 0

	validCombos     [][]int = [][]int{} // Valid combos of the three numbers
	validAnswers    []int   = []int{}   // Valid answers of the three numbers
	strippedAnswers []int   = []int{}   // Stripped answers of the three numbers
)

func main() {
	println("Hello, World!")

	for {
		matrix()
	}
}

// Check if slice contains value
// The first argument is the slice that you want to check and the second argument is the value to check it against
func Contains(slice []int, number int) bool {
	for _, value := range slice {
		if value == number {
			return true
		}
	}
	return false
}

func matrix() {

	// Pause
	//time.Sleep(100000000)

	// Add & Multiple
	multAnswer = num1 * num2 * num3
	addAnswer = num1 + num2 + num3

	// If the answers are equal
	if multAnswer == addAnswer {
		fmt.Printf("True, the numbers %d, %d, %d are valid\n", num1, num2, num3)

		// Add the valid numbers to the array
		validCombos = append(validCombos, []int{num1, num2, num3})
	}

	// If the answers are not equal
	if multAnswer != addAnswer {
		fmt.Printf("False, the numbers %d, %d, %d are invalid\n", num1, num2, num3)
	}

	// Increment the first counter
	num1 = num1 + 1

	// When the first counter reaches 10, reset it and increment the second counter
	if num1 == 10 {
		num1 = 0
		num2 = num2 + 1
	}

	// When the second counter reaches 10, reset it and increment the third counter
	if num2 == 10 {
		num2 = 0
		num3 = num3 + 1
	}

	// Once the third counter reaches 10, do the final calculations and exit the program
	if num3 == 10 {

		// Loop through the valid combos and add them together
		for i := 0; i < len(validCombos); i++ {
			validAnswers = append(validAnswers, validCombos[i][0]+validCombos[i][1]+validCombos[i][2])
		}

		// Loop through the valid answers and for each new occurance add it to a slice
		for i := 0; i < len(validAnswers); i++ {

			// If the answer is not in the slice, add it
			if !Contains(strippedAnswers, validAnswers[i]) {
				strippedAnswers = append(strippedAnswers, validAnswers[i])
			}

			// // Looping through the stripped answers to check for occurances
			// for ii := 0; ii < len(strippedAnswers); ii++ {

			// 	// If the answer is already in the stripped answers, break the loop
			// 	if validAnswers[i] == strippedAnswers[ii] {
			// 		break
			// 	}

			// 	// If the answer is not in the stripped answers, add it to the stripped answers
			// 	if validAnswers[i] != strippedAnswers[ii] {
			// 		strippedAnswers = append(strippedAnswers, validAnswers[i])
			// 	}
			// }

		}

		// Print the answers
		fmt.Printf("The valid combos are: %d\n", validCombos)
		fmt.Printf("The valid answers are: %d\n", validAnswers)
		fmt.Printf("The stripped answers are: %d\n", strippedAnswers)

		// Exit
		fmt.Println("Done!")
		os.Exit(0)
	}

}
