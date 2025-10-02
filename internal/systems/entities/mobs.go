package entities

import (
	"fmt"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/dngn"

	"github.com/diegoxter/7planet/internal/assets"
)

type MobMovementType int

const (
	Free MobMovementType = iota
	Pursuing
	ShapedStar
	Circular
	Stationary
)

type Mob struct {
	Data         Entity
	HP           int
	MovementType MobMovementType
	moveCounter  int
	dirX, dirY   float32
	Drops        Item
	circleAngle  float32
}

func pickMobDataPerFloor(f int) (string, int) {
	id := rand.Intn(f + (10 * f))
	if id == 0 {
		id = 1
	}

	var movementType MobMovementType

	switch {
	case id%2 == 0 && id%3 == 0:
		movementType = ShapedStar
	case id%2 == 0:
		movementType = Pursuing
	case id%3 == 0:
		movementType = Circular
	case isPrime(id):
		movementType = Stationary
	default:
		movementType = Free
	}

	return fmt.Sprintf("internal/assets/png/mobs/Icon%d.png", id), int(movementType)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GenerateMob(floor int, r *dngn.BSPRoom, l *dngn.Layout) *Mob {
	gX := r.X + 1 + rand.Intn(r.W-2)
	gY := r.Y + 1 + rand.Intn(r.H-2)
	mT := l.Get(gX, gY)
	if mT != 32 {
		return nil
	}
	mF, mMT := pickMobDataPerFloor(floor)

	t, err := assets.LoadTexture(mF)
	if err != nil {
		panic("mob cannot be loaded")
	}

	mob := &Mob{
		Data: Entity{
			Position: rl.NewVector2(float32(gX), float32(gY)),
			Sprite:   *NewSprite(t, 1, 1, 1),
		},
		MovementType: MobMovementType(mMT),
		HP:           50,
		moveCounter:  30 + rand.Intn(20),
		dirX:         1,
		dirY:         0,
	}

	switch MobMovementType(mMT) {
	case Circular:
		mob.circleAngle = rand.Float32() * 2 * math.Pi
	}

	return mob
}

func (m *Mob) Update(roomMaxX, roomMaxY int, playerPos rl.Vector2, layout *dngn.Layout) {
	switch m.MovementType {
	case Free:
		if m.moveCounter <= 0 {
			m.dirX = rand.Float32()*2 - 1
			m.dirY = rand.Float32()*2 - 1
			length := float32(math.Sqrt(float64(m.dirX*m.dirX + m.dirY*m.dirY)))
			if length > 0 {
				m.dirX /= length
				m.dirY /= length
			}
			m.moveCounter = 30 + rand.Intn(20) // 30-50 ticks
			fmt.Printf("  New direction: (%.2f,%.2f), Counter=%d\n", m.dirX, m.dirY, m.moveCounter)
		}
		m.moveCounter--
		newX := m.Data.Position.X + m.dirX*0.1
		newY := m.Data.Position.Y + m.dirY*0.1

		t := layout.Get(int(newX), int(newY))

		if newX <= float32(roomMaxX) && newY <= float32(roomMaxY) && t != 35 {
			m.Data.Move(m.dirX*0.1, m.dirY*0.1, layout)
		}

	case Pursuing:
		dx := playerPos.X - m.Data.Position.X
		dy := playerPos.Y - m.Data.Position.Y
		distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))

		if distance > 0 && distance < 10 {
			dx /= distance
			dy /= distance
			// TODO have a movement modifier here
			m.Data.Move(dx*0.01, dy*0.01, layout)
		}

	case ShapedStar:
		if m.moveCounter <= 0 {
			m.moveCounter = 20 + rand.Intn(10)
		}
		m.moveCounter--
		newX := m.Data.Position.X + m.dirX*0.1
		newY := m.Data.Position.Y + m.dirY*0.1

		t := layout.Get(int(newX), int(newY))
		t1 := layout.Get(int(newX), int(newY-1))
		t2 := layout.Get(int(newX), int(newY-1))

		if t != 35 && t1 != 120 && t2 != 120 {
			moved := m.Data.Move(m.dirX*0.1, m.dirY*0.1, layout)

			if !moved {
				m.seekPlayerOnCollision(playerPos, layout)
			}
		} else {
			m.seekPlayerOnCollision(playerPos, layout)
		}
	case Circular:
		m.circleAngle += 0.12
		radius := float32(2.0)

		circularX := float32(math.Cos(float64(m.circleAngle))) * radius
		circularY := float32(math.Sin(float64(m.circleAngle))) * radius

		dx := playerPos.X - m.Data.Position.X
		dy := playerPos.Y - m.Data.Position.Y
		distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))

		if distance > 0 {
			dx /= distance
			dy /= distance
			m.dirX = circularX*0.7 + dx*0.3
			m.dirY = circularY*0.7 + dy*0.3
		} else {
			m.dirX = circularX
			m.dirY = circularY
		}

		m.Data.Move(m.dirX*0.09, m.dirY*0.05, layout)

	case Stationary:
		m.Data.Sprite.Direction = Dir(rand.Intn(4))

	}
}

func (m *Mob) seekPlayerOnCollision(playerPos rl.Vector2, layout *dngn.Layout) {
	dx := playerPos.X - m.Data.Position.X
	dy := playerPos.Y - m.Data.Position.Y
	distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))

	if distance == 0 {
		return
	}

	dx /= distance
	dy /= distance

	alternativeDirs := m.generateAlternativeDirections(dx, dy)

	for _, dir := range alternativeDirs {
		if m.Data.Move(dir[0]*0.1, dir[1]*0.1, layout) {
			m.dirX, m.dirY = dir[0], dir[1]
			return
		}
	}

	oppositeDirs := [][2]float32{
		{-m.dirX, -m.dirY},
		{-m.dirY, m.dirX},
		{m.dirY, -m.dirX},
	}

	for _, dir := range oppositeDirs {
		if m.Data.Move(dir[0]*0.1, dir[1]*0.1, layout) {
			m.dirX, m.dirY = dir[0], dir[1]
			return
		}
	}
}

func (m *Mob) generateAlternativeDirections(targetX, targetY float32) [][2]float32 {
	return [][2]float32{
		{targetX, targetY},
		{targetY, -targetX},
		{-targetY, targetX},
		{
			targetX * 0.5,
			targetY * 0.5,
		},
		{targetX + targetY*0.3, targetY - targetX*0.3},
		{targetX - targetY*0.3, targetY + targetX*0.3},
	}
}
