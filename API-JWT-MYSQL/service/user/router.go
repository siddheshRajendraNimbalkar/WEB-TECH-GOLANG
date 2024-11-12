package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) LoginRoutes(r *mux.Router) {
	r.HandleFunc("/login", h.HandelLogin).Methods("POST")
	r.HandleFunc("/register", h.HandelRegister).Methods("POST")
}

func (h *Handler) HandelLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandelRegister(w http.ResponseWriter, r *http.Request) {

}
