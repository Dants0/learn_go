package concorrencia

import "fmt"

func worker(c chan string, w chan string){
	c <- "Worker done"
	w <- "Second message received"
}

func Channels(){
	ch := make(chan string)
	wh := make(chan string)

	go worker(ch, wh)

	msg := <- ch
	msg1 := <- wh

	fmt.Println(msg, msg1)
	// fmt.Println(ch1)
}