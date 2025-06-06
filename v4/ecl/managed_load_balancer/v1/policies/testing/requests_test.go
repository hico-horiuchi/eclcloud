package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/policies"
	"github.com/nttcom/eclcloud/v4/pagination"
	"github.com/nttcom/eclcloud/v4/testhelper/client"

	th "github.com/nttcom/eclcloud/v4/testhelper"
)

const TokenID = client.TokenID

func ServiceClient() *eclcloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v1.0/"

	return sc
}

func TestListPolicies(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/policies",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, listResponse)
		})

	cli := ServiceClient()
	count := 0
	listOpts := policies.ListOpts{}

	err := policies.List(cli, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := policies.ExtractPolicies(page)
		if err != nil {
			t.Errorf("Failed to extract policies: %v", err)

			return false, err
		}

		th.CheckDeepEquals(t, listResult(), actual)

		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestCreatePolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/policies",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, createResponse)
		})

	cli := ServiceClient()
	serverNameIndication1 := policies.CreateOptsServerNameIndication{
		ServerName:    "*.example.com",
		InputType:     "fixed",
		Priority:      1,
		CertificateID: "fdfed344-e8ab-4f20-bd62-a4039453a389",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	createOpts := policies.CreateOpts{
		Name:                  "policy",
		Description:           "description",
		Tags:                  tags,
		Algorithm:             "round-robin",
		Persistence:           "cookie",
		PersistenceTimeout:    525600,
		IdleTimeout:           600,
		SorryPageUrl:          "https://example.com/sorry",
		SourceNat:             "enable",
		ServerNameIndications: &[]policies.CreateOptsServerNameIndication{serverNameIndication1},
		CertificateID:         "f57a98fe-d63e-4048-93a0-51fe163f30d7",
		HealthMonitorID:       "dd7a96d6-4e66-4666-baca-a8555f0c472c",
		ListenerID:            "68633f4f-f52a-402f-8572-b8173418904f",
		DefaultTargetGroupID:  "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
		BackupTargetGroupID:   "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
		TLSPolicyID:           "4ba79662-f2a1-41a4-a3d9-595799bbcd86",
		LoadBalancerID:        "67fea379-cff0-4191-9175-de7d6941a040",
	}

	actual, err := policies.Create(cli, createOpts).Extract()

	th.CheckDeepEquals(t, createResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/policies/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showResponse)
		})

	cli := ServiceClient()
	showOpts := policies.ShowOpts{}

	actual, err := policies.Show(cli, id, showOpts).Extract()

	th.CheckDeepEquals(t, showResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdatePolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/policies/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, updateRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, updateResponse)
		})

	cli := ServiceClient()

	name := "policy"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	updateOpts := policies.UpdateOpts{
		Name:        &name,
		Description: &description,
		Tags:        &tags,
	}

	actual, err := policies.Update(cli, id, updateOpts).Extract()

	th.CheckDeepEquals(t, updateResult(), actual)
	th.AssertNoErr(t, err)
}

func TestDeletePolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/policies/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := policies.Delete(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}

func TestCreateStagedPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/policies/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createStagedRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, createStagedResponse)
		})

	cli := ServiceClient()
	serverNameIndication1 := policies.CreateStagedOptsServerNameIndication{
		ServerName:    "*.example.com",
		InputType:     "fixed",
		Priority:      1,
		CertificateID: "fdfed344-e8ab-4f20-bd62-a4039453a389",
	}
	createStagedOpts := policies.CreateStagedOpts{
		Algorithm:             "round-robin",
		Persistence:           "cookie",
		PersistenceTimeout:    525600,
		IdleTimeout:           600,
		SorryPageUrl:          "https://example.com/sorry",
		SourceNat:             "enable",
		ServerNameIndications: &[]policies.CreateStagedOptsServerNameIndication{serverNameIndication1},
		CertificateID:         "f57a98fe-d63e-4048-93a0-51fe163f30d7",
		HealthMonitorID:       "dd7a96d6-4e66-4666-baca-a8555f0c472c",
		ListenerID:            "68633f4f-f52a-402f-8572-b8173418904f",
		DefaultTargetGroupID:  "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
		BackupTargetGroupID:   "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
		TLSPolicyID:           "4ba79662-f2a1-41a4-a3d9-595799bbcd86",
	}

	actual, err := policies.CreateStaged(cli, id, createStagedOpts).Extract()

	th.CheckDeepEquals(t, createStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowStagedPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/policies/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showStagedResponse)
		})

	cli := ServiceClient()
	actual, err := policies.ShowStaged(cli, id).Extract()

	th.CheckDeepEquals(t, showStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateStagedPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/policies/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, updateStagedRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, updateStagedResponse)
		})

	cli := ServiceClient()

	serverNameIndication1ServerName := "*.example.com"
	serverNameIndication1InputType := "fixed"
	serverNameIndication1Priority := 1
	serverNameIndication1CertificateID := "fdfed344-e8ab-4f20-bd62-a4039453a389"
	serverNameIndication1 := policies.UpdateStagedOptsServerNameIndication{
		ServerName:    &serverNameIndication1ServerName,
		InputType:     &serverNameIndication1InputType,
		Priority:      &serverNameIndication1Priority,
		CertificateID: &serverNameIndication1CertificateID,
	}

	algorithm := "round-robin"
	persistence := "cookie"
	persistenceTimeout := 525600
	idleTimeout := 600
	sorryPageUrl := "https://example.com/sorry"
	sourceNat := "enable"
	certificateID := "f57a98fe-d63e-4048-93a0-51fe163f30d7"
	healthMonitorID := "dd7a96d6-4e66-4666-baca-a8555f0c472c"
	listenerID := "68633f4f-f52a-402f-8572-b8173418904f"
	defaultTargetGroupID := "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
	backupTargetGroupID := "f1a117f1-f8df-ce07-6c8c-4bbf103059b6"
	tlsPolicyID := "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
	updateStagedOpts := policies.UpdateStagedOpts{
		Algorithm:             &algorithm,
		Persistence:           &persistence,
		PersistenceTimeout:    &persistenceTimeout,
		IdleTimeout:           &idleTimeout,
		SorryPageUrl:          &sorryPageUrl,
		SourceNat:             &sourceNat,
		ServerNameIndications: &[]policies.UpdateStagedOptsServerNameIndication{serverNameIndication1},
		CertificateID:         &certificateID,
		HealthMonitorID:       &healthMonitorID,
		ListenerID:            &listenerID,
		DefaultTargetGroupID:  &defaultTargetGroupID,
		BackupTargetGroupID:   &backupTargetGroupID,
		TLSPolicyID:           &tlsPolicyID,
	}

	actual, err := policies.UpdateStaged(cli, id, updateStagedOpts).Extract()

	th.CheckDeepEquals(t, updateStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestCancelStagedPolicy(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/policies/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := policies.CancelStaged(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}
