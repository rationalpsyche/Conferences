package main

import "fmt"
import "time"

/*
go routines will compute everything they can while waiting for sys calls
which are known to be slow. Hence the output will be i=9 for every Println
*/
func main() {
	for i:=0; i <9;i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	// necessary otherwise the go routines finish before 
	// the syscall to stdout is made
	time.Sleep(time.Second)	
}