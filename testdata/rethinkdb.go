package rethinkdb

// Person struct is useless
type Person struct {
	ID   int    `gorethink:"id"`
	Name string `gorethink:"name"`
}
