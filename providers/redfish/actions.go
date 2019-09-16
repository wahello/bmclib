package redfish

import "github.com/stmcginnis/gofish/redfish"

// PowerCycle reboots the machine via bmc
func (r *Redfish) PowerCycle() (status bool, err error) {
	err = r.httpLogin()
	if err != nil {
		return status, err
	}

	entries, err := r.service.Systems()
	if err != nil {
		return status, err
	}

	if len(entries) < 1 {
		return status, ErrNoSystemEntry
	}

	for _, e := range entries {
		if e.SystemType == "Physical" {
			err = e.Reset(redfish.ForceRestartResetType)
			if err != nil {
				panic(err)
			}
		}
	}

	return status, err
}
