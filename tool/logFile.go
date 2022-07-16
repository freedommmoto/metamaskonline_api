package tool

import (
	"log"
	"os"
)

func AddErrorLogIntoFile(text string) {
	addLogToFile("error.log", text)
}
func AddApiLogIntoFile(text string) {
	addLogToFile("api.log", text)
}
func addLogToFile(fileName string, text string) {
	f, _ := os.Create(fileName)
	_, errAddString := f.WriteString(text + "\n")
	if errAddString != nil {
		log.Fatal(errAddString)
	}
	defer f.Close()
}
