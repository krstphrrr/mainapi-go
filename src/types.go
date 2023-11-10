package main

import(
	"time"
	"encoding/json"
	"reflect"
	"database/sql"
)

type NullInt16 sql.NullInt16

func(ni *NullInt16) Scan(value interface{}) error{
	var i sql.NullInt16
	if err := i.Scan(value); err != nil{
		return err
	}
	if reflect.TypeOf(value) == nil {
		*ni = NullInt16{i.Int16, false}
	} else {
		*ni = NullInt16{i.Int16, true}
	}
	return nil
}

type NullInt64 sql.NullInt64

func(ni *NullInt64) Scan(value interface{}) error{
	var i sql.NullInt64
	if err := i.Scan(value); err != nil{
		return err
	}
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}

type NullString sql.NullString

func(ns *NullString) Scan(value interface{}) error{
	var s sql.NullString
	if err := s.Scan(value); err != nil{
		return err
	}
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}
	return nil
}

type NullTime sql.NullTime

func (nt *NullTime) Scan(value interface{}) error {
	var t sql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nt = NullTime{t.Time, false}
	} else {
		*nt = NullTime{t.Time, true}
	}

	return nil
}

type NullFloat64 sql.NullFloat64

func (nf *NullFloat64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nf = NullFloat64{f.Float64, false}
	} else {
		*nf = NullFloat64{f.Float64, true}
	}

	return nil
}
// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}


// MarshalJSON for NullInt16
func (ni *NullInt16) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int16)
}

// UnmarshalJSON for NullInt16
func (ni *NullInt16) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int16)
	ni.Valid = (err == nil)
	return err
}
// MarshalJSON for NullFloat64
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON for NullFloat64
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.Float64)
	nf.Valid = (err == nil)
	return err
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

// UnmarshalJSON for NullTime
func (nt *NullTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	// s = Stripchars(s, "\"")

	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		nt.Valid = false
		return err
	}

	nt.Time = x
	nt.Valid = true
	return nil
}

type Gap struct{
	rid int `json:"rid"`
	ProjectKey NullString `json:"ProjectKey"`
	PrimaryKey NullString `json:"PrimaryKey"`
	LineKey NullString `json:"LineKey"`
	RecKey NullString `json:"RecKey"`
	FormDate NullTime `json:"FormDate"`
	DateVisited NullTime `json:"DateVisited"`
	GapStart NullFloat64 `json:"GapStart"`
	GapEnd NullFloat64 `json:"GapEnd"`
	Gap NullFloat64 `json:"Gap"`
	OtherBasal NullInt16 `json:"OtherBasal"`
	RecType NullString `json:"RecType"`
	Measure NullInt64 `json:"Measure"`
	LineLengthAmount NullFloat64 `json:"LineLengthAmount"`
	SeqNo NullInt64 `json:"SeqNo"`
	FormType NullString `json:"FormType"`
	Direction NullString `json:"Direction"`
	GapMin NullFloat64 `json:"GapMin"`
	GapData NullInt64 `json:"GapData"`
	PerennialsCanopy NullInt16 `json:"PerennialsCanopy"`
	AnnualGrassesCanopy NullInt16 `json:"AnnualGrassesCanopy"`
	AnnualForbsCanopy NullInt16 `json:"AnnualForbsCanopy"`
	OtherCanopy NullInt16 `json:"OtherCanopy"`
	Notes NullString `json:"Notes"`
	NoCanopyGaps NullInt16 `json:"NoCanopyGaps"`
	NoBasalGaps NullInt16 `json:"NoBasalGaps"`
	PerennialsBasal NullInt16 `json:"PerennialsBasal"`
	AnnualGrassesBasal NullInt16 `json:"AnnualGrassesBasal"`
	AnnualForbsBasal NullInt16 `json:"AnnualForbsBasal"`
	DBKey NullString `json:"DBKey"` 
	DateLoadedInDb NullTime `json:"DateLoadedInDb"`
	source NullString `json:"source"`
}