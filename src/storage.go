package main
import (
	"database/sql"
	_ "github.com/lib/pq"
)
//  swappable storage interface
type Storage interface{
	GetGapRID(int) (*Gap, error)
}

type PGStore struct{
	db *sql.DB
}

func newPGStore()(*PGStore, error){
	// dummy pg access
	connStr := "user=postgres dbname=gisdb password=example host=db port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err!= nil{
		return nil, err
	}
	if err := db.Ping(); err!= nil {
		return nil, err
	}

	return &PGStore{
		db: db,
	}, nil
}