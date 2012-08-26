package main

import (
	_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
)

type Attendee struct {
	name   string
	gender string
}

func (a Attendee) String() string {
	return "name: " + a.name + ", gender: " + a.gender
}

type DB struct {
	db   *sql.DB
	load int
}

func (d *DB) query(q string, ch chan<- []Attendee) {
	rows, err := d.db.Query(q)
	if err != nil {
		panic(err)
	}
	var attendees []Attendee
	for rows.Next() {
		var attendee Attendee
		rows.Scan(&attendee.name, &attendee.gender)
		attendees = append(attendees, attendee)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	d.load--
	ch <- attendees
}

func (d *DB) close() {
	d.db.Close()
}

func connect(dbname string) DB {
	conn, err := sql.Open("mysql", "root@/db1")
	if err != nil {
		panic(err)
	}
	return DB{db: conn}
}
