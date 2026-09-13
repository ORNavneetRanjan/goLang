package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextId int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exits := contactIndexByName[name]; exits {
		fmt.Println("Contact with name: ", name, " already exits")
		return
	}
	newContact := Contact {
		ID: nextId,
		Name: name,
		Email: email,
		Phone: phone,
	}

	contactList = append(contactList, newContact)
	contactIndexByName[name] = newContact.ID
	nextId++;
	fmt.Printf("Contact with name: %v is added successfully \n", name)
}

func getContactByName(name string) *Contact {
	if index, exits := contactIndexByName[name]; exits {
		return &contactList[index-1];
	}
	return nil;
}

func displayContacts(){
	fmt.Println("-----------Contact List---------")
	if len(contactList) == 0 {
		fmt.Println("No Contacts found.")
		return
	}

	for _, contact := range contactList {
		fmt.Printf("Contact Name: %v\n", contact.Name)
		fmt.Printf("Contact email: %v\n", contact.Email)
		fmt.Println("Contact Phone no. : ", contact.Phone)
	}

	fmt.Println("--------END----------")

}


func main() {
	displayContacts()
	addContact("Navneet", "email@gmail.com", "9900230243")
	addContact("Ranjan", "email@gmail.com", "9900230243")
	addContact("Kumar", "email@gmail.com", "9900230243")
	addContact("Navneet", "email@gmail.com", "9900230243")

	displayContacts()

	navneet := getContactByName("Navneet")
	if(navneet == nil) {
		fmt.Println("No contact with that name exits")
	}else{
		fmt.Println("Contact with that name is " , *navneet)
	}
}