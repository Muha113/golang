package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/Muha113/task_database/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ContactsRepositoryInMemory : represents contacts in memory
type ContactsRepositoryInMemory struct {
	sync.RWMutex
	db   *mongo.Collection
	size uint
}

// NewContactsRepositoryInMemory : creates new ContactsRepositoryInMemory struct
func NewContactsRepositoryInMemory(collection *mongo.Collection) *ContactsRepositoryInMemory {
	ctn, _ := collection.CountDocuments(context.Background(), bson.D{})
	return &ContactsRepositoryInMemory{
		db:   collection,
		size: uint(ctn),
	}
}

// Save : saves contact in database. If this contact exists in db, error will be returned
func (r *ContactsRepositoryInMemory) Save(contact model.Contact) (model.Contact, error) {
	emailFilter := bson.D{{Key: "email", Value: contact.Email}}
	phoneFilter := bson.D{{Key: "phone", Value: contact.Phone}}
	r.RLock()

	var result model.Contact

	err := r.db.FindOne(context.TODO(), emailFilter).Decode(&result)
	if err == nil {
		return model.Contact{}, fmt.Errorf("contact with email %q already exists", contact.Email)
	}
	err = r.db.FindOne(context.TODO(), phoneFilter).Decode(&result)
	if err == nil {
		return model.Contact{}, fmt.Errorf("contact with phone %q already exists", contact.Phone)
	}

	r.RUnlock()
	r.Lock()
	defer r.Unlock()

	_, err = r.db.InsertOne(context.TODO(), contact)
	if err != nil {
		return model.Contact{}, err
	}
	r.size++

	return contact, nil
}

// ListAll : return all contacts from db. If there are no contacts, error will be returned
func (r *ContactsRepositoryInMemory) ListAll() ([]model.Contact, error) {
	r.RLock()
	defer r.RUnlock()

	var result []model.Contact

	cursor, err := r.db.Find(context.TODO(), bson.D{})
	if err != nil {
		return []model.Contact{}, err
	}
	fmt.Println(cursor)
	err = cursor.Decode(&result)
	if err != nil {
		return []model.Contact{}, err
	}

	return result, nil
}

// GetByID : return contact found by id. If there is no such id, error will be returned
func (r *ContactsRepositoryInMemory) GetByID(id uint) (model.Contact, error) {
	r.RLock()
	defer r.RUnlock()

	filter := bson.D{{Key: "id", Value: id}}
	var result model.Contact
	err := r.db.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return model.Contact{}, err
	}

	return result, nil
}

// GetByPhone : return conatct by phone. If there is no such phone, error will be returned
func (r *ContactsRepositoryInMemory) GetByPhone(phone string) (model.Contact, error) {
	r.RLock()
	defer r.RUnlock()

	filter := bson.D{{Key: "phone", Value: phone}}
	var result model.Contact
	err := r.db.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return model.Contact{}, err
	}

	return result, nil
}

// GetByEmail : return contact by email. If there is no such email, error will be returned
func (r *ContactsRepositoryInMemory) GetByEmail(email string) (model.Contact, error) {
	r.RLock()
	defer r.RUnlock()

	filter := bson.D{{Key: "email", Value: email}}
	var result model.Contact
	err := r.db.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return model.Contact{}, err
	}

	return result, nil
}

// SearchByName : return slice of contacts by name. If there are no such name, error will be returned
func (r *ContactsRepositoryInMemory) SearchByName(name string) ([]model.Contact, error) {
	r.RLock()
	defer r.RUnlock()

	result := make([]model.Contact, 0, r.size)
	filter := bson.D{{Key: "name", Value: name}}

	cursor, err := r.db.Find(context.TODO(), filter)
	if err != nil {
		return []model.Contact{}, err
	}
	err = cursor.Decode(&result)
	if err != nil {
		return []model.Contact{}, err
	}

	return result, nil
}

// Delete : delete contact by id. If there no such id, error will be returned
func (r *ContactsRepositoryInMemory) Delete(id uint) error {
	r.Lock()
	defer r.Unlock()

	filter := bson.D{{Key: "id", Value: id}}

	_, err := r.db.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
