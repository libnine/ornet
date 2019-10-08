package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	mdb "../mdb"
	models "../models"
	"go.mongodb.org/mongo-driver/mongo"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/v/dist/index.html")
}

func CasesHandler(c *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := mdb.FindAllCases(c)
		json.NewEncoder(w).Encode(data)
	}

	return http.HandlerFunc(fn)
}

func CasesHandlerPost(c *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var _case models.Case
		_ = json.NewDecoder(r.Body).Decode(&_case)
		mdb.InsertCase(c, &_case)
	}

	return http.HandlerFunc(fn)
}

func CasesByUser(c *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var _user models.User
		_ = json.NewDecoder(r.Body).Decode(&_user)
		data := mdb.FindCasesByUser(c, fmt.Sprintf("%s %s", _user.FirstName, _user.LastName))
		json.NewEncoder(w).Encode(data)
	}

	return http.HandlerFunc(fn)
}

func EqHandler(c *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := mdb.FindAllEq(c)
		json.NewEncoder(w).Encode(data)
	}

	return http.HandlerFunc(fn)
}

func HistoricHandler(c *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := mdb.FindAllCasesHistorical(c)
		json.NewEncoder(w).Encode(data)
	}

	return http.HandlerFunc(fn)
}

func UsersHandler(c *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := mdb.FindAllUsers(c)
		json.NewEncoder(w).Encode(data)
	}

	return http.HandlerFunc(fn)
}

func UsersHandlerPost(c *mongo.Client) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var _user models.User
		_ = json.NewDecoder(r.Body).Decode(&_user)
		mdb.InsertUser(c, &_user)
	}

	return http.HandlerFunc(fn)
}
