package concorrencia

import (
  "fmt"
  "time"
)

func Select(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(2*time.Second)
		c1 <- "Canal 1 pronto"
	}()

	go func(){
    time.Sleep(3*time.Second)
    c2 <- "Canal 2 pronto"
  }()

	select {
	case msg1 := <- c1:
		fmt.Println(msg1)
	case msg2 := <- c2:
		fmt.Println(msg2)
	}
}