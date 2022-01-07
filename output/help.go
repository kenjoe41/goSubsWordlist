package output

import "fmt"

//PrintHelp prints the help.
func PrintHelp() {
	Beautify()
	fmt.Println(`Usage of goSubsWordlist:
	-iR bool
		Include Root Domain names in wordlist output.
	-help
		Print this help message.`)
}
