package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/philoj/goplanes/client/internal/game"
)

var (
	playerId     = flag.Int("id", 1, "Set a unique id for each client") // FIXME: generate unique id if omitted
	screenWidth  = flag.Int("w", 600, "Screen Width in pixels")
	screenHeight = flag.Int("h", 600, "Screen height in pixels")
	debug        = flag.Bool("debug", false, "Debug enabled(default false)")
	host         = flag.String("host", "localhost:8080", "Debug enabled(default false)")
	path         = flag.String("path", "/lobby", "Debug enabled(default false)")
)

func main() {
	flag.Parse()
	slog.Info("Starting client", "id", *playerId)
	*path = fmt.Sprintf("%s/%d", *path, *playerId)
	g := game.NewGame(*playerId, *debug, *host, *path)
	configureEbiten()
	if err := ebiten.RunGame(g); err != nil {
		slog.Error("Game exited with error", "err", err)
		os.Exit(1)
	}
}

func configureEbiten() {
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowSize(*screenWidth, *screenHeight)
	ebiten.SetWindowTitle("Watch you Six")
}
