package git

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/kushal2/gmd/pkg"
)

func Version() *exec.Cmd {
	return exec.Command("git", "--version")
}

func ShowBranch() *exec.Cmd {
	return exec.Command("git", "branch", "-l")
}

func CreateBranch(brName string) *exec.Cmd {
	output := ValidateBranch(brName)
	if output != nil {
		return exec.Command("echo", "\nBranch Already exist")
	}
	return exec.Command("git", "branch", brName)
}

func DeleteBranch(brName string) *exec.Cmd {
	output := ValidateBranch(brName)
	if output == nil {
		return exec.Command("echo", "\nBranch Does not exist")
	}
	return exec.Command("git", "branch", "-d", brName)
}

func ValidateBranch(brName string) []byte {
	c := exec.Command("git", "branch", "-l", brName)
	output, err := c.CombinedOutput()
	if err != nil {
		log.Printf("cmd failed with %s\n", err)
		return nil
	}
	if output != nil && string(output) == "" {
		return nil
	}
	return output
}

func BranchType() string {
	option := []string{"feat", "fix", "hotfix"}
	result := pkg.SelectOptions(option)
	return result
}

func TaskType() string {
	option := []string{"AB", "JIRA"}
	result := pkg.SelectOptions(option)
	return result
}

func ListBranch() []string {
	c := exec.Command("git", "branch", "-l")
	output, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd failed with %s\n", err)
	}
	val := strings.Split(strings.Split(string(output), "|")[0], "\n")
	//fmt.Printf("%T", val)
	//fmt.Print(val)
	brList := []string{}
	for i, s := range val {
		s = strings.TrimSpace(s)
		if s != "" {
			fmt.Println(i, s)
			if !strings.Contains(s, "main") {
				brList = append(brList, s)
			}
		}

	}

	return brList
}
