package storageManager

import "github.com/google/uuid"

type StorageInterface[T any] interface {
	Add(device *T) error
	Get(id string) (*T, error)
	List() ([]*uuid.UUID, error)
	Update(newEntry *T) error
	Delete(id string) error
}
