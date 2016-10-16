package interfaces

// Storager to abstract the storage. This would be the name of the table/collection
type Storager interface {
	AddOne(c StorageItem) string
	GetOne(id string) (StorageItem, error)
	GetAll() []StorageItem
	DeleteOne(id string) error
	UpdateOne(c StorageItem) error
	MyName() string
}

// StorageItem thing to be storged should do this
type StorageItem interface {
	GetID() string
	SetID(id string) error
}
