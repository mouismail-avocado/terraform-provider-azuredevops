//go:build (all || permissions || resource_serviceendpoint_project_permissions) && (!exclude_permissions || !exclude_resource_serviceendpoint_project_permissions)
// +build all permissions resource_serviceendpoint_project_permissions
// +build !exclude_permissions !exclude_resource_serviceendpoint_project_permissions

package serviceendpoint

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/utils/datahelper"
)

// This file contains unit tests for the azuredevops_serviceendpoint_project_permissions resource.

// TestAccServiceEndpointProjectPermissions_CreateUpdate tests the creation and update of the azuredevops_serviceendpoint_project_permissions resource.
func TestAccServiceEndpointProjectPermissions_CreateUpdate(t *testing.T) {
	projectName := testutils.GenerateResourceName()
	serviceEndpointName := testutils.GenerateResourceName()
	config := hclServiceEndpointProjectPermissions(projectName, serviceEndpointName, map[string]map[string]string{
		"project_reference": {
			"project_id":            "projectID1",
			"service_endpoint_name": "service-connection-shared",
			"description":           "Service Connection Shared by Terraform - Cluster One",
		},
		"project_reference": {
			"project_id":            "projectID2",
			"service_endpoint_name": "service-connection-shared",
			"description":           "Service Connection Shared by Terraform - Cluster Two",
		},
	})
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testutils.PreCheck(t, nil) },
		Providers:    testutils.GetProviders(),
		CheckDestroy: testutils.CheckProjectDestroyed,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testutils.CheckProjectExists(projectName),
					resource.TestCheckResourceAttrSet("azuredevops_serviceendpoint_project_permissions.example-share", "serviceendpoint_id"),
					resource.TestCheckResourceAttr("azuredevops_serviceendpoint_project_permissions.example-share", "project_reference.#", "2"),
					resource.TestCheckResourceAttr("azuredevops_serviceendpoint_project_permissions.example-share", "project_reference.0.project_id", "projectID1"),
					resource.TestCheckResourceAttr("azuredevops_serviceendpoint_project_permissions.example-share", "project_reference.0.service_endpoint_name", "service-connection-shared"),
					resource.TestCheckResourceAttr("azuredevops_serviceendpoint_project_permissions.example-share", "project_reference.0.description", "Service Connection Shared by Terraform - Cluster One"),
					resource.TestCheckResourceAttr("azuredevops_serviceendpoint_project_permissions.example-share", "project_reference.1.project_id", "projectID2"),
					resource.TestCheckResourceAttr("azuredevops_serviceendpoint_project_permissions.example-share", "project_reference.1.service_endpoint_name", "service-connection-shared"),
					resource.TestCheckResourceAttr("azuredevops_serviceendpoint_project_permissions.example-share", "project_reference.1.description", "Service Connection Shared by Terraform - Cluster Two"),
				),
			},
		},
	})
}

// hclServiceEndpointProjectPermissions generates HCL describing an azuredevops_serviceendpoint_project_permissions resource
func hclServiceEndpointProjectPermissions(projectName string, serviceEndpointName string, projectReferences map[string]map[string]string) string {
	projectReferenceBlocks := ""
	for _, projectReference := range projectReferences {
		projectReferenceBlocks += fmt.Sprintf(`
		project_reference {
			project_id            = "%s"
			service_endpoint_name = "%s"
			description           = "%s"
		}
		`, projectReference["project_id"], projectReference["service_endpoint_name"], projectReference["description"])
	}

	return fmt.Sprintf(`
resource "azuredevops_project" "example" {
	name               = "%s"
	work_item_template = "Agile"
	version_control    = "Git"
	visibility         = "private"
	description        = "Managed by Terraform"
}

resource "azuredevops_serviceendpoint_azurerm" "example" {
	project_id             = azuredevops_project.example.id
	service_endpoint_name  = "%s"
	azurerm_spn_tenantid   = "00000000-0000-0000-0000-000000000000"
	azurerm_subscription_id = "00000000-0000-0000-0000-000000000000"
	azurerm_subscription_name = "Subscription Name"
	credentials {
		serviceprincipalid  = "00000000-0000-0000-0000-000000000000"
		serviceprincipalkey = "00000000000000000000000000000000"
	}
}

resource "azuredevops_serviceendpoint_project_permissions" "example-share" {
	serviceendpoint_id = azuredevops_serviceendpoint_azurerm.example.id
	%s
}
`, projectName, serviceEndpointName, projectReferenceBlocks)
}
