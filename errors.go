package main

import (
	"fmt"
)

func main() {
	e := 9
	if e != 10 {
		fmt.Errorf("This is not okay")
	}
	fmt.Println("this is okay")
}

#What's the issue?
You're calling fmt.Errorf("This is not okay") correctly to create an error, but you're not doing anything with it. 

fmt.Errorf returns an error value — it doesn't print or stop execution on its own.

#So, what happens?
The error is created but discarded, because you didn't assign it to a variable or use it in any meaningful way (like printing or returning it).

The program continues normally and prints "this is okay".

#solution
if e != 10 {
	err := fmt.Errorf("This is not okay")
	fmt.Println(err)
}

