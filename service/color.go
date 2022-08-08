package service

const WarningText = "\u001B[31m"
const SuccessText = "\u001B[32m"
const InfoText = "\u001b[36m"

func PrintInfoText(toolName string) string {
	return InfoText + "Install " + toolName + " .... "
}

func PrintDoneText() string {
	return SuccessText + "DONE \n"
}

func PrintErrorText() string {
	return WarningText + "ERROR \n"
}
