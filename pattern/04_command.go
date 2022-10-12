package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Command - поведенческий паттерн проектирования, который превращает запросы в объекты,
	позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
	а также поддерживать отмену операций.

	Use-cases:
	- Когда вы хотите параметризовать объекты выполняемым действием.
	- Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
	- Когда вам нужна операция отмены

	Props:
	- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют
	- Позволяет реализовать простую отмену и повтор операций
	- Позволяет реализовать отложенный запуск операций
	- Позволяет собирать сложные команды из простых
	- Реализует принцип открытости/закрытости

	Cons:
	- Усложняет код программы из-за введения множества дополнительных классов

	Examples:
	- GUI button implementations
	- undo operations for transactions fallback
*/

// Command is the abstraction for all commands
type Command interface {
	execute()
}

// SwitchLightsCommand is the concrete Command
type SwitchLightsCommand struct {
	light LightProcessor
}

func (c SwitchLightsCommand) execute() {
	c.light.switchLight()
}

// CloseDoorCommand is the concrete Command
type CloseDoorCommand struct {
	door DoorProcessor
}

func (c CloseDoorCommand) execute() {
	c.door.close()
}

// SmartHouse is the invoker for a command
type SmartHouse struct {
	command Command
}

func (h *SmartHouse) setCommand(c Command) {
	h.command = c
}

func (h *SmartHouse) executeCommand() {
	h.command.execute()
}

// LightProcessor is the receiver which contains actual business logic
type LightProcessor struct {
}

func (l LightProcessor) switchLight() {
	fmt.Println("Switch the light")
}

// DoorProcessor is the receiver which contains actual business logic
type DoorProcessor struct {
}

func (d DoorProcessor) close() {
	fmt.Println("Close the door")
}

func main() {
	smartHouse := SmartHouse{}
	light := LightProcessor{}
	door := DoorProcessor{}

	smartHouse.setCommand(SwitchLightsCommand{light: light})
	smartHouse.executeCommand()

	smartHouse.setCommand(CloseDoorCommand{door: door})
	smartHouse.executeCommand()
}
