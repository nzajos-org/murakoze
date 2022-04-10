// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Member struct {
	Id        string `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	TelNumber string `json:"TelNumber"`
}

var Members []Member

func returnAllMembers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllMembers")
	json.NewEncoder(w).Encode(Members)
}

func createNewMember(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var member Member
	json.Unmarshal(reqBody, &member)
	Members = append(Members, member)

	json.NewEncoder(w).Encode(member)
}

func returnSingleMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, member := range Members {
		if member.Id == key {
			json.NewEncoder(w).Encode(member)
		}
	}
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, member := range Members {
		if member.Id == id {
			Members = append(Members[:index], Members[index+1:]...)
		}
	}

}

//test webhookkkkkklllhtkljkjlkj
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/members", returnAllMembers)
	myRouter.HandleFunc("/member", createNewMember).Methods("POST")
	myRouter.HandleFunc("/member/{id}", deleteMember).Methods("DELETE")
	myRouter.HandleFunc("/member/{id}", returnSingleMember)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to address!")
	fmt.Println("Endpoint Hit: Adress")
}

func main() {
	Members = []Member{
		{Id: "1", FirstName: "Jane", LastName: "Karangwa", TelNumber: "080988884443"},
		{Id: "2", FirstName: "James", LastName: "Kalisa", TelNumber: "080923433232"},
	}
	handleRequests()
}
