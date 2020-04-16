package main

const (
	testquery = ` SELECT * FROM items`
)

const (
	testPostQuery = ` INSERT INTO items (name, description, price) VALUES (?, ?, ?)`
)

const (
	testUpdateQuery = ` UPDATE items SET name = ?, description = ?, price = ? WHERE id = ?`
)

const (
	testDeleteQuery = ` DELETE FROM items WHERE id = ?`
)
