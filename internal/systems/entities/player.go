package entities

import (
	"math"

	"github.com/solarlune/dngn"
)

type Player struct {
	Data Entity
	HP   int
}

func (p *Player) Move(dX, dY float32, maxX, maxY int, l *dngn.Layout) {
	newX := p.Data.Position.X + dX
	newY := p.Data.Position.Y + dY
	const marginX = 0.28
	const marginY = 0.48

	var tileX, tileY int

	if dX > 0 {
		tileX = int(math.Ceil(float64(newX - marginX)))
	} else if dX < 0 {
		tileX = int(math.Floor(float64(newX + marginX)))
	} else {
		tileX = int(math.Round(float64(newX)))
	}

	if dY > 0 {
		tileY = int(math.Ceil(float64(newY + 0.28)))
	} else if dY < 0 {
		tileY = int(math.Floor(float64(newY + marginY)))
	} else {
		tileY = int(math.Round(float64(newY)))
	}

	t := l.Get(tileX, tileY)

	if t != 120 {
		p.Data.Move(newX, newY, maxX, maxY)
	}
}
