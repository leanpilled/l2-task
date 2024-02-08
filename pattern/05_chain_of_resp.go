package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// шаблон применим, когда существуют разные способы обработки одного и того же запроса

type EducationalInsitution interface {
	Study(p *Person)
	SetNext(EducationalInsitution)
}

type Person struct {
	Name         string
	Kindergarten bool
	School       bool
	Uni          bool
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

type kindergarten struct {
	next EducationalInsitution
}

func NewKindergarten() *kindergarten {
	return &kindergarten{}
}

func (r *kindergarten) Study(p *Person) {
	if p.Kindergarten {
		fmt.Println("You've been here")
		r.next.Study(p)
		return
	}
	fmt.Println("Graduating kindergarten")
	p.Kindergarten = true
	r.next.Study(p)
}

func (r *kindergarten) SetNext(next EducationalInsitution) {
	r.next = next
}

type school struct {
	next EducationalInsitution
}

func NewSchool() *school {
	return &school{}
}

func (r *school) Study(p *Person) {
	if p.School {
		fmt.Println("You've been here")
		r.next.Study(p)
		return
	}
	fmt.Println("Graduating school")
	p.School = true
	r.next.Study(p)
}

func (r *school) SetNext(next EducationalInsitution) {
	r.next = next
}

type uni struct {
	next EducationalInsitution
}

func NewUni() *uni {
	return &uni{}
}

func (r *uni) Study(p *Person) {
	if p.Uni {
		fmt.Println("You've been here")
		r.next.Study(p)
		return
	}
	fmt.Println("Graduating uni")
	p.Uni = true
}

func (r *uni) SetNext(next EducationalInsitution) {
	r.next = next
}

func ChainOfRespPattern() {
	p := NewPerson("XXX")
	k := NewKindergarten()
	s := NewSchool()
	u := NewUni()
	k.SetNext(s)
	s.SetNext(u)
	k.Study(p)
	fmt.Println(p)

	smart := NewPerson("SmartFella")
	k.SetNext(u)
	k.Study(smart)
	fmt.Println(smart)
}
