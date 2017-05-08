package main

import "fmt"
import "runtime"

func main() {
	go func() {
		// infinite loop since byte is unsigned
		// syntactic sugar for uint8
		var i byte
		for i=0;i <= 255; i++ {
		}
	}()

	runtime.Gosched()

	// the gargabe collector needs the go routines to stop
	// for taking action but in this case we have infinite loops
	runtime.GC()
	fmt.Println("Done")
}