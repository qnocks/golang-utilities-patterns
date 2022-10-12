Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

2:  Оператор return без аргументов возвращает именованные возвращаемые значения.  
    Отложенные функции могут читать и присваивать именованные возвращаемые значения возвращаемой функции.
    x = 1
    x++ (defer fuct)
    x == 2 (result)
    
    
1:  defer сохраняет функцию с текущим состоянием x = 0, после выхода из функции инкрементирует его

```
