package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
	"fmt"
)

type Pessoa struct {
    ID        string   `json:"id,omitempty"`
    Nome string   `json:"nome,omitempty"`
    Sobrenome  string   `json:"sobrenome,omitempty"`
    Endereco   *Endereco `json:"endereco,omitempty"`
}
type Endereco struct {
    Cidade  string `json:"cidade,omitempty"`
    Estado string `json:"estado,omitempty"`
}

var people []Pessoa

// GetPeople mostra todos os contatos da variável people
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// GetPerson mostra apenas um contato
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Pessoa{})
}

// CreatePerson cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Pessoa
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// DeletePerson deleta um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

// função principal para executar a api
func main() {
    router := mux.NewRouter()
    people = append(people, Pessoa{ID: "1", Nome: "Renata", Sobrenome: "Garcia", Endereco: &Endereco{Cidade: "Campinas", Estado: "SP"}})
    people = append(people, Pessoa{ID: "2", Nome: "Julio", Sobrenome: "Teste", Endereco: &Endereco{Cidade: "Valinhos", Estado: "SP"}})
    router.HandleFunc("/contato", GetPeople).Methods("GET")
    router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
	fmt.Println("Server running in port: 8000")

    log.Fatal(http.ListenAndServe(":8000", router))
}