//where tf the function
package main
import "fmt"

func add(a int, b int)int{
	return a + b
}

func main(){
	addition := add(3,4)
	fmt.Println("addition of 3 and 4 is",addition)
}