package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Strategy - поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый
		из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

	Use-cases:
	- Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
	- Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
	- Когда вы не хотите обнажать детали реализации алгоритмов для других классов. (Стратегия позволяет изолировать код,
		данные и зависимости алгоритмов от других объектов, скрыв эти детали внутри классов-стратегий)
	- Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора.
		Каждая ветка такого оператора представляет собой вариацию алгоритма.

	Props:
	- Горячая замена алгоритмов на лету.
	- Изолирует код и данные алгоритмов от остальных классов.
	- Уход от наследования к делегированию.
	- Реализует принцип открытости/закрытости.

	Cons:
	- Усложняет программу за счёт дополнительных классов.
	- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

	Examples:

*/

// Cache is the Context which choose a strategy and do some work
type Cache struct {
	evictionAlgo EvictionAlgo
}

func (c *Cache) setEvictionAlgo(algo EvictionAlgo) {
	c.evictionAlgo = algo
}

func (c *Cache) flush() {
	c.evictionAlgo.evict()
}

// EvictionAlgo is the interface for all strategies
type EvictionAlgo interface {
	evict()
}

// SecondChance is the concrete strategy (EvictionAlgo)
type SecondChance struct {
}

func (s *SecondChance) evict() {
	fmt.Println("Evicting using second chance")
}

// LRU is the concrete strategy (EvictionAlgo)
type LRU struct {
}

func (L *LRU) evict() {
	fmt.Println("Evicting using second LRU")
}

func main() {
	secondChance := SecondChance{}
	lru := LRU{}

	cache := Cache{}

	cache.setEvictionAlgo(&secondChance)
	cache.flush()

	cache.setEvictionAlgo(&lru)
	cache.flush()
}
