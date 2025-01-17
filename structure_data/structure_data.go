package structuredata

import (
  "fmt"
)

type Person struct {
  Name string
  Age  int
}

type SuperPerson struct{
  Person
  SuperPower string
}

func (p *Person) Greet() string {
  return fmt.Sprintf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}


func (p *Person) SayHello(){
	fmt.Println("Hello, World!")
}

func (sp *SuperPerson) HeroPower() string {
  return fmt.Sprintf("Hello, my name is %s and I am %d years old. I have the power of %s.\n", sp.Person.Name, sp.Person.Age, sp.SuperPower)
}