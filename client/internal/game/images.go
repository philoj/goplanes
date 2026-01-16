package game

import "github.com/hajimehoshi/ebiten/v2"

type imageInfo struct {
	path  string
	image *ebiten.Image
}

var (
	images = map[string]*imageInfo{
		bgImageAssetId: {
			path: "/bg.png",
		},
		iconImageAssetId: {
			path: "/icon.png",
		},
		blipImageAssetId: {
			path: "/blip.png",
		},
	}
)
