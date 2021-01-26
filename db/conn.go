package db

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Conn holds information about our database connection
type Conn struct {
	URI    string
	Driver *neo4j.Driver
}

// QueryContext is the data type meant to be passed in
// neo4j transactions for dynamic queries
type QueryContext map[string]interface{}

// NewConn will return a new database connection
func NewConn(uri, username, password string) (*Conn, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}

	conn := &Conn{
		URI:    uri,
		Driver: &driver,
	}

	return conn, nil
}
