package main

import (
	"fmt"

	"github.com/kushal2/gmd/pkg"
	"github.com/kushal2/gmd/pkg/git"
)

func main() {

	for {
		option := []string{"version", "git", "exit"}
		result := pkg.SelectOptions(option)

		switch result {
		case "version":
			pkg.GMDVersion()
		case "git":
			git.GitCMD()
		case "exit":
			pkg.ExitGDM()
			break
		default:
			fmt.Println("Select a valid option")
		}

	}

}
