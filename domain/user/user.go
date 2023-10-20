package user 

import "errors"
       "github.com/google/uuid"

var (
	ErrInvalidID=errors.New("invalid id providec")
)
type User struct{
	ID        ID
	username  username 
	email     email
}

func New (email Email, username Username) user {
	 return User{
		ID: ID(uuid.Newstring()),
		Email: email,
		username: username ,
	 }
} 

type ID string 

func NewID(id string ) (ID, error) {
     //domain logic comes in  
	_, err := uuid.Parse(i)
    if err !-= nil {
		return "", errors.Join(err, ErrInvalidID)
}
	return ID(i), nil
}

func (i ID) string() string {
	return string (i)
}

type Email string 

func NewEmail(e string ) (Email,error) {
	//domain logic comes in 
	// validate it is an valid email
	return Email(e), nil
}

func (e Email) string() string {
	return string(e)
}

type username string 

func NewUsername(u string ) (username, error) {
	// domain logic comes in 
	// validate it is an valid email
	return username (u), nil 
}
func (u Username ) string () string {
	return string (u)
}