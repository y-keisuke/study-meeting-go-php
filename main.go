package main

import (
	"encoding/json"
	"free/user"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", resisterUser).Methods(http.MethodPost)
	r.HandleFunc("/users", fetchUser).Methods(http.MethodGet)

	http.ListenAndServe(":8878", r)
}

// curl -X POST -H "Content-Type: application/json" -d '{"first_name":"Suzuki", "last_name":"Ichiro", "address":"Tokyo", "email":"ichiro@example.com"}' localhost:8878/users
func resisterUser(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var u user.User
	json.Unmarshal(b, &u)

	u.Resister()

	msg := `{"msg":"success"}`
	json, _ := json.Marshal(msg)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// curl localhost:8878/users
func fetchUser(w http.ResponseWriter, r *http.Request) {
	users := user.Fetch()
	json, _ := json.Marshal(users)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}