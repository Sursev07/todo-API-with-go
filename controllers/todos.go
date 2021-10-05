package controllers


import (
	 "fmt"
	"todo-API-with-go/models"
	"net/http"
	"encoding/json"

	// "gorm.io/gorm"
	// "gorm.io/driver/mysql"
	"github.com/gorilla/mux"
	"todo-API-with-go/database"
)

// var DB *gorm.DB
var err error


func GetTodos(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var todos []models.Todo
	DB := database.GetDB()
	DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
	
}

func WelcomeAPI(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	result := "Hello"
	json.NewEncoder(w).Encode(result)
	
}


func GetTodo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo models.Todo
	DB := database.GetDB()
	DB.First(&todo, params["id"])
	json.NewEncoder(w).Encode(todo)
	
}

func CreateTodo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var todo models.Todo
	fmt.Print(r.Body, ":>>>>>")
	json.NewDecoder(r.Body).Decode(&todo)
	DB := database.GetDB()
	DB.Create(&todo)
	json.NewEncoder(w).Encode(todo)
	res := models.Result{Code: 200, Data: todo, Message: "Success create todo"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	
}

func UpdateTodo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo models.Todo
	DB := database.GetDB()
	DB.First(&todo, params["id"])
	json.NewDecoder(r.Body).Decode(&todo)
	DB.Save(&todo)
	json.NewEncoder(w).Encode(todo)
	
}

func DeleteTodo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var todo models.Todo
	DB := database.GetDB()
	DB.Delete(&todo, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
	
}