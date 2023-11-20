package flink

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig/stateupgrader"
	"github.com/aiven/terraform-provider-aiven/internal/sdkprovider/userconfig/service"
)

func aivenFlinkSchema() map[string]*schema.Schema {
	aivenFlinkSchema := schemautil.ServiceCommonSchema()
	aivenFlinkSchema[schemautil.ServiceTypeFlink] = &schema.Schema{
		Type:        schema.TypeList,
		MaxItems:    1,
		Computed:    true,
		Description: "Flink server provided values",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"host_ports": {
					Type:        schema.TypeList,
					Computed:    true,
					Description: "Host and Port of a Flink server",
					Optional:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
	aivenFlinkSchema[schemautil.ServiceTypeFlink+"_user_config"] = service.GetUserConfig(schemautil.ServiceTypeFlink)

	return aivenFlinkSchema
}

func ResourceFlink() *schema.Resource {
	return &schema.Resource{
		Description:   "The Flink resource allows the creation and management of Aiven Flink services.",
		CreateContext: schemautil.ResourceServiceCreateWrapper(schemautil.ServiceTypeFlink),
		ReadContext:   schemautil.ResourceServiceRead,
		UpdateContext: schemautil.ResourceServiceUpdate,
		DeleteContext: schemautil.ResourceServiceDelete,
		CustomizeDiff: customdiff.Sequence(
			schemautil.SetServiceTypeIfEmpty(schemautil.ServiceTypeFlink),
			schemautil.CustomizeDiffDisallowMultipleManyToOneKeys,
			customdiff.IfValueChange("tag",
				schemautil.TagsShouldNotBeEmpty,
				schemautil.CustomizeDiffCheckUniqueTag,
			),
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

		Schema:         aivenFlinkSchema(),
		SchemaVersion:  1,
		StateUpgraders: stateupgrader.Flink(),
	}
}
