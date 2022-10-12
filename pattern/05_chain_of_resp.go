package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Chain of responsibility - поведенческий паттерн проектирования, который позволяет передавать запросы последовательно
	по цепочке обработчиков. Каждый последующий обработчик решает,
	может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

	Use-cases:
	- Когда программа должна обрабатывать разнообразные запросы несколькими способами,
		но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
	- Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	- Когда набор объектов, способных обработать запрос, должен задаваться динамически.

	Props:
	- Уменьшает зависимость между клиентом и обработчиками.
	- Реализует принцип единственной обязанности.
	- Реализует принцип открытости/закрытости.
	- Дополнительная гибкость при распределении обязанносте между объектами

	Cons:
	- Запрос может остаться никем не обработанным.

	Examples:

*/

// RequestFilter is the abstract handler for concrete handlers
type RequestFilter interface {
	execute(request string)
	setNext(handler RequestFilter)
}

// AuthFilter is the concrete RequestFilter
type AuthFilter struct {
	next RequestFilter
}

func (f *AuthFilter) execute(request string) {
	fmt.Println("Extracting and validating JWT for request:", request)
	f.next.execute(request)
	return
}

func (f *AuthFilter) setNext(handler RequestFilter) {
	f.next = handler
}

// ErrorFilter is the concrete RequestFilter
type ErrorFilter struct {
	next RequestFilter
}

func (f *ErrorFilter) execute(request string) {
	fmt.Println("Error handling for request:", request)
	f.next.execute(request)
	return
}

func (f *ErrorFilter) setNext(handler RequestFilter) {
	f.next = handler
}

// LoggingFilter is the concrete RequestFilter
type LoggingFilter struct {
	next RequestFilter
}

func (f *LoggingFilter) execute(request string) {
	fmt.Println("Logging request:", request)
	f.next.execute(request)
	return
}

func (f *LoggingFilter) setNext(handler RequestFilter) {
	f.next = handler
}

// ValidationFiler is the concrete RequestFilter
type ValidationFiler struct {
	next RequestFilter
}

func (f *ValidationFiler) execute(request string) {
	fmt.Println("Validating request headers, params and body:", request)
	return
}

func (f *ValidationFiler) setNext(handler RequestFilter) {
	f.next = handler
}

func main() {
	errorFilter := ErrorFilter{}
	loggingFilter := LoggingFilter{}
	authFilter := AuthFilter{}
	validationFiler := ValidationFiler{}

	errorFilter.setNext(&loggingFilter)
	loggingFilter.setNext(&authFilter)
	authFilter.setNext(&validationFiler)

	errorFilter.execute("/auth/login")

	fmt.Println()

	errorFilter.setNext(&loggingFilter)
	loggingFilter.setNext(&validationFiler)

	errorFilter.execute("/api/public")
}
