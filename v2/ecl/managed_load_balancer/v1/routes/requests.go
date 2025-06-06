package routes

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

/*
List Routes
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the route attributes you want to see returned.
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

	// - CIDR of destination for the (static) route
	DestinationCidr string `q:"destination_cidr"`

	// - IP address of next hop for the (static) route
	NextHopIPAddress string `q:"next_hop_ip_address"`

	// - ID of the load balancer which the resource belongs to
	LoadBalancerID string `q:"load_balancer_id"`

	// - ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`
}

// ToRouteListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRouteListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToRouteListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of routes.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToRouteListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return RoutePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Create Route
*/

// CreateOpts represents options used to create a new route.
type CreateOpts struct {

	// - Name of the (static) route
	// - This field accepts UTF-8 characters up to 3 bytes
	Name string `json:"name,omitempty"`

	// - Description of the (static) route
	// - This field accepts UTF-8 characters up to 3 bytes
	Description string `json:"description,omitempty"`

	// - Tags of the (static) route
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags map[string]interface{} `json:"tags,omitempty"`

	// - CIDR of destination for the (static) route
	// - If you configure `destination_cidr` as default gateway, set `0.0.0.0/0`
	// - `destination_cidr` can not be changed once configured
	//   - If you want to change `destination_cidr`, recreate the (static) route again
	// - Set a unique CIDR for all (static) routes which belong to the same load balancer
	// - Set a CIDR which is not included in subnet of load balancer interfaces that the (static) route belongs to
	// - Cannot use a CIDR in the following networks
	//   - This host on this network (0.0.0.0/8)
	//   - Shared Address Space (100.64.0.0/10)
	//   - Loopback (127.0.0.0/8)
	//   - Link Local (169.254.0.0/16)
	//   - Multicast (224.0.0.0/4)
	//   - Reserved (240.0.0.0/4)
	//   - Limited Broadcast (255.255.255.255/32)
	DestinationCidr string `json:"destination_cidr"`

	// - IP address of next hop for the (static) route
	// - Set a CIDR which is included in subnet of load balancer interfaces that the (static) route belongs to
	// - Must not set a network address and a broadcast address
	NextHopIPAddress string `json:"next_hop_ip_address"`

	// - ID of the load balancer which the (static) route belongs to
	LoadBalancerID string `json:"load_balancer_id"`
}

// ToRouteCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToRouteCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "route")
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToRouteCreateMap() (map[string]interface{}, error)
}

// Create accepts a CreateOpts struct and creates a new route using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToRouteCreateMap()
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
Show Route
*/

// ShowOpts represents options used to show a route.
type ShowOpts struct {

	// - If `true` is set, `current` and `staged` are returned in response body
	Changes bool `q:"changes"`
}

// ToRouteShowQuery formats a ShowOpts into a query string.
func (opts ShowOpts) ToRouteShowQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ShowOptsBuilder allows extensions to add additional parameters to the Show request.
type ShowOptsBuilder interface {
	ToRouteShowQuery() (string, error)
}

// Show retrieves a specific route based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string, opts ShowOptsBuilder) (r ShowResult) {
	url := showURL(c, id)

	if opts != nil {
		query, _ := opts.ToRouteShowQuery()
		url += query
	}

	_, r.Err = c.Get(url, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Route Attributes
*/

// UpdateOpts represents options used to update a existing route.
type UpdateOpts struct {

	// - Name of the (static) route
	// - This field accepts UTF-8 characters up to 3 bytes
	Name *string `json:"name,omitempty"`

	// - Description of the (static) route
	// - This field accepts UTF-8 characters up to 3 bytes
	Description *string `json:"description,omitempty"`

	// - Tags of the (static) route
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags *map[string]interface{} `json:"tags,omitempty"`
}

// ToRouteUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToRouteUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "route")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToRouteUpdateMap() (map[string]interface{}, error)
}

// Update accepts a UpdateOpts struct and updates a existing route using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToRouteUpdateMap()
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
Delete Route
*/

// Delete accepts a unique ID and deletes the route associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

/*
Create Staged Route Configurations
*/

// CreateStagedOpts represents options used to create new route configurations.
type CreateStagedOpts struct {

	// - IP address of next hop for the (static) route
	// - Set a CIDR which is included in subnet of load balancer interfaces that the (static) route belongs to
	// - Must not set a network address and a broadcast address
	NextHopIPAddress string `json:"next_hop_ip_address,omitempty"`
}

// ToRouteCreateStagedMap builds a request body from CreateStagedOpts.
func (opts CreateStagedOpts) ToRouteCreateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "route")
}

// CreateStagedOptsBuilder allows extensions to add additional parameters to the CreateStaged request.
type CreateStagedOptsBuilder interface {
	ToRouteCreateStagedMap() (map[string]interface{}, error)
}

// CreateStaged accepts a CreateStagedOpts struct and creates new route configurations using the values provided.
func CreateStaged(c *eclcloud.ServiceClient, id string, opts CreateStagedOptsBuilder) (r CreateStagedResult) {
	b, err := opts.ToRouteCreateStagedMap()
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
Show Staged Route Configurations
*/

// ShowStaged retrieves specific route configurations based on its unique ID.
func ShowStaged(c *eclcloud.ServiceClient, id string) (r ShowStagedResult) {
	_, r.Err = c.Get(showStagedURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Staged Route Configurations
*/

// UpdateStagedOpts represents options used to update existing Route configurations.
type UpdateStagedOpts struct {

	// - IP address of next hop for the (static) route
	// - Set a CIDR which is included in subnet of load balancer interfaces that the (static) route belongs to
	// - Must not set a network address and a broadcast address
	NextHopIPAddress *string `json:"next_hop_ip_address,omitempty"`
}

// ToRouteUpdateStagedMap builds a request body from UpdateStagedOpts.
func (opts UpdateStagedOpts) ToRouteUpdateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "route")
}

// UpdateStagedOptsBuilder allows extensions to add additional parameters to the UpdateStaged request.
type UpdateStagedOptsBuilder interface {
	ToRouteUpdateStagedMap() (map[string]interface{}, error)
}

// UpdateStaged accepts a UpdateStagedOpts struct and updates existing Route configurations using the values provided.
func UpdateStaged(c *eclcloud.ServiceClient, id string, opts UpdateStagedOptsBuilder) (r UpdateStagedResult) {
	b, err := opts.ToRouteUpdateStagedMap()
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
Cancel Staged Route Configurations
*/

// CancelStaged accepts a unique ID and deletes route configurations associated with it.
func CancelStaged(c *eclcloud.ServiceClient, id string) (r CancelStagedResult) {
	_, r.Err = c.Delete(cancelStagedURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
