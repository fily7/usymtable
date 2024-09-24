package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// For JetBrains Mono

const (
	LINES      = 15
	SYMBOLS    = 15
	LINE_LEN   = SYMBOLS * 7
	PAGE_SIZE  = LINES * SYMBOLS
	END_SYMBOL = 0xffff
	PAGES      = 1 + END_SYMBOL/PAGE_SIZE
)

func clean_screen() {
	fmt.Printf("\033[0;0H")
	for i := 0; i < LINES*3; i++ {
		fmt.Print("\033[2K")
		fmt.Print("\033[1B")
	}
	fmt.Printf("\033[0;0H")
}

func print_contorl_panel() {
	control_panel := fmt.Sprintf("\n- [u]09af to symbol, 1-%d to page, [p]revious, [q]uit, [n]ext ", PAGES)
	control_panel = fmt.Sprintf("%s%s", control_panel, strings.Repeat("-", LINE_LEN-len(control_panel)))
	fmt.Println(control_panel)
}

func print_page(n int) {
	if n < 1 || n > PAGES {
		fmt.Printf("%d page out of range 1-%d", n, PAGES)
		print_contorl_panel()
		return
	}
	start := (n - 1) * PAGE_SIZE
	page_title := fmt.Sprintf("- PAGE %d of %d ", n, PAGES)
	fmt.Printf("%s%s\n\n", page_title, strings.Repeat("-", LINE_LEN-len(page_title)))
	for j := 0; j < LINES; j++ {
		for i := 0; i < SYMBOLS; i++ {
			if start+i > END_SYMBOL {
				break
			}
			fmt.Printf("   %c   ", start+i)
		}
		fmt.Println(" ")
		for i := 0; i < SYMBOLS; i++ {
			if start+i > END_SYMBOL {
				print_contorl_panel()
				return
			}
			fmt.Printf(" u%04x ", start+i)
		}
		fmt.Println(" ")
		start += SYMBOLS
	}
	print_contorl_panel()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	page := 274
	clean_screen()
	print_page(page)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if text == "q\n" {
			clean_screen()
			os.Exit(0)
		}
		if text == "p\n" {
			if page > 1 {
				page--
			}
		} else if text == "n\n" || text == "\n" {
			if page < PAGES {
				page++
			}
		} else if text[0] == 'u' && len(text) == 6 {
			char_num, err := strconv.ParseUint(text[1:len(text)-1], 16, 64)
			if err == nil {
				page = 1 + int(char_num)/PAGE_SIZE
			} else {
				fmt.Println("use u0000-uffff to find char")
				continue
			}
		} else if len(text) > 1 {
			to_page, err := strconv.Atoi(text[:len(text)-1])
			if err == nil {
				if 0 > to_page || to_page > PAGES {
					fmt.Printf("%d page out of range 1-%d", to_page, PAGES)
					print_contorl_panel()
					continue
				}
				page = to_page
			} else {
				print_contorl_panel()
				continue
			}
		}
		clean_screen()
		print_page(page)
	}
}
