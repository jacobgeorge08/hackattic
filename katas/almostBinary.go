package hackattic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
        input := scanner.Text()
		sixteenBit := ""
		for _, r := range input {
			if r == '.' {
				sixteenBit += "0"
			} else if r == '#' {
				sixteenBit += "1"
			}
		}
		base10, _ := strconv.ParseInt(sixteenBit, 2, 64)
		fmt.Println(base10)
	}
}
