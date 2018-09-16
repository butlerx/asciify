package asciify

import (
	"bufio"
	"html/template"
	"os"
	"strings"
)

// Readme template
type readme struct {
	Image string
}

// InjectReadMe injects ascii in to readme where {{ . }} is
func InjectReadMe(readmePath, img string) error {
	tmpl := template.Must(template.ParseFiles(readmePath))
	output, err := os.Create(readmePath)
	if err != nil {
		return err
	}
	defer output.Close()
	return tmpl.Execute(output, readme{wrapCode(img)})
}

func wrapCode(code string) string {
	return "\n```\n" + trimWhitespace(code) + "\n```\n"
}

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
