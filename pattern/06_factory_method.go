package pattern

import (
	"errors"
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// паттерн позволяет скрыть логику создания генерируемых экземпляров

type Clothes interface {
	SetSize(size string)
	SetPrice(price int)
	GetSize() string
	GetPrice() int
}

type cloth struct {
	size  string
	price int
}

func (c *cloth) SetSize(size string) {
	c.size = size
}

func (c *cloth) GetSize() string {
	return c.size
}

func (c *cloth) SetPrice(price int) {
	c.price = price
}

func (c *cloth) GetPrice() int {
	return c.price
}

type shirt struct {
	cloth
}

func NewLargeShirt() Clothes {
	return &shirt{
		cloth: cloth{
			size: "L", price: 100,
		},
	}
}

type jeans struct {
	cloth
}

func NewSmallJeans() Clothes {
	return &shirt{
		cloth: cloth{
			size: "S", price: 190,
		},
	}
}

func GetCloth(cloth string) (Clothes, error) {
	if cloth == "Large shirt" {
		return NewLargeShirt(), nil
	}
	if cloth == "Small jeans" {
		return NewSmallJeans(), nil
	}
	return nil, errors.New("wrong cloth")
}

func FacadePattern() {
	shirtL, err := GetCloth("Large shirt")
	if err != nil {
		log.Println(err)
	}
	jeansS, err := GetCloth("Small jeans")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(shirtL)
	fmt.Println(jeansS)
}
