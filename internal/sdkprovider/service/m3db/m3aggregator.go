package m3db

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig/stateupgrader"
	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/service"
)

func aivenM3AggregatorSchema() map[string]*schema.Schema {
	schemaM3 := schemautil.ServiceCommonSchema()
	schemaM3[schemautil.ServiceTypeM3Aggregator] = &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: "M3 aggregator specific server provided values",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{},
		},
	}
	schemaM3[schemautil.ServiceTypeM3Aggregator+"_user_config"] = service.GetUserConfig(schemautil.ServiceTypeM3Aggregator)

	return schemaM3
}
func ResourceM3Aggregator() *schema.Resource {
	return &schema.Resource{
		Description:   "The M3 Aggregator resource allows the creation and management of Aiven M3 Aggregator services.",
		CreateContext: schemautil.ResourceServiceCreateWrapper(schemautil.ServiceTypeM3Aggregator),
		ReadContext:   schemautil.ResourceServiceRead,
		UpdateContext: schemautil.ResourceServiceUpdate,
		DeleteContext: schemautil.ResourceServiceDelete,
		CustomizeDiff: customdiff.Sequence(
			schemautil.SetServiceTypeIfEmpty(schemautil.ServiceTypeM3Aggregator),
			schemautil.CustomizeDiffDisallowMultipleManyToOneKeys,
			customdiff.IfValueChange("disk_space",
				schemautil.DiskSpaceShouldNotBeEmpty,
				schemautil.CustomizeDiffCheckDiskSpace,
			),
			customdiff.IfValueChange("additional_disk_space",
				schemautil.DiskSpaceShouldNotBeEmpty,
				schemautil.CustomizeDiffCheckDiskSpace,
			),
			customdiff.IfValueChange("service_integrations",
				schemautil.ServiceIntegrationShouldNotBeEmpty,
				schemautil.CustomizeDiffServiceIntegrationAfterCreation,
			),
			customdiff.Sequence(
				schemautil.CustomizeDiffCheckPlanAndStaticIpsCannotBeModifiedTogether,
				schemautil.CustomizeDiffCheckStaticIPDisassociation,
			),
		),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: schemautil.DefaultResourceTimeouts(),

		Schema:         aivenM3AggregatorSchema(),
		SchemaVersion:  1,
		StateUpgraders: stateupgrader.M3Aggregator(),
	}
}
