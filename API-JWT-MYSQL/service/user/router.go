package user

import (
	"net/http"

	"github.com/gorilla/mux"
	types "github.com/siddheshRajendraNimbalkar/API-JWT-MYSQL/TYPES"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) LoginRoutes(r *mux.Router) {
	r.HandleFunc("/login", h.HandelLogin).Methods("POST")
	r.HandleFunc("/register", h.HandelRegister).Methods("POST")
}

func (h *Handler) HandelLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandelRegister(w http.ResponseWriter, r *http.Request) {

}
