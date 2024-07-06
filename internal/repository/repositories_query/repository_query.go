package repositories_query

import (
	_ "embed"
)

var (
	// go:embed repositories_query/select_by_name.sql
	GetUserByUsername string
	// go:embed repositories_query/insert.sql
	InsertUser string
	// go:embed repositories_query/select.sql
	GetUserByID string
	// go:embed repositories_query/select_all.sql
	GetAllUsers string
)
