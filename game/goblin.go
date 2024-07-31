package game

import (
	"fmt"
	"game_by_ebiten/config"
	"game_by_ebiten/definition"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	uuid "github.com/satori/go.uuid"
	"log"
	"math"
	"math/rand"
)

type GoblinManager struct {
	goblinMap map[string]*Goblin
	count     int
	game      *Game
	runImage  []*ebiten.Image
	diedImage []*ebiten.Image
}

type Goblin struct {
	id              string
	x, y            float64
	w, h            float64
	speed           float64
	runIdx, diedIdx int
	direction       int
	die             bool
}

func NewGoblinManager(game *Game) *GoblinManager {
	var runImgList = make([]*ebiten.Image, 0)
	var diedImgList = make([]*ebiten.Image, 0)
	for i := 0; i < config.Cfg.GoblinConfig.RunImg.ImgNum; i++ {
		tmpImg := fmt.Sprintf(config.Cfg.GoblinConfig.RunImg.Img, i)
		img, _, err := ebitenutil.NewImageFromFile(tmpImg)
		if err != nil {
			log.Fatal(err)
		}

		runImgList = append(runImgList, img)
	}
	for i := 0; i < config.Cfg.GoblinConfig.DiedImg.ImgNum; i++ {
		tmpImg := fmt.Sprintf(config.Cfg.GoblinConfig.DiedImg.Img, i)
		img, _, err := ebitenutil.NewImageFromFile(tmpImg)
		if err != nil {
			log.Fatal(err)
		}

		diedImgList = append(diedImgList, img)
	}

	goblinManager := &GoblinManager{
		goblinMap: make(map[string]*Goblin),
		game:      game,
		runImage:  runImgList,
		diedImage: diedImgList,
	}
	return goblinManager
}

func (goblinManager *GoblinManager) addGoblin() {
	if len(goblinManager.goblinMap) < 3 {
		goblin := NewGoblin(float64(goblinManager.runImage[0].Bounds().Dx()), float64(goblinManager.runImage[0].Bounds().Dy()), config.Cfg.GoblinConfig.Speed)
		goblinManager.goblinMap[goblin.id] = goblin
	}
}

func NewGoblin(w, h, speed float64) *Goblin {

	goblin := &Goblin{
		id:     uuid.NewV4().String(),
		w:      w,
		h:      h,
		runIdx: 0,
		speed:  speed,
	}
	// 不超过屏幕的随机位置
	goblin.x = rand.Float64() * float64(config.Cfg.ScreenWidth)
	goblin.y = rand.Float64() * float64(config.Cfg.ScreenHeight)

	return goblin
}

func (goblinManager *GoblinManager) Update() {
	for _, goblin := range goblinManager.goblinMap {
		if goblin.isHit(goblinManager.game.sword) {
			goblin.die = true
		}
		if goblin.isDied() {
			delete(goblinManager.goblinMap, goblin.id)
		}
		goblin.Update()

	}
	goblinManager.addGoblin()
}

func (goblinManager *GoblinManager) Draw(screen *ebiten.Image) {
	for _, goblin := range goblinManager.goblinMap {
		goblin.Draw(screen, goblinManager.runImage, goblinManager.diedImage)
	}
}

func (goblin *Goblin) Update() {
	if goblin.die {
		if goblin.diedIdx == config.Cfg.GoblinConfig.DiedImg.ImgNum-1 {
			return
		}
		goblin.diedIdx++
		return
	}
	goblin.run()

}
func (goblin *Goblin) Draw(screen *ebiten.Image, runImgList, diedImgList []*ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if goblin.direction != definition.Left {
		op.GeoM.Scale(-1, 1)
	}

	op.GeoM.Translate(goblin.x, goblin.y)

	if goblin.die {
		screen.DrawImage(diedImgList[goblin.diedIdx], op)
	} else {
		screen.DrawImage(runImgList[goblin.runIdx], op)
	}
}

func (goblin *Goblin) run() {
	goblin.runIdx++
	if goblin.runIdx == config.Cfg.GoblinConfig.RunImg.ImgNum-1 {
		goblin.runIdx = 0
	}
	x, y := ebiten.CursorPosition()
	if goblin.x < float64(x) {
		goblin.x += goblin.speed
		goblin.direction = definition.Left
	} else {
		goblin.x -= goblin.speed
		goblin.direction = definition.Right
	}
	if goblin.y < float64(y) {
		goblin.y += goblin.speed
	} else {
		goblin.y -= goblin.speed
	}

}

func (goblin *Goblin) isDied() bool {
	return goblin.diedIdx == config.Cfg.GoblinConfig.DiedImg.ImgNum
}

func (goblin *Goblin) isHit(sword *Sword) bool {
	var hit bool
	// 计算 sword 的四个边界坐标
	r := math.Sqrt(sword.w*sword.w+sword.h*sword.h) / 2
	angle := math.Atan2(sword.h, sword.w)
	x1 := sword.x + r*math.Cos(sword.theta+angle)
	y1 := sword.y + r*math.Sin(sword.theta+angle)
	x2 := sword.x + r*math.Cos(sword.theta-angle)
	y4 := sword.y - r*math.Sin(sword.theta-angle)

	// 计算敌人的四个边界坐标
	ex1 := goblin.x - goblin.w/2
	ey1 := goblin.y - goblin.h/2
	ex2 := goblin.x + goblin.w/2
	ey4 := goblin.y + goblin.h/2

	// 判断敌人的边界是否与 sword 的边界相交
	if x1 <= ex2 && x2 >= ex1 && y1 <= ey4 && y4 <= ey1 {
		hit = true
	}

	goblin.die = hit
	return hit
}
