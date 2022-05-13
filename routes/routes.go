package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leanderp/golang_rest_web/handlers"
	"github.com/leanderp/golang_rest_web/middleware"
	"github.com/leanderp/golang_rest_web/server"
)

func BindRoutes(s server.Server, r *mux.Router) {

	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/signup", handlers.SingUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.CheckAuthMiddleware(s))

	// POSTS - Authorized
	api.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)
	api.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)
	api.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods(http.MethodPut)
	api.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods(http.MethodDelete)
	api.HandleFunc("/posts", handlers.GetPostsHandler(s)).Methods(http.MethodGet)

	// WEB SOCKET
	r.HandleFunc("/ws", s.Hub().HandleWebSocket)
}
