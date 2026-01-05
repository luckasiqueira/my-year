package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "modernc.org/sqlite"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Icon     string
	Class    string
}

var (
	DB = connect()
)

func connect() *sql.DB {
	dsn := "file:db.sqlite"
	conn, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func List() ([]Item, error) {
	query := fmt.Sprint("SELECT * FROM acoes ORDER BY id ASC")
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	var items []Item
	for rows.Next() {
		var item Item
		rows.Scan(
			&item.ID,
			&item.Name,
			&item.Quantity,
			&item.Icon,
			&item.Class,
		)
		item.Name = strings.ToUpper(item.Name)
		items = append(items, item)
	}
	//list := make(map[string]Item)
	//for _, item := range items {
	//	list[item.Name] = item
	//}
	return items, nil
}

func Sum(item string) {
	query := `UPDATE acoes	SET quantity = quantity + 1	WHERE name LIKE ?;`
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	item = "%" + item + "%"
	_, err = stmt.Exec(item)
	if err != nil {
		log.Fatal(err)
	}
}
