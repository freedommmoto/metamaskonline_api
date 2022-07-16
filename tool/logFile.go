package tool

import (
	"log"
	"os"
)

func AddErrorLogIntoFile(text string) {
	f, err := os.Create("error.log")
	if err != nil {
		log.Fatal(err)
	}

	_, errAddString := f.WriteString(text + "\n")
	if errAddString != nil {
		log.Fatal(errAddString)
	}
	defer f.Close()
}
