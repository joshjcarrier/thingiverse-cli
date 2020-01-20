package thingiverse

import (
	"context"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Options for configuring the client.
type Options struct {
	Token string
}

// InstallAPICommandFlags ensures the command can use the API
func InstallAPICommandFlags(c *cobra.Command, opts *Options) {
	c.Flags().StringVar(&opts.Token, "token", "", "a Thingiverse authentication token")
	viper.BindPFlag("token", c.Flags().Lookup("token"))
}

// GetThingByID gets a thing by its database id
func GetThingByID(id string, opts *Options) (*Thing, error) {
	client, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	response, err := client.GetThingByIdWithResponse(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if (*response).StatusCode() >= 300 {
		return nil, nil
	}
	return response.JSON200, nil
}

// GetThingFilesByThingID gets a thing by its database id
func GetThingFilesByThingID(id string, opts *Options) (Files, error) {
	client, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	response, err := client.GetThingFilesByIdWithResponse(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if (*response).StatusCode() >= 300 {
		return nil, nil
	}
	return *response.JSON200, nil
}

// ListThingsByUsername will list things by username
func ListThingsByUsername(username string, opts *Options) (Things, error) {
	client, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	response, err := client.ListThingsByUsernameWithResponse(context.Background(), username)
	if err != nil {
		return nil, err
	}
	if (*response).StatusCode() >= 300 {
		return nil, nil
	}
	return *response.JSON200, nil
}

// SearchThingsByTerm searchs Thingiverse for items matching all terms
func SearchThingsByTerm(terms []string, opts *Options) (Things, error) {
	client, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	response, err := client.SearchThingsByTermWithResponse(context.Background(), strings.Join(terms, "+"))
	if err != nil {
		return nil, err
	}
	return *response.JSON200, nil
}

func newClient(opts *Options) (*ClientWithResponses, error) {
	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken(opts.Token)
	if err != nil {
		return nil, err
	}
	return NewClientWithResponses("https://api.thingiverse.com", WithRequestEditorFn(bearerTokenProvider.Intercept))
}
