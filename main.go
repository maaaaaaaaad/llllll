package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	seedMoney = 1000
	loseCount = 100
	winCount  = 100
	endPoint  = 5000
	overPoint = 0
)

var stdin = bufio.NewReader(os.Stdin)

func Playing() (uint8, error) {
	var n uint8

	_, err := fmt.Scanln(&n)

	if err != nil {
		stdin.ReadString('\n')
	}

	return n, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var gameNumber uint8 = uint8(rand.Intn(5+1)) + 1
	var initMoney uint32 = seedMoney

L:
	for {
		fmt.Print("please you enter number: ")
		result, err := Playing()
		if err != nil {
			fmt.Println("error:", err)
		} else {
			if result != gameNumber {
				switch {
				case initMoney == overPoint:
					fmt.Println("seed money 0 game over")
					break L
				case initMoney >= loseCount:
					initMoney -= loseCount
					fmt.Println("lose! current money:", initMoney)
				}
			}
		}
	}
}
