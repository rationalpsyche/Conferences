package main 

import "fmt"
import "time"

/*
Go routines are multiplexed on OS threads
This program behaves differently depending on the number of cores:
with multiple cores one threads will perform `sleep` and reach the end.
With a single core you will get the infinite loop as expected.
*/
func main() {
	go func() {
		for i:=0;true;i++ {

		}
	}()
	time.Sleep(2*time.Second)
	fmt.Println("Done")	
}