package health_monitors

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

/*
List Health Monitors
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the health monitor attributes you want to see returned.
type ListOpts struct {

	// - ID of the resource
	ID string `q:"id"`

	// - Name of the resource
	// - This field accepts UTF-8 characters up to 3 bytes
	Name string `q:"name"`

	// - Description of the resource
	// - This field accepts UTF-8 characters up to 3 bytes
	Description string `q:"description"`

	// - Configuration status of the resource
	ConfigurationStatus string `q:"configuration_status"`

	// - Operation status of the resource
	OperationStatus string `q:"operation_status"`

	// - Port number of the resource for healthchecking or listening
	Port int `q:"port"`

	// - Protocol of the resource for healthchecking or listening
	Protocol string `q:"protocol"`

	// - Interval of healthchecking (in seconds)
	Interval int `q:"interval"`

	// - Retry count of healthchecking
	Retry int `q:"retry"`

	// - Timeout of healthchecking (in seconds)
	Timeout int `q:"timeout"`

	// - URL path of healthchecking
	// - Must be started with `"/"`
	Path string `q:"path"`

	// - HTTP status codes expected in healthchecking
	// - Format: `"xxx"` or `"xxx-xxx"` ( `xxx` between [100, 599])
	HttpStatusCode string `q:"http_status_code"`

	// - ID of the load balancer which the resource belongs to
	LoadBalancerID string `q:"load_balancer_id"`

	// - ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`
}

// ToHealthMonitorListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToHealthMonitorListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToHealthMonitorListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of health monitors.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToHealthMonitorListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return HealthMonitorPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Create Health Monitor
*/

// CreateOpts represents options used to create a new health monitor.
type CreateOpts struct {

	// - Name of the health monitor
	// - This field accepts UTF-8 characters up to 3 bytes
	Name string `json:"name,omitempty"`

	// - Description of the health monitor
	// - This field accepts UTF-8 characters up to 3 bytes
	Description string `json:"description,omitempty"`

	// - Tags of the health monitor
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags map[string]interface{} `json:"tags,omitempty"`

	// - Port number of the health monitor for healthchecking
	// - If 'protocol' is 'icmp', value must be set `0`
	Port int `json:"port"`

	// - Protocol of the health monitor for healthchecking
	Protocol string `json:"protocol"`

	// - Interval of healthchecking (in seconds)
	Interval int `json:"interval,omitempty"`

	// - Retry count of healthchecking
	// - Initial monitoring is not included
	// - Retry is executed at the interval set in `interval`
	Retry int `json:"retry,omitempty"`

	// - Timeout of healthchecking (in seconds)
	// - Value must be less than or equal to `interval`
	Timeout int `json:"timeout,omitempty"`

	// - URL path of healthchecking
	// - If `protocol` is `"http"` or `"https"`, URL path can be set
	//   - If `protocol` is neither `"http"` nor `"https"`, URL path must not be set
	// - Must be started with /
	Path string `json:"path,omitempty"`

	// - HTTP status codes expected in healthchecking
	// - If `protocol` is `"http"` or `"https"`, HTTP status code (or range) can be set
	//   - If `protocol` is neither `"http"` nor `"https"`, HTTP status code (or range) must not be set
	// - Format: `"xxx"` or `"xxx-xxx"` ( `xxx` between [100, 599])
	HttpStatusCode string `json:"http_status_code,omitempty"`

	// - ID of the load balancer which the health monitor belongs to
	LoadBalancerID string `json:"load_balancer_id"`
}

// ToHealthMonitorCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToHealthMonitorCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "health_monitor")
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToHealthMonitorCreateMap() (map[string]interface{}, error)
}

// Create accepts a CreateOpts struct and creates a new health monitor using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToHealthMonitorCreateMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Post(createURL(c), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Show Health Monitor
*/

// ShowOpts represents options used to show a health monitor.
type ShowOpts struct {

	// - If `true` is set, `current` and `staged` are returned in response body
	Changes bool `q:"changes"`
}

// ToHealthMonitorShowQuery formats a ShowOpts into a query string.
func (opts ShowOpts) ToHealthMonitorShowQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ShowOptsBuilder allows extensions to add additional parameters to the Show request.
type ShowOptsBuilder interface {
	ToHealthMonitorShowQuery() (string, error)
}

// Show retrieves a specific health monitor based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string, opts ShowOptsBuilder) (r ShowResult) {
	url := showURL(c, id)

	if opts != nil {
		query, _ := opts.ToHealthMonitorShowQuery()
		url += query
	}

	_, r.Err = c.Get(url, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Health Monitor Attributes
*/

// UpdateOpts represents options used to update a existing health monitor.
type UpdateOpts struct {

	// - Name of the health monitor
	// - This field accepts UTF-8 characters up to 3 bytes
	Name *string `json:"name,omitempty"`

	// - Description of the health monitor
	// - This field accepts UTF-8 characters up to 3 bytes
	Description *string `json:"description,omitempty"`

	// - Tags of the health monitor
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags *map[string]interface{} `json:"tags,omitempty"`
}

// ToHealthMonitorUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToHealthMonitorUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "health_monitor")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToHealthMonitorUpdateMap() (map[string]interface{}, error)
}

// Update accepts a UpdateOpts struct and updates a existing health monitor using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToHealthMonitorUpdateMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Patch(updateURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Delete Health Monitor
*/

// Delete accepts a unique ID and deletes the health monitor associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

/*
Create Staged Health Monitor Configurations
*/

// CreateStagedOpts represents options used to create new health monitor configurations.
type CreateStagedOpts struct {

	// - Port number of the health monitor for healthchecking
	// - If 'protocol' is 'icmp', value must be set `0`
	Port int `json:"port,omitempty"`

	// - Protocol of the health monitor for healthchecking
	Protocol string `json:"protocol,omitempty"`

	// - Interval of healthchecking (in seconds)
	Interval int `json:"interval,omitempty"`

	// - Retry count of healthchecking
	// - Initial monitoring is not included
	// - Retry is executed at the interval set in `interval`
	Retry int `json:"retry,omitempty"`

	// - Timeout of healthchecking (in seconds)
	// - Value must be less than or equal to `interval`
	Timeout int `json:"timeout,omitempty"`

	// - URL path of healthchecking
	// - If `protocol` is `"http"` or `"https"`, URL path can be set
	//   - If `protocol` is neither `"http"` nor `"https"`, URL path must not be set
	// - Must be started with /
	Path string `json:"path,omitempty"`

	// - HTTP status codes expected in healthchecking
	// - If `protocol` is `"http"` or `"https"`, HTTP status code (or range) can be set
	//   - If `protocol` is neither `"http"` nor `"https"`, HTTP status code (or range) must not be set
	// - Format: `"xxx"` or `"xxx-xxx"` ( `xxx` between [100, 599])
	HttpStatusCode string `json:"http_status_code,omitempty"`
}

// ToHealthMonitorCreateStagedMap builds a request body from CreateStagedOpts.
func (opts CreateStagedOpts) ToHealthMonitorCreateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "health_monitor")
}

// CreateStagedOptsBuilder allows extensions to add additional parameters to the CreateStaged request.
type CreateStagedOptsBuilder interface {
	ToHealthMonitorCreateStagedMap() (map[string]interface{}, error)
}

// CreateStaged accepts a CreateStagedOpts struct and creates new health monitor configurations using the values provided.
func CreateStaged(c *eclcloud.ServiceClient, id string, opts CreateStagedOptsBuilder) (r CreateStagedResult) {
	b, err := opts.ToHealthMonitorCreateStagedMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Post(createStagedURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Show Staged Health Monitor Configurations
*/

// ShowStaged retrieves specific health monitor configurations based on its unique ID.
func ShowStaged(c *eclcloud.ServiceClient, id string) (r ShowStagedResult) {
	_, r.Err = c.Get(showStagedURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Staged Health Monitor Configurations
*/

// UpdateStagedOpts represents options used to update existing Health Monitor configurations.
type UpdateStagedOpts struct {

	// - Port number of the health monitor for healthchecking
	// - If 'protocol' is 'icmp', value must be set `0`
	Port *int `json:"port,omitempty"`

	// - Protocol of the health monitor for healthchecking
	Protocol *string `json:"protocol,omitempty"`

	// - Interval of healthchecking (in seconds)
	Interval *int `json:"interval,omitempty"`

	// - Retry count of healthchecking
	// - Initial monitoring is not included
	// - Retry is executed at the interval set in `interval`
	Retry *int `json:"retry,omitempty"`

	// - Timeout of healthchecking (in seconds)
	// - Value must be less than or equal to `interval`
	Timeout *int `json:"timeout,omitempty"`

	// - URL path of healthchecking
	// - If `protocol` is `"http"` or `"https"`, URL path can be set
	//   - If `protocol` is neither `"http"` nor `"https"`, URL path must not be set
	// - Must be started with /
	Path *string `json:"path,omitempty"`

	// - HTTP status codes expected in healthchecking
	// - If `protocol` is `"http"` or `"https"`, HTTP status code (or range) can be set
	//   - If `protocol` is neither `"http"` nor `"https"`, HTTP status code (or range) must not be set
	// - Format: `"xxx"` or `"xxx-xxx"` ( `xxx` between [100, 599])
	HttpStatusCode *string `json:"http_status_code,omitempty"`
}

// ToHealthMonitorUpdateStagedMap builds a request body from UpdateStagedOpts.
func (opts UpdateStagedOpts) ToHealthMonitorUpdateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "health_monitor")
}

// UpdateStagedOptsBuilder allows extensions to add additional parameters to the UpdateStaged request.
type UpdateStagedOptsBuilder interface {
	ToHealthMonitorUpdateStagedMap() (map[string]interface{}, error)
}

// UpdateStaged accepts a UpdateStagedOpts struct and updates existing Health Monitor configurations using the values provided.
func UpdateStaged(c *eclcloud.ServiceClient, id string, opts UpdateStagedOptsBuilder) (r UpdateStagedResult) {
	b, err := opts.ToHealthMonitorUpdateStagedMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Patch(updateStagedURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Cancel Staged Health Monitor Configurations
*/

// CancelStaged accepts a unique ID and deletes health monitor configurations associated with it.
func CancelStaged(c *eclcloud.ServiceClient, id string) (r CancelStagedResult) {
	_, r.Err = c.Delete(cancelStagedURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
