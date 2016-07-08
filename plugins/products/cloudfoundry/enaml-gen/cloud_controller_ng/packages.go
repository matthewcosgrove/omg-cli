package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Packages struct {

	/*MaxValidPackagesStored - Descr: Number of recent, valid packages stored per app (not including package for current droplet) Default: 5
*/
	MaxValidPackagesStored interface{} `yaml:"max_valid_packages_stored,omitempty"`

	/*Cdn - Descr: URI for a CDN to used for app package downloads Default: 
*/
	Cdn *PackagesCdn `yaml:"cdn,omitempty"`

	/*MaxPackageSize - Descr: Maximum size of application package Default: 1073741824
*/
	MaxPackageSize interface{} `yaml:"max_package_size,omitempty"`

	/*WebdavConfig - Descr: The location of the webdav server eg: https://blobstore.internal Default: https://blobstore.service.cf.internal
*/
	WebdavConfig *PackagesWebdavConfig `yaml:"webdav_config,omitempty"`

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

	/*AppPackageDirectoryKey - Descr: Directory (bucket) used store app packages.  It does not have be pre-created. Default: cc-packages
*/
	AppPackageDirectoryKey interface{} `yaml:"app_package_directory_key,omitempty"`

}