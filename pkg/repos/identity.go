package repos

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"identityserver/pkg/models/orm"
)

type Identity interface {
	Create(ctx context.Context, identity orm.Identity) (orm.Identity, error)
	GetByUserName(ctx context.Context, username string) (orm.Identity, error)
	GetByEthAddress(ctx context.Context, ethAddress []byte) (orm.Identity, error)
	Update(ctx context.Context, identity orm.Identity) (orm.Identity, error)
}

var ErrAlreadyExist = errors.New("identity already exist for given username")

type identity struct {
	db *gorm.DB
}

func NewIdentityRepo(db *gorm.DB) Identity {
	return identity{db: db}
}

func (p identity) Create(ctx context.Context, identity orm.Identity) (orm.Identity, error) {
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

func (p identity) GetByUserName(ctx context.Context, username string) (orm.Identity, error) {
	var identity orm.Identity

	results := p.db.WithContext(ctx).
		Where(&orm.Identity{Username: username}).
		First(&identity)

	return identity, results.Error
}

func (p identity) GetByEthAddress(ctx context.Context, ethAddress []byte) (orm.Identity, error) {
	var identity orm.Identity

	results := p.db.WithContext(ctx).
		Where(&orm.Identity{EthAddress: ethAddress, State: orm.StateSubmitted}).
		First(&identity)

	return identity, results.Error
}

func (p identity) Update(ctx context.Context, identity orm.Identity) (orm.Identity, error) {
	result := p.db.WithContext(ctx).
		Select("eth_address", "country", "state", "updated_at").
		Where(orm.Identity{Username: identity.Username}).
		Updates(&identity)

	return identity, result.Error
}
