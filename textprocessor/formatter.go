package textprocessor

import (
	"regexp"
	"strconv"
	"strings"
	// "unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Применяет изменение регистра к заданному количеству слов
func ApplyCaseChange(words []string, index int, mode string) []string {
	count := extractNumber(words[index])
	start := max(0, index-count)

	for i := start; i < index; i++ {
		switch mode {
		case "upper":
			words[i] = strings.ToUpper(words[i])
		case "lower":
			words[i] = strings.ToLower(words[i])
		case "capital":
			words[i] = cases.Title(language.Und).String(words[i])
		}
	}

	return append(words[:index], words[index+1:]...)
}

// Исправляет пунктуацию в тексте
func FixPunctuation(text string) string {
	lines := strings.Split(text, "\n") // Разбиваем по строкам, чтобы сохранить переносы
	for i, line := range lines {
		re := regexp.MustCompile(`\s*([,!.?:;])\s*`)
		line = re.ReplaceAllString(line, "$1 ")

		reEllipsis := regexp.MustCompile(`\s*\.\.\.\s*`)
		line = reEllipsis.ReplaceAllString(line, "...")

		lines[i] = strings.TrimSpace(line)
	}
	return strings.Join(lines, "\n") // Собираем текст обратно с переносами строк
}

// Исправляет использование артиклей "a" и "an"
func FixArticle(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" {
			nextWord := words[i+1]
			if strings.ContainsAny(strings.ToLower(string(nextWord[0])), "aeiouh") {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}

// Применяет Capitalize к X словам перед маркером (cap, X)
func ApplyCapitalize(words []string) []string {
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], "(cap") {
			count := extractNumber(words[i])

			for j := max(0, i-count); j < i; j++ {
				words[j] = Capitalize(words[j])
			}

			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return words
}

// Делает первую букву каждого слова заглавной
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

// Извлекает число из строки, например: (cap, 3) → 3
func extractNumber(word string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(word)
	if match == "" {
		return 1
	}
	num, err := strconv.Atoi(match)
	if err != nil {
		return 1
	}
	return num
}

// Возвращает максимальное значение из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
