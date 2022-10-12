package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	Builder - порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово
	Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов

	Use-cases:
	- Когда вы хотите избавиться от «телескопического конструктора»
	- Когда ваш код должен создавать разные представления какого-то объекта. Например, деревянные и железобетонные дома

	Props:
	- Позволяет создавать объект пошагов
	- Позволяет использовать один и тот же код для создания различных представлений объектов
	- Изолирует код создания класса/структуры от его основной бизнес-логики.

	Cons:
	- Усложняет код программы из-за введения дополнительных классов.
	- Builder classes must be mutable

	Examples:
	- Построение xml/excel документа/таблицы (последнее: различные шаблоны форматирования таблиц)
	- Построение объекта http ответа (код, тело, хедеры)
*/

// House is the struct to build
type House struct {
	length int
	width  int
	height int
}

// Builder declare abstract steps to build House type and its subtypes
type Builder interface {
	setLength()
	setWidth()
	setHeight()
	getHouse() House
}

// FirstTypeHouseBuilder is the concrete Builder
type FirstTypeHouseBuilder struct {
	length int
	width  int
	height int
}

func (b *FirstTypeHouseBuilder) setLength() {
	b.length = 1
}

func (b *FirstTypeHouseBuilder) setWidth() {
	b.width = 1
}

func (b *FirstTypeHouseBuilder) setHeight() {
	b.height = 1
}

func (b *FirstTypeHouseBuilder) getHouse() House {
	return House{
		length: b.length,
		width:  b.width,
		height: b.height,
	}
}

// SecondTypeHouseBuilder is the concrete Builder
type SecondTypeHouseBuilder struct {
	length int
	width  int
	height int
}

func (b *SecondTypeHouseBuilder) setLength() {
	b.length = 2
}

func (b *SecondTypeHouseBuilder) setWidth() {
	b.width = 2
}

func (b *SecondTypeHouseBuilder) setHeight() {
	b.height = 2
}

func (b *SecondTypeHouseBuilder) getHouse() House {
	return House{
		length: b.length,
		width:  b.width,
		height: b.height,
	}
}

// Director defines build step order
type Director struct {
	builder Builder
}

func newDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b Builder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setLength()
	d.builder.setWidth()
	d.builder.setHeight()
	return d.builder.getHouse()
}

func main() {
	b1 := FirstTypeHouseBuilder{}
	b2 := SecondTypeHouseBuilder{}

	director := newDirector(&b1)
	house1 := director.buildHouse()
	fmt.Println(house1)

	director.setBuilder(&b2)
	house2 := director.buildHouse()
	fmt.Println(house2)
}
