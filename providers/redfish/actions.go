package redfish

// PowerCycle reboots the machine via bmc
func (r *RedFish) PowerCycle() (status bool, err error) {
	err = r.httpLogin()
	if err != nil {
		return status, err
	}

	return status, err
}
