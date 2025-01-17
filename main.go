package main

import (
	"fmt"
	// "strings"
)

// import "learning/ping"
import "learning/math"
// import "learning/router"
// import "learning/structure_data"

func main() {

	// person := structuredata.Person{
	// 	Name:    "John Doe",
  //   Age:      30,
	// }

	// person.Greet()

	// router.Server()

	// fmt.Println(math.Factorial(6))
	// ping.Pong("www.google.com")
	// fmt.Println(strings.Compare("www.google.com", "www.google.com"))

	// text := "www.google.com"

	// sameText := strings.Clone(text)

	// fmt.Println(strings.Clone(sameText))

	for i := 0; i <= 10 ; i++ {
		fmt.Println(math.Factorial(i), i)
	}
}
