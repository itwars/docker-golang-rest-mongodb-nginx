package dao

import (
	"log"

	. "github.com/user/app/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ContactsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "contacts"
)

// Establish a connection to database
func (m *ContactsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of contacts
func (m *ContactsDAO) FindAll() ([]Contact, error) {
	var contacts []Contact
	err := db.C(COLLECTION).Find(bson.M{}).All(&contacts)
	return contacts, err
}

// Find a contact by its id
func (m *ContactsDAO) FindById(id string) (Contact, error) {
	var contact Contact
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&contact)
	return contact, err
}

// Insert a contact into database
func (m *ContactsDAO) Insert(contact Contact) error {
	err := db.C(COLLECTION).Insert(&contact)
	return err
}

// Delete an existing contact
func (m *ContactsDAO) Delete(contact Contact) error {
	err := db.C(COLLECTION).Remove(&contact)
	return err
}

// Update an existing contact
func (m *ContactsDAO) Update(contact Contact) error {
	err := db.C(COLLECTION).UpdateId(contact.ID, &contact)
	return err
}
