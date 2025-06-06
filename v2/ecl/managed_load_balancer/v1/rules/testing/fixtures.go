package testing

import (
	"encoding/json"
	"fmt"

	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/rules"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "rules": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "rule",
            "description": "description",
            "tags": {
                "key": "value"
            },
            "configuration_status": "ACTIVE",
            "operation_status": "COMPLETE",
            "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
            "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
            "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
            "priority": 1,
            "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
            "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
            "conditions": {
                "path_patterns": [
                    "^/statics/"
                ]
            }
        }
    ]
}`)

func listResult() []rules.Rule {
	var rule1 rules.Rule

	condition1 := rules.ConditionInResponse{
		PathPatterns: []string{"^/statics/"},
	}

	var tags1 map[string]interface{}
	tags1Json := `{"key":"value"}`
	err := json.Unmarshal([]byte(tags1Json), &tags1)
	if err != nil {
		panic(err)
	}

	rule1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	rule1.Name = "rule"
	rule1.Description = "description"
	rule1.Tags = tags1
	rule1.ConfigurationStatus = "ACTIVE"
	rule1.OperationStatus = "COMPLETE"
	rule1.PolicyID = "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"
	rule1.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	rule1.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	rule1.Priority = 1
	rule1.TargetGroupID = "29527a3c-9e5d-48b7-868f-6442c7d21a95"
	rule1.BackupTargetGroupID = "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"
	rule1.Conditions = condition1

	return []rules.Rule{rule1}
}

var createRequest = fmt.Sprintf(`
{
    "rule": {
        "name": "rule",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
            "path_patterns": [
                "^/statics/"
            ]
        }
    }
}`)

var createResponse = fmt.Sprintf(`
{
    "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": null,
        "target_group_id": null,
        "backup_target_group_id": null,
        "conditions": null
    }
}`)

func createResult() *rules.Rule {
	var rule rules.Rule

	var condition rules.ConditionInResponse

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	rule.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	rule.Name = "rule"
	rule.Description = "description"
	rule.Tags = tags
	rule.ConfigurationStatus = "CREATE_STAGED"
	rule.OperationStatus = "NONE"
	rule.PolicyID = "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"
	rule.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	rule.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	rule.Priority = 0
	rule.TargetGroupID = ""
	rule.BackupTargetGroupID = ""
	rule.Conditions = condition

	return &rule
}

var showResponse = fmt.Sprintf(`
{
    "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
            "path_patterns": [
                "^/statics/"
            ]
        },
        "current": {
            "priority": 1,
            "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
            "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
            "conditions": {
                "path_patterns": [
                    "^/statics/"
                ]
            }
        },
        "staged": null
    }
}`)

func showResult() *rules.Rule {
	var rule rules.Rule

	condition := rules.ConditionInResponse{
		PathPatterns: []string{"^/statics/"},
	}
	var staged rules.ConfigurationInResponse
	current := rules.ConfigurationInResponse{
		Priority:            1,
		TargetGroupID:       "29527a3c-9e5d-48b7-868f-6442c7d21a95",
		BackupTargetGroupID: "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
		Conditions:          condition,
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	rule.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	rule.Name = "rule"
	rule.Description = "description"
	rule.Tags = tags
	rule.ConfigurationStatus = "ACTIVE"
	rule.OperationStatus = "COMPLETE"
	rule.PolicyID = "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"
	rule.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	rule.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	rule.Priority = 1
	rule.TargetGroupID = "29527a3c-9e5d-48b7-868f-6442c7d21a95"
	rule.BackupTargetGroupID = "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"
	rule.Conditions = condition
	rule.Current = current
	rule.Staged = staged

	return &rule
}

var updateRequest = fmt.Sprintf(`
{
    "rule": {
        "name": "rule",
        "description": "description",
        "tags": {
            "key": "value"
        }
    }
}`)

var updateResponse = fmt.Sprintf(`
{
    "rule": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "rule",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "policy_id": "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "priority": null,
        "target_group_id": null,
        "backup_target_group_id": null,
        "conditions": null
    }
}`)

func updateResult() *rules.Rule {
	var rule rules.Rule

	var condition rules.ConditionInResponse

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	rule.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	rule.Name = "rule"
	rule.Description = "description"
	rule.Tags = tags
	rule.ConfigurationStatus = "CREATE_STAGED"
	rule.OperationStatus = "NONE"
	rule.PolicyID = "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4"
	rule.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	rule.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	rule.Priority = 0
	rule.TargetGroupID = ""
	rule.BackupTargetGroupID = ""
	rule.Conditions = condition

	return &rule
}

var createStagedRequest = fmt.Sprintf(`
{
    "rule": {
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
            "path_patterns": [
                "^/statics/"
            ]
        }
    }
}`)

var createStagedResponse = fmt.Sprintf(`
{
    "rule": {
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
            "path_patterns": [
                "^/statics/"
            ]
        }
    }
}`)

func createStagedResult() *rules.Rule {
	var rule rules.Rule

	condition := rules.ConditionInResponse{
		PathPatterns: []string{"^/statics/"},
	}

	rule.Priority = 1
	rule.TargetGroupID = "29527a3c-9e5d-48b7-868f-6442c7d21a95"
	rule.BackupTargetGroupID = "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"
	rule.Conditions = condition

	return &rule
}

var showStagedResponse = fmt.Sprintf(`
{
    "rule": {
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
            "path_patterns": [
                "^/statics/"
            ]
        }
    }
}`)

func showStagedResult() *rules.Rule {
	var rule rules.Rule

	condition := rules.ConditionInResponse{
		PathPatterns: []string{"^/statics/"},
	}

	rule.Priority = 1
	rule.TargetGroupID = "29527a3c-9e5d-48b7-868f-6442c7d21a95"
	rule.BackupTargetGroupID = "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"
	rule.Conditions = condition

	return &rule
}

var updateStagedRequest = fmt.Sprintf(`
{
    "rule": {
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
            "path_patterns": [
                "^/statics/"
            ]
        }
    }
}`)

var updateStagedResponse = fmt.Sprintf(`
{
    "rule": {
        "priority": 1,
        "target_group_id": "29527a3c-9e5d-48b7-868f-6442c7d21a95",
        "backup_target_group_id": "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52",
        "conditions": {
            "path_patterns": [
                "^/statics/"
            ]
        }
    }
}`)

func updateStagedResult() *rules.Rule {
	var rule rules.Rule

	condition := rules.ConditionInResponse{
		PathPatterns: []string{"^/statics/"},
	}

	rule.Priority = 1
	rule.TargetGroupID = "29527a3c-9e5d-48b7-868f-6442c7d21a95"
	rule.BackupTargetGroupID = "dfa2dbb6-e2f8-4a9d-a8c1-e1a578ea0a52"
	rule.Conditions = condition

	return &rule
}
