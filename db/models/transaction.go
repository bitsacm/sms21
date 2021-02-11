package models

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/dush-t/sms21/db"
	"github.com/dush-t/sms21/db/query"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Transactions struct {
	conn *db.Conn
	DataType reflect.Type
}

type Transaction struct {
	Price    float64 `neoKey:"Price" json:"price"`
	Quantity uint    `neoKey:"Quantity" json:"quantity"`
	Type     string	 `neoKey:"Type" json:"type"`
}

func (tr *Transactions) BuyStock(user User, stock Stock, quantity int64) (Transaction, error) {
	driver := *(tr.conn.Driver)
	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return Transaction{}, err
	}
	defer session.Close()

	log.Println("Quantity = ", quantity, " ", stock.Quantity)
	if(quantity <= 0) {
		return Transaction{}, errors.New("Invalid quantity")
	}

	if(quantity > stock.Quantity) {
		return Transaction{}, errors.New(fmt.Sprint("Invalid Quantity. Maximum Quantity avaialable %d", stock.Quantity))
	}

	if(user.Balance < (float64(quantity) * float64(stock.Price))) {
		return Transaction{}, errors.New("Insufficient Balance")
	}

	newUserBalance := user.Balance - (float64(quantity) * float64(stock.Price))
	newStockQuantity := stock.Quantity - quantity

	buyTransaction, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		buyResult, err := transaction.Run(
			query.BuyStock,
			query.Context {
				"userID":   user.ID,
				"stockID":  stock.ID,
				"quantity": quantity,
				"price":    stock.Price,
				"newStockQuantity": newStockQuantity,
				"newUserBalance": newUserBalance,
			},
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		if buyResult.Next() {
			return buyResult.Record().GetByIndex(0), nil
		}

		return nil, buyResult.Err()
	})

	if err != nil {
		log.Println(err)
		return Transaction{}, err
	}

	log.Println("Created new buy transaction ", buyTransaction)

	return  Transaction{}, nil
}