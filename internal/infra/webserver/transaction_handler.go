package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/josimarz/fc-eda-challenge/internal/usecase"
)

type CreateTransactionHandler struct {
	uc *usecase.CreateTransactionUseCase
}

func NewCreateTransactionHandler(uc *usecase.CreateTransactionUseCase) *CreateTransactionHandler {
	return &CreateTransactionHandler{uc}
}

func (h *CreateTransactionHandler) GetMethod() string {
	return "POST"
}

func (h *CreateTransactionHandler) GetPattern() string {
	return "/transactions"
}

func (h *CreateTransactionHandler) GetHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input usecase.CreateTransactionInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
