package messages

const (
	WelcomeMsg = "Welcome to Number Guessing Game!"

	SetMinMsg = "Please select the minimal number of the range: "

	SetMaxMsg = "Please select the maximal number of the range: "

	WrongMinMax = "Min number should be less than max number."

	SetDiffMsg = "Please select the difficulty level by its number:\n" +
		"1. Easy (%d chances)\n" +
		"2. Medium (%d chances)\n" +
		"3. Hard (%d chances)\n\n" +
		"Your choise: "

	WrongDiffMsg = "There is no difficulty level with number: %d\n\n"

	StartRoundMsg = "I'm thinking of a number between %d and %d.\n" +
		"You have %d attempts to guess it.\n\n"

	NumberReqMsg = "Enter your guess: "

	SecretLessMsg = "Incorrect! The number is less than %d. Attempts left: %d\n\n"

	SecretGreaterMsg = "Incorrect! The number is greater than %d. Attempts left: %d\n\n"

	SecretGuessedMsg = "Congratulations! You guessed the correct number in %d attempts. It was %d\n\n"

	LosingMsg = "You lose. The hidden number was: %d\n\n"

	AskForNewRoundMsg = "Do you want to play one more round?\n"

	ReaskForNewRoundMsg = "Write 'yes' or 'no'. Do you want to play one more round?\n"

	GoodbyeMsg = "Thanks for playing. Bye!\n"
)
