package concorrencia

import (
	"fmt"
	"sync"
	"time"
)

func workergroup(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d começou \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d terminou \n", id)
}

func WaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
  go workergroup(i, &wg)
	}

	wg.Wait()
	fmt.Println("Todos os workers terminaram")

}
