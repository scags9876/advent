package main

import (
	"fmt"

	"github.com/scags9876/adventOfCode/lib"
)

const inputFilename = "input.txt"
const testInputFilename = "testinput.txt"

const verbose = false

func main() {
	//fmt.Println("start")
	input := lib.GetInputStrings(inputFilename)
	part1(input)
	part2(input)
}

type gameState struct {
	id      int
	player1 []int
	player2 []int
	round   int
}

func part1(input []string) {
	game := parseInput(input)

	//fmt.Printf("Initial Game State: %#v", game)
	previousHands := make(map[string]bool)
	var winner int

	for {
		game.round++
		//fmt.Printf("-- Round %d --\n", game.round)
		//fmt.Printf("Player 1's deck: %s\n", lib.JoinInts(game.player1, ", "))
		//fmt.Printf("Player 2's deck: %s\n", lib.JoinInts(game.player2, ", "))

		thisHand := lib.JoinInts(game.player1, ",") + "#" + lib.JoinInts(game.player2, ", ")
		if _, ok := previousHands[thisHand]; ok {
			//fmt.Println("Player 1 automatically wins to avoid infinite recursion")
			winner = 1
			break
		}

		previousHands[thisHand] = true

		var card1, card2 int
		card1, game.player1 = game.player1[0], game.player1[1:]
		card2, game.player2 = game.player2[0], game.player2[1:]
		//fmt.Printf("Player 1 plays: %d\n", card1)
		//fmt.Printf("Player 2 plays: %d\n", card2)

		if card1 > card2 {
			//fmt.Println("Player 1 wins the round!")
			game.player1 = append(game.player1, card1, card2)
		} else {
			//fmt.Println("Player 2 wins the round!")
			game.player2 = append(game.player2, card2, card1)
		}
		//fmt.Println()
		if len(game.player1) == 0 || len(game.player2) == 0 {
			break
		}
	}

	//fmt.Print("== Post-game results ==\n")
	//fmt.Printf("Player 1's deck: %s\n", lib.JoinInts(game.player1, ", "))
	//fmt.Printf("Player 2's deck: %s\n", lib.JoinInts(game.player2, ", "))
	if winner == 0 {
		if len(game.player1) == 0 {
			winner = 2
		}
		if len(game.player2) == 0 {
			winner = 1
		}
	}
	//fmt.Printf("Player %d wins game %d!\n", winner, game.id)

	var winningDeck []int
	if winner == 2 {
		winningDeck = game.player2
	}
	if winner == 1 {
		winningDeck = game.player1
	}
	//fmt.Println()

	score := scoreDeck(winningDeck)
	fmt.Printf("Part 1: %d\n", score)
}
func part2(input []string) {
	game := parseInput(input)

	//fmt.Printf("Initial Game State: %#v", game)
	winner, game := playRecursiveCombat(game)

	var winningDeck []int
	if winner == 1 {
		winningDeck = game.player1
	} else {
		winningDeck = game.player2
	}
	score := scoreDeck(winningDeck)
	fmt.Printf("Part 2: %d\n", score)
}

var gameID int

func playRecursiveCombat(game gameState) (int, gameState) {
	//fmt.Printf("\n==== Game %d ====\n", game.id)
	previousHands := make(map[string]bool)
	var winner int

	for {
		game.round++
		//fmt.Printf("-- Round %d --\n", game.round)
		//fmt.Printf("Player 1's deck: %s\n", lib.JoinInts(game.player1, ", "))
		//fmt.Printf("Player 2's deck: %s\n", lib.JoinInts(game.player2, ", "))

		thisHand := lib.JoinInts(game.player1, ",") + "#" + lib.JoinInts(game.player2, ", ")
		if _, ok := previousHands[thisHand]; ok {
			//fmt.Println("Player 1 automatically wins to avoid infinite recursion")
			winner = 1
			break
		}

		previousHands[thisHand] = true

		var card1, card2 int
		card1, game.player1 = game.player1[0], game.player1[1:]
		card2, game.player2 = game.player2[0], game.player2[1:]
		//fmt.Printf("Player 1 plays: %d\n", card1)
		//fmt.Printf("Player 2 plays: %d\n", card2)

		var winner int
		if len(game.player1) >= card1 && len(game.player2) >= card2 {
			gameID++
			player1 := make([]int, card1)
			player2 := make([]int, card2)
			for i, card := range game.player1[:card1] {
				player1[i] = card
			}
			for i, card := range game.player2[:card2] {
				player2[i] = card
			}
			subGame := gameState{
				id:      gameID,
				player1: player1,
				player2: player2,
			}
			//fmt.Println("Playing a sub-game to determine the winner...")
			winner, subGame = playRecursiveCombat(subGame)
			//fmt.Printf("...anyway, back to game %d.\n", game.id)
		} else {
			if card1 > card2 {
				winner = 1
			} else {
				winner = 2
			}
		}
		//fmt.Printf("Player %d wins round %d of game %d!\n", winner, game.round, game.id)
		if winner == 1 {
			game.player1 = append(game.player1, card1, card2)
		} else {
			game.player2 = append(game.player2, card2, card1)
		}
		//fmt.Println()
		if len(game.player1) == 0 || len(game.player2) == 0 {
			break
		}
	}
	//fmt.Printf("== Post-game %d results ==\n", game.id)
	//fmt.Printf("Player 1's deck: %s\n", lib.JoinInts(game.player1, ", "))
	//fmt.Printf("Player 2's deck: %s\n", lib.JoinInts(game.player2, ", "))
	if winner == 0 {
		if len(game.player1) == 0 {
			winner = 2
		}
		if len(game.player2) == 0 {
			winner = 1
		}
	}
	//fmt.Printf("Player %d wins game %d!\n", winner, game.id)

	return winner, game
}

func parseInput(input []string) gameState {
	game := gameState{
		id:      1,
		player1: make([]int, 0),
		player2: make([]int, 0),
	}
	var player int
	for _, line := range input {
		switch line {
		case "Player 1:":
			player = 1
		case "Player 2:":
			player = 2
		case "":
			continue
		default:
			card := lib.ToInt(line)
			if player == 1 {
				game.player1 = append(game.player1, card)
			} else if player == 2 {
				game.player2 = append(game.player2, card)
			}
		}
	}
	return game
}

func scoreDeck(deck []int) int {
	score := 0
	multiplier := len(deck)
	for _, card := range deck {
		score += card * multiplier
		multiplier--
	}
	return score
}
