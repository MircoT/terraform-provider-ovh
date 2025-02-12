package ovh

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testAccCloudProjectDatabasePostgresqlUserConfig = `
resource "ovh_cloud_project_database" "db" {
	service_name = "%s"
	engine       = "postgresql"
	version      = "%s"
	plan         = "essential"
	nodes {
		region     = "%s"
	}
	flavor = "%s"
}

resource "ovh_cloud_project_database_postgresql_user" "user" {
	service_name = ovh_cloud_project_database.db.service_name
	cluster_id   = ovh_cloud_project_database.db.id
	name		 = "%s"
	roles = ["replication"]
}
`

func TestAccCloudProjectDatabasePostgresqlUser_basic(t *testing.T) {
	serviceName := os.Getenv("OVH_CLOUD_PROJECT_SERVICE_TEST")
	version := os.Getenv("OVH_CLOUD_PROJECT_DATABASE_POSTGRESQL_VERSION_TEST")
	if version == "" {
		version = os.Getenv("OVH_CLOUD_PROJECT_DATABASE_VERSION_TEST")
	}
	region := os.Getenv("OVH_CLOUD_PROJECT_DATABASE_REGION_TEST")
	flavor := os.Getenv("OVH_CLOUD_PROJECT_DATABASE_FLAVOR_TEST")
	name := "johndoe"

	config := fmt.Sprintf(
		testAccCloudProjectDatabasePostgresqlUserConfig,
		serviceName,
		version,
		region,
		flavor,
		name,
	)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCloudDatabaseNoEngine(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"ovh_cloud_project_database_postgresql_user.user", "created_at"),
					resource.TestCheckResourceAttrSet(
						"ovh_cloud_project_database_postgresql_user.user", "password"),
					resource.TestCheckResourceAttrSet(
						"ovh_cloud_project_database_postgresql_user.user", "roles.#"),
					resource.TestCheckResourceAttrSet(
						"ovh_cloud_project_database_postgresql_user.user", "status"),
					resource.TestCheckResourceAttr(
						"ovh_cloud_project_database_postgresql_user.user", "name", name,
					),
				),
			},
		},
	})
}
