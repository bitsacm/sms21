package models

import (
	"fmt"
	"log"
	"reflect"

	"github.com/dush-t/sms21/db"
	"github.com/dush-t/sms21/db/query"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Transactions struct {
	conn     *db.Conn
	DataType reflect.Type
}

type Transaction struct {
	ID        string
	Username  string
	StockID   string
	Timestamp float64
	Quantity  int
	Price     float64
}

func (tr *Transactions) SerializeFromEdge(e neo4j.Relationship) Transaction {
	transaction := db.SerializeEdge(tr.DataType, e)

	return transaction.(Transaction)
}

func (tr *Transactions) Add(t Transaction) (string, error) {
	driver := *(tr.conn.Driver)
	session, err := driver.Session(neo4j.AccessModeWrite)

	if err != nil {
		return "", err
	}

	transaction, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		fmt.Println(t.Username, t.StockID)
		result, err := tx.Run(
			query.CreateTransaction,
			query.Context{
				"username": t.Username,
				"stockID":  t.StockID,
				"quantity": t.Quantity,
			})
		if err != nil {
			log.Println(err)
			return nil, err
		}

		if result.Next() {
			record, _ := result.Record().Get("t")
			fmt.Println("here", record)
			return record, nil
		}
		return nil, result.Err()
	})

	id := transaction.(neo4j.Relationship).Id

	return fmt.Sprintf("%d", id), nil
}
