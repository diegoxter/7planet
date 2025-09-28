package assets

import (
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	DrawSize = int32(26)

	TileSize = int32(32)

	TileCoords = map[int32]rl.Vector2{
		' ': {X: float32(TileSize * 2), Y: float32(TileSize * 2)}, // espacio ' ' → suelo
		'#': {X: float32(TileSize * 2), Y: float32(TileSize)},     // 'x' → pared
		'x': {X: float32(TileSize), Y: 0},                         // '#' → puerta
	}
)

func Tileset() (*rl.Image, error) {
	t, err := LoadImage("internal/assets/png/tileset_complete.png")
	if err != nil {
		return nil, err
	}

	return t, nil
}

func LoadImage(path string) (*rl.Image, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	image := rl.LoadImage(absPath) // Loaded in CPU memory (RAM)

	return image, nil
}

func LoadTexture(i *rl.Image) rl.Texture2D {
	texture := rl.LoadTextureFromImage(
		i,
	) // Image converted to texture, GPU memory (VRAM)
	rl.UnloadImage(i)

	return texture
}
