package main

import (
	"database/sql"
	// "encoding/json"

	_ "github.com/lib/pq"
)

//  swappable storage interface
type Storage interface{
	GetGaps() ([]*Gap, error)
	// GetGapRID(int) (*Gap, error)
}

type PGStore struct{
	db *sql.DB
}

func newPGStore()(*PGStore, error){
	// dummy pg access
	connStr := "user=postgres dbname=gisdb password=example host=localhost port=5433 sslmode=disable"
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

func (s *PGStore) GetGaps() ([]*Gap, error) {
	rows, err := s.db.Query("select * from public_test.\"dataGap\" limit 10; ")
	if err != nil {
		return nil, err
	}
	gaps := []*Gap{} // why interface{ }? 
	for rows.Next() {
		gap, err := scanIntoGap(rows)
		if err != nil {
			return nil, err
		}

		gaps = append(gaps,gap)
	}
	return gaps, nil
}

func scanIntoGap(rows *sql.Rows) (*Gap, error){
	gap := new(Gap)
	err:= rows.Scan(
		&gap.rid,
		&gap.ProjectKey,
		&gap.PrimaryKey,
		&gap.LineKey,
		&gap.RecKey,
		&gap.FormDate,
		&gap.DateVisited,
		&gap.GapStart,
		&gap.GapEnd,
		&gap.Gap,
		&gap.OtherBasal,
		&gap.RecType,
		&gap.Measure,
		&gap.LineLengthAmount,
		&gap.SeqNo,
		&gap.FormType,
		&gap.Direction,
		&gap.GapMin,
		&gap.GapData,
		&gap.PerennialsCanopy,
		&gap.AnnualGrassesCanopy,
		&gap.AnnualForbsCanopy, 
		&gap.OtherCanopy,
		&gap.Notes,
		&gap.NoCanopyGaps,
		&gap.NoBasalGaps,
		&gap.PerennialsBasal,
		&gap.AnnualGrassesBasal,
		&gap.AnnualForbsBasal,
		&gap.DBKey,
		&gap.DateLoadedInDb,
		&gap.source,
	)

	// jsonRet,err := json.Marshal(&gap)
	return gap, err
}