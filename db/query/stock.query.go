package query

// CreateStock is the neo4j query to add a new Stock
var CreateStock string = `
	CREATE (stock:Stock)
		SET stock.ID = $id
		SET stock.Name = $name
		SET stock.Price = $price
		SET stock.Quantity = 10000
	RETURN stock
`

// GetStockByID is the neo4j query to fetch stock by ID
var GetStockByID string = `
	MATCH (stock:Stock {ID: $id})
	RETURN stock
`