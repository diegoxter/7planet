package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/dngn"

	"github.com/diegoxter/7planet/internal/assets"
	"github.com/diegoxter/7planet/internal/systems/entities"
	"github.com/diegoxter/7planet/internal/systems/mapgen"
	"github.com/diegoxter/7planet/internal/systems/render"
)

type Game struct {
	Renderer *render.Render
	W, H     int32
	Map      *mapgen.Map
	Camera   *rl.Camera2D
	Player   *entities.Player
}

func Init(w, h int32) *Game {
	m := mapgen.CreateMap(40, 60)
	sX, sY := float32(m.StartingPoint.X), float32(m.StartingPoint.Y)

	tiles, err := assets.Tileset()
	if err != nil {
		return nil
	}

	r := &render.Render{
		ScreenWidth:  w,
		ScreenHeight: h,
		Tileset:      tiles,
	}
	r.Init(int32(m.Layout.Width), int32(m.Layout.Height))

	m.Texture = mapgen.LayoutToTexture2D(m.Layout, r.Tileset)

	game := &Game{
		W: w,
		H: h,
		Player: &entities.Player{
			Data: entities.Entity{
				Position: entities.Position{X: sX, Y: sY},
				HP:       100,
			},
		},
		Renderer: r,
		Map:      m,
		Camera:   &rl.Camera2D{},
	}

	game.Camera.Target = rl.NewVector2(
		float32(game.Player.Data.Position.X+20),
		float32(game.Player.Data.Position.Y+10),
	)
	game.Camera.Offset = rl.NewVector2(float32(w/4), float32(h/4))
	game.Camera.Rotation = 0.0
	game.Camera.Zoom = 1.0

	return game
}

func (g *Game) Unload() {
	g.Renderer.Unload()
}

func (g *Game) handleInput() {
	if rl.IsKeyDown(rl.KeyRight) {
		g.Player.Data.Move(0.2, 0, g.Map.Layout.Width, g.Map.Layout.Height)
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		g.Player.Data.Move(-0.2, 0, g.Map.Layout.Width, g.Map.Layout.Height)
	}
	if rl.IsKeyDown(rl.KeyDown) {
		g.Player.Data.Move(0, 0.2, g.Map.Layout.Width, g.Map.Layout.Height)
	}

	if rl.IsKeyDown(rl.KeyUp) {
		g.Player.Data.Move(0, -0.2, g.Map.Layout.Width, g.Map.Layout.Height)
	}
}

func (g *Game) render() {
	g.Renderer.Render(g.Map.Texture, g.Player)
}

func (g *Game) currentRoom() *dngn.BSPRoom {
	for _, r := range g.Map.Rooms {
		if g.Player.Data.Position.X >= float32(r.X) &&
			g.Player.Data.Position.X < float32(r.X+r.W) &&
			g.Player.Data.Position.Y >= float32(r.Y) &&
			g.Player.Data.Position.Y < float32(r.Y+r.H) {
			return r
		}
	}

	return nil
}

func (g *Game) updateCameraForRoom() {
	room := g.currentRoom()
	if room == nil {
		return
	}

	drawSize := float32(assets.DrawSize)
	worldCenterX := (float32(room.X) + float32(room.W)/2) * drawSize
	worldCenterY := (float32(room.Y) + float32(room.H)/2) * drawSize

	lerp := func(a, b, t float32) float32 { return a + (b-a)*t }
	g.Camera.Target.X = lerp(g.Camera.Target.X, worldCenterX, 0.1)
	g.Camera.Target.Y = lerp(g.Camera.Target.Y, worldCenterY, 0.1)

	targetCoverage := float32(0.8)

	zoomX := (float32(g.W) * targetCoverage) / (float32(room.W) * drawSize)
	zoomY := (float32(g.H) * targetCoverage) / (float32(room.H) * drawSize)

	zoom := min(zoomX, zoomY)
	zoom = min(max(zoom, 0.3), 3.0)

	g.Camera.Zoom = lerp(g.Camera.Zoom, zoom, 0.1)
	g.Camera.Offset = rl.NewVector2(float32(g.W)/2, float32(g.H)/2)
}

func (g *Game) Run() {
	g.handleInput()
	g.updateCameraForRoom()
	g.render()
}
