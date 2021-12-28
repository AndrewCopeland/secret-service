package secretservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/infamousjoeg/go-keyconfig"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id not provided", http.StatusBadRequest)
		return
	}

	var config interface{}
	err := keyconfig.GetConfig(id, &config)
	if err != nil {
		http.Error(w, fmt.Sprintf("'%s' not found", id), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send body. '%s'", err.Error()), http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id not provided", http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}
	var config interface{}
	err = json.Unmarshal(body, &config)
	if err != nil {
		http.Error(w, fmt.Sprintf("body is not a valid json. %s", err.Error()), http.StatusBadRequest)
	}

	err = keyconfig.SetConfig(id, config)
	if err != nil {
		http.Error(w, fmt.Sprintf("'%s' could not be set. %s", id, err.Error()), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send body. '%s'", err.Error()), http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id not provided", http.StatusBadRequest)
		return
	}
	err := keyconfig.DeleteConfig(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to send body. '%s'", err.Error()), http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
