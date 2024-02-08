package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// паттерн позволяет динамически изменять поведение объекта при смене его состояния

type Hero struct {
	currentState State
	walkingState State
	jumpingState State
	x            int
}

func NewHero() *Hero {
	hero := &Hero{}
	hero.walkingState = &WalkingState{hero: hero}
	hero.jumpingState = &JumpingState{hero: hero}
	hero.setState(hero.walkingState)
	return hero
}

func (p *Hero) setState(state State) {
	p.currentState = state
}

func (p *Hero) moveForward() {
	p.currentState.moveForward()
}

func (p *Hero) moveBackward() {
	p.currentState.moveBackward()
}

func (p *Hero) jump() {
	p.currentState.jump()
}

type State interface {
	moveForward()
	moveBackward()
	jump()
}

type WalkingState struct {
	hero *Hero
}

func (s *WalkingState) moveForward() {
	s.hero.x += 1
}

func (s *WalkingState) moveBackward() {
	s.hero.x -= 1
}

func (s *WalkingState) jump() {
	s.hero.setState(s.hero.jumpingState)
}

type JumpingState struct {
	hero *Hero
}

func (s *JumpingState) moveForward() {
	s.hero.x += 2
}

func (s *JumpingState) moveBackward() {
	s.hero.x -= 2
}

func (s *JumpingState) jump() {}

func StatePattern() {
	hero := NewHero()
	fmt.Println(hero.x)
	hero.moveForward()
	fmt.Println(hero.x)
	hero.jump()
	hero.moveBackward()
	fmt.Println(hero.x)
}
