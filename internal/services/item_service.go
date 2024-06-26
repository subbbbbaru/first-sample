package services

import (
	"context"

	"github.com/subbbbbaru/first-sample/internal/models"
	"github.com/subbbbbaru/first-sample/internal/repositories"
	"github.com/subbbbbaru/first-sample/pkg/log"
)

type ItemService interface {
	GetItems(ctx context.Context) ([]models.Item, error)
	CreateItem(ctx context.Context, name string) (models.Item, error)
}

type itemService struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{repo: repo}
}

func (s *itemService) GetItems(ctx context.Context) ([]models.Item, error) {
	items, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Error().Printf("Error getting items from repository: %v", err)
		return nil, err
	}

	return items, nil

}

func (s *itemService) CreateItem(ctx context.Context, name string) (models.Item, error) {
	item := models.Item{Name: name}
	createdItem, err := s.repo.Create(ctx, item)
	if err != nil {
		log.Error().Printf("Error creating item in repository: %v", err)
		return models.Item{}, err
	}
	return createdItem, nil
}
