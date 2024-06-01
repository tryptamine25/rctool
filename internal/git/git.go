package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteGitCommands(repoPath string, commands []string) error {
	for _, cmd := range commands {
		fmt.Println("Executing command:", cmd)
		parts := strings.Fields(cmd)
		head := parts[0]
		parts = parts[1:]

		command := exec.Command(head, parts...)
		command.Dir = repoPath
		output, err := command.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error executing command '%s': %v - %s", cmd, err, output)
		}
		fmt.Printf("Result: %s\n", string(output))
	}
	return nil
}

func UpdateRepository(repoPath, devBranch, releaseBranch, mainBranch string) error {
	commands := []string{
		fmt.Sprintf("git checkout %s", devBranch),
		fmt.Sprintf("git pull origin %s", devBranch),
		fmt.Sprintf("git checkout %s", releaseBranch),
		fmt.Sprintf("git pull origin %s", releaseBranch),
		fmt.Sprintf("git checkout %s", mainBranch),
		fmt.Sprintf("git pull origin %s", mainBranch),
		fmt.Sprintf("git checkout %s", releaseBranch),
		fmt.Sprintf("git merge %s", mainBranch),
		fmt.Sprintf("git merge %s", devBranch),
		fmt.Sprintf("git push origin %s", releaseBranch),
	}

	return ExecuteGitCommands(repoPath, commands)
}

func UpdateDevReleaseBranches(repoPath, devBranch, releaseBranch, mainBranch string) error {
	commands := []string{
		fmt.Sprintf("git checkout %s", mainBranch),
		fmt.Sprintf("git pull origin %s", mainBranch),
		fmt.Sprintf("git checkout %s", releaseBranch),
		fmt.Sprintf("git pull origin %s", releaseBranch),
		fmt.Sprintf("git merge %s", mainBranch),
		fmt.Sprintf("git push origin %s", releaseBranch),
		fmt.Sprintf("git checkout %s", devBranch),
		fmt.Sprintf("git pull origin %s", devBranch),
		fmt.Sprintf("git merge %s", mainBranch),
		fmt.Sprintf("git push origin %s", devBranch),
	}

	return ExecuteGitCommands(repoPath, commands)
}
