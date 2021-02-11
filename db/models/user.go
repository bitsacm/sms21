package models

import (
	"log"
	"reflect"
	"time"

	"github.com/dush-t/sms21/db"
	"github.com/dush-t/sms21/db/query"
	"github.com/dush-t/sms21/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Users represents the User model in general
type Users struct {
	conn     *db.Conn
	DataType reflect.Type
}

// User represents an instance of the User model
type User struct {
	ID       string `neoKey:"ID" json:"id"`
	Password string `neoKey:"Password" json:"password"`
	Username string `neoKey:"Username" json:"username"`
	Name     string `neoKey:"Name" json:"name"`
	RegToken string `neoKey:"RegToken" json:"regToken"`
	Balance  float64 `neoKey:"Balance" json:"balance"`
}

// Claims is the information that we'll encode in a user's JWT
type Claims struct {
	ID string
	jwt.StandardClaims
}

// SerializeFromNode will parse a neo4j User node based on the
// User struct's "neoKey" values and return a corresponding User
// instance struct
func (us *Users) SerializeFromNode(n neo4j.Node) User {
	user := db.Serialize(us.DataType, n)
	return user.(User)
}

// Add can be used to add a new user to the database
func (us *Users) Add(u User) error {
	driver := *(us.conn.Driver)
	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return err
	}
	defer session.Close()

	passwordHash := util.HashPassword(u.Password)
	u.Password = passwordHash
	anonID := util.GenerateID(u.Username)
	u.ID = anonID

	user, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			query.AddUser,
			query.Context{
				"id":       u.ID,
				"password": passwordHash,
				"username": u.Username,
				"name":     u.Name,
				"regToken": u.RegToken,
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
		return err
	}

	savedUser := us.SerializeFromNode(user.(neo4j.Node))
	log.Println("Created new user with username", savedUser.Username)

	return nil
}

// GetByID fetches a User with the given ID.
func (us *Users) GetByID(ID string) (User, error) {
	driver := *(us.conn.Driver)
	session, err := driver.Session(neo4j.AccessModeRead)
	if err != nil {
		return User{}, err
	}
	defer session.Close()

	userNode, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			query.GetUserByID,
			query.Context{
				"id": ID,
			},
		)
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}

		return nil, result.Err()
	})

	if err != nil {
		log.Println("Error getting user:", err)
		return User{}, err
	}

	user := us.SerializeFromNode(userNode.(neo4j.Node))
	return user, nil
}

// GenerateJWT will generate a JWT that expires in 2 months
// for the User
func (u *User) GenerateJWT() (string, error) {
	expirationTime := time.Now().Add(60 * 24 * time.Hour).Unix()
	claims := &Claims{
		ID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	jwtKey := []byte("lolmao12345")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetUserByUsername fetches a User with the given username.
func (us *Users) GetUserByUsername(username string) (User, error) {
	driver := *(us.conn.Driver)
	session, err := driver.Session(neo4j.AccessModeRead)
	if err != nil {
		return User{}, err
	}
	defer session.Close()

	userNode, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error){
		result, err := transaction.Run(
			query.GetUserByUsername,
			query.Context{
				"username": username,
			},
		)

		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().GetByIndex(0), nil
		}
		return nil, result.Err()
	})

	if err != nil {
		log.Println("Error getting user by username:", err)
		return User{}, err
	}

	user := us.SerializeFromNode(userNode.(neo4j.Node))
	return user, nil
}