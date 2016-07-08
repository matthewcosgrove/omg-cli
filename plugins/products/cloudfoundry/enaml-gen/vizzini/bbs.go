package vizzini 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Bbs struct {

	/*ClientCert - Descr: PEM-encoded client certificate Default: <nil>
*/
	ClientCert interface{} `yaml:"client_cert,omitempty"`

	/*ClientKey - Descr: PEM-encoded client key Default: <nil>
*/
	ClientKey interface{} `yaml:"client_key,omitempty"`

	/*ApiLocation - Descr: The address of the BBS Default: bbs.service.cf.internal:8889
*/
	ApiLocation interface{} `yaml:"api_location,omitempty"`

	/*RequireSsl - Descr: enable ssl for all communication with the bbs Default: true
*/
	RequireSsl interface{} `yaml:"require_ssl,omitempty"`

}