// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/thetreep/go-swagger/examples/composed-auth/restapi"
	"github.com/thetreep/go-swagger/examples/composed-auth/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewMultiAuthExampleAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Composing authorizations"
	parser.LongDescription = "This sample API demonstrates how to compose several authentication schemes\nand configure complex security requirements for your operations.\n\nThis API simulates a very simple market place with customers and resellers\nof items.\n\nPersonas:\n  - as a first time user, I want to see all items on sales\n  - as a registered customer, I want to post orders for items and\n    consult my past orders\n  - as a registered reseller, I want to see all pending orders on the items\n    I am selling on the market place\n  - as a reseller managing my own inventories, I want to post replenishment orders for the items I provide\n  - as a register user, I want to consult my personal account infos\n\nThe situation we defined on the authentication side is as follows:\n  - every known user is authenticated using a basic token\n  - resellers are authenticated using API keys - we let the option to authenticate using a header or query param\n  - any registered user (customer or reseller) will add a signed JWT to access more API endpoints\n\nObviously, there are several ways to achieve the same result. We just wanted to demonstrate here how\nsecurity requirements may compose several schemes.\n\nNote that we used the \"OAuth2\" declaration here but don't implement a real\nOAuth2 workflow: our intend here is just to be able to extract scopes from a passed JWT token (the\nonly way to manipulate scoped authorizers with Swagger 2.0 is to declare them with type \"oauth2\").\n"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
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
		log.Fatalln(err)
	}

}
