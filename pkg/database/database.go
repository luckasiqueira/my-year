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

func List() (map[string]Item, error) {
	query := fmt.Sprint("SELECT * FROM acoes")
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	var items []Item
	for rows.Next() {
		var item Item
		rows.Scan(&item.ID, &item.Name, &item.Quantity)
		item.Name = strings.ToUpper(item.Name)
		items = append(items, item)
	}
	list := make(map[string]Item)
	for _, item := range items {
		list[item.Name] = item
	}
	return list, nil
}

func Sum(item string) {
	query := `UPDATE acoes	SET quantity = quantity + 1	WHERE name LIKE ?;`
	_, err := DB.Exec(query, "%"+item+"%")
	if err != nil {
		log.Fatal(err)
	}
}
