package devices

// Fan represents percentage of fan request per blade
type FanRequest struct {
	BladePosition int
	Percentage    int
}
