package main

import (
	"fmt"
)

func query(q string, dbs []DB) []Attendee {
	ch := make(chan []Attendee, len(dbs))
	for _, db := range dbs {
		go db.query(q, ch)
	}
	return <-ch
}

func main() {
	dbnames := []string{"db1", "db2", "db3", "db4", "db5"}
	dbs := make([]DB, len(dbnames))
	for i, name := range dbnames {
		dbs[i] = connect(name)
		defer dbs[i].close()
	}
	q := "select name, gender from attendees where name like 'A%'"
	for _, a := range query(q, dbs) {
		fmt.Println(a)
	}
}
