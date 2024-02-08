package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// паттерн позволяет добавлять поведение в структуру без фактического её изменения
// можно использовать при добавлении функционала к сторонним либам

type Bike struct{}

type Car struct{}

type VehicleVisitor interface {
	VisitCar(*Car)
	VisitBike(*Bike)
}

type RepairVisitor struct{}

func (rv *RepairVisitor) VisitCar(c *Car) {
	fmt.Println("Repairing car")
}

func (rv *RepairVisitor) VisitBike(b *Bike) {
	fmt.Println("Repairing bike")
}

func (c *Car) Accept(visitor RepairVisitor) {
	visitor.VisitCar(c)
}

func (b *Bike) Accept(visitor RepairVisitor) {
	visitor.VisitBike(b)
}

func VisitorPattern() {
	car := &Car{}
	bike := &Bike{}
	repairman := RepairVisitor{}
	car.Accept(repairman)
	bike.Accept(repairman)
}
