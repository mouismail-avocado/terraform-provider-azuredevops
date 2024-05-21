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
    // parameters
    sharedServiceEndpoint := serviceendpoint.SharedServiceEndpoint{
      ProjectReference: &serviceendpoint.ProjectReference{
        Id: &projectID,
      },
      ServiceEndpointName: &serviceEndpointName,
      Description: &description,
	}

	d.SetId(fmt.Sprintf("%s:%s", serviceEndpointID, projectID))
	return resourceServiceEndpointProjectPermissionsRead(d, m)
}

func resourceServiceEndpointProjectPermissionsRead(d *schema.ResourceData, m interface{}) error {
  // Set the ID to the service endpoint ID and project ID
  serviceEndpointID, projectID := converter.SplitID(d.Id())
  d.Set("serviceendpoint_id", serviceEndpointID)
  d.Set("project_id", projectID)

  // Get the shared service endpoint
  sharedServiceEndpoint, err := client.GetClient().ServiceEndpoint.GetSharedServiceEndpoint(serviceendpoint.GetSharedServiceEndpointArgs{
    ServiceEndpointId: &serviceEndpointID,
    ProjectId: &projectID,
  })

  if err != nil {
    return err
  }

  // Set the project reference
  projectReference := map[string]interface{}{
    "project_id": projectID,
    "service_endpoint_name": sharedServiceEndpoint.ServiceEndpointName,
    "description": sharedServiceEndpoint.Description,
  }

  d.Set("project_reference", projectReference)
  return nil
}

func resourceServiceEndpointProjectPermissionsUpdate(d *schema.ResourceData, m interface{}) error {
  // Update the project reference
  projectReference := d.Get("project_reference").([]interface{})[0].(map[string]interface{})
  serviceEndpointName := projectReference["service_endpoint_name"].(string)

  // Update the shared service endpoint
  serviceEndpointID, projectID := converter.SplitID(d.Id())
  sharedServiceEndpoint := serviceendpoint.SharedServiceEndpoint{
    ProjectReference: &serviceendpoint.ProjectReference{
      Id: &projectID,
    },
    ServiceEndpointName: &serviceEndpointName,
  }

  _, err := client.GetClient().ServiceEndpoint.UpdateSharedServiceEndpoint(serviceendpoint.UpdateSharedServiceEndpointArgs{
    ServiceEndpointId: &serviceEndpointID,
    ProjectId: &projectID,
    SharedServiceEndpoint: &sharedServiceEndpoint,
  })

  if err != nil {
    return err
  }

  return resourceServiceEndpointProjectPermissionsRead(d, m)
}

func resourceServiceEndpointProjectPermissionsDelete(d *schema.ResourceData, m interface{}) error {
	// Logic to unshare the service connection from the projects
	serviceEndpointID, projectID := converter.SplitID(d.Id())
	_, err := client.GetClient().ServiceEndpoint.DeleteSharedServiceEndpoint(serviceendpoint.DeleteSharedServiceEndpointArgs{
		ServiceEndpointId: &serviceEndpointID,
		ProjectId: &projectID,
	})

	if err != nil {
		return err
	}


	return nil
}
