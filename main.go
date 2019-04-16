package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func spongeMockText(text []rune) string {
	upper := false
	for i, r := range text {
		if unicode.IsLetter(r) {
			if upper = !upper; upper {
				text[i] = unicode.ToUpper(r)
			} else {
				text[i] = unicode.ToLower(r)
			}
		}
	}
	return string(text)
}

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo \"hello\" | spongemock")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	fmt.Println(spongeMockText(output))

}
