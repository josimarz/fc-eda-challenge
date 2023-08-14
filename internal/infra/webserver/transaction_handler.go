package webserver

import (
	"encoding/json"
	"net/http"

	eventhandling "github.com/josimarz/fc-eda-challenge/internal/event_handling"
	"github.com/josimarz/fc-eda-challenge/internal/usecase"
	"github.com/josimarz/fc-eda-challenge/pkg/events"
)

type CreateTransactionHandler struct {
	ed *events.EventDispatcher
}

func NewCreateTransactionHandler(ed *events.EventDispatcher) *CreateTransactionHandler {
	return &CreateTransactionHandler{ed}
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
		event := eventhandling.NewTransactionCreatedEvent()
		event.SetPayload(input)
		h.ed.Dispatch(event)
		w.WriteHeader(http.StatusNoContent)
	}
}
