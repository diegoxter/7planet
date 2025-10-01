package entities

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/dngn"

	"github.com/diegoxter/7planet/internal/assets"
)

type Mob struct {
	Data  Entity
	HP    int
	Drops Item
}

func pickMobPerFloor(f int) string {
	id := rand.Intn(f + (10 * f))
	return fmt.Sprintf("internal/assets/png/mobs/Icon%d.png", id)
}

func GenerateMob(floor int, r *dngn.BSPRoom, l *dngn.Layout) *Mob {
	gX := r.X + 1 + rand.Intn(r.W-2)
	gY := r.Y + 1 + rand.Intn(r.H-2)
	mT := l.Get(gX, gY)
	if mT != 32 { // return nil if its not in a valid position
		return nil
	}
	mF := pickMobPerFloor(floor)

	t, err := assets.LoadTexture(mF)
	if err != nil {
		panic("mob cannot be loaded")
	}

	return &Mob{
		Data: Entity{
			Sprite:   *NewSprite(t, 1, 1, 1),
			Position: rl.NewVector2(float32(gX), float32(gY)),
		},
		HP: 10,
		// TODO add drops
	}
}

func (m *Mob) Move(dX, dY float32, maxX, maxY int) {
	newX := m.Data.Position.X + dX
	newY := m.Data.Position.Y + dY

	m.Data.Move(newX, newY, maxX, maxY)
}
