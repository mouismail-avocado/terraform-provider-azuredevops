---
layout: "azuredevops"
page_title: "AzureDevops: azuredevops_serviceendpoint_project_permissions"
description: |-
  Manages permissions for sharing a AzureDevOps Service Endpoint with multiple projects.
---

# azuredevops_serviceendpoint_project_permissions

Manages permissions for sharing a Service Endpoint with multiple projects.

~> **Note** Permissions can be assigned to group principals and not to single user principals.

## Permission levels

Permission for Service Endpoints within Azure DevOps can be applied on two different levels.
Those levels are reflected by specifying (or omitting) values for the arguments `project_id` and `serviceendpoint_id`.

## Example Usage

```hcl
resource "azuredevops_project" "example" {
  name               = "Example Project"
  work_item_template = "Agile"
  version_control    = "Git"
  visibility         = "private"
  description        = "Managed by Terraform"
}

data "azuredevops_group" "example-readers" {
  project_id = azuredevops_project.example.id
  name       = "Readers"
}

resource "azuredevops_serviceendpoint_project_permissions" "example-share" {
  serviceendpoint_id = azuredevops_serviceendpoint_azurerm.example.id

  project_reference = {
    project_id            = azuredevops_project.example_one.id
    service_endpoint_name = "service-connection-shared"
    description           = "Service Connection Shared by Terraform - Cluster One"
  }

  project_reference = {
    project_id            = azuredevops_project.example_two.id
    service_endpoint_name = "service-connection-shared"
    description           = "Service Connection Shared by Terraform - Cluster Two"
  }
}
```

## Argument Reference

The following arguments are supported:

* `serviceendpoint_id` - (Required) The ID of the service endpoint to share.
* `project_reference` - (Required) A list of `project_reference` blocks as defined below. Objects describing which projects the service connection will be shared with.

An `project_reference` block supports the following:

* `project_id` - (Required) Project id which service endpoint will be shared.
* `service_endpoint_name` - (Optional) Name for service connection in the shared project. Default keep the same name.
* `description` - (Optional) Description for service connection in the shared project. Default keep the same description.

## Relevant Links

* [Azure DevOps Services REST API 7.1 - Endpoints](https://learn.microsoft.com/en-us/rest/api/azure/devops/serviceendpoint/endpoints/share-service-endpoint?view=azure-devops-rest-7.1&tabs=HTTP)

## Import

The resource does not support import.

## PAT Permissions Required

- **Project & Team**: vso.security_manage - Grants the ability to read, write, and manage security permissions.
