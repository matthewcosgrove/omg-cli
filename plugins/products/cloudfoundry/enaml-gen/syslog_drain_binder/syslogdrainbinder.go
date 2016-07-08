package syslog_drain_binder 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type SyslogDrainBinder struct {

	/*DrainUrlTtlSeconds - Descr: Time to live for drain urls in seconds Default: 60
*/
	DrainUrlTtlSeconds interface{} `yaml:"drain_url_ttl_seconds,omitempty"`

	/*PollingBatchSize - Descr: Batch size for the poll from cloud controller Default: 1000
*/
	PollingBatchSize interface{} `yaml:"polling_batch_size,omitempty"`

	/*UpdateIntervalSeconds - Descr: Interval on which to poll cloud controller in seconds Default: 15
*/
	UpdateIntervalSeconds interface{} `yaml:"update_interval_seconds,omitempty"`

	/*Debug - Descr: boolean value to turn on verbose logging for syslog_drain_binder Default: false
*/
	Debug interface{} `yaml:"debug,omitempty"`

}