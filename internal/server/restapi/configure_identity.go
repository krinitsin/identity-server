// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"identityserver/internal/log"
	"identityserver/internal/server/auth"
	"identityserver/internal/server/handlers"
	"identityserver/pkg/repos"
	"identityserver/pkg/repos/identity"
	"identityserver/pkg/services"
	"identityserver/pkg/utils"
	"net/http"

	"github.com/go-openapi/swag"

	"go.uber.org/zap"

	"identityserver/internal/server/restapi/operations"
	"identityserver/internal/server/restapi/operations/private"
	"identityserver/internal/server/restapi/operations/public"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
)

func configureFlags(api *operations.IdentityAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "db",
			LongDescription:  "db",
			Options:          &repos.PgConfig{},
		},
	}
}

func configureAPI(api *operations.IdentityAPI) http.Handler {
	logger := log.New(true)
	api.Logger = logger.Sugar().Debugf

	api.ServeError = errors.ServeError

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	dbOps := api.CommandLineOptionsGroups[0].Options.(*repos.PgConfig)
	conn, err := repos.GetDBInstance(repos.PgConfig{
		UserName: dbOps.UserName,
		Password: dbOps.Password,
		Dbname:   dbOps.Dbname,
		Port:     dbOps.Port,
		Host:     dbOps.Host,
	})
	if err != nil {
		logger.Fatal("db init error", zap.Error(err))
	}

	repo := identity.NewIdentityRepo(conn)

	tSvc := utils.NewTimeService()

	auth := auth.NewBasicAuthInterceptor(repo, logger)
	// Applies when the Authorization header is set with the Basic scheme
	api.BasicAuthAuth = auth.Auth

	// Set your custom authorizer if needed. Default one is security.Authorized()
	api.APIAuthorizer = auth.Authorized()

	svc := services.NewIdentity(conn, repo, tSvc)
	identityCtrl := handlers.NewIdentityController(svc, logger)

	api.PrivateGetPrivateIdentityHandler = private.GetPrivateIdentityHandlerFunc(identityCtrl.GetPrivateIdentity)
	api.PublicGetPublicCountryHandler = public.GetPublicCountryHandlerFunc(identityCtrl.GetCountry)
	api.PublicRegistrationHandler = public.RegistrationHandlerFunc(identityCtrl.Registration)
	api.PrivateSetIdentityHandler = private.SetIdentityHandlerFunc(identityCtrl.SetIdentity)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
