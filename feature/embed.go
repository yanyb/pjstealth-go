package feature

import (
	"embed"
)

//go:embed *.js
var f embed.FS

func FromFile(name string) string {
	data, _ := f.ReadFile(name)
	return string(data)
}
