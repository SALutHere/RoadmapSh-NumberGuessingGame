package game

import (
	"fmt"
	"guess-number/internal/input"
	"math/rand"
)

func Start() {
	var difficulty int

	printWelcome()
	setDifficulty(&difficulty)
	startRound(difficulty)
}

func printWelcome() {
	message := "Welcome to Number Guessing Game!\n" +
		"I'm thinking of a number between 1 and 100.\n"

	fmt.Println(message)
}

func setDifficulty(difficulty *int) {
	message := "Please select the difficulty level by its number:\n" +
		"1. Easy (10 chances)\n" +
		"2. Medium (5 chances)\n" +
		"3. Hard (3 chances)\n\n" +
		"Your choise: "

	var value int
	for value = input.GetInt(message); value < 1 || value > 3; {
		fmt.Printf("There is no difficulty level with number: %d\n\n", value)
		value = input.GetInt(message)
	}

	*difficulty = value
}

func startRound(difficulty int) {
	hiddenNumber := rand.Intn(100) + 1

	var maxAttempts int
	switch difficulty {
	case 1:
		maxAttempts = 10
	case 2:
		maxAttempts = 5
	case 3:
		maxAttempts = 3
	}

	attempts := 1
	for attempts < maxAttempts {
		assumption := input.GetInt("Enter your guess: ")
		if assumption > hiddenNumber {
			fmt.Printf("Incorrect! The number is less than %d. Attempts left: %d\n\n", assumption, maxAttempts-attempts)
			attempts++
			continue
		} else if assumption < hiddenNumber {
			fmt.Printf("Incorrect! The number is greater than %d. Attempts left: %d\n\n", assumption, maxAttempts-attempts)
			attempts++
			continue
		} else {
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts. It was %d\n", attempts, hiddenNumber)
			return
		}
	}

	fmt.Printf("You lose. The hidden number was: %d", hiddenNumber)
}
