package orm

import (
	"identityserver/pkg/models/orm/validators"
	"time"

	"github.com/gofrs/uuid"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Identity struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username   string    `json:"username" gorm:"unique_index:idx_identity_username"`
	PassHash   string    `json:"-"`
	EthAddress []byte    `json:"eth_address" gorm:"type:bytea"`
	Country    string    `json:"country"`
	State      state     `json:"state"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const StateSubmitted state = "SUBMITTED"
const StatePending state = "PENDING"

type state string

func (i *Identity) TableName() string {
	return "identity"
}

func (i *Identity) Validate() error {
	return validation.ValidateStruct(i,
		validation.Field(&i.ID, validation.By(validators.ValidateUUID)),
		validation.Field(&i.State, validation.In(StateSubmitted, StatePending)),
		validation.Field(&i.Username, validation.Required),
		validation.Field(&i.PassHash, validation.Required),
		validation.Field(&i.CreatedAt, validation.Required),
	)
}
