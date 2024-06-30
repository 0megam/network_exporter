package ping

import "time"

const defaultTimeout = 5 * time.Second
const defaultPackerSize = 56
const defaultCount = 10
const defaultTTL = 128
const defaultDontFragment = false

// PingResult Calculated results
type PingResult struct {
	Success              bool          `json:"success"`
	DestAddr             string        `json:"dest_address"`
	DestIp               string        `json:"dest_ip"`
	DropRate             float64       `json:"drop_rate"`
	SumTime              time.Duration `json:"sum"`
	BestTime             time.Duration `json:"best"`
	AvgTime              time.Duration `json:"avg"`
	WorstTime            time.Duration `json:"worst"`
	SquaredDeviationTime time.Duration `json:"sd"`
	UncorrectedSDTime    time.Duration `json:"usd"`
	CorrectedSDTime      time.Duration `json:"csd"`
	RangeTime            time.Duration `json:"range"`
	SntSummary           int           `json:"snt_summary"`
	SntFailSummary       int           `json:"snt_fail_summary"`
	SntTimeSummary       time.Duration `json:"snt_time_summary"`
}

// PingReturn ICMP Response
type PingReturn struct {
	success   bool
	succSum   int
	allTime   []time.Duration
	sumTime   time.Duration
	bestTime  time.Duration
	avgTime   time.Duration
	worstTime time.Duration
}

// PingOptions ICMP Options
type PingOptions struct {
	count      int
	timeout    time.Duration
	packetSize int
	dontFragment bool
}

// Count Getter
func (options *PingOptions) Count() int {
	if options.count == 0 {
		options.count = defaultCount
	}
	return options.count
}

// SetCount Setter
func (options *PingOptions) SetCount(count int) {
	options.count = count
}

// Timeout Getter
func (options *PingOptions) Timeout() time.Duration {
	if options.timeout == 0 {
		options.timeout = defaultTimeout
	}
	return options.timeout
}

// SetTimeout Setter
func (options *PingOptions) SetTimeout(timeout time.Duration) {
	options.timeout = timeout
}

// PacketSize Getter
func (options *PingOptions) PacketSize() int {
	if options.packetSize == 0 {
		options.packetSize = defaultPackerSize
	}
	return options.packetSize
}
// DontFragment Getter
func (options *PingOptions) DontFragment() bool {
	if options.dontFragment == false {
		options.dontFragment = defaultDontFragment
	}
	return options.dontFragment
}

// SetPacketSize Setter
func (options *PingOptions) SetPacketSize(packetSize int) {
	options.packetSize = packetSize
}

// SetDontFragment Setter
func (options *PingOptions) SetDontFragment(dontFragment bool) {
	options.dontFragment = dontFragment
}
