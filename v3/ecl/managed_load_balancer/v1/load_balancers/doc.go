/*
Package load_balancers contains functionality for working with ECL Managed Load Balancer resources.

Example to list load balancers

	listOpts := load_balancers.ListOpts{}

	allPages, err := load_balancers.List(managedLoadBalancerClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allLoadBalancers, err := load_balancers.ExtractLoadBalancers(allPages)
	if err != nil {
		panic(err)
	}

	for _, loadBalancer := range allLoadBalancers {
		fmt.Printf("%+v\n", loadBalancer)
	}

Example to create a load balancer

	reservedFixedIP1 := load_balancers.CreateOptsReservedFixedIP{
		IPAddress: "192.168.0.2",
	}
	reservedFixedIP2 := load_balancers.CreateOptsReservedFixedIP{
		IPAddress: "192.168.0.3",
	}
	reservedFixedIP3 := load_balancers.CreateOptsReservedFixedIP{
		IPAddress: "192.168.0.4",
	}
	reservedFixedIP4 := load_balancers.CreateOptsReservedFixedIP{
		IPAddress: "192.168.0.5",
	}
	interface1 := load_balancers.CreateOptsInterface{
		NetworkID: "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
		VirtualIPAddress: "192.168.0.1",
		ReservedFixedIPs: &[]load_balancers.CreateOptsReservedFixedIP{reservedFixedIP1, reservedFixedIP2, reservedFixedIP3, reservedFixedIP4},
	}
	syslogServer1 := load_balancers.CreateOptsSyslogServer{
		IPAddress: "192.168.0.6",
		Port: 514,
		Protocol: "udp",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	createOpts := load_balancers.CreateOpts{
		Name: "load_balancer",
		Description: "description",
		Tags: tags,
		PlanID: "00713021-9aea-41da-9a88-87760c08fa72",
		SyslogServers: &[]load_balancers.CreateOptsSyslogServer{syslogServer1},
		Interfaces: &[]load_balancers.CreateOptsInterface{interface1},
	}

	loadBalancer, err := load_balancers.Create(managedLoadBalancerClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", loadBalancer)

Example to show a load balancer

	showOpts := load_balancers.ShowOpts{}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	loadBalancer, err := load_balancers.Show(managedLoadBalancerClient, id, showOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", loadBalancer)

Example to update a load balancer

	name := "load_balancer"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	updateOpts := load_balancers.UpdateOpts{
		Name: &name,
		Description: &description,
		Tags: &tags,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	loadBalancer, err := load_balancers.Update(managedLoadBalancerClient, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", loadBalancer)

Example to delete a load balancer

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := load_balancers.Delete(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to perform apply-configurations action on a load balancer

	actionOpts := load_balancers.ActionOpts{
		ApplyConfigurations: true,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := load_balancers.Action(cli, id, actionOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to perform system-update action on a load balancer

	systemUpdate := load_balancers.ActionOptsSystemUpdate{
		SystemUpdateID: "31746df7-92f9-4b5e-ad05-59f6684a54eb",
	}
	actionOpts := load_balancers.ActionOpts{
		SystemUpdate: &systemUpdate,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := load_balancers.Action(cli, id, actionOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to perform apply-configurations and system-update action on a load balancer

	systemUpdate := load_balancers.ActionOptsSystemUpdate{
		SystemUpdateID: "31746df7-92f9-4b5e-ad05-59f6684a54eb",
	}
	actionOpts := load_balancers.ActionOpts{
		ApplyConfigurations: true,
		SystemUpdate: &systemUpdate,
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := load_balancers.Action(cli, id, actionOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to perform cancel-configurations action on a load balancer

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := load_balancers.CancelConfigurations(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to create staged load balancer configurations

	reservedFixedIP1 := load_balancers.CreateStagedOptsReservedFixedIP{
		IPAddress: "192.168.0.2",
	}
	reservedFixedIP2 := load_balancers.CreateStagedOptsReservedFixedIP{
		IPAddress: "192.168.0.3",
	}
	reservedFixedIP3 := load_balancers.CreateStagedOptsReservedFixedIP{
		IPAddress: "192.168.0.4",
	}
	reservedFixedIP4 := load_balancers.CreateStagedOptsReservedFixedIP{
		IPAddress: "192.168.0.5",
	}
	interface1 := load_balancers.CreateStagedOptsInterface{
		NetworkID: "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
		VirtualIPAddress: "192.168.0.1",
		ReservedFixedIPs: &[]load_balancers.CreateStagedOptsReservedFixedIP{reservedFixedIP1, reservedFixedIP2, reservedFixedIP3, reservedFixedIP4},
	}
	syslogServer1 := load_balancers.CreateStagedOptsSyslogServer{
		IPAddress: "192.168.0.6",
		Port: 514,
		Protocol: "udp",
	}
	createStagedOpts := load_balancers.CreateStagedOpts{
		SyslogServers: &[]load_balancers.CreateStagedOptsSyslogServer{syslogServer1},
		Interfaces: &[]load_balancers.CreateStagedOptsInterface{interface1},
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	loadBalancerConfigurations, err := load_balancers.CreateStaged(managedLoadBalancerClient, id, createStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", loadBalancerConfigurations)

Example to show staged load balancer configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	loadBalancerConfigurations, err := load_balancers.ShowStaged(managedLoadBalancerClient, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", loadBalancerConfigurations)

Example to update staged load balancer configurations

	reservedFixedIP1IPAddress := "192.168.0.2"
	reservedFixedIP1 := load_balancers.UpdateStagedOptsReservedFixedIP{
		IPAddress: &reservedFixedIP1IPAddress,
	}

	reservedFixedIP2IPAddress := "192.168.0.3"
	reservedFixedIP2 := load_balancers.UpdateStagedOptsReservedFixedIP{
		IPAddress: &reservedFixedIP2IPAddress,
	}

	reservedFixedIP3IPAddress := "192.168.0.4"
	reservedFixedIP3 := load_balancers.UpdateStagedOptsReservedFixedIP{
		IPAddress: &reservedFixedIP3IPAddress,
	}

	reservedFixedIP4IPAddress := "192.168.0.5"
	reservedFixedIP4 := load_balancers.UpdateStagedOptsReservedFixedIP{
		IPAddress: &reservedFixedIP4IPAddress,
	}

	interface1NetworkID := "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"
	interface1VirtualIPAddress := "192.168.0.1"
	interface1 := load_balancers.UpdateStagedOptsInterface{
		NetworkID: &interface1NetworkID,
		VirtualIPAddress: &interface1VirtualIPAddress,
		ReservedFixedIPs: &[]load_balancers.UpdateStagedOptsReservedFixedIP{reservedFixedIP1, reservedFixedIP2, reservedFixedIP3, reservedFixedIP4},
	}

	syslogServer1IPAddress := "192.168.0.6"
	syslogServer1Port := 514
	syslogServer1Protocol := "udp"
	syslogServer1 := load_balancers.UpdateStagedOptsSyslogServer{
		IPAddress: &syslogServer1IPAddress,
		Port: &syslogServer1Port,
		Protocol: &syslogServer1Protocol,
	}

	updateStagedOpts := load_balancers.UpdateStagedOpts{
		SyslogServers: &[]load_balancers.UpdateStagedOptsSyslogServer{syslogServer1},
		Interfaces: &[]load_balancers.UpdateStagedOptsInterface{interface1},
	}

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	loadBalancerConfigurations, err := load_balancers.UpdateStaged(managedLoadBalancerClient, updateStagedOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", loadBalancerConfigurations)

Example to cancel staged load balancer configurations

	id := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := load_balancers.CancelStaged(managedLoadBalancerClient, id).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package load_balancers
