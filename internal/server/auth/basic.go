package auth

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"identityserver/pkg/models/rest"
	"identityserver/pkg/repos/identity"
)

var errForbidden = errors.New("access denied")

type Basic interface {
	Auth(username string, password string) (*rest.Principal, error)
	Authorized() runtime.Authorizer
}

type basic struct {
	repo   identity.Identity
	logger *zap.Logger
}

func NewBasicAuthInterceptor(
	repo identity.Identity,
	logger *zap.Logger) Basic {
	return basic{repo: repo, logger: logger}
}

func (b basic) Auth(username string, password string) (*rest.Principal, error) {
	return &rest.Principal{
		Password: password,
		Username: username,
	}, nil
}

func (b basic) Authorized() runtime.Authorizer {
	return runtime.AuthorizerFunc(func(r *http.Request, principal interface{}) error {
		pr := principal.(*rest.Principal)

		iden, err := b.repo.GetByUserName(r.Context(), pr.Username)
		if err != nil {
			b.logger.Debug("Authorization failed for: ", zap.String("username", pr.Username), zap.Error(err))
			return errForbidden
		}

		return bcrypt.CompareHashAndPassword([]byte(iden.PassHash), []byte(pr.Password))
	})
}
