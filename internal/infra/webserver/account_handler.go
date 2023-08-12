package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/josimarz/fc-eda-challenge/internal/usecase"
)

type CreateAccountHandler struct {
	uc *usecase.CreateAccountUseCase
}

func NewCreateAccountHandler(uc *usecase.CreateAccountUseCase) *CreateAccountHandler {
	return &CreateAccountHandler{uc}
}

func (h *CreateAccountHandler) GetMethod() string {
	return "POST"
}

func (h *CreateAccountHandler) GetPattern() string {
	return "/customers/{id}/accounts"
}

func (h *CreateAccountHandler) GetHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := usecase.CreateAccountInput{
			CustomerId: chi.URLParam(r, "id"),
		}
		output, err := h.uc.Execute(input)
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

type ListCustomerAccountsHandler struct {
	uc *usecase.ListCustomerAccountsUseCase
}

func NewListCustomerAccountsHandler(uc *usecase.ListCustomerAccountsUseCase) *ListCustomerAccountsHandler {
	return &ListCustomerAccountsHandler{uc}
}

func (h *ListCustomerAccountsHandler) GetMethod() string {
	return "GET"
}

func (h *ListCustomerAccountsHandler) GetPattern() string {
	return "/customers/{id}/accounts"
}

func (h *ListCustomerAccountsHandler) GetHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := usecase.ListCustomerAccountsInput{
			CustomerId: chi.URLParam(r, "id"),
		}
		output, err := h.uc.Execute(input)
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

type DepositHandler struct {
	uc *usecase.DepositUseCase
}

func NewDepositHandler(uc *usecase.DepositUseCase) *DepositHandler {
	return &DepositHandler{uc}
}

func (h *DepositHandler) GetMethod() string {
	return "POST"
}

func (h *DepositHandler) GetPattern() string {
	return "/accounts/{id}/deposit"
}

func (h *DepositHandler) GetHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input usecase.DepositInput
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
		}
	}
}

type WithdrawHandler struct {
	uc *usecase.WithdrawUseCase
}

func NewWithdrawHandler(uc *usecase.WithdrawUseCase) *WithdrawHandler {
	return &WithdrawHandler{uc}
}

func (h *WithdrawHandler) GetMethod() string {
	return "POST"
}

func (h *WithdrawHandler) GetPattern() string {
	return "/accounts/{id}/withdraw"
}

func (h *WithdrawHandler) GetHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input usecase.WithdrawInput
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
		if err := json.NewEncoder(w).Encode(output); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
