package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)


type Contact struct {
	ID int
	Name  string
	PhoneNumber int
	Email string
}

var contacts []Contact
var scanner = bufio.NewScanner(os.Stdin)
var choice string

func main() {
	fmt.Println("In memory Contact Management System")
	dataSeeding()
	for {
		fmt.Printf("\nEnter your choice\n\n")
		fmt.Printf("1. Show all contacts\n2. Get contact by id\n3. Store contact\n4. Update contact\n5. Delete contact\n6. Exit\n")
		scanner.Scan()
		choice = scanner.Text()

		switch(choice) {
		case "1" :
			GetAllContacts()
		case "2" :
			GetContactByID()
		case "3" :
			StoreContact()
		case "4" :
			UpdateContact()
		case "5" :
			DeleteContact()
		case "6" :
			os.Exit(0)
		}
	}
}

func NewContact(name string, phoneNumber int, email string) *Contact {
	id := len(contacts) + 1
	return &Contact{ID:id,Name:name,PhoneNumber:phoneNumber,Email:email}
}

func StoreContact() {
	fmt.Println("")
	fmt.Println("Enter contact Name :")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Enter contact Phonenumber :")
	scanner.Scan()
	num := scanner.Text()
	phoneNumber, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Please enter integer number.")
		return
	}

	fmt.Println("Enter contact Email :")
	scanner.Scan()
	email := scanner.Text()

	contact := NewContact(name,phoneNumber,email)
	storeContact(contact)
	fmt.Printf("Contact saved.\n\n")
}

func GetAllContacts() {
	fmt.Printf("\nYour contact list :\n\n")
	for _,contact := range contacts {
		fmt.Printf("ID : %d, Name : %s , Phone Number : %d , Email : %s\n",contact.ID,contact.Name,contact.PhoneNumber,contact.Email)
	}
	fmt.Println("")
}

func GetContactByID() {
	fmt.Println("")
	fmt.Println("Enter contact id :")
	scanner.Scan()
	num := scanner.Text()
	id, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Please enter integer number.")
		return
	}
	contact := showContactDetailsByID(id)
	fmt.Printf("ID : %d, Name : %s , Phone Number : %d , Email : %s\n",contact.ID,contact.Name,contact.PhoneNumber,contact.Email)
}

func UpdateContact() {
	fmt.Println("")
	fmt.Println("Enter id of contact that you want to update :")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Please enter integer number.")
		return
	}

	contact := showContactDetailsByID(id)
	if contact == nil {
		fmt.Printf("Contact not found with id %d\n\n",id)
		return
	}

	fmt.Println("Enter contact Name :")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Enter contact Phonenumber :")
	scanner.Scan()
	num := scanner.Text()
	phoneNumber, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Please enter integer number.")
		return
	}

	fmt.Println("Enter contact Email :")
	scanner.Scan()
	email := scanner.Text()

	if isOk := updateDetailsByID(id,name,phoneNumber,email); !isOk {
		fmt.Println("Failed to update the contact details.")
		return
	}

	fmt.Println("Contact details are updated.")
}

func DeleteContact() {
	fmt.Println("")
	fmt.Println("Enter contact id you want to delete :")
	scanner.Scan()
	num := scanner.Text()
	id, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Please enter integer number.")
		return
	}

	if isOk := deleteContactDetailsByID(id); !isOk {
		fmt.Printf("Contact not found with id %d\n\n",id)
		return
	}
	fmt.Printf("Contact deleted with id %d\n\n",id)
}

func storeContact(c *Contact) {
	contacts = append(contacts,*c)
}

func showContactDetailsByID(id int) *Contact {
	for _,contact := range contacts {
		if contact.ID == id {
			return &contact
		}
	}
	return nil
}

func deleteContactDetailsByID(id int) bool {
	for i,contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:i],contacts[i+1:]...)
			return true
		}
	}
	return false
}

func updateDetailsByID(id int,name string, phoneNumber int, email string) bool {
	for i,contact := range contacts {
		if contact.ID == id {
			contacts[i].Name = name
			contacts[i].PhoneNumber = phoneNumber
			contacts[i].Email = email
			return true
		}
	}
	return false
}

func dataSeeding() {
	c1 := Contact{
		ID : 1,
		Name : "Suresh",
		PhoneNumber : 9874587960,
		Email : "suresh@gmail.com",
	}

	c2 := Contact{
		ID : 2,
		Name : "Raj",
		PhoneNumber : 7878787878,
		Email : "raj@gmail.com",
	}

	contacts = append(contacts,c1,c2)
}