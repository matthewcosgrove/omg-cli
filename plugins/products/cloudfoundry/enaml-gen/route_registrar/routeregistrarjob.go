package route_registrar 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type RouteRegistrarJob struct {

	/*Nats - Descr: TCP port of NATS servers Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*RouteRegistrar - Descr: * Array of hashes determining which routes will be registered.
* Each hash should have 'port', 'uris', 'registration_interval'
  and 'name' keys.
* 'registration_interval' is the delay between
  routing updates. It must be a time duration represented as a string
  (e.g. "10s").
  It must parse to a positive time duration i.e. "-5s" is not permitted.
* Additionally, the 'tags' and 'health_check' keys are optional.
* 'uris' is an array of URIs to register for the 'port'.
* 'tags' are included in metrics that gorouter emits to support filtering.
* 'health_check' is a hash which should have 'name' and 'script_path'.
* 'health_check.timeout' is optional.
  If the health_check timeout is not provided, it defaults to half of the
  value of `registration_interval`.
  If it is provided it must be a time duration represented as a string (e.g. "10s"),
  and less than the value of `registration_interval`.
  It must parse to a positive time duration i.e. "-5s" is not permitted.
* if the healthcheck is not configured, the route is continually registered.
* if the healthcheck script exits with success, the route is registered.
* if the healthcheck script exits with error, the route is unregistered.
* if a timeout is configured, the healthcheck script must exit within the timeout,
  otherwise it is terminated (with `SIGKILL`) and the route is unregistered.
 Default: <nil>
*/
	RouteRegistrar *RouteRegistrar `yaml:"route_registrar,omitempty"`

}