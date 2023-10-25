package storageManager

import (
	"errors"
	"github.com/AnEventTechInventory/Backend/pkg/database"
	"github.com/AnEventTechInventory/Backend/pkg/registry"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

func VerifyContents(baseId string, contents string, db *gorm.DB) (bool, error) {
	// verify that the device contents are valid
	// can only a list of valid uuids followed by exactly one ', ' or be the end
	match, err := regexp.Match(`^([a-f\d]{8}(-[a-f\d]{4}){4}[a-f\d]{8}, )*[a-f\d]{8}(-[a-f\d]{4}){4}[a-f\d]{8}$`, []byte(contents))
	if err != nil {
		return false, err
	}
	if !match {
		return false, database.InsertError{ErrorMessage: "Device contents must be a list of valid uuids separated by exactly one ', '"}
	}

	// Verify that the device contents exists
	for _, content := range strings.Split(contents, ", ") {
		// check if the content ids already exist
		db.Find(&registry.Device{}, "id = ?", content)
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return false, database.InsertError{ErrorMessage: "Device content id does not exist"}
		}

		// prevent circular references
		if content == baseId {
			return false, database.InsertError{ErrorMessage: "Device cannot contain itself"}
		}
	}

	return true, nil
}
