package identity

import (
	"context"
	"identityserver/pkg/models/orm"

	"gorm.io/gorm"
)

type postgres struct {
	db *gorm.DB
}

func NewIdentityRepo(db *gorm.DB) Identity {
	return postgres{db: db}
}

func (p postgres) Create(ctx context.Context, identity orm.Identity) (orm.Identity, error) {
	if err := identity.Validate(); err != nil {
		return orm.Identity{}, err
	}

	var exists orm.Identity
	p.db.WithContext(ctx).
		Where(orm.Identity{Username: identity.Username}).
		First(&exists)
	if err := exists.Validate(); err == nil {
		return exists, ErrAlreadyExist
	}

	result := p.db.Create(&identity)
	if result.Error != nil {
		return identity, result.Error
	}

	return identity, nil
}

func (p postgres) GetByUserName(ctx context.Context, username string) (orm.Identity, error) {
	var identity orm.Identity

	results := p.db.WithContext(ctx).
		Where(&orm.Identity{Username: username}).
		First(&identity)

	return identity, results.Error
}

func (p postgres) GetByEthAddress(ctx context.Context, ethAddress []byte) (orm.Identity, error) {
	var identity orm.Identity

	results := p.db.WithContext(ctx).
		Where(&orm.Identity{EthAddress: ethAddress, State: orm.StateSubmitted}).
		First(&identity)

	return identity, results.Error
}

func (p postgres) Update(ctx context.Context, identity orm.Identity) (orm.Identity, error) {
	//identity.UpdatedAt = p.timeService.GetUTCTimeNow()

	result := p.db.WithContext(ctx).
		Select("eth_address", "country", "state", "updated_at").
		Where(orm.Identity{Username: identity.Username}).
		Updates(&identity)

	return identity, result.Error
}
