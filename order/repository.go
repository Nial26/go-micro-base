package order

type Repository interface {
	Create(order Entity) (Entity, error)
	FindById(id string) (Entity, error)
}
