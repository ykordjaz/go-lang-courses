// Floating point number entered by user, trunctated, printed (int)

package main

import "fmt"

func main() {
	var floatNum float32
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&floatNum)
	truncNum := int(floatNum)
	fmt.Println(truncNum)
}