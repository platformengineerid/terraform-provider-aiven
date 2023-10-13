package kafkatopic_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/stretchr/testify/assert"

	acc "github.com/aiven/terraform-provider-aiven/internal/acctest"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/service/kafkatopic"
)

func TestAccAivenKafkaTopic_basic(t *testing.T) {
	resourceName := "aiven_kafka_topic.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenKafkaTopicResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccKafkaTopicResource(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAivenKafkaTopicAttributes("data.aiven_kafka_topic.topic"),
					resource.TestCheckResourceAttr(resourceName, "project", os.Getenv("AIVEN_PROJECT_NAME")),
					resource.TestCheckResourceAttr(resourceName, "service_name", fmt.Sprintf("test-acc-sr-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "topic_name", fmt.Sprintf("test-acc-topic-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "partitions", "3"),
					resource.TestCheckResourceAttr(resourceName, "replication", "2"),
					resource.TestCheckResourceAttr(resourceName, "termination_protection", "false"),
				),
			},
		},
	})
}

func TestAccAivenKafkaTopic_many_topics(t *testing.T) {
	resourceName := "aiven_kafka_topic.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenKafkaTopicResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccKafka451TopicResource(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAivenKafkaTopicAttributes("data.aiven_kafka_topic.topic"),
					resource.TestCheckResourceAttr(resourceName, "project", os.Getenv("AIVEN_PROJECT_NAME")),
					resource.TestCheckResourceAttr(resourceName, "service_name", fmt.Sprintf("test-acc-sr-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "topic_name", fmt.Sprintf("test-acc-topic-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "partitions", "3"),
					resource.TestCheckResourceAttr(resourceName, "replication", "2"),
				),
			},
		},
	})
}

func TestAccAivenKafkaTopic_termination_protection(t *testing.T) {
	resourceName := "aiven_kafka_topic.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenKafkaTopicResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:                    testAccKafkaTopicTerminationProtectionResource(rName),
				PreventPostDestroyRefresh: true,
				ExpectNonEmptyPlan:        true,
				PlanOnly:                  true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "project", os.Getenv("AIVEN_PROJECT_NAME")),
					resource.TestCheckResourceAttr(resourceName, "service_name", fmt.Sprintf("test-acc-sr-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "topic_name", fmt.Sprintf("test-acc-topic-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "partitions", "3"),
					resource.TestCheckResourceAttr(resourceName, "replication", "2"),
					resource.TestCheckResourceAttr(resourceName, "termination_protection", "true"),
				),
			},
		},
	})
}

func TestAccAivenKafkaTopic_custom_timeouts(t *testing.T) {
	resourceName := "aiven_kafka_topic.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenKafkaTopicResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccKafkaTopicCustomTimeoutsResource(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAivenKafkaTopicAttributes("data.aiven_kafka_topic.topic"),
					resource.TestCheckResourceAttr(resourceName, "project", os.Getenv("AIVEN_PROJECT_NAME")),
					resource.TestCheckResourceAttr(resourceName, "service_name", fmt.Sprintf("test-acc-sr-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "topic_name", fmt.Sprintf("test-acc-topic-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "partitions", "3"),
					resource.TestCheckResourceAttr(resourceName, "replication", "2"),
					resource.TestCheckResourceAttr(resourceName, "termination_protection", "false"),
				),
			},
		},
	})
}

func testAccKafka451TopicResource(name string) string {
	return testAccKafkaTopicResource(name) + `
resource "aiven_kafka_topic" "more" {
  count = 450

  project      = data.aiven_project.foo.project
  service_name = aiven_kafka.bar.service_name
  topic_name   = "test-acc-topic-${count.index}"
  partitions   = 3
  replication  = 2
}`

}

func testAccKafkaTopicResource(name string) string {
	return fmt.Sprintf(`
data "aiven_project" "foo" {
  project = "%s"
}

resource "aiven_kafka" "bar" {
  project                 = data.aiven_project.foo.project
  cloud_name              = "google-europe-west1"
  plan                    = "startup-2"
  service_name            = "test-acc-sr-%s"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"
}

resource "aiven_kafka_topic" "foo" {
  project      = data.aiven_project.foo.project
  service_name = aiven_kafka.bar.service_name
  topic_name   = "test-acc-topic-%s"
  partitions   = 3
  replication  = 2

  config {
    flush_ms                  = 10
    cleanup_policy            = "compact"
    min_cleanable_dirty_ratio = 0.01
    delete_retention_ms       = 50000
  }
}

data "aiven_kafka_topic" "topic" {
  project      = aiven_kafka_topic.foo.project
  service_name = aiven_kafka_topic.foo.service_name
  topic_name   = aiven_kafka_topic.foo.topic_name

  depends_on = [aiven_kafka_topic.foo]
}`, os.Getenv("AIVEN_PROJECT_NAME"), name, name)
}

func testAccKafkaTopicCustomTimeoutsResource(name string) string {
	return fmt.Sprintf(`
data "aiven_project" "foo" {
  project = "%s"
}

resource "aiven_kafka" "bar" {
  project                 = data.aiven_project.foo.project
  cloud_name              = "google-europe-west1"
  plan                    = "startup-2"
  service_name            = "test-acc-sr-%s"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"

  timeouts {
    create = "25m"
    update = "20m"
  }
}

resource "aiven_kafka_topic" "foo" {
  project      = data.aiven_project.foo.project
  service_name = aiven_kafka.bar.service_name
  topic_name   = "test-acc-topic-%s"
  partitions   = 3
  replication  = 2

  timeouts {
    create = "15m"
    read   = "15m"
  }
}

data "aiven_kafka_topic" "topic" {
  project      = aiven_kafka_topic.foo.project
  service_name = aiven_kafka_topic.foo.service_name
  topic_name   = aiven_kafka_topic.foo.topic_name

  depends_on = [aiven_kafka_topic.foo]
}`, os.Getenv("AIVEN_PROJECT_NAME"), name, name)
}

func testAccKafkaTopicTerminationProtectionResource(name string) string {
	return fmt.Sprintf(`
data "aiven_project" "foo" {
  project = "%s"
}

resource "aiven_kafka" "bar" {
  project                 = data.aiven_project.foo.project
  cloud_name              = "google-europe-west1"
  plan                    = "startup-2"
  service_name            = "test-acc-sr-%s"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"

  kafka_user_config {
    kafka {
      group_max_session_timeout_ms = 70000
      log_retention_bytes          = 1000000000
    }
  }
}

resource "aiven_kafka_topic" "foo" {
  project                = data.aiven_project.foo.project
  service_name           = aiven_kafka.bar.service_name
  topic_name             = "test-acc-topic-%s"
  partitions             = 3
  replication            = 2
  termination_protection = true
}

data "aiven_kafka_topic" "topic" {
  project      = aiven_kafka_topic.foo.project
  service_name = aiven_kafka_topic.foo.service_name
  topic_name   = aiven_kafka_topic.foo.topic_name

  depends_on = [aiven_kafka_topic.foo]
}`, os.Getenv("AIVEN_PROJECT_NAME"), name, name)
}

func testAccCheckAivenKafkaTopicAttributes(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		r := s.RootModule().Resources[n]
		a := r.Primary.Attributes

		log.Printf("[DEBUG] kafka topic attributes %v", a)

		if a["project"] == "" {
			return fmt.Errorf("expected to get a project name from Aiven")
		}

		if a["service_name"] == "" {
			return fmt.Errorf("expected to get a service_name from Aiven")
		}

		if a["topic_name"] == "" {
			return fmt.Errorf("expected to get a topic_name from Aiven")
		}

		if a["partitions"] == "" {
			return fmt.Errorf("expected to get a partitions from Aiven")
		}

		if a["replication"] == "" {
			return fmt.Errorf("expected to get a replication from Aiven")
		}

		return nil
	}
}

func testAccCheckAivenKafkaTopicResourceDestroy(s *terraform.State) error {
	c := acc.GetTestAivenClient()

	ctx := context.Background()

	// loop through the resources in state, verifying each kafka topic is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aiven_kafka_topic" {
			continue
		}

		project, serviceName, topicName, err := schemautil.SplitResourceID3(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = c.Services.Get(ctx, project, serviceName)
		if err != nil {
			if aiven.IsNotFound(err) {
				return nil
			}
			return err
		}

		t, err := c.KafkaTopics.Get(ctx, project, serviceName, topicName)
		if err != nil {
			if aiven.IsNotFound(err) {
				return nil
			}
			return err
		}

		if t != nil {
			return fmt.Errorf("kafka topic (%s) still exists, id %s", topicName, rs.Primary.ID)
		}
	}

	return nil
}

func TestPartitions(t *testing.T) {
	type args struct {
		numPartitions int
	}
	tests := []struct {
		name           string
		args           args
		wantPartitions []*aiven.Partition
	}{
		{
			"basic",
			args{numPartitions: 3},
			[]*aiven.Partition{{}, {}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPartitions := partitions(tt.args.numPartitions); !reflect.DeepEqual(gotPartitions, tt.wantPartitions) {
				t.Errorf("partitions() = %v, want %v", gotPartitions, tt.wantPartitions)
			}
		})
	}
}

// TestAccAivenKafkaTopic_recreate validates that topic is recreated if it is missing
// Kafka looses all topics on turn off/on, then TF recreates topics. This test imitates the case.
func TestAccAivenKafkaTopic_recreate_missing(t *testing.T) {
	project := os.Getenv("AIVEN_PROJECT_NAME")

	prefix := "test-tf-acc-" + acctest.RandString(7)
	kafkaResource := "aiven_kafka.kafka"
	topicResource := "aiven_kafka_topic.topic"
	kafkaName := prefix + "-kafka"
	topicName := "topic"
	kafkaID := fmt.Sprintf("%s/%s", project, kafkaName)
	topicID := kafkaID + "/topic"

	config := testAccAivenKafkaTopicResourceRecreateMissing(prefix, project)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenKafkaTopicResourceDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: setups resources, creates the state
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(kafkaResource, "id", kafkaID),
					resource.TestCheckResourceAttr(topicResource, "id", topicID),
				),
			},
			{
				// Step 2: deletes topic, then runs apply, same config & checks
				PreConfig: func() {
					client := acc.GetTestAivenClient()

					ctx := context.Background()

					// deletes
					err := client.KafkaTopics.Delete(ctx, project, kafkaName, topicName)
					assert.NoError(t, err)

					// Makes sure topic does not exist
					tc, err := client.KafkaTopics.Get(ctx, project, kafkaName, topicName)
					assert.Nil(t, tc)
					assert.True(t, aiven.IsNotFound(err))

					// Invalidates cache for the topic
					kafkatopic.DeleteTopicFromCache(project, kafkaName, topicName)
				},
				// Now plan shows a diff
				ExpectNonEmptyPlan: true,
				RefreshState:       true,
			},
			{
				ExpectNonEmptyPlan: true,
				// Step 3: recreates the topic
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					// Saved in state
					resource.TestCheckResourceAttr(kafkaResource, "id", kafkaID),
					resource.TestCheckResourceAttr(topicResource, "id", topicID),
					func(state *terraform.State) error {
						// Topic exists and active
						client := acc.GetTestAivenClient()
						return retry.RetryContext(
							context.Background(),
							time.Minute,
							func() *retry.RetryError {
								ctx := context.Background()

								tc, err := client.KafkaTopics.Get(ctx, project, kafkaName, topicName)
								if err != nil {
									return &retry.RetryError{
										Err:       err,
										Retryable: aiven.IsNotFound(err),
									}
								}
								assert.Equal(t, tc.State, "ACTIVE")
								return nil
							},
						)
					},
				),
			},
		},
	})
}

func testAccAivenKafkaTopicResourceRecreateMissing(prefix, project string) string {
	return fmt.Sprintf(`
data "aiven_project" "project" {
  project = %[2]q
}

resource "aiven_kafka" "kafka" {
  project                 = data.aiven_project.project.project
  cloud_name              = "google-europe-west1"
  plan                    = "startup-2"
  service_name            = "%[1]s-kafka"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"
}

resource "aiven_kafka_topic" "topic" {
  project      = aiven_kafka.kafka.project
  service_name = aiven_kafka.kafka.service_name
  topic_name   = "topic"
  partitions   = 5
  replication  = 2
}`, prefix, project)
}

// TestAccAivenKafkaTopic_import_missing tests that simple import doesn't create a new topic
func TestAccAivenKafkaTopic_import_missing(t *testing.T) {
	project := os.Getenv("AIVEN_PROJECT_NAME")
	prefix := "test-tf-acc-" + acctest.RandString(7)
	kafkaName := prefix + "-kafka"
	topicName := "topic"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenKafkaTopicResourceDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: setups resources, creates the state
				Config: testAccAivenKafkaTopicResourceImportMissing(prefix, project),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("aiven_kafka.kafka", "id", fmt.Sprintf("%s/%s", project, kafkaName)),
				),
			},
			{
				// Step 2:
				// Tries to import non-existing topic
				// Must fail not create
				Config:        testAccAivenKafkaTopicResourceImportMissingStep2(prefix, project),
				ResourceName:  "aiven_kafka_topic.topic",
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s/%s/%s", project, kafkaName, topicName),
				ExpectError:   regexp.MustCompile(`While attempting to import an existing object to "aiven_kafka_topic.topic"`),
			},
		},
	})
}

func testAccAivenKafkaTopicResourceImportMissing(prefix, project string) string {
	result := fmt.Sprintf(`
resource "aiven_kafka" "kafka" {
  project                 = %[2]q
  cloud_name              = "google-europe-west1"
  plan                    = "startup-2"
  service_name            = "%[1]s-kafka"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"
}
`, prefix, project)
	return result
}

func testAccAivenKafkaTopicResourceImportMissingStep2(prefix, project string) string {
	result := fmt.Sprintf(`
resource "aiven_kafka" "kafka" {
  project                 = %[2]q
  cloud_name              = "google-europe-west1"
  plan                    = "startup-2"
  service_name            = "%[1]s-kafka"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"
}

resource "aiven_kafka_topic" "topic" {
  project      = aiven_kafka.kafka.project
  service_name = aiven_kafka.kafka.service_name
  topic_name   = "topic"
  partitions   = 5
  replication  = 2
}
`, prefix, project)
	return result
}

func TestAccAivenKafkaTopic_conflicts_if_exists(t *testing.T) {
	project := os.Getenv("AIVEN_PROJECT_NAME")
	prefix := "test-tf-acc-" + acctest.RandString(7)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenKafkaTopicResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccAivenKafkaTopicConflictsIfExists(prefix, project),
				ExpectError: regexp.MustCompile(`Topic conflict, already exists: conflict`),
			},
		},
	})
}

func testAccAivenKafkaTopicConflictsIfExists(prefix, project string) string {
	return fmt.Sprintf(`data "aiven_project" "project" {
  project = %[2]q
}

resource "aiven_kafka" "kafka" {
  project                 = data.aiven_project.project.project
  cloud_name              = "google-europe-west1"
  plan                    = "startup-2"
  service_name            = "%[1]s-kafka"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"
}

resource "aiven_kafka_topic" "topic" {
  project      = aiven_kafka.kafka.project
  service_name = aiven_kafka.kafka.service_name
  topic_name   = "conflict"
  partitions   = 5
  replication  = 2
}

resource "aiven_kafka_topic" "topic_conflict" {
  project      = aiven_kafka.kafka.project
  service_name = aiven_kafka.kafka.service_name
  topic_name   = "conflict"
  partitions   = 5
  replication  = 2

  depends_on = [
    aiven_kafka_topic.topic
  ]
}
`, prefix, project)
}

// partitions returns a slice, of empty aiven.Partition, of specified size
func partitions(numPartitions int) (partitions []*aiven.Partition) {
	for i := 0; i < numPartitions; i++ {
		partitions = append(partitions, &aiven.Partition{})
	}
	return
}
