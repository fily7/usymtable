package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	LINES   = 15
	SYMBOLS = 15
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	start := 0xf000
	end := 0xffff
	page_size := LINES * SYMBOLS
	start = (start / page_size) * page_size
	line_len := 7 * SYMBOLS
	for {
		page_title := fmt.Sprintf("- PAGE %d of %d ", 1+start/page_size, 1+end/page_size)
		fmt.Printf("%s%s\n\n", page_title, strings.Repeat("-", line_len-len(page_title)))
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
		page_footer := fmt.Sprintf("\n- [u]09af to symbol, 1-%d to page, [p]revious, [q]uit, [n]ext ", 1+end/page_size)
		fmt.Printf("%s%s\n", page_footer, strings.Repeat("-", line_len-len(page_footer)))
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if text == "q\n" {
			os.Exit(0)
		}
		if text == "p\n" {
			start -= page_size * 2
			if start < 0 {
				start = 0
			}
		} else if text[0] == 'u' && len(text) == 6 {
			char_num, err := strconv.ParseUint(text[1:len(text)-1], 16, 64)
			if err == nil {
				start = (int(char_num) / page_size) * page_size
			} else {
				fmt.Println("use u0000-uffff to find char")
			}
		} else if len(text) > 1 {
			to_page, err := strconv.Atoi(text[:len(text)-1])
			if err == nil {
				to_page--
				start = to_page * page_size
			}
		}
	}
}
