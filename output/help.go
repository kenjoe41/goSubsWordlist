package output

import "fmt"

//PrintHelp prints the help.
func PrintHelp() {
	Beautify()
	fmt.Println(`Usage of goSubsWordlist:
	-iR Bool
		Include Root Domain names in wordlist output.
	-t Int
		Threads for Concurrency. Default is 20.
	-help
		Print this help message.`)
}
