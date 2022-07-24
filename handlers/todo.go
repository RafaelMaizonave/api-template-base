package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RafaelMaizonave/api-template-base/entities"
	"github.com/RafaelMaizonave/api-template-base/models"
	"github.com/go-chi/chi/v5"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var todo entities.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso, ID : %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func Update(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo entities.Todo

	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(int64(id), todo)

	if err != nil {
		fmt.Println("erro ao atualizar registro ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		fmt.Printf("Error: Foram atualizados %d registros ", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo atualizado com sucesso, ID : %d", id),
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))

	if err != nil {
		fmt.Println("erro ao remover registro ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		fmt.Printf("Error: Foram removidos %d registros ", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo removido com sucesso, ID : %d", id),
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()

	if err != nil {
		fmt.Println("erro ao buscar registros", err)
	}

	w.Header().Add("Content-Type", "application/json")

	if len(todos) == 0 {
		json.NewEncoder(w).Encode("Nenhum registro encontrado")
		return
	}

	json.NewEncoder(w).Encode(todos)

}

func GetById(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id))

	if err != nil {
		fmt.Println("erro ao buscar registro pelo id ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	if todo.ID == 0 {
		json.NewEncoder(w).Encode("Nenhum registro encontrado, verifique o parametro enviado")
		return
	}

	json.NewEncoder(w).Encode(todo)

}
