package cassandra_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acctest3 "github.com/aiven/terraform-provider-aiven/internal/acctest"
)

func TestAccAiven_cassandra(t *testing.T) {
	resourceName := "aiven_cassandra.bar"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest3.TestAccPreCheck(t) },
		ProviderFactories: acctest3.TestAccProviderFactories,
		CheckDestroy:      acctest3.TestAccCheckAivenServiceResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCassandraResource(rName),
				Check: resource.ComposeTestCheckFunc(
					acctest3.TestAccCheckAivenServiceCommonAttributes("data.aiven_cassandra.common"),
					testAccCheckAivenServiceCassandraAttributes("data.aiven_cassandra.common"),
					resource.TestCheckResourceAttr(resourceName, "service_name", fmt.Sprintf("test-acc-sr-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
					resource.TestCheckResourceAttr(resourceName, "project", os.Getenv("AIVEN_PROJECT_NAME")),
					resource.TestCheckResourceAttr(resourceName, "service_type", "cassandra"),
					resource.TestCheckResourceAttr(resourceName, "cloud_name", "google-europe-west1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window_dow", "monday"),
					resource.TestCheckResourceAttr(resourceName, "maintenance_window_time", "10:00:00"),
					resource.TestCheckResourceAttr(resourceName, "state", "RUNNING"),
					resource.TestCheckResourceAttr(resourceName, "termination_protection", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "service_username"),
					resource.TestCheckResourceAttrSet(resourceName, "service_password"),
					resource.TestCheckResourceAttrSet(resourceName, "service_host"),
					resource.TestCheckResourceAttrSet(resourceName, "service_port"),
					resource.TestCheckResourceAttrSet(resourceName, "service_uri"),
				),
			},
			{
				Config:             testAccCassandraDoubleTagKeyResource(rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				ExpectError:        regexp.MustCompile("tag keys should be unique"),
			},
		},
	})
}

func testAccCassandraResource(name string) string {
	return fmt.Sprintf(`
data "aiven_project" "foo" {
  project = "%s"
}

resource "aiven_cassandra" "bar" {
  project                 = data.aiven_project.foo.project
  cloud_name              = "google-europe-west1"
  plan                    = "startup-4"
  service_name            = "test-acc-sr-%s"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"

  cassandra_user_config {
    migrate_sstableloader = true

    public_access {
      prometheus = true
    }
  }

  tag {
    key   = "test"
    value = "val"
  }
}

data "aiven_cassandra" "common" {
  service_name = aiven_cassandra.bar.service_name
  project      = data.aiven_project.foo.project

  depends_on = [aiven_cassandra.bar]
}`, os.Getenv("AIVEN_PROJECT_NAME"), name)
}

func testAccCassandraDoubleTagKeyResource(name string) string {
	return fmt.Sprintf(`
data "aiven_project" "foo" {
  project = "%s"
}

resource "aiven_cassandra" "bar" {
  project                 = data.aiven_project.foo.project
  cloud_name              = "google-europe-west1"
  plan                    = "startup-4"
  service_name            = "test-acc-sr-%s"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"

  cassandra_user_config {
    migrate_sstableloader = true

    public_access {
      prometheus = true
    }
  }

  tag {
    key   = "test"
    value = "val"
  }
  tag {
    key   = "test"
    value = "val2"
  }
}

data "aiven_cassandra" "common" {
  service_name = aiven_cassandra.bar.service_name
  project      = data.aiven_project.foo.project

  depends_on = [aiven_cassandra.bar]
}`, os.Getenv("AIVEN_PROJECT_NAME"), name)
}

func testAccCheckAivenServiceCassandraAttributes(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		if a["service_type"] != "cassandra" {
			return fmt.Errorf("expected to get a correct service type from Aiven, got :%s", a["service_type"])
		}

		if a["cassandra_user_config.0.public_access.0.prometheus"] != "true" {
			return fmt.Errorf("expected to get a correct public_access.prometheus from Aiven")
		}

		if a["cassandra_user_config.0.service_to_fork_from"] != "" {
			return fmt.Errorf("expected to get a correct service_to_fork_from from Aiven")
		}

		if a["cassandra_user_config.0.migrate_sstableloader"] != "true" {
			return fmt.Errorf("expected to get a correct migrate_sstableloader from Aiven")
		}

		return nil
	}
}
