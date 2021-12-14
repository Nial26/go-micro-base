package item

import "github.com/segmentio/ksuid"

type Service interface {
	CreateItem(item Item) (Item, error)
	FindItem(id string) (Item, error)
	UpdateItem(id string, item Item) (Item, error)
	DeleteItem(id string) error
}

type service struct {
	repository Repository
}

func (s service) CreateItem(item Item) (Item, error) {
	item.Id = ksuid.New().String()
	ie, err := s.repository.Create(item.ToEntity())
	if err != nil {
		return Item{}, err
	}
	return ie.ToDTO(), nil
}

func (s service) FindItem(id string) (Item, error) {
	ie, err := s.repository.FindById(id)
	if err != nil {
		return Item{}, err
	}
	return ie.ToDTO(), nil
}

func (s service) UpdateItem(id string, item Item) (Item, error) {
	panic("implement me")
}

func (s service) DeleteItem(id string) error {
	panic("implement me")
}

func NewService(repo Repository) Service {
	return service{repository: repo}
}
