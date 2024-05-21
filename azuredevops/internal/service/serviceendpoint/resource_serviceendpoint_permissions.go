// Package serviceendpoint
// Implementation of the azuredevops_serviceendpoint_project_permissions resource.
package serviceendpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

	return nil
}

func resourceServiceEndpointProjectPermissionsRead(d *schema.ResourceData, m interface{}) error {
	// Set the ID to the service endpoint ID and project ID

	return nil
}

func resourceServiceEndpointProjectPermissionsUpdate(d *schema.ResourceData, m interface{}) error {
	// Update the project reference

	return resourceServiceEndpointProjectPermissionsRead(d, m)
}

func resourceServiceEndpointProjectPermissionsDelete(d *schema.ResourceData, m interface{}) error {

	return nil
}
