package save

import (
	"os"
	"path/filepath"
	"strconv"
)

// ToFiles saves a slice of images data into numbered files.
func ToFiles(images []string, outputFolder string) {
	_ = os.MkdirAll(outputFolder, os.ModePerm)
	for i := range images {
		filename := "page" + strconv.Itoa(i+1) + ".jpg"
		file, err := os.Create(filepath.Join(outputFolder, filename))
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = file.WriteString(images[i])
		if err != nil {
			panic(err)
		}
	}
}
