/**
*
*  /\_/\
* ( o.o )
*  > ^ <
*
* Main
*
* Main entry point for the Void Engine 2D
*
 */

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowTitle("Void Engine 2D")
	ebiten.SetWindowResizingMode(2)

	ebiten.SetWindowSize(640, 480)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
