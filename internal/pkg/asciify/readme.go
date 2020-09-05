package asciify

import (
	"html/template"
	"os"
)

// Readme template.
type readme struct {
	Image string
}

// InjectReadMe injects ascii in to readme where {{ . }} is.
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
	return "\n```\n" + code + "\n```\n"
}
