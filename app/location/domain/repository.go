package domain

type AddressRepository interface {
	GetByID(id string) (*Address, error)
}
