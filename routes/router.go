package routes


import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"todo-API-with-go/controllers"
	"todo-API-with-go/middlewares"
	_ "todo-API-with-go/docs" // docs is generated by Swag CLI, you have to import it.

	httpSwagger "github.com/swaggo/http-swagger"
)

func HandleRequest()  {
	r := mux.NewRouter()


	r.HandleFunc("/api/v1/register", controllers.UserRegister).Methods("POST")
	r.HandleFunc("/api/v1/login", controllers.UserLogin).Methods("POST")
	
	r.Use(middlewares.Authentication)
	r.HandleFunc("/api/v1", (controllers.WelcomeAPI)).Methods("GET")
	r.HandleFunc("/api/v1/todos", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/api/v1/todos/{id}", controllers.GetTodo).Methods("GET")
	r.HandleFunc("/api/v1/todos", controllers.CreateTodo).Methods("POST")
	r.Use(middlewares.UserAuthorization)
	r.HandleFunc("/api/v1/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/api/v1/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	fmt.Println("Server is starting at 8080")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}