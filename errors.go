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

//Basic error handling

err, variable := some function() call
if err != nil {
	some Print statement or handle the error however you want
	fmt.Println("please use a float64 value")
}
return a, nil

// error handling using a function 

func add (int a,b) (int, error) --> a and b are the arguments for the add() function and (int, error) are the return type for the add () function and the error
{
	if a+b != 100 {
		return 0, errors.New("Use a number that adds to 100")
	}
	return a+b
}


