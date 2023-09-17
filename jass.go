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
	points1        int
	points2        int
	startingPlayer int
}

type stich struct {
	cardsPlayed    []card
	startingPlayer int
	points         int
	winningPlayer  int
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

func getTrump(currentPlayer player) int {
	var trump int
	showCards(currentPlayer)
	fmt.Println("What should be trump(0:E, 1:R, 2:Sche, 3:Schi, 4:O, 5:U): ")
	fmt.Scanln(&trump)
	fmt.Println(getColor(trump), "is trump")
	return trump
}

func playCards(currentStich stich, players [4]player) stich {
	var currentCardNumb int
	var currentCard card
	var currentPlayer player

	for i := 0; i < 4; i++ {
		currentPlayer = players[(currentStich.startingPlayer+i)%4]
		showCards(currentPlayer)
		fmt.Println("What card to play: ")
		fmt.Scanln(&currentCardNumb)
		currentCard = currentPlayer.hand[currentCardNumb]
		currentPlayer.hand[currentCardNumb].isPlayed = true
		fmt.Println(getCardName(currentCard), "is played")
		currentStich.cardsPlayed = append(currentStich.cardsPlayed, currentCard)
	}

	return currentStich
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

func countCard(currentCard card, trump int) int {
	if currentCard.color == trump {
		if currentCard.figure == 5 {
			return 20
		} else if currentCard.figure == 3 {
			return 14
		}
	}
	if currentCard.figure == 8 {
		return 11
	} else if currentCard.figure == 7 {
		return 4
	} else if currentCard.figure == 6 {
		return 3
	} else if currentCard.figure == 5 {
		return 2
	} else if currentCard.figure == 4 {
		return 10
	}
	return 0

}

func countStich(cardsPlayed []card, trump int) int {
	var score int
	score = 0
	for i := 0; i < 4; i++ {
		score += countCard(cardsPlayed[i], trump)
	}
	return score
}

func evalStich(currentStich stich, trump int) stich {
	var outcardColor int
	var currentBest card

	outcardColor = currentStich.cardsPlayed[0].color
	currentBest = currentStich.cardsPlayed[0]

	for i := 1; i < 4; i++ {
		if isHigherCard(currentBest, currentStich.cardsPlayed[i], trump, outcardColor) != currentBest {
			currentBest = currentStich.cardsPlayed[i]
			currentStich.winningPlayer = i
		}
	}

	currentStich.points = countStich(currentStich.cardsPlayed, trump)

	fmt.Println("Player", currentStich.winningPlayer, "sticht with", getCardName(currentBest), "counting", currentStich.points, "points")

	return currentStich
}

func playRound(players [4]player, currentRound round) {
	var currentStich stich
	var startingPlayer int
	startingPlayer = currentRound.startingPlayer
	for i := 0; i < 9; i++ {
		currentStich = stich{nil, startingPlayer, 0, startingPlayer}
		currentStich = playCards(currentStich, players)
		currentStich = evalStich(currentStich, currentRound.trump)
		if currentStich.winningPlayer == 0 || currentStich.winningPlayer == 2 {

		}
	}
}

func main() {
	var startPlayer int
	var trump int
	var currentRound round

	var cards = createCardStack()
	var shuffledStack = shuffleCardStack(cards)
	var players = distributeCards(shuffledStack)
	sortCards(players)
	var team1 = team{0, 0}
	var team2 = team{1, 0}

	startPlayer = 0
	for team1.points < 2500 && team2.points < 2500 {
		trump = getTrump(players[startPlayer])
		currentRound = round{trump, 0, 0, startPlayer}
		playRound(players, currentRound)
		startPlayer += 1
	}
}
