/*
Version:1.16
author: Muves
I like pizza
created date: July/18/2020
Gimme my pizza to get copyright
all the codes are copyrighted
*/
package main

import(
    "fmt"
    "strconv"
)

func main() {
    
    //Im a gamer
    fmt.Println("I love gaming and im pro"); 
    fmt.Println("This code is covered by pizza")
    myinput()
}
func myinput()  {
    var input string
    fmt.Println("Please give me number of your credit card to continue")
    fmt.Scanln(&input)
    a , err := strconv.ParseFloat(input,32)
    
    if err != nil{
       fmt.Println(err)
    } else {
       fmt.Print(input + " + 7 = " + strconv.FormatFloat(a + 7, 'f', -1, 32))
    }
}