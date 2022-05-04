package validators

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofrs/uuid"
)

// ValidateUUID - uuid v4 validation
func ValidateUUID(value interface{}) error {
	uuidValue, _ := value.(uuid.UUID)
	return validation.Validate(uuidValue, validation.Required, is.UUIDv4)
}
