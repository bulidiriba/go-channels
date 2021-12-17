package main

import "fmt"

func main() {
	var a chan int // empty channel created
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int) // channel defined
		fmt.Printf("Type of a is %T", a)
		fmt.Println("Channel is: ", a)
	}
}
