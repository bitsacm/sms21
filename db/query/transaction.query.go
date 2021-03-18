package query

var CreateTransaction = `
	MATCH (u:User {Username: $username}), (s:Stock {ID: $stockID})
	CREATE (u)-[t:TRANSACTION]->(s)
	SET t.Timestamp = localdatetime(), t.Quantity = $quantity, t.Price = s.Price
	MERGE (u)-[r:OWNS]->(s)
	ON CREATE SET r.Quantity = t.Quantity
	ON MATCH SET r.Quantity = r.Quantity + t.Quantity 
	RETURN t
`
