package models

import (
	"reflect"

	"github.com/dush-t/sms21/db"
)

// Models will hold references to all the db models,
// to be passed to API handlers for DB access
type Models struct {
	Users        *Users
	Stocks       *Stocks
	Transactions *Transactions
}

// Init returns the Models which can be used in the rest
// of the app for database access
func Init(conn *db.Conn) *Models {
	models := Models{}

	// USERS
	users := Users{
		conn:     conn,
		DataType: reflect.TypeOf(User{}),
	}
	models.Users = &users

	// Stocks
	stocks := Stocks{
		conn:     conn,
		DataType: reflect.TypeOf(Stock{}),
	}
	models.Stocks = &stocks

	// Transactions
	transactions := Transactions{
		conn:     conn,
		DataType: reflect.TypeOf(Transaction{}),
	}
	models.Transactions = &transactions

	return &models
}
