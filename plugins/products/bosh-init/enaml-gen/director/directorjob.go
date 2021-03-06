package director 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DirectorJob struct {

	/*Openstack - Descr: OpenStack endpoint type (optional, by default publicURL) Default: publicURL
*/
	Openstack *Openstack `yaml:"openstack,omitempty"`

	/*Ntp - Descr: List of ntp server IPs. pool.ntp.org attempts to return IPs closest to your location, but you can still specify if needed. Default: [0.pool.ntp.org 1.pool.ntp.org]
*/
	Ntp interface{} `yaml:"ntp,omitempty"`

	/*Aws - Descr: The number of seconds before the aws cpi should timeout while waiting for response Default: 60
*/
	Aws *Aws `yaml:"aws,omitempty"`

	/*Nats - Descr: Username to connect to nats with Default: nats
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*Blobstore - Descr: Whether the simple blobstore plugin should use SSL to connect to the blobstore server Default: true
*/
	Blobstore *Blobstore `yaml:"blobstore,omitempty"`

	/*Vcd - Descr: The name of the calalog for media files Default: <nil>
*/
	Vcd *Vcd `yaml:"vcd,omitempty"`

	/*Agent - Descr: AWS KMS key ID to use for object encryption. All GET and PUT requests for an object protected by AWS KMS will fail if not made via SSL or using SigV4. Default: <nil>
*/
	Agent *Agent `yaml:"agent,omitempty"`

	/*Vcenter - Descr: Address of vCenter server used by vsphere cpi Default: <nil>
*/
	Vcenter *Vcenter `yaml:"vcenter,omitempty"`

	/*Registry - Descr: User to access the Registry Default: <nil>
*/
	Registry *Registry `yaml:"registry,omitempty"`

	/*Dns - Descr: DNS Database host Default: 127.0.0.1
*/
	Dns *Dns `yaml:"dns,omitempty"`

	/*CompiledPackageCache - Descr: Signature version of the blobstore used by s3 blobstore plugin (optional, if not provided the s3 client decides which version to use) Default: <nil>
*/
	CompiledPackageCache *CompiledPackageCache `yaml:"compiled_package_cache,omitempty"`

	/*Director - Descr: Symmetric key to verify Uaa token Default: <nil>
*/
	Director *Director `yaml:"director,omitempty"`

	/*Env - Descr: List of comma-separated hosts that should skip connecting to the proxy in the director, scheduler and workers Default: <nil>
*/
	Env *Env `yaml:"env,omitempty"`

}