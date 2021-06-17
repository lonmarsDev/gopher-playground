package main

import (
	"github.com/go-chi/chi"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	r.Post("/login", loginHandler)
}


type Login struct {
	username string `json:"username"`
	password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r http.Request){
	var login Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if login.username != "admin"{
		JSONResponse(w, http.StatusBadRequest, "Invalid username")
		return
	}
	if login.password != "admin" {
		JSONResponse(w, http.StatusBadRequest, "Invalid passowrd")
		return
	}
	JSONResponse(w, http.StatusOK, "login success")
	return
}


func JSONResponse((w http.ResponseWriter, code int, message string){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if message != "" {
		fmt.Fprint(w,message)
	}
	return
}