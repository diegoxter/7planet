package render

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/diegoxter/7planet/internal/assets"
	"github.com/diegoxter/7planet/internal/systems/entities"
)

type Offset struct {
	X, Y int32
}

type Render struct {
	ScreenWidth, ScreenHeight int32
	Offset                    Offset
	Tileset                   *rl.Image
}

func (r *Render) renderPlayer(p *entities.Player) {
	p.Data.RenderSelf(r.Offset.X, r.Offset.Y)
}

func (r *Render) renderRoom(t *rl.Texture2D) {
	rl.DrawTexture(
		*t,
		r.ScreenWidth/2-t.Width/2,
		r.ScreenHeight/2-t.Height/2,
		rl.White,
	)
}

func (r *Render) Unload() {
	rl.UnloadImage(r.Tileset)
}

func (r *Render) Init(lW, lH int32) {
	mapPixelW := lW * assets.DrawSize
	mapPixelH := lH * assets.DrawSize

	r.Offset = Offset{
		X: (r.ScreenWidth - mapPixelW) / 2,
		Y: (r.ScreenHeight - mapPixelH) / 2,
	}
}

func (r *Render) Render(t *rl.Texture2D, p *entities.Player) {
	r.renderRoom(t)
	r.renderPlayer(p)
}
