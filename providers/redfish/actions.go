package redfish

import (
	"github.com/stmcginnis/gofish/redfish"
)

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
		if e.SystemType == redfish.PhysicalSystemType {
			err = e.Reset(redfish.ForceRestartResetType)
			if err != nil {
				return status, err
			}
		}
	}

	return true, err
}

// PowerCycleBmc reboots the bmc we are connected to
func (r *Redfish) PowerCycleBmc() (status bool, err error) {
	err = r.httpLogin()
	if err != nil {
		return status, err
	}

	entries, err := r.service.Managers()
	if err != nil {
		return status, err
	}

	if len(entries) < 1 {
		return status, ErrNoSystemEntry
	}

	for _, e := range entries {
		if e.ManagerType == redfish.BMCManagerType {
			err = e.Reset(redfish.ForceRestartResetType)
			if err != nil {
				return status, err
			}
		}
	}

	return true, err
}

// PowerOn power on the machine via bmc
func (r *Redfish) PowerOn() (status bool, err error) {
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
		if e.SystemType == redfish.PhysicalSystemType {
			err = e.Reset(redfish.OnResetType)
			if err != nil {
				return status, err
			}
		}
	}

	return true, err
}

// PowerOff power off the machine via bmc
func (r *Redfish) PowerOff() (status bool, err error) {
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
		if e.SystemType == redfish.PhysicalSystemType {
			err = e.Reset(redfish.ForceOffResetType)
			if err != nil {
				return status, err
			}
		}
	}

	return true, err
}

// PxeOnce makes the machine to boot via pxe once
func (r *Redfish) PxeOnce() (status bool, err error) {
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
		if e.SystemType == redfish.PhysicalSystemType {
			b := redfish.Boot{
				BootSourceOverrideTarget:  redfish.PxeBootSourceOverrideTarget,
				BootSourceOverrideEnabled: redfish.OnceBootSourceOverrideEnabled,
			}
			err = e.SetBoot(b)
			if err != nil {
				return status, err
			}

			return r.PowerCycle()

		}
	}

	return true, err
}

// IsOn tells if a machine is currently powered on
func (r *Redfish) IsOn() (status bool, err error) {
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
		if e.SystemType == redfish.PhysicalSystemType && e.PowerState == redfish.OnPowerState {
			return true, err
		}
	}

	return status, err
}
