package director 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Vcd struct {

	/*User - Descr: The user name of the target vCloud Director Default: <nil>
*/
	User interface{} `yaml:"user,omitempty"`

	/*Password - Descr: The password of the target vCloud Director Default: <nil>
*/
	Password interface{} `yaml:"password,omitempty"`

	/*Url - Descr: The endpoint of the target vCloud Director Default: <nil>
*/
	Url interface{} `yaml:"url,omitempty"`

	/*Entities - Descr: The name of the calalog for media files Default: <nil>
*/
	Entities *Entities `yaml:"entities,omitempty"`

}