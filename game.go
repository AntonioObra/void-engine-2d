package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	gopherImage *ebiten.Image
}

func NewGame(gopherImage *ebiten.Image) *Game {
	return &Game{gopherImage: gopherImage}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gopherWidth := float64(g.gopherImage.Bounds().Dx())
	gopherHeight := float64(g.gopherImage.Bounds().Dy())

	x := (float64(Settings.ScreenWidth) - gopherWidth) / 2
	y := (float64(Settings.ScreenHeight) - gopherHeight) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)
	screen.DrawImage(g.gopherImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return Settings.ScreenWidth, Settings.ScreenHeight
}
