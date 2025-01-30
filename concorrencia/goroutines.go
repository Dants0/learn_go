package concorrencia

import (
  "fmt"
  "time"
)

func hello(){
	fmt.Println("Hello, goroutine!")
}

func Goroutine(){
	go hello()
	fmt.Println("Function main running")

	time.Sleep(time.Second)
}