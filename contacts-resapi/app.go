package main

import (
	"os"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"

	. "github.com/user/app/config"
	. "github.com/user/app/dao"
	. "github.com/user/app/models"
)

var config = Config{}
var dao = ContactsDAO{}

// GET list of contacts
func AllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusOK, contacts)
}

// GET a contact by its ID
func FindContactEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	contact, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Contact ID")
		return
	}
	retResponse(w, http.StatusOK, contact)
}

// POST a new contact
func CreateContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	contact.ID = bson.NewObjectId()
	if err := dao.Insert(contact); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusCreated, contact)
}

// PUT update an existing contact
func UpdateContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(contact); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing contact
func DeleteContact(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(contact); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	retResponse(w, http.StatusOK, map[string]string{"result": "success"})
}

func StatusContact(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	retResponse(w, http.StatusOK, map[string]string{"server":name,"result": "success"})
} 

func respondWithError(w http.ResponseWriter, code int, msg string) {
	retResponse(w, code, map[string]string{"error": msg})
}

func retResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/app-back-status", StatusContact)
	r.HandleFunc("/contacts", AllContacts).Methods("GET")
	r.HandleFunc("/contacts", CreateContact).Methods("POST")
	r.HandleFunc("/contacts", UpdateContact).Methods("PUT")
	r.HandleFunc("/contacts", DeleteContact).Methods("DELETE")
	r.HandleFunc("/contacts/{id}", FindContactEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
