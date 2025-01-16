package integer

import "fmt"


// Add takes two integers and returns the sum of them
func Add(x int, y int) int {
	return x+y
}

func main(){
	fmt.Println(Add(1,2))
	// Output: 5
}