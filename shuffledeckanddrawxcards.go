package main

import (
	"fmt"
	"math/rand"
)

var place int
var drawn []int

func main() {
	var cfizz, z int

	fmt.Println("size of the deck?")
	fmt.Scanln(&cfizz)
	fmt.Println("how many cards to draw out?")
	fmt.Scanln(&z)

	fibu(cfizz, z)

}

func fibu(x, z int) {
	var i int
	var deck []int
	for i = 1; i < x+1; i++ {
		deck = append(deck, i)
	}
	suffle(deck)
	for i = 1; i < z+1; i++ {
		draw(deck)
	}

}

func areyouhere(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func draw(x []int) {
	place = x[rand.Intn(len(x))]
	if areyouhere(drawn, place) {
		draw(x)
	} else {
		fmt.Println(place)
	}
}

func suffle(x []int) {
	for i := range x {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}

}
