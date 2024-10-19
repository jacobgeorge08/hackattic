package hackattic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parens() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			return
		}

		stack := []rune{} 
		balanced := true

		for _, char := range input {
			if char == '(' {
				stack = append(stack, char)
			} else if char == ')' {
				if len(stack) == 0 {
					balanced = false 
					break
				}
				stack = stack[:len(stack)-1] // Pop from stack
			}
		}

		if balanced && len(stack) == 0 { // Balanced and stack is empty
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
