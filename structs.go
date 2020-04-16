package main

//Error struct
type Error struct {
	Text string
}

//Error struct
type Responce struct {
	Text  string
	Hello string
}

//test db
type TestDb struct {
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Price       string `db:"price"`
}

type Results struct {
	Result []*TestDb
}

//new account struct

//part struct

//pro structs
