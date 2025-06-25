package logic

import (
	"fmt"
	g "guess-number/internal/game"
	i "guess-number/internal/input"
	msg "guess-number/internal/messages"
	s "guess-number/internal/service"
	"strings"
)

func NewGame() {
	game := g.Game{}

	fmt.Println(msg.WelcomeMsg)

	var min, max int
	for {
		min = i.GetInt(msg.SetMinMsg)
		max = i.GetInt(msg.SetMaxMsg)
		if min <= max {
			break
		}
		fmt.Println(msg.WrongMinMax)
	}

	game.SetRange(min, max)

	eAtts, mAtts, hAtts := s.CalculateAttempts(min, max)
	setDiffMsg := fmt.Sprintf(msg.SetDiffMsg, eAtts, mAtts, hAtts)
	diff := i.GetInt(setDiffMsg)
	for diff < 1 || diff > 3 {
		fmt.Printf(msg.WrongDiffMsg, diff)
		diff = i.GetInt(setDiffMsg)
	}
	game.SetAttempts(diff)

	NewRound(game)
	for PlayAnotherRound() {
		fmt.Println()
		NewRound(game)
	}
	fmt.Printf(msg.GoodbyeMsg)
}

func NewRound(game g.Game) {
	game.SetSecret()
	fmt.Printf(msg.StartRoundMsg, game.MinNumber, game.MaxNumber, game.Attempts)
	attempts := 1
	for attempts <= game.Attempts {
		try := i.GetInt(msg.NumberReqMsg)
		if try > game.Secret {
			fmt.Printf(msg.SecretLessMsg, try, game.Attempts-attempts)
			attempts++
			continue
		} else if try < game.Secret {
			fmt.Printf(msg.SecretGreaterMsg, try, game.Attempts-attempts)
			attempts++
			continue
		} else {
			fmt.Printf(msg.SecretGuessedMsg, attempts, game.Secret)
			return
		}
	}
	fmt.Printf(msg.LosingMsg, game.Secret)
}

func PlayAnotherRound() bool {
	answer := strings.ToLower(i.GetString(msg.AskForNewRoundMsg))
	for {
		switch answer {
		case "yes":
			return true
		case "no":
			return false
		default:
			answer = strings.ToLower(i.GetString(msg.ReaskForNewRoundMsg))
		}
	}
}
