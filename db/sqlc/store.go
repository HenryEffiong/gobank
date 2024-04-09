package db

import "database/sql"

// Store provides all the functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}
