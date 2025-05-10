package dogsay

import (
	"fmt"
	"strings"
)

const DOG = ` /)-_-(\  /
  (o o)
   \o/\__-----.
    \      __  \
     \| /_/  \ /\__/
      ||      \\
      ||      //
      /|     /|`

func bubble(text string) string {
	lines := strings.Split(text, "\n")
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	top := "           " + " " + strings.Repeat("_", maxLen+2)
	bottom := "           " + " " + strings.Repeat("-", maxLen+2)
	var middle []string
	for _, line := range lines {
		padded := line + strings.Repeat(" ", maxLen-len(line))
		middle = append(middle, "           "+fmt.Sprintf("< %s >", padded))
	}

	connector := "           /"
	return fmt.Sprintf("%s\n%s\n%s\n%s", top, strings.Join(middle, "\n"), bottom, connector)
}

func DogSay(text string) string {
	return bubble(text) + "\n" + DOG
}

func PrintDogSay(text string) {
	fmt.Println(DogSay(text))
}
