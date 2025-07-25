package user

import (
	"encoding/json"
	"github.com/SavinDevelop/techcrm-go/internal/model"
	repository "github.com/SavinDevelop/techcrm-go/internal/repository/user"
	service "github.com/SavinDevelop/techcrm-go/internal/service/user"
	"github.com/SavinDevelop/techcrm-go/pkg/db"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	service  *service.Service
	validate *validator.Validate
}

func NewHandler(db *db.Postgres) *Handler {
	repo := repository.NewRepository(db)
	s := service.NewService(repo)

	return &Handler{
		service:  s,
		validate: validator.New(),
	}
}
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input model.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.validate.Struct(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := model.User{
		Email:    input.Email,
		Password: input.Password,
		IsActive: input.IsActive,
	}
	if err := h.service.Create(r.Context(), user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
