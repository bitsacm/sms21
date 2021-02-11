package models

import (
	"log"
	"reflect"

	"github.com/dush-t/sms21/db"
	"github.com/dush-t/sms21/db/query"
	"github.com/dush-t/sms21/util"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Stocks represents the Stock model in general
type Stocks struct {
	conn     *db.Conn
	DataType reflect.Type
}

// Stock represents an instance of the Stock model
type Stock struct {
	ID       string  `neoKey:"ID" json:"id"`
	Name     string  `neoKey:"Name" json:"name"`
	Price    float64 `neoKey:"Price" json:"price"`
	Quantity int64   `neoKey:"Quantity" json:"quantity"`
}

// SerializeFromNode will parse a neo4j Stock node based on the
// Stock struct's "neoKey" values and return a corresponding Stock
// instance struct
func (st *Stocks) SerializeFromNode(n neo4j.Node) Stock {
	stock := db.Serialize(st.DataType, n)
	return stock.(Stock)
}

// Add creates a stock in the database
func (st *Stocks) Add(s Stock) (string, error) {
	driver := *(st.conn.Driver)
	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return "", err
	}
	defer session.Close()

	s.ID = util.GenerateID(s.Name)

	stock, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			query.CreateStock,
			query.Context{
				"id":    s.ID,
				"name":  s.Name,
				"price": s.Price,
			},
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}
		return nil, result.Err()
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	savedStock := st.SerializeFromNode(stock.(neo4j.Node))
	log.Println("Created new stock with name", savedStock.Name)
	return savedStock.ID, nil
}

// GetByID fetches a stock by the given ID
func (st *Stocks) GetByID(ID string) (Stock, error) {
	driver := *(st.conn.Driver)
	session, err := driver.Session(neo4j.AccessModeRead)
	if err != nil {
		return Stock{}, err
	}
	defer session.Close()

	stockNode, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			query.GetStockByID,
			query.Context{
				"id": ID,
			},
		)
		if err != nil {
			return nil, err
		}

		if result.Next() {
			log.Println("Inside")
			return result.Record().GetByIndex(0), nil
		}

		return nil, result.Err()
	})

	if err != nil {
		log.Println("Error getting stock:", err)
		return Stock{}, err
	}

	stock := st.SerializeFromNode(stockNode.(neo4j.Node))
	return stock, nil
}
