package handlers

import (
	"encoding/hex"
	"errors"
	"identityserver/internal/server/restapi/operations/private"
	"identityserver/internal/server/restapi/operations/public"
	"identityserver/pkg/models/rest"
	"identityserver/pkg/repos/identity"
	"identityserver/pkg/services"
	"identityserver/pkg/utils/helpers"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
)

type identityController struct {
	service services.Identity
	log     *zap.Logger
}

type IdentityController interface {
	GetPrivateIdentity(params private.GetPrivateIdentityParams, principal *rest.Principal) middleware.Responder
	GetCountry(params public.GetPublicCountryParams) middleware.Responder
	Registration(params public.RegistrationParams) middleware.Responder
	SetIdentity(params private.SetIdentityParams, principal *rest.Principal) middleware.Responder
}

func (i identityController) GetPrivateIdentity(params private.GetPrivateIdentityParams, principal *rest.Principal) middleware.Responder {
	if principal == nil {
		return private.NewGetPrivateIdentityUnauthorized().WithWWWAuthenticate("Basic")
	}

	id, err := i.service.GetByUserName(params.HTTPRequest.Context(), principal.Username)
	if err != nil {
		i.log.Error("getting private identity", zap.Error(err))
		return public.NewGetPublicCountryDefault(http.StatusInternalServerError).WithPayload(&rest.Error{
			Code:    500,
			Message: "internal server error",
		})
	}

	return private.NewGetPrivateIdentityOK().WithPayload(&rest.IdentityResponse{
		Country:    id.Country,
		EthAddress: "0x" + hex.EncodeToString(id.EthAddress),
		Username:   id.Username,
	})
}

func (i identityController) GetCountry(params public.GetPublicCountryParams) middleware.Responder {
	if !helpers.IsValidAddress(params.Address) {
		return public.NewGetPublicCountryBadRequest().
			WithPayload(&rest.Error{
				Code:    400,
				Message: "provided eth address is invalid",
			})
	}

	id, err := i.service.GetCountry(params.HTTPRequest.Context(), params.Address)
	if err != nil && !errors.Is(err, services.ErrUserNotfound) {
		i.log.Error("getting id by eth address", zap.Error(err))
		return public.NewGetPublicCountryDefault(http.StatusInternalServerError).WithPayload(&rest.Error{
			Code:    500,
			Message: "internal server error",
		})
	}

	return public.NewGetPublicCountryOK().
		WithPayload(&rest.CountryResponse{Country: id.Country})
}

func (i identityController) Registration(params public.RegistrationParams) middleware.Responder {
	_, err := i.service.Create(params.HTTPRequest.Context(), *params.Body.Username, *params.Body.Password)
	if errors.Is(err, identity.ErrAlreadyExist) {
		return public.NewRegistrationConflict().WithPayload(&rest.Error{
			Code:    409,
			Message: "username already taken",
		})
	}

	if err != nil {
		i.log.Error("creating new identity", zap.Error(err))
		return public.NewRegistrationDefault(http.StatusInternalServerError).WithPayload(&rest.Error{
			Code:    500,
			Message: "internal server error",
		})
	}

	return public.NewRegistrationCreated()
}

func (i identityController) SetIdentity(params private.SetIdentityParams, principal *rest.Principal) middleware.Responder {
	if principal == nil {
		return private.NewSetIdentityUnauthorized().WithWWWAuthenticate("Basic")
	}

	if !helpers.IsValidAddress(*params.Body.EthAddress) {
		return private.NewSetIdentityDefault(http.StatusBadRequest).WithPayload(&rest.Error{
			Code:    400,
			Message: "eth-address is not valid",
		})
	}

	if !helpers.IsValidCountry(*params.Body.Country) {
		return private.NewSetIdentityDefault(http.StatusBadRequest).WithPayload(&rest.Error{
			Code:    400,
			Message: "country are accepted in Aplha-3 format only",
		})
	}

	_, err := i.service.SetPublicInfo(params.HTTPRequest.Context(),
		principal.Username, *params.Body.EthAddress, *params.Body.Country)

	if errors.Is(err, services.ErrEthAddressAlreadyAssigned) {
		return private.NewSetIdentityConflict().WithPayload(&rest.Error{
			Code:    409,
			Message: "eth-address already assigned to another identity",
		})
	}

	if err != nil {
		i.log.Error("getting private identity", zap.Error(err))
		return private.NewSetIdentityDefault(http.StatusInternalServerError).WithPayload(&rest.Error{
			Code:    500,
			Message: "internal server error",
		})
	}

	return private.NewSetIdentityOK()
}

func NewIdentityController(svc services.Identity, logger *zap.Logger) IdentityController {
	return &identityController{service: svc, log: logger}
}
