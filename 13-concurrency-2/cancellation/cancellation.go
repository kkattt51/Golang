package main

import (
	"fmt"
	"time"
)

func worker(cancel chan struct{}, in <-chan string, out chan<- string) {
	for {
		select {
		case <-cancel:
			fmt.Println("получена команда отмены - поток завершил работу")
			close(out)
			return
		case val := <-in:
			fmt.Println("Обработка сообщения:", val)
			out <- fmt.Sprintf("Сообщение %s обработано", val)
		}
	}
}

func main() {
	data := make(chan string)
	result := make(chan string)
	cancel := make(chan struct{})
	go worker(cancel, data, result)

	go func() {
		for i := 0; i < 10; i++ {
			if i > 8 {
				close(cancel)
				break
			}
			data <- fmt.Sprintf("сообщение #%d", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	for val := range result {
		fmt.Println(val)
	}
}
