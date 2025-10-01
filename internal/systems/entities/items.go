package entities

type ItemType int

const (
	Weapon = iota
	Consumable
	Gear
)

type Item struct {
	Data Entity
	Type ItemType
}
