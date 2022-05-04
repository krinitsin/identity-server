package services

import (
	"encoding/hex"

	"identityserver/pkg/models/orm"
	identity2 "identityserver/pkg/repos/identity"
	"identityserver/pkg/utils"
	"identityserver/pkg/utils/helpers"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Identity interface {
	Create(ctx context.Context, username, password string) (orm.Identity, error)
	SetPublicInfo(ctx context.Context, username, ethAddress, country string) (orm.Identity, error)
	GetByUserName(ctx context.Context, username string) (orm.Identity, error)
	GetCountry(ctx context.Context, ethAddress string) (orm.Identity, error)
}

type identity struct {
	db          *gorm.DB
	repo        identity2.Identity
	timeService utils.ITimeService
}

func NewIdentity(db *gorm.DB, repo identity2.Identity, timeService utils.ITimeService) Identity {
	return identity{
		db:          db,
		repo:        repo,
		timeService: timeService,
	}
}

var ErrEthAddressAlreadyAssigned = errors.New("eth address already assigned to existing identity")
var ErrUserNotfound = errors.New("identity not found")

func (i identity) Create(ctx context.Context, username, password string) (orm.Identity, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return orm.Identity{}, errors.Wrap(err, "generating password hash")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return orm.Identity{}, errors.Wrap(err, "generating uuid")
	}

	identity, err := i.repo.Create(ctx, orm.Identity{
		ID:         id,
		Username:   username,
		PassHash:   string(hash),
		EthAddress: nil,
		Country:    "",
		State:      orm.StatePending,
		CreatedAt:  i.timeService.GetUTCTimeNow(),
		UpdatedAt:  i.timeService.GetUTCTimeNow(),
	})

	return identity, errors.Wrap(err, "creating identity")
}

func (i identity) SetPublicInfo(ctx context.Context, username, ethAddress, country string) (updIdentity orm.Identity, err error) {
	bAddr, err := hex.DecodeString(strings.TrimPrefix(ethAddress, "0x"))
	if err != nil {
		return updIdentity, errors.Wrap(err, "ethAddress is a not valid hex")
	}

	// start TX
	tx := i.db.Begin()
	defer func() { err = helpers.FinalizeTx(tx, err) }()

	txRepo := identity2.NewIdentityRepo(tx)

	// check of eth-address already in use
	identity, err := txRepo.GetByEthAddress(ctx, bAddr)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return updIdentity, errors.Wrap(err, "getting eth address")
	}

	// another user holds this eth-address
	if err == nil && identity.Username != username {
		return updIdentity, ErrEthAddressAlreadyAssigned
	}

	updIdentity, err = txRepo.Update(ctx, orm.Identity{
		Username:   username,
		EthAddress: bAddr,
		Country:    country,
		State:      orm.StateSubmitted,
	})

	return updIdentity, errors.Wrap(err, "updating identity")
}

func (i identity) GetByUserName(ctx context.Context, username string) (orm.Identity, error) {
	identity, err := i.repo.GetByUserName(ctx, username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return identity, ErrUserNotfound
	}
	return identity, errors.Wrap(err, "getting identity by username")
}

func (i identity) GetCountry(ctx context.Context, ethAddress string) (orm.Identity, error) {
	bAddr, err := hex.DecodeString(strings.TrimPrefix(ethAddress, "0x"))
	if err != nil {
		return orm.Identity{}, errors.Wrap(err, "ethAddress is a not valid hex")
	}

	identity, err := i.repo.GetByEthAddress(ctx, bAddr)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return identity, ErrUserNotfound
	}
	return identity, errors.Wrap(err, "getting identity by username")
}
