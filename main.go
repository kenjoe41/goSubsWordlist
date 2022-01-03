package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kenjoe41/goSubsWordlist/utils"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		scDomain := strings.TrimSpace(strings.ToLower(sc.Text()))

		domain := utils.CleanDomain(scDomain)

		fmt.Println(domain)

		subdomain := utils.ExtractSubdomain(domain)
		fmt.Println(subdomain)

		// check there were no errors reading stdin (unlikely)
		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}
	}

}
