package helpers

import (
	"go.uber.org/multierr"
	"gorm.io/gorm"
)

// FinalizeTx commit or rollback tx whether err is empty.
func FinalizeTx(dbTx *gorm.DB, err error) error {
	if err != nil {
		if dbErr := dbTx.Rollback().Error; dbErr != nil {
			return multierr.Append(err, dbErr)
		}
		return err
	}

	return dbTx.Commit().Error
}
