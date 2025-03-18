package textprocessor

import (
	"strings"
)

func ProcessText(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		switch {
		case words[i] == "(hex)" && i > 0:
			words[i-1] = ConvertHexToDec(words[i-1])
			words = append(words[:i], words[i+1:]...)

		case words[i] == "(bin)" && i > 0:
			words[i-1] = ConvertBinToDec(words[i-1])
			words = append(words[:i], words[i+1:]...)

		case strings.HasPrefix(words[i], "(up"):
			words = ApplyCaseChange(words, i, "upper")

		case strings.HasPrefix(words[i], "(low"):
			words = ApplyCaseChange(words, i, "lower")

		case strings.HasPrefix(words[i], "(cap"):
			words = ApplyCaseChange(words, i, "capital")
		}
	}

	text = strings.Join(words, " ")
	text = FixPunctuation(text)
	text = FixArticle(text)
	return text
}
