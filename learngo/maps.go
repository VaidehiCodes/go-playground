//key-value pair storage
package main
import (
	"fmt"
)	
func main() {

	menu := map [string]float64{
		"coffee": 2.99,
		"tea":    1.99,
		"cake":   4.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["cake"])
	// loop thru
	for k, v := range menu{
		fmt.Println(k, "-", v)
	}
	// ints as key
	phonebook := map[int]string{
		1234567890 : " Alice",
		9876543210 : " Bob",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[1234567890])

	phonebook[9876543210] = " Charlie"
	fmt.Println(phonebook)
}