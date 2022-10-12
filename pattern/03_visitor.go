package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Visitor - поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции,
	не изменяя классы объектов, над которыми эти операции могут выполняться.

	Use-cases:
	- В проекте присутствуют объекты многих классов с различными интерфейсами
		и нам необходимо выполнить над ними операции, которые зависят от конкретных классов;
	- Необходимо выполнять не связанные между собой операции над объектами, которые
  		входят в состав структуры и мы не хотим добавлять эти операции в классы;
	- Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.

	Props:
	- Упрощает добавление операций, работающих со сложными структурами объектов.
	- Объединяет родственные операции в одном классе.
	- Посетитель может накапливать состояние при обходе структуры элементов.

	Cons:
	- Паттерн не оправдан, если иерархия элементов часто меняется (меняется логика визиторов)
	- Может привести к нарушению инкапсуляции элементов.

	Examples:

*/

// Document is the element is going to be visited
type Document interface {
	getData() string

	// each element have to accept visitors
	accept(v DocumentVisitor)
}

// JSONDocument is the concrete element
type JSONDocument struct {
	data string
}

func (j JSONDocument) getData() string {
	return fmt.Sprintf("{\"data\": \"%s\"}", j.data)
}

func (j JSONDocument) accept(v DocumentVisitor) {
	v.visitJSONDocument(j)
}

// XMLDocument is the concrete element
type XMLDocument struct {
	data string
}

func (x XMLDocument) accept(v DocumentVisitor) {
	v.visitXMLDocument(x)
}

func (x XMLDocument) getData() string {
	return fmt.Sprintf("<data>%s</data>", x.data)
}

// DocumentVisitor is the abstraction for visiting Document element
type DocumentVisitor interface {
	visitJSONDocument(document JSONDocument)
	visitXMLDocument(document XMLDocument)
}

// DocumentExporter is the concrete DocumentVisitor
type DocumentExporter struct {
}

func (e DocumentExporter) visitJSONDocument(d JSONDocument) {
	fmt.Println("Exporting json document: ", d.getData())
}

func (e DocumentExporter) visitXMLDocument(d XMLDocument) {
	fmt.Println("Exporting xml document: ", d.getData())
}

// DocumentCompressor is the concrete DocumentVisitor
type DocumentCompressor struct {
}

func (c DocumentCompressor) visitJSONDocument(d JSONDocument) {
	fmt.Println("Compressing json document: ", d.getData())
}

func (c DocumentCompressor) visitXMLDocument(d XMLDocument) {
	fmt.Println("Compressing xml document:", d.getData())
}

func main() {
	jsonDocument := JSONDocument{data: "my-json"}
	xmlDocument := XMLDocument{data: "my-xml"}

	exporter := DocumentExporter{}
	compressor := DocumentCompressor{}

	jsonDocument.accept(exporter)
	xmlDocument.accept(exporter)
	fmt.Println()
	jsonDocument.accept(compressor)
	xmlDocument.accept(compressor)
}
