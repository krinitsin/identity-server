package handlers

import (
	"encoding/hex"
	"identityserver/internal/log"
	"identityserver/internal/server/auth"
	"identityserver/internal/server/restapi/operations/private"
	"identityserver/internal/server/restapi/operations/public"
	"identityserver/pkg/models/orm"
	"identityserver/pkg/models/rest"
	"identityserver/pkg/repos/identity"
	"identityserver/pkg/services"
	testutils "identityserver/pkg/test_utils"
	mock_utils "identityserver/pkg/test_utils/mocks/utils"
	"identityserver/pkg/utils"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"context"
	"testing"
)

func TestIdentityController(t *testing.T) {
	// set up db
	db := testutils.NewDbTestSetUp(
		t,
		testutils.WithSchemaFiles(
			"../../../pkg/test_utils/fixtures/migrations/identity.sql",
		),
		testutils.WithFixtureFiles(
			"../../../pkg/test_utils/fixtures/repos/identity/identity.yml",
		),
	)
	defer db.TearDown()

	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	tServiceMock := mock_utils.NewMockITimeService(mockCtrl)
	tNow := utils.NewTimeService().GetUTCTimeNow()
	tServiceMock.EXPECT().GetUTCTimeNow().Return(tNow).AnyTimes()
	defer mockCtrl.Finish()

	identityRepo := identity.NewIdentityRepo(db.Conn)
	svc := services.NewIdentity(db.Conn, identityRepo, tServiceMock)
	logger := log.New(true)

	bAuth := auth.NewBasicAuthInterceptor(identityRepo, log.New(true))

	ctrl := identityController{
		service: svc,
		log:     logger,
	}

	id, err := uuid.FromString("9ef240d5-b079-4cd1-99cf-280f2f6eea08")
	require.NoError(t, err, err)
	createdAt, err := utils.ParseStringPGFormat("2020-01-29 21:22:51.680713")
	require.NoError(t, err, err)
	updatedAt, err := utils.ParseStringPGFormat("2020-01-29 21:22:51.701372")
	require.NoError(t, err, err)
	bAddr, err := hex.DecodeString(strings.TrimPrefix("0x1236A4Efd84dE5Fb0cCb21969C2db775B3891b65", "0x"))
	assert.NoError(t, err, err)

	VladIdentity := orm.Identity{
		ID:         id,
		Username:   "Vlad",
		PassHash:   "123",
		EthAddress: bAddr,
		Country:    "RU",
		State:      orm.StateSubmitted,
		CreatedAt:  createdAt.Local(),
		UpdatedAt:  updatedAt.Local(),
	}

	keanu := "Keanu"
	reeves := "Reeves"

	t.Run("Registration", func(t *testing.T) {
		req := &http.Request{}
		resp := ctrl.Registration(public.RegistrationParams{
			HTTPRequest: req.WithContext(ctx),
			Body: &rest.RegistrationRequest{
				Password: &reeves,
				Username: &keanu,
			},
		})

		assert.Equal(t, public.NewRegistrationCreated(), resp, "responses is not OK")

		actual, err := identityRepo.GetByUserName(ctx, keanu)

		expected := orm.Identity{
			ID:        actual.ID,
			Username:  keanu,
			PassHash:  actual.PassHash,
			State:     orm.StatePending,
			CreatedAt: tNow.Local(),
			UpdatedAt: tNow.Local(),
		}
		assert.NoError(t, err, err)

		err = bAuth.Authorized().Authorize(&http.Request{}, &rest.Principal{
			Password: reeves,
			Username: actual.Username,
		})

		assert.NoError(t, err, err)

		assert.EqualValues(t, expected, actual, "identity created in database is unexpected")
	})

	t.Run("Registration already exist", func(t *testing.T) {
		username := "Vlad"
		password := "New pass"

		req := &http.Request{}
		resp := ctrl.Registration(public.RegistrationParams{
			HTTPRequest: req.WithContext(ctx),
			Body: &rest.RegistrationRequest{
				Password: &password,
				Username: &username,
			},
		})

		assert.Equal(t, public.NewRegistrationConflict().WithPayload(&rest.Error{
			Code:    409,
			Message: "username already taken",
		}), resp, "responses is not OK")
	})

	t.Run("Set Identity", func(t *testing.T) {
		country := "USA"
		ethAddress := "0x71C7656EC7ab88b098defB751B7401B5f6d8976F"

		req := &http.Request{}
		resp := ctrl.SetIdentity(private.SetIdentityParams{
			HTTPRequest: req.WithContext(ctx),
			Body: &rest.SetIdentityRequest{
				Country:    &country,
				EthAddress: &ethAddress,
			},
		},
			&rest.Principal{
				Password: reeves,
				Username: keanu,
			})

		assert.Equal(t, private.NewSetIdentityOK(), resp, "responses is not OK")

		actual, err := identityRepo.GetByUserName(ctx, keanu)
		assert.NoError(t, err, err)

		bAddr, err := hex.DecodeString(strings.TrimPrefix(ethAddress, "0x"))
		assert.NoError(t, err, err)

		expected := orm.Identity{
			ID:         actual.ID,
			Username:   keanu,
			PassHash:   actual.PassHash,
			EthAddress: bAddr,
			Country:    country,
			State:      orm.StateSubmitted,
			CreatedAt:  tNow.Local(),
			UpdatedAt:  actual.UpdatedAt,
		}
		assert.NoError(t, err, err)

		assert.EqualValues(t, expected, actual, "identity created in database is unexpected")
	})

	t.Run("Get Private Identity", func(t *testing.T) {
		req := &http.Request{}
		resp := ctrl.GetPrivateIdentity(private.GetPrivateIdentityParams{
			HTTPRequest: req.WithContext(ctx),
		},
			&rest.Principal{
				Password: "",
				Username: VladIdentity.Username,
			})

		expected := &private.GetPrivateIdentityOK{
			Payload: &rest.IdentityResponse{
				Country:    VladIdentity.Country,
				EthAddress: "0x" + hex.EncodeToString(VladIdentity.EthAddress),
				Username:   VladIdentity.Username,
			},
		}
		assert.Equal(t, expected, resp, "unexpected response")
	})

	t.Run("Get Country when exist", func(t *testing.T) {
		req := &http.Request{}
		resp := ctrl.GetCountry(public.GetPublicCountryParams{
			HTTPRequest: req.WithContext(ctx),
			Address:     "0x" + hex.EncodeToString(VladIdentity.EthAddress),
		})

		expected := &public.GetPublicCountryOK{
			Payload: &rest.CountryResponse{
				Country: VladIdentity.Country,
			},
		}
		assert.Equal(t, expected, resp, "unexpected response")
	})

	t.Run("Get Country when doesn't exist", func(t *testing.T) {
		ethAddress := "0x1236E3e05169B46D25d692C7b35599C1A3b36354"

		req := &http.Request{}
		resp := ctrl.GetCountry(public.GetPublicCountryParams{
			HTTPRequest: req.WithContext(ctx),
			Address:     ethAddress,
		})

		expected := &public.GetPublicCountryOK{
			Payload: &rest.CountryResponse{
				Country: "",
			},
		}
		assert.Equal(t, expected, resp, "unexpected response")
	})
}
