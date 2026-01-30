// TODO
// Implement fish movement
// fish spawner randomly
// fish catch mechanic: tap left/right (A/D) to counteract fish pull, do it -
// quick to avoid it getting to the threshold, which causes to fail.

package main

import (
	"log"
	"strconv"

	"go-fishing/collision"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

const (
	gameScreenSize int = 512
	gameWindowSize int = 720
)

var (
	bg           *ebiten.Image
	muskie 		 *ebiten.Image
	rollSpeed    int = 97
	bgTranslateY int = 0

	bgLength int = 5120

	Ylayer int = 0

	hookMaxLength int = -bgLength + gameScreenSize
	hookCurrentLength int = -gameScreenSize + 100
	hookSpeed int = 5

	muskiePositionX int = 200
	muskiePositionY int = 400
)

// Gameloop
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyS){
		if bgTranslateY > (hookCurrentLength) {
		bgTranslateY -= hookSpeed
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyW){
		if bgTranslateY != (0) {
		bgTranslateY += hookSpeed
		}
	}

	

	return nil
}

func init() {
	var err error
	bg, _, err = ebitenutil.NewImageFromFile("assets/bg.png")
	if err != nil {
		log.Fatal(err)
	}
	
	muskie, _, err = ebitenutil.NewImageFromFile("assets/fishes/muskie.png")
	if err != nil {
		log.Fatal(err)
	}
}

// Render to screen
func (g *Game) Draw(screen *ebiten.Image) {
	bgop := &ebiten.DrawImageOptions{}
	muskieop := &ebiten.DrawImageOptions{}
	bgop.GeoM.Translate(0, float64(bgTranslateY))
	muskieop.GeoM.Translate(float64(muskiePositionX), float64(bgTranslateY + muskiePositionY))
	screen.DrawImage(bg, bgop)
	screen.DrawImage(muskie, muskieop)


	ebitenutil.DebugPrint(screen, strconv.Itoa((bgTranslateY * -1) / 8))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return gameScreenSize, gameScreenSize
}

func main() {
	ebiten.SetWindowSize(gameWindowSize, gameWindowSize)
	ebiten.SetWindowTitle("Go Fishing")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	collision.Test()
}