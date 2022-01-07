package output

import (
	"fmt"

	"github.com/fatih/color"
)

//Beautify prints the banner
func Beautify() {
	banner1 := "          __                           \n"
	banner2 := "    _  _ (_    |_  _|  | _  _ _||. _|_ \n"
	banner3 := "   (_)(_)__)|_||_)_)|/\\|(_)| (_|||_)|_ \n"
	banner4 := "   _/                                  \n"
	banner5 := ""
	banner6 := " > github.com/kenjoe41/goSubsWordlist\n"
	banner7 := " > evilzone.org\n"
	banner8 := "======================================"
	bannerPart1 := banner1 + banner2 + banner3 + banner4
	bannerPart2 := banner5 + banner6 + banner7 + banner8
	color.Cyan("%s\n", bannerPart1)
	fmt.Println(bannerPart2)

}
