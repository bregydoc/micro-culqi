package littleid

import "github.com/matoous/go-nanoid"

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func New() string {
	id, _ := gonanoid.Generate(alphabet, 5)
	return id
}
