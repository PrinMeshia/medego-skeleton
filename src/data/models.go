package data

import (
	"database/sql"
	"os"

	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper db2.Session

// Models is the wrapper for all database models
type Models struct {
	// any models inserted here (and in the New function)
	// are easily accessible throughout the entire application
}

// New initializes the models package for use
func New(databasePool *sql.DB) Models {
	db = databasePool

	switch os.Getenv("DATABASE_TYPE") {
		case "mysql", "mariadb":
			upper, _ = mysql.New(databasePool)
		case "postgres","postgresql":	
			upper, _ = postgresql.New(databasePool)
		default :
		 //nothing
	}
	return Models{
	}
}

// getInsertID returns the integer value of a newly inserted id (using upper)
func getInsertID(i interface{}) int {
	switch v := i.(type) {
	case int:
		return v
	case int64:
		return int(v)
	default:
		// Handle other types or return an appropriate default value
		return 0
	}
}
