package main

import (
	"fmt"
)

func main() {
	var names []string
	var values []int
	var lower, higher int
	//set up is made by the user himself
	names = getnames()
	values = getvalues(names)
	fmt.Println("Where should the counting start?")
	fmt.Scanln(&lower)
	fmt.Println("Where should the counting stop?")
	fmt.Scanln(&higher)
	//launch the fizzbuzz function
	fibu(names, values, lower, higher)

}

func fibu(names []string, values []int, lower int, higher int) {
	//loop from lower end to higher end
	for i := lower; i < higher+1; i++ {
		findinarray(i, values, names)
	}

}

func findinarray(arrow int, target []int, name []string) {
	var y int
	//the sentence array contains either the number or the fizz, buzz or other
	var sentence []interface{}
	y = len(name)
	//looks through the value array to see if any number in it is divisible by i
	for i := 0; i < y; i++ {
		if arrow%target[i] == 0 {
			//sends the equivalent word to the sentence array
			sentence = append(sentence, name[i])
		}
	}
	//if sentence array is empty, prints the number instead
	if len(sentence) == 0 {
		sentence = append(sentence, arrow)
	}
	fmt.Println(sentence)

}

func getnames() []string {
	var namesgot []string
	var i int
	i = 1
	var namewritten string
	fmt.Println("writre [stop] when you entered all the names.")
	//sets up the names of the variables (fizz, buzz, mazz, ...)
	for true {
		fmt.Println("what should be the name of value number ", i, "?")
		fmt.Scanln(&namewritten)
		i++
		if namewritten == "stop" {
			return namesgot
		} else {
			namesgot = append(namesgot, namewritten)
		}
	}
	return namesgot
}

func getvalues(x []string) []int {
	var i int
	var y []int
	var valuegot int
	i = 0
	//fizz = 3, buzz = 5, mazz = 4, ...
	for i < len(x) {
		fmt.Println("what should be the value of ", x[i], " to be?")
		fmt.Scanln(&valuegot)
		y = append(y, valuegot)
		i++
	}
	return y

}
