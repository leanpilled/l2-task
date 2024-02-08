package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// паттерн используется для создания сложных/однотипных объектов
// не подходит для создания частично инициализированного объекта

type Builder interface {
	SetEngineType()
	SetTransmissionType()
	SetCountry()
	GetCar() Car1
}

type Car1 struct {
	engine       string
	transmission string
	country      string
}

type VolvoBuilder struct {
	engine       string
	transmission string
	country      string
}

func NewVolvoBuilder() *VolvoBuilder {
	return &VolvoBuilder{}
}

func (b *VolvoBuilder) SetEngineType() {
	b.engine = "V6"
}

func (b *VolvoBuilder) SetTransmissionType() {
	b.transmission = "Manual"
}

func (b *VolvoBuilder) SetCountry() {
	b.country = "Sweden"
}

func (b *VolvoBuilder) GetCar() Car1 {
	return Car1{engine: b.engine, transmission: b.transmission, country: b.country}
}

type PeugeotBuilder struct {
	engine       string
	transmission string
	country      string
}

func NewPeugeotBuilder() *PeugeotBuilder {
	return &PeugeotBuilder{}
}

func (b *PeugeotBuilder) SetEngineType() {
	b.engine = "V3"
}

func (b *PeugeotBuilder) SetTransmissionType() {
	b.transmission = "Auto"
}

func (b *PeugeotBuilder) SetCountry() {
	b.country = "France"
}

func (b *PeugeotBuilder) GetCar() Car1 {
	return Car1{engine: b.engine, transmission: b.transmission, country: b.country}
}

type performer struct {
	builder Builder
}

func NewPerformer(b Builder) *performer {
	return &performer{
		builder: b,
	}
}

func (d *performer) SetBuilder(b Builder) {
	d.builder = b
}

func (d *performer) BuildCar() Car1 {
	d.builder.SetEngineType()
	d.builder.SetTransmissionType()
	d.builder.SetCountry()
	return d.builder.GetCar()
}

func BuilderPattern() {
	peugeot := NewPerformer(&PeugeotBuilder{}).BuildCar()
	volvo := NewPerformer(&VolvoBuilder{}).BuildCar()
	fmt.Println(peugeot.engine)
	fmt.Println(volvo.engine)
}
