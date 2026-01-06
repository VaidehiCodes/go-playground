package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}	
	return total
}
func main(){
	sum(1,3,5)
	sum(1,2,3)

	nums := []int{1,2,3,4,5}

	sum(nums...)

}