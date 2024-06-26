/*
Package routes contains functionality for working with ECL Managed Load Balancer resources.

Example to list routes

	listOpts := routes.ListOpts{}

	allPages, err := routes.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allRoutes, err := routes.ExtractRoutes(allPages)
	if err != nil {
		panic(err)
	}

	for _, route := range allRoutes {
		fmt.Printf("%+v\n", route)
	}

Example to create a route


	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	createOpts := routes.CreateOpts{
		Name: "route",
		Description: "description",
		Tags: tags,
		DestinationCidr: "172.16.0.0/24",
		NextHopIPAddress: "192.168.0.254",
		LoadBalancerID: "67fea379-cff0-4191-9175-de7d6941a040",
	}

	route, err := routes.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", route)

Example to show a route

	showOpts := routes.ShowOpts{}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	route, err := routes.Show(managedLoadBalancerClient, id, showOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", route)

Example to update a route

	name := "route"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	updateOpts := routes.UpdateOpts{
		Name: &name,
		Description: &description,
		Tags: &tags,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	route, err := routes.Update(managedLoadBalancerClient, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", route)

Example to delete a route

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := routes.Delete(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to create staged route configurations

	createStagedOpts := routes.CreateStagedOpts{
		NextHopIPAddress: "192.168.0.254",
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	routeConfigurations, err := routes.CreateStaged(managedLoadBalancerClient, id, createStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", routeConfigurations)

Example to show staged route configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	routeConfigurations, err := routes.ShowStaged(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", routeConfigurations)

Example to update staged route configurations

	nextHopIPAddress := "192.168.0.254"
	updateStagedOpts := routes.UpdateStagedOpts{
		NextHopIPAddress: &nextHopIPAddress,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	routeConfigurations, err := routes.UpdateStaged(managedLoadBalancerClient, updateStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", routeConfigurations)

Example to cancel staged route configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := routes.CancelStaged(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package routes
