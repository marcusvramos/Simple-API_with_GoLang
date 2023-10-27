package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

type Aluno struct {
	Pessoa
	RA string `json:"RA"`
}

var alunos []Aluno

func printHelloWorld(w http.ResponseWriter, r *http.Request) {
	var retorno string = "Hello World!"
	json.NewEncoder(w).Encode(retorno)
}

func postAluno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var aluno Aluno
	err := json.NewDecoder(r.Body).Decode(&aluno)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	alunos = append(alunos, aluno)
	json.NewEncoder(w).Encode("Inserido")
}

func getAluno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(alunos)
}

func deleteAluno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	ra := vars["RA"]

	for index, aluno := range alunos {
		if aluno.RA == ra {
			alunos = append(alunos[:index], alunos[index+1:]...)
			json.NewEncoder(w).Encode("Excluído")
			return
		}
	}

	http.Error(w, "Aluno não encontrado", http.StatusNotFound)
}

func putAluno(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	ra := vars["RA"]

	for index, aluno := range alunos {
		if aluno.RA == ra {
			var novoAluno Aluno
			err := json.NewDecoder(r.Body).Decode(&novoAluno)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			alunos[index] = novoAluno
			json.NewEncoder(w).Encode("Atualizado")
			return
		}
	}

	http.Error(w, "Aluno não encontrado", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", printHelloWorld).Methods("GET")
	router.HandleFunc("/aluno", postAluno).Methods("POST")
	router.HandleFunc("/aluno", getAluno).Methods("GET")
	router.HandleFunc("/aluno/{RA}", putAluno).Methods("PUT")
	router.HandleFunc("/aluno/{RA}", deleteAluno).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("to rodante")
	http.ListenAndServe(":8000", nil)
}
