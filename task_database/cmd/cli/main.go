package main

import (
	"fmt"
	"log"

	"github.com/Muha113/golang/task_database/pkg/model"
	"github.com/Muha113/golang/task_database/pkg/mongodb"
	"github.com/Muha113/golang/task_database/pkg/repository/memory"
)

const (
	// CommandSave : Save to db
	CommandSave = iota + 1
	// CommandListAll : Get all contacts
	CommandListAll
	// CommandGetByID : Get contact by id
	CommandGetByID
	// CommandGetByPhone : Get contact by phone
	CommandGetByPhone
	// CommandGetByEmail : Get contact by Email
	CommandGetByEmail
	// CommandSearchByName : Get slice of contacts by name
	CommandSearchByName
	// CommandDelete : Delete contact by id
	CommandDelete
)

func main() {
	fmt.Println("Connecting to database...")
	client, err := mongodb.ConnectToDB("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	db := mongodb.GetAccessToDB(client, "mydb")
	collection := mongodb.GetAccessToCollection(db, "contacts")
	repository := memory.NewContactsRepositoryInMemory(collection)

	for {
		fmt.Print(menu)
		var command int
		if _, err := fmt.Scanf("%d", &command); err != nil {
			log.Println(err)
		}

		switch command {
		case CommandSave:
			if err := Save(repository); err != nil {
				log.Println(err)
			}
		case CommandListAll:
			if err := ListAll(repository); err != nil {
				log.Println(err)
			}
		case CommandGetByID:
			if err := GetByID(repository); err != nil {
				log.Println(err)
			}
		case CommandGetByPhone:
			if err := GetByPhone(repository); err != nil {
				log.Println(err)
			}
		case CommandGetByEmail:
			if err := GetByEmail(repository); err != nil {
				log.Println(err)
			}
		case CommandSearchByName:
			if err := SearchByName(repository); err != nil {
				log.Println(err)
			}
		case CommandDelete:
			if err := Delete(repository); err != nil {
				log.Println(err)
			}
		default:
			log.Printf("command not foumd for value %d\n", command)
		}

		printSeparator()
	}
}

// ListAll : get all contacts
func ListAll(rep model.ContactsRepository) error {
	records, err := rep.ListAll()
	if err != nil {
		return fmt.Errorf("error in ListAll: %q", err.Error())
	}

	fmt.Println("ListAll:")
	for _, r := range records {
		fmt.Println(r)
	}

	return nil
}

// GetByID : get by id
func GetByID(rep model.ContactsRepository) error {
	id := readUint("Please enter an 'ID' field and press Enter")

	record, err := rep.GetByID(id)
	if err != nil {
		return fmt.Errorf("error in GetByID: %q", err.Error())
	}

	fmt.Println("GetByID")
	fmt.Println(record)

	return nil
}

// GetByPhone : get by phone
func GetByPhone(rep model.ContactsRepository) error {
	phone := readString("Please enter an 'Phone' field and press Enter")

	record, err := rep.GetByPhone(phone)
	if err != nil {
		return fmt.Errorf("error in GetByPhone: %q", err.Error())
	}

	fmt.Println("GetByPhone:")
	fmt.Println(record)

	return nil
}

// GetByEmail : get by email
func GetByEmail(rep model.ContactsRepository) error {
	email := readString("Please enter an 'Email' field and press Enter")

	record, err := rep.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("error in GetByEmail: %q", err.Error())
	}

	fmt.Println("GetByEmail:")
	fmt.Println(record)

	return nil
}

// SearchByName : search by name
func SearchByName(rep model.ContactsRepository) error {
	email := readString("Please enter prefix for 'Name' field and press Enter")

	records, err := rep.SearchByName(email)
	if err != nil {
		return fmt.Errorf("error in SearchByName: %q", err.Error())
	}

	fmt.Println("SearchByName:")
	for _, r := range records {
		fmt.Println(r)
	}

	return nil
}

// Delete : delete by id
func Delete(rep model.ContactsRepository) error {
	id := readUint("Please enter an 'ID' field and press Enter")

	if err := rep.Delete(id); err != nil {
		return fmt.Errorf("error in GetByID: %q", err.Error())
	}

	fmt.Printf("Delete:\nRecord with ID %d successfylly deleted\n", id)
	return nil
}

// Save : save contact
func Save(rep model.ContactsRepository) error {
	contact := model.Contact{
		FirstName: readString("Please enter an 'FirstName' field and press Enter"),
		LastName:  readString("Please enter an 'LastName' field and press Enter"),

		Phone: readString("Please enter an 'Phone' field and press Enter"),
		Email: readString("Please enter an 'Email' field and press Enter"),
	}

	result, err := rep.Save(contact)
	if err != nil {
		return err
	}

	fmt.Println("Save Contact:")
	fmt.Println(result)

	return nil
}

const menu = `
Please enter operation number:
  * 1 - Save
  * 2 - ListAll
  * 3 - GetByID
  * 4 - GetByPhone
  * 5 - GetByEmail
  * 6 - SearchByName
  * 7 - Delete 
  * Control + C - to exit 
`

func readString(message string) string {
	var r string

	for r == "" {
		fmt.Println(message)
		if _, err := fmt.Scanf("%s", &r); err != nil {
			fmt.Printf("Error in process of reading string from console\n\t%q\n please try again\n", err.Error())
			printSeparator()
		}

	}

	return r
}

func readUint(message string) uint {
	var r uint

	for {
		fmt.Println(message)
		_, err := fmt.Scanf("%d", &r)
		if err == nil {
			break
		}

		fmt.Printf("Error in process of reading string from console\n\t%q\n please try again\n", err.Error())
		printSeparator()
	}

	return r
}

func printSeparator() {
	for i := 0; i < 50; i++ {
		fmt.Print("*")
	}

	fmt.Println()
}
