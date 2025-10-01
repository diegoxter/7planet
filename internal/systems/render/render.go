package render

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/diegoxter/7planet/internal/systems/entities"
)

type Render struct {
	ScreenWidth, ScreenHeight int32
	Tileset                   *rl.Image
}

func renderPlayer(p *entities.Player) {
	p.Data.RenderSelf()
}

func renderMobs(ms []*entities.Mob) {
	for _, m := range ms {
		m.Data.RenderSelf()
	}
}

func renderRoom(t *rl.Texture2D) {
	rl.DrawTexture(
		*t, 0, 0,
		rl.White,
	)
}

func (r *Render) Unload() {
	rl.UnloadImage(r.Tileset)
}

func (r *Render) Init(lW, lH int32) {
}

func (r *Render) Render(t *rl.Texture2D, p *entities.Player, ms []*entities.Mob) {
	renderRoom(t)
	renderMobs(ms)
	renderPlayer(p)
}
