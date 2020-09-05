package ascii

import (
	"bufio"
	"strings"
)

func trimWhitespace(txt string) string {
	var formatted []string

	scanner := bufio.NewScanner(strings.NewReader(txt))
	for scanner.Scan() {
		formatted = append(formatted, strings.TrimRight(scanner.Text(), " "))
	}

	return strings.Join(trimLeft(trimRight(formatted)), "\n")
}

func trimLeft(sa []string) []string {
	for i := 0; i < len(sa); i++ {
		if sa[i] != "" {
			return sa[i:]
		}
	}

	return sa
}

func trimRight(sa []string) []string {
	for i := len(sa) - 1; i >= 0; i-- {
		if sa[i] != "" {
			return sa[0 : i+1]
		}
	}

	return sa
}
