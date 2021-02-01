package query

// AddUser is the neo4j query to add a new user
var AddUser string = `
	CREATE (user:User)
		SET user.ID = $id
		SET user.Password = $password
		SET user.Username = $username
		SET user.Name = $name
		SET user.RegToken = $regToken
	RETURN user
`

// GetUserByID is the neo4j query to fetch user by ID
var GetUserByID string = `
	MATCH (user:User {ID: $id})
	RETURN user
`

var GetUserByUsername string = `
	MATCH (user:User {Username: $username})
	RETURN user
`