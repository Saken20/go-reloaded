package textprocessor

import (
	"strconv"
	"strings"
)

// Изменение регистра
func ApplyCaseChange(words []string, index int, mode string) []string {
	count := 1
	if len(words[index]) > 4 {
		count = extractNumber(words[index])
	}

	start := max(0, index-count)
	for i := start; i < index; i++ {
		switch mode {
		case "upper":
			words[i] = strings.ToUpper(words[i])
		case "lower":
			words[i] = strings.ToLower(words[i])
		case "capital":
			words[i] = strings.Title(words[i])
		}
	}

	return append(words[:index], words[index+1:]...)
}

// Исправление пунктуации
func FixPunctuation(text string) string {
	replacements := []struct {
		old string
		new string
	}{
		{" ,", ","},
		{" .", "."},
		{" !", "!"},
		{" ?", "?"},
		{" :", ":"},
		{" ;", ";"},
		{" ...", "..."},
		{" !?", "!?"},
	}

	for _, r := range replacements {
		text = strings.ReplaceAll(text, r.old, r.new)
	}
	return text
}

// Исправление "a"/"an"
func FixArticle(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" {
			nextWord := words[i+1]
			if strings.ContainsAny(strings.ToLower(nextWord[:1]), "aeiouh") {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}

// Извлечение числа из (up, 3)
func extractNumber(word string) int {
	start := strings.Index(word, ",")
	end := strings.Index(word, ")")
	if start == -1 || end == -1 {
		return 1
	}

	num, err := strconv.Atoi(word[start+1 : end])
	if err != nil {
		return 1
	}
	return num
}

// Функция max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
