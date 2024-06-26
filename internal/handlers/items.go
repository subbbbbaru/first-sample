package handlers

import (
	"encoding/json"
	"io"

	"context"
	"net/http"

	"github.com/subbbbbaru/first-sample/internal/services"
	"github.com/subbbbbaru/first-sample/pkg/log"
)

type ItemHandler struct {
	service services.ItemService
}

func NewItemHandler(service services.ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (h *ItemHandler) HandleItems(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	log.Info().Printf("Received %s request for %s", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		h.getItems(ctx, w, r)
	case http.MethodPost:
		h.createItem(ctx, w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *ItemHandler) getItems(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetItems(ctx)
	if err != nil {
		log.Error().Printf("Error getting items: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *ItemHandler) createItem(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var newItem struct {
		Name string `json:"name"`
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Printf("Error reading request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &newItem); err != nil {
		log.Error().Printf("Error unmarshaling request body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item, err := h.service.CreateItem(ctx, newItem.Name)
	if err != nil {
		log.Error().Printf("Error creating item: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Info().Printf("Created item: %v", item)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
