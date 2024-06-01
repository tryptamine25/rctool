package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const createTableSQL = `CREATE TABLE IF NOT EXISTS branches (
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"repo_path" TEXT,
	"dev_branch" TEXT,
	"release_branch" TEXT,
	"main_branch" TEXT
);`

type DB struct {
	conn *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	conn, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	if _, err := conn.Exec(createTableSQL); err != nil {
		return nil, err
	}

	return &DB{conn: conn}, nil
}

func (db *DB) Close() {
	db.conn.Close()
}

func (db *DB) AddRepository(repoPath, devBranch, releaseBranch, mainBranch string) error {
	insertSQL := `INSERT INTO branches (repo_path, dev_branch, release_branch, main_branch) VALUES (?, ?, ?, ?);`
	statement, err := db.conn.Prepare(insertSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(repoPath, devBranch, releaseBranch, mainBranch)
	return err
}

func (db *DB) GetRepositories() (map[int]string, error) {
	rows, err := db.conn.Query("SELECT id, repo_path FROM branches")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	repositories := make(map[int]string)
	for rows.Next() {
		var id int
		var repoPath string
		if err := rows.Scan(&id, &repoPath); err != nil {
			return nil, err
		}
		repositories[id] = repoPath
	}
	return repositories, rows.Err()
}

func (db *DB) GetRepositoryBranches(id int) (string, string, string, error) {
	row := db.conn.QueryRow("SELECT dev_branch, release_branch, main_branch FROM branches WHERE id = ?", id)
	var devBranch, releaseBranch, mainBranch string
	err := row.Scan(&devBranch, &releaseBranch, &mainBranch)
	return devBranch, releaseBranch, mainBranch, err
}
