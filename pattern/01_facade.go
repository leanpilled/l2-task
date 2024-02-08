package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// суть паттерна скрыть сложности системы от пользователя и предоставить ему простой интерфейс
// допустим у нас есть игра, в которой можно только ходить, реализуем на этом примере паттерн фасад ниже

type Point struct {
	X, Y float64
}

type Player1Interface interface {
	getPos() Point
	walk([]Point)
}

type GameWorldInterface interface {
	getObsticles() []Point
}

type PathFinderInterface interface {
	findPath(start, end Point, obsticles []Point) ([]Point, error)
}

type gameFacade struct {
	player Player1Interface
	world  GameWorldInterface
	path   PathFinderInterface
}

func NewGameFacade(p Player1Interface, w GameWorldInterface, pf PathFinderInterface) *gameFacade {
	return &gameFacade{player: p, world: w, path: pf}
}

func (g *gameFacade) Move(destPoint Point) bool {
	currPoint := g.player.getPos()
	if currPoint == destPoint {
		return true
	}
	obst := g.world.getObsticles()
	path, err := g.path.findPath(currPoint, destPoint, obst)
	if err != nil {
		return false
	}
	g.player.walk(path)
	return true
}
