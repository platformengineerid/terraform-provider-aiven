package organization

import (
	"context"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig"
)

var aivenOrganizationUserGroupSchema = map[string]*schema.Schema{
	"organization_id": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: userconfig.Desc("The ID of the organization.").ForceNew().Build(),
	},
	"name": {
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
		Description: userconfig.Desc("The name of the user group.").ForceNew().Build(),
	},
	"description": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: userconfig.Desc("The description of the user group.").ForceNew().Build(),
	},
	"create_time": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Time of creation.",
	},
	"update_time": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Time of last update.",
	},
	"group_id": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The ID of the user group.",
	},
}

func ResourceOrganizationUserGroup() *schema.Resource {
	return &schema.Resource{
		Description:   "Creates and manages a user group in an organization.",
		CreateContext: resourceOrganizationUserGroupCreate,
		ReadContext:   resourceOrganizationUserGroupRead,
		UpdateContext: resourceOrganizationUserGroupUpdate,
		DeleteContext: resourceOrganizationUserGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: schemautil.DefaultResourceTimeouts(),

		Schema: aivenOrganizationUserGroupSchema,
	}
}

func resourceOrganizationUserGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	orgID := d.Get("organization_id").(string)
	r, err := client.OrganizationUserGroups.Create(
		ctx,
		orgID,
		aiven.OrganizationUserGroupRequest{
			UserGroupName: d.Get("name").(string),
			Description:   d.Get("description").(string),
		},
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(schemautil.BuildResourceID(orgID, r.UserGroupID))

	return resourceOrganizationUserGroupRead(ctx, d, m)
}

func resourceOrganizationUserGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	orgID, userGroupID, err := schemautil.SplitResourceID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.OrganizationUserGroups.Get(ctx, orgID, userGroupID)
	if err != nil {
		return diag.FromErr(schemautil.ResourceReadHandleNotFound(err, d))
	}

	if err := d.Set("name", r.UserGroupName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", r.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("create_time", r.CreateTime.String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("update_time", r.UpdateTime.String()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("group_id", r.UserGroupID); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceOrganizationUserGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	orgID, userGroupID, err := schemautil.SplitResourceID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.OrganizationUserGroups.Update(ctx, orgID, userGroupID, aiven.OrganizationUserGroupRequest{
		UserGroupName: d.Get("name").(string),
		Description:   d.Get("description").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceOrganizationUserGroupRead(ctx, d, m)
}

func resourceOrganizationUserGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	orgID, userGroupID, err := schemautil.SplitResourceID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if err = client.OrganizationUserGroups.Delete(ctx, orgID, userGroupID); err != nil && !aiven.IsNotFound(err) {
		return diag.FromErr(err)
	}

	return nil
}
