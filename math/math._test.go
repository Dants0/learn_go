package math

import (
	"testing"
)

func TestFactorial(t *testing.T){
	result := Factorial(5)
  equals := 120
 if result!= equals {
    t.Error("Expected", equals, "got", result)
  }
}

func TestSum(t *testing.T) {
	result := Sum(5,5)
	equals := 10
 if result != float64(equals) {
    t.Error("Expected", equals, "got", result)
  }
}
