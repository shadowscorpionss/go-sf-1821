package main

import (
	"fmt"
	"sync"
)

// Создадим программу, которая запускает 5 рутин,
// каждая из которых печатает свой порядковый номер 10 раз.
// Добиться такого результата за один вызов wg.Add.
func main() {
	var wg sync.WaitGroup

	//задаем константами количество рутин
	//и количество повторов печати
	const routinesN = 5
	const printTimes = 10

	//добавляем  в "группу ожидания" количество рутин
	wg.Add(routinesN)

	//создаем рутины которые печатают свой номер printTimes раз
	for routine := 0; routine < routinesN; routine++ {
		go func(routine int) {
			//после завершения работы рутины сообщаем о завершении потока
			defer wg.Done()
			//печатаем printTimes раз номер рутины
			for i := 0; i < printTimes; i++ {
				fmt.Println(routine + 1)
			}
		}(routine)
	}
	//ожидаем завершения всех созданных рутин
	wg.Wait()
}
