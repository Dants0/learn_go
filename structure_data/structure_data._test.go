package structuredata

import (
	"testing"
)

func TestPerson_Greet(t *testing.T) {
	expect := &Person{
		Name: "John Doe",
		Age: 30,
	}
	phrase := "Hello, my name is John Doe and I am 30 years old.\n"

	if expect.Greet() != phrase{
		t.Errorf("Expected '%s', got '%s'", phrase, expect.Greet())
	}
}
