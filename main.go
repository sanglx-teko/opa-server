package main

import (
	"flag"
	"fmt"
	"opatutorial/utils"
)

func main() {
	flag.Parse() // get the arguments from command line

	destinationfile := flag.Arg(0)
	sourcedir := flag.Arg(1)

	if err := utils.Compress_tarball(destinationfile, sourcedir); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}
