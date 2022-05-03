package git

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/kushal2/gmd/pkg"
)

func GitCMD() {
	option := []string{"version", "show-branch", "create-branch", "delete-branch"}
	result := pkg.SelectOptions(option)

	var cmd *exec.Cmd
	var msg string
	switch result {
	case "version":
		msg = "Git version: "
		cmd = Version()
	case "show-branch":
		msg = "Git show branch: \n"
		cmd = ShowBranch()
	case "create-branch":
		msg = "Git create branch:"
		brName := pkg.StringPrompt("name:")
		brType := BranchType()
		taskType := TaskType()
		tId := pkg.StringPrompt("Task ID:")

		fullBrName := brType + "/" + taskType + "#" + tId + "-" + strings.Join(strings.Split(brName, " "), "-")
		msg += fullBrName
		fmt.Println(fullBrName)
		cmd = CreateBranch(fullBrName)
	case "delete-branch":
		msg = "Git delete branch: "
		option := ListBranch()
		result := pkg.SelectOptions(option)
		msg += result
		cmd = DeleteBranch(result)
	default:
		fmt.Println("Select a valid option")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd failed with %s\n", err)
	}
	if output != nil {
		fmt.Println(string(output))
		err = beeep.Notify("Hi", msg+string(output), "")
		if err != nil {
			panic(err)
		}
	}

}
