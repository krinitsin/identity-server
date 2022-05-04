package identity

import (
	"context"
	"errors"
	"identityserver/pkg/models/orm"
)

type Identity interface {
	Create(ctx context.Context, identity orm.Identity) (orm.Identity, error)
	GetByUserName(ctx context.Context, username string) (orm.Identity, error)
	GetByEthAddress(ctx context.Context, ethAddress []byte) (orm.Identity, error)
	Update(ctx context.Context, identity orm.Identity) (orm.Identity, error)
}

var ErrAlreadyExist = errors.New("identity already exist for given username")
