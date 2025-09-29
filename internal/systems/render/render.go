package render

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/diegoxter/7planet/internal/systems/entities"
)

type Render struct {
	ScreenWidth, ScreenHeight int32
	Tileset                   *rl.Image
}

func (r *Render) renderPlayer(p *entities.Player) {
	p.Data.RenderSelf()
}

func (r *Render) renderRoom(t *rl.Texture2D) {
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

func (r *Render) Render(t *rl.Texture2D, p *entities.Player) {
	r.renderRoom(t)
	r.renderPlayer(p)
}
