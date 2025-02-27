package serviceintegration

import (
	"context"
	"fmt"
	"os"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/aiven/terraform-provider-aiven/internal/common"
	"github.com/aiven/terraform-provider-aiven/internal/sweep"
)

func init() {
	ctx := context.Background()

	sweep.AddTestSweepers("aiven_service_integration", &resource.Sweeper{
		Name: "aiven_service_integration",
		F:    sweepServiceIntegrations(ctx),
	})

	sweep.AddTestSweepers("aiven_service_integration_endpoint", &resource.Sweeper{
		Name: "aiven_service_integration_endpoint",
		F:    sweepServiceIntegrationEndpoints(ctx),
	})
}

func sweepServiceIntegrations(ctx context.Context) func(region string) error {
	return func(region string) error {
		projectName := os.Getenv("AIVEN_PROJECT_NAME")
		client, err := sweep.SharedClient()
		if err != nil {
			return err
		}

		services, err := client.Services.List(ctx, projectName)
		if common.IsCritical(err) {
			return fmt.Errorf("error retrieving a list of service for a project `%s`: %s", projectName, err)
		}

		for _, service := range services {
			if len(service.Integrations) == 0 {
				continue
			}

			serviceIntegrations, err := client.ServiceIntegrations.List(ctx, projectName, service.Name)
			if err != nil {
				return fmt.Errorf("error retrieving a list of service integration for service `%s`: %s", service.Name, err)
			}
			for _, serviceIntegration := range serviceIntegrations {
				if err := client.ServiceIntegrations.Delete(
					ctx,
					projectName,
					serviceIntegration.ServiceIntegrationID,
				); err != nil {
					if !aiven.IsNotFound(err) {
						return fmt.Errorf(
							"unable to delete service integration `%s`: %s",
							serviceIntegration.ServiceIntegrationID,
							err,
						)
					}
				}
			}
		}

		return nil
	}
}

func sweepServiceIntegrationEndpoints(ctx context.Context) func(region string) error {
	return func(region string) error {
		projectName := os.Getenv("AIVEN_PROJECT_NAME")
		client, err := sweep.SharedClient()
		if err != nil {
			return err
		}

		endpoints, err := client.ServiceIntegrationEndpoints.List(ctx, projectName)
		if err != nil {
			return err
		}

		for _, endpoint := range endpoints {
			err = client.ServiceIntegrationEndpoints.Delete(ctx, projectName, endpoint.EndpointID)
			if common.IsCritical(err) {
				return err
			}
		}

		return nil
	}
}
