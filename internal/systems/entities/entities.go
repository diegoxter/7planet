package entities

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/dngn"

	"github.com/diegoxter/7planet/internal/assets"
)

type Dir int

const (
	None Dir = iota
	North
	South
	East
	West
)

func (d Dir) toInt() int {
	switch d {
	case None:
		return 0
	case North:
		return 0
	case South:
		return 2
	case East:
		return 1
	case West:
		return 1
	default:
		return 0
	}
}

type Sprite struct {
	Direction                                                Dir
	Texture                                                  *rl.Texture2D
	FrameRec                                                 rl.Rectangle
	CurrentFrame                                             float32
	Moving                                                   bool
	SpritesPerRow, SpritesPerCol, FramesCounter, FramesSpeed int
}

type Entity struct {
	ID       int
	Sprite   Sprite
	Position rl.Vector2
}

func NewSprite(p *rl.Texture2D, sC, sR float32, fS int) *Sprite {
	return &Sprite{
		Texture:       p,
		CurrentFrame:  float32(0),
		FramesCounter: 0,
		FramesSpeed:   fS,
		SpritesPerRow: int(sR),
		SpritesPerCol: int(sC),
		FrameRec: rl.NewRectangle(
			0, 0,
			float32(p.Width)/sR,
			float32(p.Height)/sC,
		),
	}
}

// entity.go
func (e *Entity) Move(dX, dY float32, layout *dngn.Layout) bool {
	e.Sprite.Moving = true
	newX := e.Position.X + dX
	newY := e.Position.Y + dY

	const marginX = 0.08
	const marginY = 0.48

	maxX := layout.Width
	maxY := layout.Height
	leftOK := newX >= 0+(-0.28)
	rightOK := newX <= float32(maxX-1)-(-0.28)
	topOK := newY >= 0+(-0.28)
	bottomOK := newY <= float32(maxY-1)-(0.48)

	var tileX, tileY int

	if dX > 0 {
		tileX = int(math.Ceil(float64(newX - marginX)))
	} else if dX < 0 {
		tileX = int(math.Floor(float64(newX + marginX)))
	} else {
		tileX = int(math.Round(float64(newX)))
	}

	if dY > 0 {
		tileY = int(math.Ceil(float64(newY + 0.18)))
	} else if dY < 0 {
		tileY = int(math.Floor(float64(newY + marginY)))
	} else {
		tileY = int(math.Round(float64(newY)))
	}

	tile := layout.Get(tileX, tileY)

	if tile == 120 { // 120 = 'x' (pared)
		return false
	}

	if leftOK && rightOK && topOK && bottomOK {
		e.Position.X = newX
		e.Position.Y = newY
		return true
	}

	return false
}

func (e *Entity) updateFrameCounter() {
	e.Sprite.FramesCounter++
	e.Sprite.FrameRec.Y = float32(
		e.Sprite.Direction.toInt(),
	) * float32(
		e.Sprite.Texture.Height,
	) / float32(
		e.Sprite.SpritesPerCol,
	)

	if e.Sprite.FramesCounter >= (60 / e.Sprite.FramesSpeed) {

		e.Sprite.FramesCounter = 0
		e.Sprite.CurrentFrame++

		if e.Sprite.CurrentFrame > float32(e.Sprite.SpritesPerCol) {
			e.Sprite.CurrentFrame = 0
		}

		isMoving := 0
		if e.Sprite.Moving {
			isMoving = 1
		}

		e.Sprite.FrameRec.X = e.Sprite.CurrentFrame * float32(isMoving) * float32(
			e.Sprite.Texture.Width,
		) / float32(e.Sprite.SpritesPerRow)

	}
}

func (e *Entity) RenderSelf() {
	if e.Sprite.Texture == nil {
		return
	}

	e.updateFrameCounter()

	pixelX := e.Position.X * float32(assets.DrawSize)
	pixelY := e.Position.Y * float32(assets.DrawSize)

	// offset para ajustar el sprite respecto al tile
	offsetX := float32(0)  // si el sprite está centrado, no tocamos X
	offsetY := float32(-8) // este valor depende del tamaño del frame

	position := rl.NewVector2(pixelX+offsetX, pixelY+offsetY)

	frameRec := e.Sprite.FrameRec
	if e.Sprite.Direction == West {
		frameRec.Width = -frameRec.Width
	}

	rl.DrawTextureRec(
		*e.Sprite.Texture,
		frameRec,
		position,
		rl.White,
	)

	// Debug: tile box
	rl.DrawRectangleLines(
		int32(pixelX),
		int32(pixelY),
		int32(assets.DrawSize),
		int32(assets.DrawSize),
		rl.Green,
	)

	// Debug: centro del tile
	centerX := int32(pixelX) + int32(assets.DrawSize/2)
	centerY := int32(pixelY) + int32(assets.DrawSize/2)
	rl.DrawCircle(centerX, centerY, 3, rl.Red)
}
