package models

import "database/sql"

type UserHandler struct {
	DB *sql.DB
}
