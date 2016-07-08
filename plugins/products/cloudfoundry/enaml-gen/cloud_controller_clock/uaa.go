package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Uaa struct {

	/*Cc - Descr: Symmetric secret used to decode uaa tokens. Used for testing. Default: <nil>
*/
	Cc *UaaCc `yaml:"cc,omitempty"`

	/*Jwt - Descr: ssl cert defined in the manifest by the UAA, required by the cc to communicate with UAA Default: 
*/
	Jwt *Jwt `yaml:"jwt,omitempty"`

	/*Url - Descr: URL of the UAA server Default: <nil>
*/
	Url interface{} `yaml:"url,omitempty"`

	/*Clients - Descr: Used to grant scope for SSO clients for service brokers Default: openid,cloud_controller_service_permissions.read
*/
	Clients *Clients `yaml:"clients,omitempty"`

}