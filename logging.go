package main

import (
	"fmt"
)

func title(text string) {
	fmt.Print("\n   ~~~", text, "~~~   \n\n")
}

func info(text ...any) {
	fmt.Print("[Info] ")
	fmt.Println(text...)
}
