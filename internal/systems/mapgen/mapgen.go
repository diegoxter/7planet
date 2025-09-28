package mapgen

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/dngn"

	"github.com/diegoxter/7planet/internal/assets"
)

type Map struct {
	Layout        *dngn.Layout
	Rooms         []*dngn.BSPRoom
	StartingRoom  *dngn.BSPRoom
	StartingPoint dngn.Position
	Texture       *rl.Texture2D
}

func CreateMap(w, h int) *Map {
	var GameMap Map

	GameMap.Layout = dngn.NewLayout(w, h)

	GameMap.Rooms = GameMap.Layout.GenerateBSP(
		dngn.BSPOptions{
			WallValue:       'x',
			SplitCount:      5,
			DoorValue:       '#',
			MinimumRoomSize: 6,
		},
	)

	GameMap.StartingRoom = GameMap.Rooms[rand.Intn(len(GameMap.Rooms))]
	GameMap.StartingPoint = GameMap.StartingRoom.Center()

	fmt.Println(GameMap.Layout.DataToString())

	return &GameMap
}

func drawTile(
	dst *rl.Image,
	tileset rl.Image,
	tileX, tileY, destX, destY, tileSize int32,
) {
	scale := float32(assets.DrawSize) / float32(tileSize)

	for y := range assets.DrawSize {
		for x := range assets.DrawSize {
			srcX := tileX + int32(float32(x)/scale)
			srcY := tileY + int32(float32(y)/scale)
			col := rl.GetImageColor(tileset, (srcX), (srcY))
			rl.ImageDrawRectangle(dst, destX+x, destY+y, 1, 1, col)
		}
	}
}

func LayoutToTexture2D(
	m *dngn.Layout,
	tileset *rl.Image,
) *rl.Texture2D {
	tileCoords := assets.TileCoords
	tileSize := assets.TileSize

	rows := m.Height
	cols := m.Width
	imgWidth := int32(cols) * assets.DrawSize
	imgHeight := int32(rows) * assets.DrawSize

	img := rl.GenImageColor(int(imgWidth), int(imgHeight), rl.Black)

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			c := int32(m.Data[y][x])

			coords, _ := tileCoords[c]
			// if !ok {
			// 	coords = tileCoords[32] // default tile
			// }
			drawTile(
				img,
				*tileset,
				int32(coords.X),
				int32(coords.Y),
				int32(x)*assets.DrawSize,
				int32(y)*assets.DrawSize,
				tileSize,
			)
		}
	}

	tex := rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)

	return &tex
}
