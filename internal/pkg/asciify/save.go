package asciify

import "os"

// Save writes a string too file, creates the file if it doesnt exist.
func Save(path, img string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()
	_, err = file.WriteString(img)

	return err
}
