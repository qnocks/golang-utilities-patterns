Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Значения 1...8 в случайном порядке, затем бесконечно 0

for-range цикл в main горутине читает из канала c до тех пор, пока он не закрыт, c никогда не будет закрыт И
т.к. есть горутины (a и b), которые пишут в канал c (эти каналы закрыты, в них уже не пишут, 
НО т.к. проверки на закрытость их нет в select'е функции merge(), то они будут выдавать zero value int), то не будет дедлока в for-range цикле
```
