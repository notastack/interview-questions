package main

// not a lot to import
import (
	"fmt"
)

func main() {
	//all possible bills
	//has to be in order but you can add a reverse shell sorting algo beforehand to make sure
	bills := []int{1000, 500, 200, 100, 50, 10, 5, 2, 1}
	var usrplay int
	// set the user total money count
	fmt.Println("How much money you got?")
	fmt.Scanln(&usrplay)
	var i int
	i = 0
	var nbbills int
	fmt.Println("That's")
	// as long as there is money to sort
	for usrplay > 0 {
		// divide by bigest bill to find number of bills
		nbbills = (usrplay / bills[i])
		// print number of bills only is there is at least one
		if nbbills != 0 {
			fmt.Println(nbbills, " $", bills[i], " bills")
		}
		// remove excess money and start counting again
		usrplay = usrplay - (nbbills * bills[i])
		// move one up in the array
		i++

	}
}
