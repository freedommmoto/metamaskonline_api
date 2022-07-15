package tool

func GetLineText(caseNumber int) string {

	lineText := [5]string{
		"you account is not register yet. please go to this url for register =>" + configValue.RegisterUrl + " ",
		"you send the correct code! you account is now active. \U0001F973 🎉",
		"you send the wrong code please try again you new code have been re-generated.",
		"you send a wrong format code please try again with code number format 4digit example: `1234` ",
		"you account is already register!. no need to send message to this line group amy more. :)",
	}
	return lineText[caseNumber-1]
}
