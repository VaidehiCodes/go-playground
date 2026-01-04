//append, change, norm slice
package main

import "fmt"

func main() {

	var scores = []int{100, 200, 20, 30}
	fmt.Println(scores, len(scores))

	scores = append(scores, 90)
	fmt.Println(scores, len(scores))

	scores[4] = 99
	fmt.Println(scores, len(scores))
}
