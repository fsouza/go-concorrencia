package main

import (
	"container/heap"
	"fmt"
)

type DBPool []*DB

func (p *DBPool) Less(i, j int) bool {
	return (*p)[i].load < (*p)[j].load
}

func (p *DBPool) Len() int {
	return len(*p)
}

func (p *DBPool) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *DBPool) Push(db interface{}) {
	*p = append(*p, db.(*DB))
}

func (p *DBPool) Pop() interface{} {
	l := p.Len()
	db := (*p)[l-1]
	*p = (*p)[:l-1]
	return db
}

type Balancer struct {
	pool DBPool
}

func NewBalancer(dbnames []string) Balancer {
	lb := Balancer{make(DBPool, 0)}
	for _, name := range dbnames {
		db := connect(name)
		heap.Push(&lb.pool, &db)
	}
	return lb
}

func (lb Balancer) query(q string) chan []Attendee {
	db := heap.Pop(&lb.pool).(*DB)
	ch := make(chan []Attendee)
	go db.query(q, ch)
	db.load++
	heap.Push(&lb.pool, db)
	return ch
}

func main() {
	dbnames := []string{"db1", "db2", "db3", "db4", "db5"}
	lb := NewBalancer(dbnames)
	ch := lb.query("select name, gender from attendees where name like 'C%'")
	for _, a := range <-ch {
		fmt.Println(a)
	}
}
