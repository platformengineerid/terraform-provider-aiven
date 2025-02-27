package project

import (
	"context"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil/userconfig"
)

var aivenProjectSchema = map[string]*schema.Schema{
	"ca_cert": {
		Type:        schema.TypeString,
		Computed:    true,
		Sensitive:   true,
		Description: "The CA certificate of the project. This is required for configuring clients that connect to certain services like Kafka.",
	},
	"account_id": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: userconfig.Desc("An optional property to link a project to an already existing account by using account ID.").Referenced().Build(),
		Deprecated:  "Use parent_id instead. This field will be removed in the next major release.",
		DiffSuppressFunc: func(_, _, _ string, d *schema.ResourceData) bool {
			_, ok := d.GetOk("parent_id")
			return ok
		},
	},
	"parent_id": {
		Type:     schema.TypeString,
		Optional: true,
		Description: userconfig.Desc(
			"An optional property to link a project to an already existing organization or account by using its ID.",
		).Referenced().Build(),
		DiffSuppressFunc: func(_, _, _ string, d *schema.ResourceData) bool {
			_, ok := d.GetOk("account_id")
			return ok
		},
	},
	"copy_from_project": {
		Type:             schema.TypeString,
		Optional:         true,
		DiffSuppressFunc: schemautil.CreateOnlyDiffSuppressFunc,
		Description:      userconfig.Desc("is the name of another project used to copy billing information and some other project attributes like technical contacts from. This is mostly relevant when an existing project has billing type set to invoice and that needs to be copied over to a new project. (Setting billing is otherwise not allowed over the API.) This only has effect when the project is created.").Referenced().Build(),
	},
	"use_source_project_billing_group": {
		Type:             schema.TypeBool,
		Optional:         true,
		DiffSuppressFunc: schemautil.CreateOnlyDiffSuppressFunc,
		Description:      "Use the same billing group that is used in source project.",
		Deprecated:       "This field is deprecated and will be removed in the next major release.",
	},
	"add_account_owners_admin_access": {
		Type:        schema.TypeBool,
		Optional:    true,
		Description: userconfig.Desc("If parent_id is set, grant account owner team admin access to the new project.").DefaultValue(true).Build(),
		Deprecated: "This field is deprecated and will be removed in the next major release. " +
			"Currently, it will always be set to true, regardless of the value set in the manifest.",
		DiffSuppressFunc: func(_ string, _ string, _ string, _ *schema.ResourceData) bool {
			return true
		},
	},
	"project": {
		Type:        schema.TypeString,
		Required:    true,
		Description: "Defines the name of the project. Name must be globally unique (between all Aiven customers) and cannot be changed later without destroying and re-creating the project, including all sub-resources.",
	},
	"technical_emails": {
		Type:        schema.TypeSet,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Optional:    true,
		Description: "Defines the email addresses that will receive alerts about upcoming maintenance updates or warnings about service instability. It is a good practice to keep this up-to-date to be aware of any potential issues with your project.",
	},
	"default_cloud": {
		Type:             schema.TypeString,
		Optional:         true,
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFunc,
		Description:      "Defines the default cloud provider and region where services are hosted. This can be changed freely after the project is created. This will not affect existing services.",
	},
	"billing_group": {
		Type:             schema.TypeString,
		Optional:         true,
		Description:      userconfig.Desc("The id of the billing group that is linked to this project.").Referenced().Build(),
		DiffSuppressFunc: schemautil.EmptyObjectDiffSuppressFunc,
	},
	"tag": {
		Description: "Tags are key-value pairs that allow you to categorize projects.",
		Type:        schema.TypeSet,
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Description: "Project tag key",
					Type:        schema.TypeString,
					Required:    true,
				},
				"value": {
					Description: "Project tag value",
					Type:        schema.TypeString,
					Required:    true,
				},
			},
		},
	},

	// computed fields
	"payment_method": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The method of invoicing used for payments for this project, e.g. `card`.",
	},
	"available_credits": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The amount of platform credits available to the project. This could be your free trial or other promotional credits.",
	},
	"estimated_balance": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The current accumulated bill for this project in the current billing period.",
	},
}

func ResourceProject() *schema.Resource {
	return &schema.Resource{
		Description:   "The Project resource allows the creation and management of Aiven Projects.",
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: schemautil.DefaultResourceTimeouts(),

		Schema: aivenProjectSchema,
		CustomizeDiff: customdiff.IfValueChange("tag",
			schemautil.TagsShouldNotBeEmpty,
			schemautil.CustomizeDiffCheckUniqueTag,
		),
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	projectName := d.Get("project").(string)

	req := aiven.CreateProjectRequest{
		Cloud:                        schemautil.OptionalStringPointer(d, "default_cloud"),
		CopyFromProject:              d.Get("copy_from_project").(string),
		Project:                      projectName,
		TechnicalEmails:              contactEmailListForAPI(d, "technical_emails", true),
		UseSourceProjectBillingGroup: d.Get("use_source_project_billing_group").(bool),
		BillingGroupId:               d.Get("billing_group").(string),
		AddAccountOwnersAdminAccess:  schemautil.OptionalBoolPointer(d, "add_account_owners_admin_access"),
		Tags:                         schemautil.GetTagsFromSchema(d),
	}

	ptrAccountID, err := accountIDPointer(ctx, client, d)
	if err != nil {
		return diag.FromErr(err)
	}

	req.AccountId = ptrAccountID

	_, err = client.Projects.Create(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	if _, ok := d.GetOk("billing_group"); !ok {
		// if billing_group is not set but copy_from_project is not empty,
		// copy billing group from source project
		if sourceProject, ok := d.GetOk("copy_from_project"); ok {
			dia := resourceProjectCopyBillingGroupFromProject(ctx, client, sourceProject.(string), d)
			if dia.HasError() {
				diag.FromErr(err)
			}
		}
	}

	d.SetId(projectName)

	return resourceProjectRead(ctx, d, m)
}

func resourceProjectCopyBillingGroupFromProject(
	ctx context.Context,
	client *aiven.Client,
	sourceProjectName string,
	d *schema.ResourceData,
) diag.Diagnostics {
	list, err := client.BillingGroup.ListAll(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, bg := range list {
		projects, err := client.BillingGroup.GetProjects(ctx, bg.Id)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, pr := range projects {
			if pr == sourceProjectName {
				log.Printf("[DEBUG] Source project `%s` has billing group `%s`", sourceProjectName, bg.Id)
				return resourceProjectAssignToBillingGroup(ctx, sourceProjectName, bg.Id, client, d)
			}
		}
	}

	log.Printf("[DEBUG] Source project `%s` is not associated to any billing group", sourceProjectName)
	return nil
}

func resourceProjectAssignToBillingGroup(
	ctx context.Context,
	projectName string,
	billingGroupID string,
	client *aiven.Client,
	d *schema.ResourceData,
) diag.Diagnostics {
	log.Printf("[DEBUG] Associating project `%s` with the billing group `%s`", projectName, billingGroupID)
	_, err := client.BillingGroup.Get(ctx, billingGroupID)
	if err != nil {
		return diag.Errorf("cannot get a billing group by id: %s", err)
	}

	var isAlreadyAssigned bool
	assignedProjects, err := client.BillingGroup.GetProjects(ctx, billingGroupID)
	if err != nil {
		return diag.Errorf("cannot get a billing group assigned projects list: %s", err)
	}
	for _, p := range assignedProjects {
		if p == projectName {
			isAlreadyAssigned = true
		}
	}

	if !isAlreadyAssigned {
		err = client.BillingGroup.AssignProjects(ctx, billingGroupID, []string{projectName})
		if err != nil {
			return diag.Errorf("cannot assign project to a billing group: %s", err)
		}
	}

	if err := d.Set("billing_group", billingGroupID); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// nolint:staticcheck // TODO: Migrate to helper/retry package to avoid deprecated resource.StateRefreshFunc.
func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	//goland:noinspection GoDeprecation
	conf := &resource.StateChangeConf{
		Pending:    []string{"pending"},
		Target:     []string{"target"},
		Timeout:    time.Minute,
		MinTimeout: time.Second,
		Delay:      time.Second,
		Refresh: func() (result interface{}, state string, err error) {
			p, err := client.Projects.Get(ctx, d.Id())
			if isNotProjectMember(err) {
				return nil, "pending", nil
			}
			if err != nil {
				return nil, "", err
			}
			return p, "target", nil
		},
	}

	project, err := conf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(schemautil.ResourceReadHandleNotFound(err, d))
	}
	return setProjectTerraformProperties(ctx, d, client, project.(*aiven.Project))
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	projectName := d.Get("project").(string)

	req := aiven.UpdateProjectRequest{
		Name:                        projectName,
		Cloud:                       schemautil.OptionalStringPointer(d, "default_cloud"),
		TechnicalEmails:             contactEmailListForAPI(d, "technical_emails", false),
		Tags:                        schemautil.GetTagsFromSchema(d),
		AddAccountOwnersAdminAccess: schemautil.OptionalBoolPointer(d, "add_account_owners_admin_access"),
	}

	ptrAccountID, err := accountIDPointer(ctx, client, d)
	if err != nil {
		return diag.FromErr(err)
	}

	// TODO: Perhaps req.AccountId should also be a pointer here?
	if ptrAccountID == nil {
		req.AccountId = ""
	} else {
		req.AccountId = *ptrAccountID
	}

	project, err := client.Projects.Update(ctx, d.Id(), req)
	if err != nil {
		return diag.FromErr(err)
	}

	if billingGroupID, ok := d.GetOk("billing_group"); ok {
		dia := resourceProjectAssignToBillingGroup(
			ctx,
			d.Get("project").(string),
			billingGroupID.(string),
			client,
			d,
		)
		if dia.HasError() {
			return dia
		}
	}

	d.SetId(project.Name)

	return nil
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	err := client.Projects.Delete(ctx, d.Id())

	// Silence "Project with open balance cannot be deleted" error
	// to make long acceptance tests pass which generate some balance
	re := regexp.MustCompile("Project with open balance cannot be deleted")
	if err != nil && os.Getenv("TF_ACC") != "" {
		if re.MatchString(err.Error()) && err.(aiven.Error).Status == 403 {
			return nil
		}
	}

	if err != nil {
		if aiven.IsNotFound(err) {
			return nil
		}

		return diag.FromErr(err)
	}

	return nil
}

func resourceProjectGetCACert(
	ctx context.Context,
	project string,
	client *aiven.Client,
	d *schema.ResourceData,
) diag.Diagnostics {
	ca, err := client.CA.Get(ctx, project)
	if err == nil {
		if err := d.Set("ca_cert", ca); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func getLongCardID(ctx context.Context, client *aiven.Client, cardID string) (*string, error) {
	if cardID == "" {
		return nil, nil
	}

	card, err := client.CardsHandler.Get(ctx, cardID)
	if err != nil {
		return nil, err
	}
	if card != nil {
		return &card.CardID, nil
	}
	return &cardID, nil
}

func contactEmailListForAPI(d *schema.ResourceData, field string, newResource bool) *[]*aiven.ContactEmail {
	var results []*aiven.ContactEmail
	// We don't want to send empty list for new resource if data is copied from other
	// project to prevent accidental override of the emails being copied. Empty array
	// should be sent if user has explicitly defined that even when copy_from_project
	// is set but Terraform does not support checking that; d.GetOkExists returns false
	// even if the value is set (to empty).
	if _, ok := d.GetOk("copy_from_project"); ok || !newResource {
		results = []*aiven.ContactEmail{}
	}
	valuesInterface, ok := d.GetOk(field)
	if ok && valuesInterface != nil {
		for _, emailInterface := range valuesInterface.(*schema.Set).List() {
			results = append(results, &aiven.ContactEmail{Email: emailInterface.(string)})
		}
	}
	if results == nil {
		return nil
	}
	return &results
}

func contactEmailListForTerraform(d *schema.ResourceData, field string, contactEmails []*aiven.ContactEmail) error {
	_, existsBefore := d.GetOk(field)
	if !existsBefore && len(contactEmails) == 0 {
		return nil
	}

	var results []string
	for _, contactEmail := range contactEmails {
		results = append(results, contactEmail.Email)
	}

	return d.Set(field, results)
}

func setProjectTerraformProperties(
	ctx context.Context,
	d *schema.ResourceData,
	client *aiven.Client,
	project *aiven.Project,
) diag.Diagnostics {
	if stateID, ok := d.GetOk("parent_id"); ok {
		idToSet, err := schemautil.DetermineMixedOrganizationConstraintIDToStore(
			ctx,
			client,
			stateID.(string),
			project.AccountId,
		)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("parent_id", idToSet); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("project", project.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("account_id", project.AccountId); err != nil {
		return diag.FromErr(err)
	}
	if err := contactEmailListForTerraform(d, "technical_emails", project.TechnicalEmails); err != nil {
		return diag.FromErr(err)
	}
	if d := resourceProjectGetCACert(ctx, project.Name, client, d); d != nil {
		return d
	}
	if err := d.Set("default_cloud", project.DefaultCloud); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("available_credits", project.AvailableCredits); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("estimated_balance", project.EstimatedBalance); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("payment_method", project.PaymentMethod); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("billing_group", project.BillingGroupId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tag", schemautil.SetTagsTerraformProperties(project.Tags)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// isNotProjectMember This happens right after project created
func isNotProjectMember(err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(aiven.Error)
	return ok && e.Status == http.StatusForbidden && strings.Contains(e.Message, "Not a project member")
}
