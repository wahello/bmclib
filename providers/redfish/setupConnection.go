package redfish

import (
	log "github.com/sirupsen/logrus"
	"github.com/stmcginnis/gofish"
)

// HttpLogin initiates the connection to a bmc device
func (r *RedFish) httpLogin() (err error) {

	if r.sessionID != nil && r.service != nil {
		return nil
	}

	r.apiClient, err = gofish.Connect(*r.clientConfig)
	if err != nil {
		return err
	}

	service, err := gofish.ServiceRoot(r.apiClient)
	if err != nil {
		return err
	}

	auth, err := service.CreateSession(r.clientConfig.Username, r.clientConfig.Password)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"step": "bmc connection", "provider": "redfish", "ip": r.ip}).Debug("connecting to bmc")

	r.service = service
	r.sessionID = &auth.Session

	return err
}

// Close closes the current session
func (r *RedFish) Close() (err error) {
	if r.sessionID != nil && r.service != nil {
		return r.service.DeleteSession(*r.sessionID)
	}
	return err
}
