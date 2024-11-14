package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siddheshRajendraNimbalkar/API-JWT-MYSQL/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	// init routes witn gorilla mux
	router := mux.NewRouter()
	subRouter := router.PathPrefix("v1/api").Subrouter()

	// handling route
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.LoginRoutes(subRouter)

	log.Println("Server is Listing on Port: ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
