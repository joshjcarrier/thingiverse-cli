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
func GetThingByID(id string, opts *Options) (Thing, error) {
	client, _ := newClient(opts)
	response, _ := client.GetThingByIdWithResponse(context.Background(), id)
	return *response.JSON200, nil
}

// ListThingsByUsername will list things by username
func ListThingsByUsername(username string, opts *Options) (Things, error) {
	client, _ := newClient(opts)
	response, _ := client.ListThingsByUsernameWithResponse(context.Background(), username)
	return *response.JSON200, nil
}

// SearchThingsByTerm searchs Thingiverse for items matching all terms
func SearchThingsByTerm(terms []string, opts *Options) (Things, error) {
	client, _ := newClient(opts)
	response, _ := client.SearchThingsByTermWithResponse(context.Background(), strings.Join(terms, "+"))
	return *response.JSON200, nil
}

func newClient(opts *Options) (*ClientWithResponses, error) {
	bearerTokenProvider, _ := securityprovider.NewSecurityProviderBearerToken(opts.Token)
	return NewClientWithResponses("https://api.thingiverse.com", WithRequestEditorFn(bearerTokenProvider.Intercept))
}
