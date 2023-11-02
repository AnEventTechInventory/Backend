package storageManager

type StorageInterface interface {
	Add(device *interface{}) error
	Get(id string) (*interface{}, error)
	List() ([]*interface{}, error)
	Update(device *interface{}) error
	Delete(id string) error
}
