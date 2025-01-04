package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	Svc service
}

func (h Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body\n"))
		return
	}
	// Call the AddUser function
	message, err := h.Svc.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		w.Write([]byte("Failed to add user\n"))
	}

	// Return a success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h Handler) Test(w http.ResponseWriter, r *http.Request) {
	m := "Okay"
	w.WriteHeader(200)
	w.Write([]byte(m))
}
