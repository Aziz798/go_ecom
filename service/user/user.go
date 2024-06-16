package user

import (
	"net/http"

	"github.com/Aziz798/ecom/types"
	"github.com/Aziz798/ecom/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	err := utils.ParseJson(r, payload)
	if err != nil {

	}

}
