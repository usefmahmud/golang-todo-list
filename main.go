package main

import "fmt"

func main() {
    var name string = "yousef"
    println("the name is ", name)

    full_name := name + " mahmoud"
    fmt.Println("the full name is %s", full_name)

    // while loop
    i := 0
    for i < 10{
        println("Current line is: ", i)
        i++
    }

    // normal for loop
    for i := 1; i <= 10; i++ {
        fmt.Printf("the new line: %d\n", i)
    }

}


