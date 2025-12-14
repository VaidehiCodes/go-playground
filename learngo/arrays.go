package main
 import "fmt"

 func main(){
	var a [5]int
	fmt.Println("array is:", a);
	//preloaded with 0 value
	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get: ", a[4])
//displays all values, including default 0 and assigned 100
	fmt.Println("length:", len(a))
// shows size of array, cool.

	 b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

	 b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
	
	  b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)

	//create a 2d array, sick
	var twoD [2][3]int
	for i:= range 

}
