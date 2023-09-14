package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

type card struct {
	color int
	figure int
	isPlayed bool
}

type player struct {
	id int
	hand []card
}

type team struct {
	id int
	points int
}

func createCardStack()([36]card) {
	var cardStack [36]card 
	for color := 0; color < 4; color++ {
		for figure := 0; figure < 9; figure++ {
			cardStack[color*9+figure] = card{color,figure,false}
		}
	}
	return cardStack
}

func shuffleCardStack(cardStack [36]card)([36]card){
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(36, func(i, j int){
		cardStack[i], cardStack[j] = cardStack[j], cardStack[i]
	})
	return cardStack
}

func distributeCards(shuffledStack [36]card)([4]player){
	var players [4]player
	var j int
	var hand []card

	for i := 0; i < 4; i++ {
		j = i*9
		hand = shuffledStack[0+j:9+j]
		players[i] = player{i, hand}
	}

	return players
}

func sortCards(players [4]player){
	for i := 0; i < 4; i++ {
		sort.Slice(players[i].hand, func(j, k int) bool { return players[i].hand[j].figure < players[i].hand[k].figure})
		sort.Slice(players[i].hand, func(j, k int) bool { return players[i].hand[j].color < players[i].hand[k].color})
	}
}


func getColor(color int)(string){
	switch {
	case color == 0:
		return "Eicheln"
	case color == 1:
		return "Rosen"
	case color == 2:
		return "Schellen"
	case color == 3:
		return "Schilten"
	}
	return ""
}

func getFigure(figure int)(string){
	switch {
	case figure == 0:
		return "6"
	case figure == 1:
		return "7"
	case figure == 2:
		return "8"
	case figure == 3:
		return "9"
	case figure == 4:
		return "10"
	case figure == 5:
		return "Under"
	case figure == 6:
		return "Ober"
	case figure == 7:
		return "König"
	case figure == 8:
		return "Ass"
	}
	return ""
}

func getCardName(cardInfo card)(string){
	var name string
	name = getColor(cardInfo.color)
	return name + " " + getFigure(cardInfo.figure)
} 

func showCards(players [4]player) {
	fmt.Println("This are your cards:")
	for i := 0; i<9; i++ {
		if players[0].hand[i].isPlayed == false{
			fmt.Println(getCardName(players[0].hand[i]))
		}
	}
}

func getTrump()(int) {
	var trump int
	fmt.Println("What should be trump(0:E, 1:R, 2:Sche, 3:Schi): ")
	fmt.Scanln(&trump)
	return trump
}

func main() {
	var cards = createCardStack()

	var shuffledStack = shuffleCardStack(cards)
	var players = distributeCards(shuffledStack)
	sortCards(players)

	showCards(players)
	var trump = getTrump()
	fmt.Println(getColor(trump), "is trump")
	//playCard()

	
}
