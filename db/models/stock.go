package models

import (
	"reflect"

	"github.com/dush-t/sms21/db"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Stocks represents the Stock model in general
type Stocks struct {
	conn *db.Conn
	DataType reflect.Type
}

// Stock represents an instance of the Stock model
type Stock struct {
	ID string `neoKey:"ID" json:"id"`
	Name string `neoKey:"Name" json:"name"`
	Price float64 `neoKey:"Price" json:"price"`
}

// SerializeFromNode will parse a neo4j Stock node based on the
// Stock struct's "neoKey" values and return a corresponding Stock
// instance struct
func (s *Stocks) SerializeFromNode(n neo4j.Node) Stock {
	stock := db.Serialize(s.DataType, n)
	return stock.(Stock)
}
