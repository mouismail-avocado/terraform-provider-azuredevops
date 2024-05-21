// Implementation of the azuredevops_serviceendpoint_project_permissions resource.
package serviceendpoint

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/serviceendpoint"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/client"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/utils/converter"
)

func ResourceServiceEndpointProjectPermissions() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceEndpointProjectPermissionsCreate,
		Read:   resourceServiceEndpointProjectPermissionsRead,
		Update: resourceServiceEndpointProjectPermissionsUpdate,
		Delete: resourceServiceEndpointProjectPermissionsDelete,

		Schema: map[string]*schema.Schema{
			"serviceendpoint_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_reference": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service_endpoint_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
					},
				},
			},
		},
	}
}

func resourceServiceEndpointProjectPermissionsCreate(d *schema.ResourceData, m interface{}) error {
	clients := m.(*client.AggregatedClient)
	serviceEndpointID := d.Get("serviceendpoint_id").(string)
	projectReferences := d.Get("project_reference").([]interface{})

	for _, projectReference := range projectReferences {
		projectRef := projectReference.(map[string]interface{})
		projectID := projectRef["project_id"].(string)
		serviceEndpointName := projectRef["service_endpoint_name"].(string)
		description := projectRef["description"].(string)

		// Logic to share service connection with specified projects
		// including handling optional service_endpoint_name and description
	}

	d.SetId(fmt.Sprintf("%s:%s", serviceEndpointID, projectID))
	return resourceServiceEndpointProjectPermissionsRead(d, m)
}

func resourceServiceEndpointProjectPermissionsRead(d *schema.ResourceData, m interface{}) error {
	// Logic to read the shared service connection permissions
	return nil
}

func resourceServiceEndpointProjectPermissionsUpdate(d *schema.ResourceData, m interface{}) error {
	// Logic to update the shared service connection permissions
	return resourceServiceEndpointProjectPermissionsRead(d, m)
}

func resourceServiceEndpointProjectPermissionsDelete(d *schema.ResourceData, m interface{}) error {
	// Logic to unshare the service connection from the projects
	return nil
}
