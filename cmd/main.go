package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"modules/internal/database"
	"modules/internal/repository"
)

func main() {
	db, err := database.NewDB("./branches.db")
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select an action: (a) add repository, (u) update repository, (d) update develop and release branches, (e) exit")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)

		switch action {
		case "a":
			repository.AddRepository(db, reader)
		case "u":
			repository.UpdateRepository(db, reader)
		case "d":
			repository.UpdateDevReleaseBranches(db, reader)
		case "e":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid action. Please try again.")
		}
	}
}
