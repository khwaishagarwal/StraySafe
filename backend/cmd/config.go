package main

import (
	"github.com/aadi-1024/StraySafe/backend/pkg/database"
)

// App carries the app wide config and units
type App struct {
	Db        *database.Database
	JwtSecret []byte
}
