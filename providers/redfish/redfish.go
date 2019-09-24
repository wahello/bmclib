package redfish

import (
	"strings"

	"github.com/bmc-toolbox/bmclib/internal/httpclient"

	"github.com/stmcginnis/gofish"
	// this make possible to setup logging and properties at any stage
	_ "github.com/bmc-toolbox/bmclib/logging"
)

// Redfish holds the status and properties of a connection to an Redfish service
// it wraps the gofish.APIClient
type Redfish struct {
	ip           string
	apiClient    *gofish.APIClient
	clientConfig *gofish.ClientConfig
	service      *gofish.Service
}

// New returns a new Redfish instance ready for use
func New(ip string, username string, password string) (r *Redfish, err error) {

	client, err := httpclient.Build()
	if err != nil {
		return r, err
	}

	config := &gofish.ClientConfig{
		Username:   username,
		Password:   password,
		Endpoint:   "https://" + ip,
		HTTPClient: client,
	}

	return &Redfish{clientConfig: config, ip: ip}, err
}

// CheckCredentials verify whether the credentials are valid or not
func (r *Redfish) CheckCredentials() (err error) {
	return r.httpLogin()
}

// Serial returns the device serial
func (r *Redfish) Serial() (serial string, err error) {
	err = r.httpLogin()
	if err != nil {
		return serial, err
	}

	entries, err := r.service.Chassis()
	if err != nil {
		return serial, err
	}

	for _, e := range entries {
		if e.ChassisType == "Blade" || e.ChassisType == "RackMount" {
			return strings.ToLower(e.SerialNumber), err
		}
	}
	return serial, err
}

// ChassisSerial returns the serial number of the chassis where the blade is attached
func (r *Redfish) ChassisSerial() (serial string, err error) {
	err = r.httpLogin()
	if err != nil {
		return serial, err
	}

	entries, err := r.service.Chassis()
	if err != nil {
		return serial, err
	}

	for _, e := range entries {
		if e.ChassisType == "Enclosure" {
			return strings.ToLower(e.SerialNumber), err
		}
	}
	return serial, err
}

// Model returns the device model
func (r *Redfish) Model() (model string, err error) {
	err = r.httpLogin()
	if err != nil {
		return model, err
	}

	entries, err := r.service.Chassis()
	if err != nil {
		return model, err
	}

	for _, e := range entries {
		if e.ChassisType == "Blade" || e.ChassisType == "RackMount" {
			return e.Model, err
		}
	}
	return model, err
}

// HardwareType returns the type of bmc we are talking to
func (r *Redfish) HardwareType() (model string, err error) {
	err = r.httpLogin()
	if err != nil {
		return model, err
	}

	entries, err := r.service.Managers()
	if err != nil {
		return model, err
	}

	for _, e := range entries {
		if e.ManagerType == "BMC" {
			return e.Model, err
		}
	}
	return model, err
}

// Version returns the version of the bmc we are running
func (r *Redfish) Version() (version string, err error) {
	err = r.httpLogin()
	if err != nil {
		return version, err
	}

	entries, err := r.service.Managers()
	if err != nil {
		return version, err
	}

	for _, e := range entries {
		if e.ManagerType == "BMC" {
			return e.FirmwareVersion, err
		}
	}
	return version, err
}

// Vendor returns the vendor of the bmc we are running
func (r *Redfish) Vendor() (vendor string, err error) {
	err = r.httpLogin()
	if err != nil {
		return vendor, err
	}

	return r.service.Vendor, nil
}

// Name returns the version of the bmc we are running
func (r *Redfish) Name() (name string, err error) {
	err = r.httpLogin()
	if err != nil {
		return name, err
	}

	entries, err := r.service.Systems()
	if err != nil {
		return name, err
	}

	for _, e := range entries {
		if e.SystemType == "Physical" {
			return e.HostName, err
		}
	}
	return name, err
}

// Status returns health string status from the bmc
func (r *Redfish) Status() (status string, err error) {
	err = r.httpLogin()
	if err != nil {
		return status, err
	}

	entries, err := r.service.Systems()
	if err != nil {
		return status, err
	}

	for _, e := range entries {
		if e.SystemType == "Physical" {
			return string(e.Status.Health), err
		}
	}
	return status, err
}

// Memory returns the total amount of memory of the server
func (r *Redfish) Memory() (memory int, err error) {
	err = r.httpLogin()
	if err != nil {
		return memory, err
	}

	entries, err := r.service.Systems()
	if err != nil {
		return memory, err
	}

	for _, e := range entries {
		if e.SystemType == "Physical" {
			return int(e.MemorySummary.TotalSystemMemoryGiB), err
		}
	}
	return memory, err
}

// CPU returns the cpu, cores and hyperthreads of the server
func (r *Redfish) CPU() (cpu string, cpuCount int, coreCount int, hyperthreadCount int, err error) {
	err = r.httpLogin()
	if err != nil {
		return cpu, cpuCount, coreCount, hyperthreadCount, err
	}

	entries, err := r.service.Systems()
	if err != nil {
		return cpu, cpuCount, coreCount, hyperthreadCount, err
	}

	for _, e := range entries {
		if e.SystemType == "Physical" {
			return e.ProcessorSummary.Model, e.ProcessorSummary.Count, 0, 0, err
		}
	}
	return cpu, cpuCount, coreCount, hyperthreadCount, err
}

// BiosVersion returns the current version of the bios
func (r *Redfish) BiosVersion() (version string, err error) {
	err = r.httpLogin()
	if err != nil {
		return version, err
	}

	entries, err := r.service.Systems()
	if err != nil {
		return version, err
	}

	if len(entries) < 1 {
		return version, ErrNoSystemEntry
	}

	for _, e := range entries {
		if e.SystemType == "Physical" {
			return e.BIOSVersion, nil
		}
	}

	return version, nil
}

// PowerState returns the current power state of the machine
func (r *Redfish) PowerState() (state string, err error) {

	err = r.httpLogin()
	if err != nil {
		return state, err
	}

	entries, err := r.service.Systems()
	if err != nil {
		return state, err
	}

	if len(entries) < 1 {
		return state, ErrNoSystemEntry
	}

	for _, e := range entries {
		if e.SystemType == "Physical" {
			return string(e.PowerState), nil
		}
	}

	return state, nil
}

// PowerKw returns the current power usage in Kw
func (r *Redfish) PowerKw() (power float64, err error) {
	err = r.httpLogin()
	if err != nil {
		return power, err
	}

	entries, err := r.service.Chassis()
	if err != nil {
		return power, err
	}

	if len(entries) < 1 {
		return power, ErrNoSystemEntry
	}

	for _, entry := range entries {
		e, err := entry.Power()
		if err != nil {
			return power, err
		}

		// This means we can't collect data from this device or part
		if e == nil || len(e.PowerControl) == 0 {
			continue
		}

		return float64(e.PowerControl[0].PowerConsumedWatts) / 1024, err
	}

	return power, err
}
