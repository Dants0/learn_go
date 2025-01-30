package main

import (
	// "fmt"
	// "strings"
)

// import "learning/ping"
// import "learning/math"
// import "learning/router"
// import "learning/structure_data"
import "learning/concorrencia"

func main() {

	// person := structuredata.Person{
	// 	Name:    "John Doe",
  //   Age:      30,
	// }

	// person.Greet()

	// hero := structuredata.SuperPerson{
	// 	Person: person,
  //   SuperPower: "Super Suit",
	// }

	// fmt.Println(hero.HeroPower())
	
	// router.Server()


	// fmt.Println(math.Factorial(6))
	// ping.Pong("www.google.com")
	// fmt.Println(strings.Compare("www.google.com", "www.google.com"))

	// text := "www.google.com"

	// sameText := strings.Clone(text)

	// fmt.Println(strings.Clone(sameText))

	// for i := 0; i <= 10 ; i++ {
	// 	fmt.Println(math.Factorial(i), i)
	// }

	// numbers := [5] int{1,2,6,23,4}

	// var more_numbers [20]int

	// fmt.Println(more_numbers)

	// concorrencia.Goroutine()
	// concorrencia.Channels()
	// concorrencia.Select()
	concorrencia.WaitGroup()


}
