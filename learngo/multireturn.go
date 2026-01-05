//return multiple values from a function
package main
import "fmt"

func arith(a int, b int) (int, int) {
	return a + b, a - b
}

func main(){
nums1, nums2 := arith(10, 5)
fmt.Println("Addition:", nums1)
fmt.Println("Subtraction:", nums2)
}