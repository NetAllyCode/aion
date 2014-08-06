package aion

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/gocql/gocql"
)

// Represents a Cassandra cache
type CQLCache struct {
	Session *gocql.Session
}

// CQLCache implements the SeriesStore interface
func (self *CQLCache) Insert(series uuid.UUID, entry Entry) error {
	seriesUUID, err := gocql.UUIDFromBytes(series)
	if err != nil {
		return err
	}
	err = self.Session.Query("INSERT INTO cache (series, time, value) VALUES (?, ?, ?)", seriesUUID, entry.Timestamp, entry.Attributes["raw"]).Exec()
	if err != nil {
		return err
	}
	return nil
}
