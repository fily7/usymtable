package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	LINES   = 20
	SYMBOLS = 15
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	start := 0xf000
	end := 0xffff
	fmt.Printf("For print next %d lines press Enter to quit press q\n", LINES)
	for {
		text, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}
		if text == 'q' {
			os.Exit(0)
		}
		for j := 0; j < LINES; j++ {
			for i := 0; i < SYMBOLS; i++ {
				if start+i > end {
					break
				}
				fmt.Printf("   %c   ", start+i)
			}
			fmt.Println(" ")
			for i := 0; i < SYMBOLS; i++ {
				if start+i > end {
					os.Exit(0)
				}
				fmt.Printf(" u%04x ", start+i)
			}
			fmt.Println(" ")
			start += SYMBOLS
		}
	}
}
