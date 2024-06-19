package main
import (
	"fmt"
	"myapp/packageone"
)
var one = "One"

func main() {
	var sometingElse = "this is a block leve variable"
	fmt.Println(sometingElse)
	myFunc()
	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)
	packageone.Exported()
}

func myFunc() {
	
	fmt.Println(one)
}