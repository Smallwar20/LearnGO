package routes

import (
    "github.com/gorilla/mux"
    "github.com/Smallwar20/MySQL_BookManagement/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
    router.HandleFunc("/api/bookstore", controllers.CreateBookStore).Methods("POST")
    router.HandleFunc("/api/bookstore", controllers.GetBookStores).Methods("GET")
    router.HandleFunc("/api/bookstore/{id}", controllers.GetBookStore).Methods("GET")
    router.HandleFunc("/api/bookstore/{id}", controllers.UpdateBookStore).Methods("PUT")
    router.HandleFunc("/api/bookstore/{id}", controllers.DeleteBookStore).Methods("DELETE")
}

