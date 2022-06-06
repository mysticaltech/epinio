// Package client connects to the Epinio API's endpoints
package client

import (
	"github.com/epinio/epinio/helpers/tracelog"
	apiv1 "github.com/epinio/epinio/internal/api/v1"
	"github.com/go-logr/logr"
)

// Client provides functionality for talking to an Epinio API
// server
type Client struct {
	log      logr.Logger
	URL      string
	WsURL    string // only stored here for the memo, the websocket client is not part of the epinioapi, yet.
	user     string
	password string
}

// New returns a new Epinio API client
func New(url string, wsURL string, user string, password string) *Client {
	log := tracelog.NewLogger().WithName("EpinioApiClient").V(3)

	// init routes
	// we don't need the controllers in the CLI, but we just need the routes endpoints
	apiv1.Routes.SetRoutes(apiv1.MakeRoutes()...)
	apiv1.Routes.SetRoutes(apiv1.MakeNamespaceRoutes(nil)...)
	apiv1.Routes.SetRoutes(apiv1.MakeWsRoutes()...)

	return &Client{
		log:      log,
		URL:      url,
		WsURL:    wsURL,
		user:     user,
		password: password,
	}
}
