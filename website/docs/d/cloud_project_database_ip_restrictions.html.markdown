---
layout: "ovh"
page_title: "OVH: cloud_project_database_ip_restrictions"
sidebar_current: "docs-ovh-datasource-cloud-project-database-ip-restrictions"
description: |-
  Get the list of IP restrictions associated with a public cloud project.
---

# ovh_cloud_project_database_ip_restrictions (Data Source)

Use the list of IP restrictions associated with a public cloud project.

## Example Usage

To get the list of IP restriction on a database cluster service:
```hcl
data "ovh_cloud_project_database_ip_restrictions" "iprestrictions" {
  service_name  = "XXXXXX"
  engine        = "YYYY"
  cluster_id    = "ZZZZ"
}

output "ips" {
  value = data.ovh_cloud_project_database_ip_restrictions.iprestrictions.ips
}
```

## Argument Reference

* `service_name` - The id of the public cloud project. If omitted,
  the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.

* `engine` - The engine of the database cluster you want to list IP restrictions. To get a full list of available engine visit.
[public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).

* `cluster_id` - Cluster ID


## Attributes Reference

`id` is set to the md5 sum of the list of all IP restrictions. In addition,
the following attributes are exported:

* `ips` - The list of IP restriction of the database associated with the project.
