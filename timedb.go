package timedb

import (
	"code.google.com/p/go-uuid/uuid"
	"time"
)

// One entry in a time series
type Entry struct {
	Timestamp  time.Time
	Attributes map[string]float64
}

// Interface of filtering data in to a level
type Filter interface {
	Insert(series uuid.UUID, entry Entry) error
	SetHandler(handler func(uuid.UUID, Entry) error)
}

// Interface for storing time series data
type SeriesStore interface {
	Insert(series uuid.UUID, entry Entry) error
	Query(series uuid.UUID, start time.Time, end time.Time, attributes []string, entries chan Entry, errors chan error)
}

// A level represents one granularity of data storage in timedb
type Level struct {
	Filter Filter
	Store  SeriesStore
}

// Root of the TimeDB API
type TimeDB struct {
	Levels []Level
}

func (self *TimeDB) createHandlers() {
	for i := 0; i < len(self.Levels)-1; i++ {
		thisLevel := self.Levels[i]
		nextLevel := self.Levels[i+1]
		thisLevel.Filter.SetHandler(func(series uuid.UUID, entry Entry) error {
			err := thisLevel.Store.Insert(series, entry)
			if err != nil {
				return err
			}
			err = nextLevel.Filter.Insert(series, entry)
			if err != nil {
				return err
			}
			return nil
		})
	}
}

// Convenience method for inserting one data point into the first level
// (possibly triggering rollups
func (self *TimeDB) Put(series uuid.UUID, entry Entry) error {
	return self.Levels[0].Filter.Insert(series, entry)
}
