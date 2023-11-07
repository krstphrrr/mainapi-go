package main 

import(
	"time"
)

type Gap struct{
	rid int
	ProjectKey string
	PrimaryKey string 
	LineKey string 
	RecKey string 
	FormDate time.Time
	DateVisited time.Time
	GapStart float64
	GapEnd float64
	Gap float64
	OtherBasal bool
	RecType string 
	Measure int
	LineLengthAmount float64
	SeqNo int
	FormType string 
	Direction string 
	GapMin float64
	GapData int
	PerennialsCanopy bool
	AnnualGrassesCanopy bool
	AnnualForbsCanopy bool
	OtherCanopy bool
	Notes string 
	NoCanopyGaps bool
	NoBasalGaps bool
	PerennialsBasal bool
	AnnualGrassesBasal bool
	AnnualForbsBasal bool
	DBKey string 
	DateLoadedInDb time.Time
}