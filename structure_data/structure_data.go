package structuredata

import (
  "fmt"
)

type Person struct {
  Name string
  Age  int
}

func (p *Person) Greet() string {
  return fmt.Sprintf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}


func (p *Person) SayHello(){
	fmt.Println("Hello, World!")
}