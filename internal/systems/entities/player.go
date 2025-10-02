package entities

import (
	"github.com/solarlune/dngn"
)

type Player struct {
	Data Entity
	HP   int
}

func (p *Player) Move(dX, dY float32, l *dngn.Layout) {
	// TODO add movement modifiers
	p.Data.Move(dX, dY, l)
}
