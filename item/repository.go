package item

import "gorm.io/gorm"

type Repository interface {
	Create(item Entity) (Entity, error)
	UpdateById(id string, item Entity) (Entity, error)
	FindById(id string) (Entity, error)
	DeleteById(id string) error
}

type repository struct {
	db *gorm.DB
}

func (r repository) Create(item Entity) (Entity, error) {
	res := r.db.Create(&item)
	if res.Error != nil {
		return Entity{}, res.Error
	}
	return item, nil
}

func (r repository) UpdateById(id string, item Entity) (Entity, error) {
	panic("implement me")
}

func (r repository) FindById(id string) (Entity, error) {
	var e Entity
	res := r.db.Find(&e, "id = ?", id)
	if res.Error != nil {
		return Entity{}, nil
	}
	return e, nil
}

func (r repository) DeleteById(id string) error {
	panic("implement me")
}

func NewRepository(db *gorm.DB) Repository {
	return repository{db: db}
}
