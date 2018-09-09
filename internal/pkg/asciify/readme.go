package asciify

import (
	"html/template"
	"os"
)

// InjectReadMe injects ascii in to readme where {{ . }} is
func InjectReadMe(readmePath, img string) error {
	tmpl, err := template.New("readme").ParseFiles(readmePath)
	if err != nil {
		return err
	}
	output, err := os.Create(readmePath)
	if err != nil {
		return err
	}
	defer output.Close()
	return tmpl.Execute(output, img)
}
