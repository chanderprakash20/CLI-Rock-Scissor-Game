package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	rock     = "rock"
	paper    = "paper"
	scissors = "scissors"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	wins, losses, draws := 0, 0, 0
	moves := []string{rock, paper, scissors}

	for {
		fmt.Print("Enter rock, paper, or scissors ('quit' to exit): ")
		userInput, _ := reader.ReadString('\n')
		userMove := strings.TrimSpace(userInput)
		if userMove == "quit" {
			break
		}

		if userMove != rock && userMove != paper && userMove != scissors {
			fmt.Println("Invalid move! Please choose rock, paper, or scissors.")
			continue
		}

		computerMove := moves[rand.Intn(3)]
		fmt.Println("Computer chose:", computerMove)

		switch {
		case userMove == computerMove:
			fmt.Println("-> Draw")
			draws++
		case userMove == rock && computerMove == scissors,
			userMove == scissors && computerMove == paper,
			userMove == paper && computerMove == rock:
			fmt.Println("-> You win!")
			wins++
		default:
			fmt.Println("-> Machine wins!")
			losses++
		}
	}

	fmt.Printf("\nGame Over! You won %d time(s), lost %d time(s), and drew %d time(s).\n", wins, losses, draws)
}
