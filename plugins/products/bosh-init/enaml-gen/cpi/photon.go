package cpi 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Photon struct {

	/*Project - Descr: Photon project Default: 
*/
	Project interface{} `yaml:"project,omitempty"`

	/*Password - Descr: Photon user's password Default: 
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Target - Descr: Photon API-FE target Default: 
*/
	Target interface{} `yaml:"target,omitempty"`

	/*IgnoreCert - Descr: Whether to ignore certs check Default: true
*/
	IgnoreCert interface{} `yaml:"ignore_cert,omitempty"`

	/*User - Descr: Photon user Default: 
*/
	User interface{} `yaml:"user,omitempty"`

}