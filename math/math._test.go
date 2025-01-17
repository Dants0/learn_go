package math

import (
	"testing"
)


// func Factorial(a int) int{
// 	if a == 0 {
//     return 1
//   }
//   result := 1
//   for i := 1; i <= a; i++ {
//     result *= i
//   }
//   return result
// }

func TestFactorial(t *testing.T){
	result := Factorial(5)
  equals := 120
 if result!= equals {
    t.Error("Expected", equals, "got", result)
  }
}

func TestSum(t *testing.T) {
	result := Sum(5.5,5)
	equals := 10
 if result != float64(equals) {
    t.Error("Expected", equals, "got", result)
  }
}
