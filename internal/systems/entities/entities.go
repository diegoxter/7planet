package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/diegoxter/7planet/internal/assets"
)

type Position struct {
	X, Y float32
}

type Entity struct {
	ID, HP, SpriteID int
	Position         Position
}

type Player struct {
	Data Entity
}

func (e *Entity) Move(dX, dY float32, maxX, maxY int) {
	pX := &e.Position.X
	pY := &e.Position.Y

	if *pX+dX < float32(maxX) && *pX+dX > float32(0) && *pY+dY < float32(maxY) &&
		*pY+dY >= float32(0) {
		*pX += dX
		*pY += dY
	}
}

func (e *Entity) RenderSelf() {
	pixelX := int32(e.Position.X) * assets.DrawSize
	pixelY := int32(e.Position.Y) * assets.DrawSize

	// TODO temporary solution, player sprite to be set
	rl.DrawRectangle(int32(pixelX), int32(pixelY), assets.DrawSize, assets.DrawSize, rl.Red)
}
