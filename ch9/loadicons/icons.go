package icons

import (
	"sync"
)

func loadIcons() {
	icons = map[string]string{
		"spades":  "spades.png",
		"hearts":  "hearts.png",
		"diamond": "diamonds.png",
		"clubs":   "clubs.png",
	}
}

var loadIconsOnce sync.Once
var icons map[string]string

// Concurrency-safe.
func Icon(name string) string {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
