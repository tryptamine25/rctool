package repository

import (
	"bufio"
	"fmt"
	"strings"

	"modules/internal/database"
	"modules/internal/git"
)

func AddRepository(db *database.DB, reader *bufio.Reader) {
	fmt.Print("Enter the repository path: ")
	repoPath, _ := reader.ReadString('\n')
	repoPath = strings.TrimSpace(repoPath)

	fmt.Print("Enter the name of the develop branch: ")
	devBranch, _ := reader.ReadString('\n')
	devBranch = strings.TrimSpace(devBranch)

	fmt.Print("Enter the name of the release branch: ")
	releaseBranch, _ := reader.ReadString('\n')
	releaseBranch = strings.TrimSpace(releaseBranch)

	fmt.Print("Enter the name of the main branch: ")
	mainBranch, _ := reader.ReadString('\n')
	mainBranch = strings.TrimSpace(mainBranch)

	if err := db.AddRepository(repoPath, devBranch, releaseBranch, mainBranch); err != nil {
		fmt.Printf("Error saving data: %v\n", err)
	} else {
		fmt.Println("Data successfully saved to the database.")
	}
}

func UpdateRepository(db *database.DB, reader *bufio.Reader) {
	repositories, err := db.GetRepositories()
	if err != nil {
		fmt.Printf("Error fetching repositories: %v\n", err)
		return
	}

	fmt.Println("Saved repositories:")
	for id, repoPath := range repositories {
		fmt.Printf("%d: %s\n", id, repoPath)
	}

	fmt.Print("Enter the repository ID to update: ")
	var repoID int
	fmt.Scan(&repoID)

	repoPath, exists := repositories[repoID]
	if !exists {
		fmt.Println("Invalid repository ID.")
		return
	}

	devBranch, releaseBranch, mainBranch, err := db.GetRepositoryBranches(repoID)
	if err != nil {
		fmt.Printf("Error fetching branches: %v\n", err)
		return
	}

	if err := git.UpdateRepository(repoPath, devBranch, releaseBranch, mainBranch); err != nil {
		fmt.Printf("Error updating repository: %v\n", err)
	}
}

func UpdateDevReleaseBranches(db *database.DB, reader *bufio.Reader) {
	repositories, err := db.GetRepositories()
	if err != nil {
		fmt.Printf("Error fetching repositories: %v\n", err)
		return
	}

	fmt.Println("Saved repositories:")
	for id, repoPath := range repositories {
		fmt.Printf("%d: %s\n", id, repoPath)
	}

	fmt.Print("Enter the repository ID to update: ")
	var repoID int
	fmt.Scan(&repoID)

	repoPath, exists := repositories[repoID]
	if !exists {
		fmt.Println("Invalid repository ID.")
		return
	}

	devBranch, releaseBranch, mainBranch, err := db.GetRepositoryBranches(repoID)
	if err != nil {
		fmt.Printf("Error fetching branches: %v\n", err)
		return
	}

	if err := git.UpdateDevReleaseBranches(repoPath, devBranch, releaseBranch, mainBranch); err != nil {
		fmt.Printf("Error updating branches: %v\n", err)
	}
}
