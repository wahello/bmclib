package redfish

import (
	"fmt"

	"github.com/bmc-toolbox/bmclib/internal/httpclient"

	gofish "github.com/stmcginnis/gofish/school"
	// this make possible to setup logging and properties at any stage
	_ "github.com/bmc-toolbox/bmclib/logging"
)

// Ilo holds the status and properties of a connection to an iLO device
type RedFish struct {
	ip        string
	username  string
	password  string
	apiClient *gofish.ApiClient
	service   *gofish.Service
	sessionID *string
}

// New returns a new Ilo ready to be used
func New(ip string, username string, password string) (r *RedFish, err error) {
	client, err := httpclient.Build()
	if err != nil {
		return r, err
	}

	apiClient, err := gofish.APIClient(fmt.Sprintf("https://%s", ip), client)
	if err != nil {
		return r, err
	}

	return &RedFish{ip: ip, username: username, password: password, apiClient: apiClient}, err
}

func (r *RedFish) CheckCredentials() (err error) {
	return r.httpLogin()
}
