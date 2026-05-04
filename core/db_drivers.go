package core

import (
	// Import PostgreSQL driver
	_ "github.com/lib/pq"

	// Import MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// This file imports database drivers so dbx can connect to PostgreSQL and MySQL.
// The drivers register themselves on import via init() functions.
//
// Usage:
//   dbx.Open("postgres", "postgres://user:pass@localhost/db")
//   dbx.Open("mysql", "user:pass@tcp(localhost)/db")
