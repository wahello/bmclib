package redfish

import (
	log "github.com/sirupsen/logrus"
	gofish "github.com/stmcginnis/gofish/school"
)

// httpLogin initiates the connection to a bmc device
func (r *RedFish) httpLogin() (err error) {
	if r.sessionID != nil && r.service != nil {
		return
	}

	service, err := gofish.ServiceRoot(r.apiClient)
	if err != nil {
		return err
	}

	auth, err := service.CreateSession(r.username, r.password)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"step": "bmc connection", "provider": "redfish", "ip": i.ip}).Debug("connecting to bmc")

	r.apiClient.Token = auth.Token
	r.service = service
	r.sessionID = &auth.Session

	return err
}

func (r *RedFish) Close() (err error) {
	if r.sessionID != nil && r.service != nil {
		return r.service.DeleteSession(*r.sessionID)
	}
	return err
}
