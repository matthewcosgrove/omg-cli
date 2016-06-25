package hm9000 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Cc struct {

	/*InternalApiUser - Descr: Username for hm9000 API Default: internal_user
*/
	InternalApiUser interface{} `yaml:"internal_api_user,omitempty"`

	/*SrvApiUri - Descr:  Default: <nil>
*/
	SrvApiUri interface{} `yaml:"srv_api_uri,omitempty"`

	/*BulkApiPassword - Descr: Password used to access the bulk_api, health_manager uses it to connect to the cc, announced over NATS Default: <nil>
*/
	BulkApiPassword interface{} `yaml:"bulk_api_password,omitempty"`

	/*InternalApiPassword - Descr: Password for hm9000 API Default: <nil>
*/
	InternalApiPassword interface{} `yaml:"internal_api_password,omitempty"`

	/*BulkApiUser - Descr: User used to access the bulk_api, health_manager uses it to connect to the cc, announced over NATS Default: bulk_api
*/
	BulkApiUser interface{} `yaml:"bulk_api_user,omitempty"`

}