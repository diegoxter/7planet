package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"

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
	ID, HP   int
	Sprite   Sprite
	Position rl.Vector2
}

type Player struct {
	Data Entity
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

func (e *Entity) Move(dX, dY float32, maxX, maxY int) {
	e.Sprite.Moving = true
	pX := &e.Position.X
	pY := &e.Position.Y
	newX := *pX + dX
	newY := *pY + dY

	margin := float32(-0.28)
	leftOK := newX >= 0+margin
	rightOK := newX <= float32(maxX-1)-margin
	topOK := newY >= 0+margin
	bottomOK := newY <= float32(maxY-1)-(0.48)

	if leftOK && rightOK && topOK && bottomOK {
		*pX = newX
		*pY = newY
	}
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
	position := rl.NewVector2(pixelX, pixelY)
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
	//
	// rl.DrawRectangleLines(
	// 	int32(position.X),
	// 	int32(position.Y),
	// 	int32(math.Abs(float64(frameRec.Width))),
	// 	int32(frameRec.Height),
	// 	rl.Red,
	// )
}
