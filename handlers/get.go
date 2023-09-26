package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"learning-go/models"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}


	todo, err := models.Get(int64(id))
	if err != nil{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
