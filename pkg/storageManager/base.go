package storageManager

import "github.com/google/uuid"

type StorageInterface interface {
	Add(device *interface{}) error
	Get(id string) (*interface{}, error)
	List() ([]*uuid.UUID, error)
	Update(device *interface{}) error
	Delete(id string) error
}
