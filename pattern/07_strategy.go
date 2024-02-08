package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// позволяет вам изменять поведение объекта во время выполнения программы без каких-либо изменений в классе этого объекта

type point struct {
	x, y int
}

type Character interface {
	SetMoveMethod(m MoveMethod)
	Move(dest point)
}

type MoveMethod interface {
	Move(dest point)
}

type character struct {
	moveMethod MoveMethod
	position   point
}

func NewCharacter() *character {
	return &character{}
}

func (c *character) SetMoveMethod(m MoveMethod) {
	c.moveMethod = m
}

func (c *character) Move(dest point) {
	fmt.Println("You moved from", c.position, "to", dest)
	c.moveMethod.Move(dest)
	c.position = dest
}

type horse struct{}

func NewHorse() horse {
	return horse{}
}

func (h horse) Move(dest point) {
	fmt.Println("by horse")
}

type plane struct{}

func NewPlane() plane {
	return plane{}
}

func (p plane) Move(dest point) {
	fmt.Println("by plane")
}

func StrategyPattern() {
	ch := NewCharacter()
	ch.SetMoveMethod(NewHorse())
	ch.Move(point{x: 1, y: 2})
	ch.SetMoveMethod(NewPlane())
	ch.Move(point{x: 100, y: 20})
}
