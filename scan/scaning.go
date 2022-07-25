package scan

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "0" {
			break
		}
		fmt.Println(text)
	}
}