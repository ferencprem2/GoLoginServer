package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Datas struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SendResponse(w http.ResponseWriter, i any, wrapper ...string) {
	data, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(wrapper) > 0 {
		data = append([]byte("{\""+wrapper[0]+"\":"), data...)
		data = append(data, []byte("}")...)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func DecodeRequest(w http.ResponseWriter, r *http.Request, i any) bool {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("LoginHandler")
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "login.html")
	case http.MethodPost:
	}
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware", r.URL)
		next.ServeHTTP(w, r)
	})

}

func main() {

	mux := mux.NewRouter()
	mux.HandleFunc("/", HomeHandler).Methods("GET")
	mux.HandleFunc("/login", LoginHandler).Methods("GET", "POST")
	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
