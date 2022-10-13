package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Factory method - порождающий паттерн проектирования, который определяет общий интерфейс
	для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

	Use-cases:
		- Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код
		- Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки
		- Когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых

	Props:
		- Избавляет класс от привязки к конкретным классам продуктов
		- Выделяет код производства продуктов в одно место, упрощая поддержку кода
		- Упрощает добавление новых классов/структур в программу
		- Реализует принцип открытости/закрытости

	Cons:
		-  Может привести к созданию больших параллельных иерархий классов,
			так как для каждого класса продукта надо создать свой подкласс создателя

	Examples:
*/

/*
	Абстрактная фабрика содержет в себе фабричный метод. Отличие паттернов заключается в топ, что в фабрике
	конкретный продукты объединяются в некоторый набор по средствам конкретных фабрик, клиент работает с набором,
	таким образом эти продукты взаимосвязаны, в случае фабричного метода каждый продукт сам по себе, т.е. клиент может
	использовать каждый продукт по отдельности
*/

// Connection is the abstract type to be created
type Connection interface {
	execute(query string)
}

// connection is the concrete Connection
type connection struct {
	connStr string
}

func (c *connection) execute(query string) {
	fmt.Printf("Executing: %s for %s", query, c.connStr)
}

// PostgresConnection is the concrete Connection
type PostgresConnection struct {
	connection
}

func NewPostgresConnection() *PostgresConnection {
	return &PostgresConnection{
		connection: connection{
			connStr: "postgres/test",
		},
	}
}

// MySQLConnection is the concrete Connection
type MySQLConnection struct {
	connection
}

func NewMySQLConnection() *MySQLConnection {
	return &MySQLConnection{
		connection: connection{
			connStr: "mysql/test",
		},
	}
}

func getConnection(connType string) (Connection, error) {
	if connType == "postgres" {
		return NewPostgresConnection(), nil
	}

	if connType == "mysql" {
		return NewMySQLConnection(), nil
	}

	return nil, errors.New("wrong connection type")
}

func main() {
	postgresConn, _ := getConnection("postgres")
	mysqlConn, _ := getConnection("mysql")

	fmt.Println(postgresConn)
	fmt.Println(mysqlConn)
}
