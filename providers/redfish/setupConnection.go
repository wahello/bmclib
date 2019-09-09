package redfish

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/stmcginnis/gofish"
)

// HttpLogin initiates the redfish connection to the device
func (r *Redfish) httpLogin() (err error) {

	if r.service != nil {
		return nil
	}

	log.WithFields(log.Fields{"step": "bmc connection", "provider": "redfish", "ip": r.ip}).Debug("connecting to bmc")

	r.apiClient, err = gofish.Connect(*r.clientConfig)
	if err != nil {
		return err
	}

	r.service = r.apiClient.Service

	return nil
}

// Close logs out of the device
func (r *Redfish) Close() (err error) {

	if r.service != nil {
		r.apiClient.Logout()
	}

	return fmt.Errorf("Attempt to close an invalid connection")
}
