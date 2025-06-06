package plans

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

/*
List Plans
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the plan attributes you want to see returned.
type ListOpts struct {

	// - ID of the resource
	ID string `q:"id"`

	// - Name of the resource
	// - This field accepts single-byte characters only
	Name string `q:"name"`

	// - Description of the resource
	// - This field accepts single-byte characters only
	Description string `q:"description"`

	// - Bandwidth of the plan
	Bandwidth string `q:"bandwidth"`

	// - Redundancy of the plan
	Redundancy string `q:"redundancy"`

	// - Maximum number of interfaces for the plan
	MaxNumberOfInterfaces int `q:"max_number_of_interfaces"`

	// - Maximum number of health monitors for the plan
	MaxNumberOfHealthMonitors int `q:"max_number_of_health_monitors"`

	// - Maximum number of listeners for the plan
	MaxNumberOfListeners int `q:"max_number_of_listeners"`

	// - Maximum number of policies for the plan
	MaxNumberOfPolicies int `q:"max_number_of_policies"`

	// - Maximum number of routes for the plan
	MaxNumberOfRoutes int `q:"max_number_of_routes"`

	// - Maximum number of target groups for the plan
	MaxNumberOfTargetGroups int `q:"max_number_of_target_groups"`

	// - Maximum number of members for the target group of the plan
	MaxNumberOfMembers int `q:"max_number_of_members"`

	// - Maximum number of rules for the policy of the plan
	MaxNumberOfRules int `q:"max_number_of_rules"`

	// - Maximum number of conditions in the rule of the plan
	MaxNumberOfConditions int `q:"max_number_of_conditions"`

	// - Maximum number of Server Name Indications (SNIs) for the policy of the plan
	MaxNumberOfServerNameInidications int `q:"max_number_of_server_name_inidications"`

	// - Whether a new load balancer can be created with this plan
	Enabled bool `q:"enabled"`
}

// ToPlanListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPlanListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToPlanListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of plans.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToPlanListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return PlanPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Show Plan
*/

// Show retrieves a specific plan based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string) (r ShowResult) {
	_, r.Err = c.Get(showURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
