package main

import "fmt"

func main() { 
	age :=19
		name :="Stephany"

	//print
	fmt.Print("HELLO, ")
		fmt.Print("WORLD! \n")
		fmt.Print("Welcom \n")
		fmt.Println("hi")

		fmt.Println("my name is", name, "and my age is", age, )
		//printf(formates strings)
fmt.Printf("my name is %v and my age is %v \n" , name, age)
fmt.Printf("my name is %q and my age is %q \n", name, age)
fmt.Printf("age is an %T \n" , age)

//sprints(save formated strings)
 var str = fmt.Sprintf("my name is %v and my age is %v \n", name, age)
 fmt.Println("the saved string is:", str)
}



