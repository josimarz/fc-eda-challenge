package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/josimarz/fc-eda-challenge/internal/usecase"
)

type CreateCustomerHandler struct {
	uc *usecase.CreateCustomerUseCase
}

func NewCreateCustomerHandler(uc *usecase.CreateCustomerUseCase) *CreateCustomerHandler {
	return &CreateCustomerHandler{uc}
}

func (h *CreateCustomerHandler) GetMethod() string {
	return "POST"
}

func (h *CreateCustomerHandler) GetPattern() string {
	return "/customers"
}

func (h *CreateCustomerHandler) GetHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input usecase.CreateCustomerInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		output, err := h.uc.Execute(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(output); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

type FindCustomerHandler struct {
	uc *usecase.FindCustomerUseCase
}

func NewFindCustomerHandler(uc *usecase.FindCustomerUseCase) *FindCustomerHandler {
	return &FindCustomerHandler{uc}
}

func (h *FindCustomerHandler) GetMethod() string {
	return "GET"
}

func (h *FindCustomerHandler) GetPattern() string {
	return "/customers/{id}"
}

func (h *FindCustomerHandler) GetHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := usecase.FindCustomerInput{
			Id: chi.URLParam(r, "id"),
		}
		output, err := h.uc.Execute(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(output); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

type ListCustomersHandler struct {
	uc *usecase.ListCustomersUseCase
}

func NewListCustomersHandler(uc *usecase.ListCustomersUseCase) *ListCustomersHandler {
	return &ListCustomersHandler{uc}
}

func (h *ListCustomersHandler) GetMethod() string {
	return "GET"
}

func (h *ListCustomersHandler) GetPattern() string {
	return "/customers"
}

func (h *ListCustomersHandler) GetHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		output, err := h.uc.Execute(&usecase.ListCustomersInput{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(output); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

type UpdateCustomerHandler struct {
	uc *usecase.UpdateCustomerUseCase
}

func NewUpdateCustomerHandler(uc *usecase.UpdateCustomerUseCase) *UpdateCustomerHandler {
	return &UpdateCustomerHandler{uc}
}

func (h *UpdateCustomerHandler) GetMethod() string {
	return "PUT"
}

func (h *UpdateCustomerHandler) GetPattern() string {
	return "/customers/{id}"
}

func (h *UpdateCustomerHandler) GetHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input usecase.UpdateCustomerInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		input.Id = chi.URLParam(r, "id")
		output, err := h.uc.Execute(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(output); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type DeleteCustomerHandler struct {
	uc *usecase.DeleteCustomerUseCase
}

func NewDeleteCustomerHandler(uc *usecase.DeleteCustomerUseCase) *DeleteCustomerHandler {
	return &DeleteCustomerHandler{uc}
}

func (h *DeleteCustomerHandler) GetMethod() string {
	return "DELETE"
}

func (h *DeleteCustomerHandler) GetPattern() string {
	return "/customers/{id}"
}

func (h *DeleteCustomerHandler) GetHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := usecase.DeleteCustomerInput{
			Id: chi.URLParam(r, "id"),
		}
		output, err := h.uc.Execute(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(output); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
