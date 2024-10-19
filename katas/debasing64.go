// Base64 is sort of a running gag around here. It's so widely used for the
// classic challenges that this one thing is certain – base64 or die.
// Your task is to read in a Base64-encoded string, decode it and print it out.
// Easy money.
package hackattic

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func debasing64() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
        decodedbytes, _ := base64.StdEncoding.DecodeString(input)
        decodedString := string(decodedbytes)
        fmt.Println(decodedString)
	}
}
