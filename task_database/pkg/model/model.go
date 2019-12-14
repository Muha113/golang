package model

import "fmt"

// ContactsRepository : interface representing repository methods
type ContactsRepository interface {
	Save(Contact) (Contact, error)
	ListAll() ([]Contact, error)
	GetByID(uint) (Contact, error)
	GetByPhone(string) (Contact, error)
	GetByEmail(string) (Contact, error)
	SearchByName(string) ([]Contact, error)
	Delete(uint) error
}

// Contact : struct representing contact from repository
type Contact struct {
	ID uint `bson:"id"`

	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`

	Phone string `bson:"phone"`
	Email string `bson:"email"`
}

func (c Contact) String() string {
	return fmt.Sprintf("Contact:\n\tID - %d\n\tFirst Name - %q\n\tLast Name - %q\n\tPhone - %q\n\tEmail - %q\n", c.ID, c.FirstName, c.LastName, c.Phone, c.Email)
}
