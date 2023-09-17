package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type card struct {
	color    int
	figure   int
	isPlayed bool
}

type player struct {
	id   int
	hand []card
}

type team struct {
	id     int
	points int
}

type round struct {
	trump          int
	cardsPlayed    []card
	startingPlayer int
}

func createCardStack() [36]card {
	var cardStack [36]card
	for color := 0; color < 4; color++ {
		for figure := 0; figure < 9; figure++ {
			cardStack[color*9+figure] = card{color, figure, false}
		}
	}
	return cardStack
}

func shuffleCardStack(cardStack [36]card) [36]card {
	rand.Shuffle(36, func(i, j int) {
		cardStack[i], cardStack[j] = cardStack[j], cardStack[i]
	})
	return cardStack
}

func distributeCards(shuffledStack [36]card) [4]player {
	var players [4]player
	var j int
	var hand []card

	for i := 0; i < 4; i++ {
		j = i * 9
		hand = shuffledStack[0+j : 9+j]
		players[i] = player{i, hand}
	}

	return players
}

func sortCards(players [4]player) {
	for i := 0; i < 4; i++ {
		sort.Slice(players[i].hand, func(j, k int) bool { return players[i].hand[j].figure < players[i].hand[k].figure })
		sort.Slice(players[i].hand, func(j, k int) bool { return players[i].hand[j].color < players[i].hand[k].color })
	}
}

func getColor(color int) string {
	switch {
	case color == 0:
		return "Eicheln"
	case color == 1:
		return "Rosen"
	case color == 2:
		return "Schellen"
	case color == 3:
		return "Schilten"
	case color == 4:
		return "Obe"
	case color == 5:
		return "Unde"
	}

	return ""
}

func getFigure(figure int) string {
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

func getCardName(cardInfo card) string {
	var name string
	name = getColor(cardInfo.color)
	return name + " " + getFigure(cardInfo.figure)
}

func showCards(player player) {
	fmt.Println("This are the cards of player", player.id, ":")
	for i := 0; i < 9; i++ {
		if player.hand[i].isPlayed == false {
			fmt.Println(i, getCardName(player.hand[i]))
		}
	}
}

func getTrump() int {
	var trump int
	fmt.Println("What should be trump(0:E, 1:R, 2:Sche, 3:Schi, 4:O, 5:U): ")
	fmt.Scanln(&trump)
	fmt.Println(getColor(trump), "is trump")
	return trump
}

func playCards(currentRound round, players [4]player) round {
	var currentCardNumb int
	var currentCard card

	for i := 0; i < 4; i++ {
		showCards(players[(currentRound.startingPlayer+i)%4])
		fmt.Println("What card to play: ")
		fmt.Scanln(&currentCardNumb)
		currentCard = players[(currentRound.startingPlayer+i)%4].hand[currentCardNumb]
		fmt.Println(getCardName(currentCard), "is played")
		currentRound.cardsPlayed = append(currentRound.cardsPlayed, currentCard)
	}

	return currentRound
}

func isHigherCard(card1 card, card2 card, trump int, outcardColor int) (card card) {
	if card1.color == trump && card2.color == trump {
		if card1.figure == 5 {
			return card1
		} else if card2.figure == 5 {
			return card2
		} else if card1.figure == 3 {
			return card1
		} else if card2.figure == 3 {
			return card2
		} else if card1.figure > card2.figure {
			return card1
		} else {
			return card2
		}
	} else if card1.color == trump {
		return card1
	} else if card2.color == trump {
		return card2
	} else if card1.color == outcardColor && card2.color == outcardColor {
		if card1.figure > card2.figure {
			return card1
		} else {
			return card2
		}
	} else if card1.color == outcardColor {
		return card1
	} else if card2.color == outcardColor {
		return card2
	}
	return card1
}

func evalStich(currentRound round) (takeStich int) {
	var outcardColor int
	var currentBest card

	outcardColor = currentRound.cardsPlayed[0].color
	currentBest = currentRound.cardsPlayed[0]
	takeStich = 0

	for i := 1; i < 4; i++ {
		if isHigherCard(currentBest, currentRound.cardsPlayed[i], currentRound.trump, outcardColor) != currentBest {
			currentBest = currentRound.cardsPlayed[i]
			takeStich = i
		}
	}

	fmt.Println("Player", takeStich, "sticht with", getCardName(currentBest))

	return takeStich
}

func playRound(players [4]player) {
	showCards(players[0])
	var trump = getTrump()
	var round = round{trump, nil, 1}
	round = playCards(round, players)
	evalStich(round)

}

func main() {
	var cards = createCardStack()
	var shuffledStack = shuffleCardStack(cards)
	var players = distributeCards(shuffledStack)
	sortCards(players)

	playRound(players)
}
