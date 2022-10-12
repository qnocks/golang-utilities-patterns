Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
<nil>
false

Псевдокод для Foo():
    err.type = *os.PathError
    err.data = nil
    return err

<nil>: печтатает данные, не тип интерфейса и не всю структуру интерфейса
false: вся структура не равна nil (err.type = *os.PathError)
```
