package main

import (
	"fmt"
	"math/rand"
)

var place int
var drawn []int

func main() {
	var cfizz, z int
	//user sets up variables
	rand.Seed(time.Now().UnixNano())
	fmt.Println("size of the deck?")
	fmt.Scanln(&cfizz)
	fmt.Println("how many cards to draw out?")
	fmt.Scanln(&z)
	
	fibu(cfizz, z)

}

func fibu(x, z int) {
	var i int
	var deck []int
	// generate the deck
	for i = 1; i < x+1; i++ {
		deck = append(deck, i)
	}
	suffle(deck)
	for i = 1; i < z+1; i++ {
		draw(deck)
	}

}

func areyouhere(s []int, str int) bool {
	//check if the card's already been drawn
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func draw(x []int) {
	//select a random card
	place = x[rand.Intn(len(x))]
	//retry if already been drawn
	if areyouhere(drawn, place) {
		draw(x)
	} else {
		//add card to the list of cards already drawn and print it
		drawn = append(drawn, place)
		fmt.Println(place)
	}
}

func suffle(x []int) []int{
	//Fisher–Yates shuffle
	for i := range x {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}
	return x

}
