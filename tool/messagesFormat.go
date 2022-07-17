package tool

func GetLineText(caseNumber int) string {

	lineText := [4]string{
		"you send a wrong format code please try again with code number format 4digit example: `1234` or register here " + configValue.RegisterUrl,
		"you send the wrong code please try again you code is on dashboard.",
		"you send the correct code! you account is now active. \U0001F973 🎉",
		"you account is already register!. no need to send message to this line group amy more. :)",
	}
	return lineText[caseNumber-1]
}
