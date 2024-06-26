package repositories

import (
	"context"
	"sync"

	"github.com/subbbbbaru/first-sample/internal/models"
	"github.com/subbbbbaru/first-sample/pkg/log"
)

type ItemRepository interface {
	GetAll(ctx context.Context) ([]models.Item, error)
	Create(ctx context.Context, item models.Item) (models.Item, error)
}

type InMemoryItemRepository struct {
	items []models.Item
	mu    sync.Mutex
}

func NewInMemoryItemRepository() *InMemoryItemRepository {
	return &InMemoryItemRepository{
		items: []models.Item{
			{ID: 1, Name: "Item One"},
			{ID: 2, Name: "Item Two"},
		},
	}
}

func (r *InMemoryItemRepository) GetAll(ctx context.Context) ([]models.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	log.Info().Print("Getting all items from in-memory repository")
	return r.items, nil
}

func (r *InMemoryItemRepository) Create(ctx context.Context, item models.Item) (models.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	item.ID = r.getNextID()
	r.items = append(r.items, item)
	log.Info().Printf("Created item with ID %d", item.ID)
	return item, nil
}

func (r *InMemoryItemRepository) getNextID() int {
	if len(r.items) == 0 {
		return 1
	}
	lastItem := r.items[len(r.items)-1]
	return lastItem.ID + 1
}
