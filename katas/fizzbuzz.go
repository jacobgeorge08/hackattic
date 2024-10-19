// The "Fizz-Buzz test" is an interview question designed to help filter out
// the 99.5% of programming job candidates who can't seem to program their way
// out of a wet paper bag.
// Read in a single line with two numbers, N and M. Then simply print all the
// numbers from N to M. But for multiples of three print “Fizz” instead of the
// number and for the multiples of five print “Buzz”. For numbers which are
// multiples of both three and five print “FizzBuzz”.
package hackattic 

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fizzbuzz() {
	reader := bufio.NewReader(os.Stdin)
	nums, _ := reader.ReadString('\n')
	nums = strings.TrimSpace(nums)

	parts := strings.Split(nums, " ")
	num1, _ := strconv.Atoi(parts[0])
	num2, _ := strconv.Atoi(parts[1])

	for i := num1; i <= num2; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		}else {
		    fmt.Println(i)
		}
	}
}
