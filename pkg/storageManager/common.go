package storageManager

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func validateUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return errors.New(fmt.Sprintf("invalid device id: %v", id))
	}
	return nil
}
