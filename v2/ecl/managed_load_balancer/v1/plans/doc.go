/*
Package plans contains functionality for working with ECL Managed Load Balancer resources.

Example to list plans

	listOpts := plans.ListOpts{}

	allPages, err := plans.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allPlans, err := plans.ExtractPlans(allPages)
	if err != nil {
		panic(err)
	}

	for _, plan := range allPlans {
		fmt.Printf("%+v\n", plan)
	}

Example to show a plan

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	plan, err := plans.Show(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", plan)
*/
package plans
