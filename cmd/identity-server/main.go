package main

import (
	"identityserver/internal/log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"identityserver/internal/server/restapi"
	"identityserver/internal/server/restapi/operations"
)

func main() {
	logger := log.New(true)

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		logger.Fatal(err.Error())
	}

	api := operations.NewIdentityAPI(swaggerSpec)
	api.Logger = logger.Sugar().Infof
	server := restapi.NewServer(api)
	defer server.Shutdown() // nolint

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Identity API"
	parser.LongDescription = "#### API for identity management\"\n"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		logger.Fatal(err.Error())
	}
}
