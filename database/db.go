package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Data struct {
	Id         int
	Memo       string
	Finished   int
	Created_at string
	Due_at     string
	Priority   string
	Tag        string
}

func dbConnect() *sql.DB {
	db, _ := sql.Open("sqlite3", "./yaminabe.db")
	return db
}

func ListAll() []Data {
	db := dbConnect()
	sql := "SELECT * from yaminabe;"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rows.Close()
	var data []Data
	for rows.Next() {
		id, memo, finished, created_at, due_at, priority, tag := 0, "", 0, "", "", "", ""
		err := rows.Scan(&id, &memo, &finished, &created_at, &due_at, &priority, &tag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		data = append(data, Data{Id: id, Memo: memo, Finished: finished, Created_at: created_at, Due_at: due_at, Priority: priority, Tag: tag})
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}

func AddDB(data Data) {
	db := dbConnect()
	_, err := db.Exec(
		`INSERT INTO "yaminabe" ("memo", "finished","due_at", "priority","tag") values (?, ?, ?, ?, ?);`,
		data.Memo,
		data.Finished,
		data.Due_at,
		data.Priority,
		data.Tag,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Remove(id int) {
	db := dbConnect()
	_, err := db.Exec(
		`DELETE FROM "yaminabe" WHERE id = ?;`,
		id,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Finish(id int) {
	db := dbConnect()
	_, err := db.Exec(
		`UPDATE "yaminabe" SET "finished" = 1 WHERE id = ?`,
		id,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CreateDB() {
	if IsExist("./yaminabe.db") {
		fmt.Println("Database already created")
		return
	}

	path, _ := os.Getwd()
	os.Chdir(path)

	os.Create("./yaminabe.db")

	db := dbConnect()

	var q = ""

	q = "CREATE TABLE yaminabe ("
	q += " id INTEGER PRIMARY KEY AUTOINCREMENT"
	q += ", memo TEXT NOT NULL"
	q += ", finished INTEGER NOT NULL DEFAULT 0"
	q += ", created_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))"
	q += ", due_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))"
	q += ", priority TEXT NOT NULL DEFAULT green"
	q += ", tag TEXT"
	q += ")"
	db.Exec(q)

	db.Close()
}

func IsExist(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
