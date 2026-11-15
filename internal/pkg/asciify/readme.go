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
func InjectReadMe(readmePath, img string) (err error) {
	tmpl := template.Must(template.ParseFiles(readmePath))

	output, err := os.Create(readmePath)
	if err != nil {
		return err
	}

	defer func() {
		if closeErr := output.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	return tmpl.Execute(output, readme{wrapCode(img)})
}

func wrapCode(code string) string {
	return "\n```\n" + code + "\n```\n"
}
