//Name: Vincent G.
//Date: 4/13/2020
//Descritpion: MPG



package main

import "fmt"
//Create a function to calculate MPG based off of gas used and miles driven.
func MPG(a, b int)(int){
//declare variable for MilesperGallon as int
var MilesperGallon int
MilesperGallon = a / b

return MilesperGallon
}


func main() {
  //declare variable for a and b as int
  var a int
  var b int
//ask user for miles driven and gallons used
 fmt.Println("Enter in the amount of miles driven")
 fmt.Scanln(&a)
 fmt.Println("Enter in how many gallons of gas was used")
 fmt.Scanln(&b)
//call MPG(a, b) and print a message telling the MPG
  fmt.Println("You got", MPG(a,b),"miles per gallon!")
}
