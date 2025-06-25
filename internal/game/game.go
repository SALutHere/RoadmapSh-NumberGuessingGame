package game

import (
	"guess-number/internal/service"
	"math/rand"
)

const (
	Easy int = 1 + iota
	Medium
	Hard
)

type Game struct {
	Attempts  int
	MinNumber int
	MaxNumber int
	Secret    int
}

func (g *Game) SetRange(min, max int) {
	g.MinNumber = min
	g.MaxNumber = max
}

func (g *Game) SetAttempts(diff int) {
	e, m, h := service.CalculateAttempts(g.MinNumber, g.MaxNumber)
	switch diff {
	case Medium:
		g.Attempts = m
	case Hard:
		g.Attempts = h
	default:
		g.Attempts = e
	}
}

func (g *Game) SetSecret() {
	g.Secret = rand.Intn(g.MaxNumber-g.MinNumber+1) + g.MinNumber
}
