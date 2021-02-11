package query

var BuyStock string = `
	MATCH (user: User {ID: $userID}), (stock: Stock {ID: $stockID})
	CREATE (user)-[transaction: Transaction {
		Quantity: $quantity,
		Price: $price,
		Type: "Buy"
	}]->(stock)
	SET user.Balance = $newUserBalance
	SET stock.Quantity = $newStockQuantity
	RETURN transaction
`