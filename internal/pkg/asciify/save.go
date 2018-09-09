package asciify

import "os"

func Save(path, img string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(img)
	return err
}
