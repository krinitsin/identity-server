package identity

import (
	"context"
	"encoding/hex"

	"identityserver/pkg/models/orm"
	testutils "identityserver/pkg/test_utils"
	mock_utils "identityserver/pkg/test_utils/mocks/utils"
	"identityserver/pkg/utils"
	"strings"
	"testing"

	"github.com/gofrs/uuid"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIdentityRepo(t *testing.T) {
	// set up db
	db := testutils.NewDbTestSetUp(
		t,
		testutils.WithSchemaFiles(
			"../../test_utils/fixtures/migrations/identity.sql",
		),
		testutils.WithFixtureFiles(
			"../../test_utils/fixtures/repos/identity/identity.yml",
		),
	)
	defer db.TearDown()

	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	tServiceMock := mock_utils.NewMockITimeService(mockCtrl)
	tNow := utils.NewTimeService().GetUTCTimeNow()
	tServiceMock.EXPECT().GetUTCTimeNow().Return(tNow).AnyTimes()
	defer mockCtrl.Finish()

	identityRepo := postgres{
		db: db.Conn,
	}

	id, err := uuid.FromString("9ef240d5-b079-4cd1-99cf-280f2f6eea08")
	require.NoError(t, err, err)
	createAt, err := utils.ParseStringPGFormat("2020-01-29 21:22:51.680713")
	require.NoError(t, err, err)
	updateAt, err := utils.ParseStringPGFormat("2020-01-29 21:22:51.701372")
	require.NoError(t, err, err)

	bAddr, err := hex.DecodeString(strings.TrimPrefix("0x1236A4Efd84dE5Fb0cCb21969C2db775B3891b65", "0x"))
	assert.NoError(t, err, err)

	identity := orm.Identity{
		ID:         id,
		Username:   "Vlad",
		PassHash:   "123",
		EthAddress: bAddr,
		Country:    "RU",
		State:      orm.StateSubmitted,
		CreatedAt:  createAt.Local(),
		UpdatedAt:  updateAt.Local(),
	}

	t.Run("Create New", func(t *testing.T) {
		id, err := uuid.NewV4()
		require.NoError(t, err, err)

		bAddr, err := hex.DecodeString(strings.TrimPrefix("0x321605Ce77603782383d2e9250Cf299EF91c0F67", "0x"))
		assert.NoError(t, err, err)

		expected := orm.Identity{
			ID:         id,
			Username:   "Keanu",
			PassHash:   "Reeves",
			EthAddress: bAddr,
			Country:    "US",
			State:      orm.StatePending,
			CreatedAt:  tNow,
			UpdatedAt:  tNow,
		}

		actual, err := identityRepo.Create(ctx, expected)
		require.NoError(t, err, err)

		assert.EqualValues(t, expected, actual, "identities have diff")
	})

	t.Run("Create existing", func(t *testing.T) {
		id, err := uuid.NewV4()
		require.NoError(t, err, err)

		newIdentity := identity
		newIdentity.ID = id

		_, err = identityRepo.Create(ctx, newIdentity)
		require.Equal(t, ErrAlreadyExist, err, "expected error oon creating identity with taken username")
	})

	t.Run("Get by username", func(t *testing.T) {
		actual, err := identityRepo.GetByUserName(ctx, identity.Username)
		require.NoError(t, err, err)

		assert.EqualValues(t, identity, actual, "identities have diff")
	})

	t.Run("Get by eth address", func(t *testing.T) {
		actual, err := identityRepo.GetByEthAddress(ctx, identity.EthAddress)
		require.NoError(t, err, err)

		assert.EqualValues(t, identity, actual, "identities have diff")
	})

	t.Run("Update", func(t *testing.T) {
		bAddr, err := hex.DecodeString(strings.TrimPrefix("0x321e7e4715Bb0b1AB7affD6b3d429b5f70F3B342", "0x"))
		assert.NoError(t, err, err)

		newIdentity := identity
		newIdentity.EthAddress = bAddr
		newIdentity.Country = "EU"
		newIdentity.State = orm.StatePending

		actual, err := identityRepo.Update(ctx, newIdentity)
		require.NoError(t, err, err)

		newIdentity.UpdatedAt = actual.UpdatedAt

		actual, err = identityRepo.GetByUserName(ctx, newIdentity.Username)
		require.NoError(t, err, err)

		assert.EqualValues(t, newIdentity, actual, "identities have diff")
	})
}
