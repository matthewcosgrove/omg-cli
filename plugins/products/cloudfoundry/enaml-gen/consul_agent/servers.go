package consul_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Servers struct {

	/*Lan - Descr: LAN server addresses to join on start. Default: []
*/
	Lan interface{} `yaml:"lan,omitempty"`

	/*Wan - Descr: WAN server addresses to join. Default: []
*/
	Wan interface{} `yaml:"wan,omitempty"`

}