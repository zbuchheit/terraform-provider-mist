package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_webhook"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &orgWebhookResource{}
	_ resource.ResourceWithConfigure   = &orgWebhookResource{}
	_ resource.ResourceWithImportState = &orgWebhookResource{}
)

func NewOrgWebhook() resource.Resource {
	return &orgWebhookResource{}
}

type orgWebhookResource struct {
	client mistapi.ClientInterface
}

func (r *orgWebhookResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Webhook client")
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(mistapi.ClientInterface)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *mistapigo.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}
func (r *orgWebhookResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_webhook"
}

func (r *orgWebhookResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryOrg + "This resource manages Org Webhooks.",
		Attributes:          resource_org_webhook.OrgWebhookResourceSchema(ctx).Attributes,
	}
}

func (r *orgWebhookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Webhook Create")
	var plan, state resource_org_webhook.OrgWebhookModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_webhook\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	webhook, diags := resource_org_webhook.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	data, err := r.client.OrgsWebhooks().CreateOrgWebhook(ctx, orgId, &webhook)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Webhook",
			"Could not create Webhook, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_webhook.SdkToTerraform(ctx, &data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgWebhookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_org_webhook.OrgWebhookModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_webhook\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	webhookId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_webhook\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Webhook Read: webhook_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsWebhooks().GetOrgWebhook(ctx, orgId, webhookId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error getting Webhook",
			"Could not get Webhook, unexpected error: "+err.Error(),
		)
		return
	}
	state, diags = resource_org_webhook.SdkToTerraform(ctx, &httpr.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *orgWebhookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_webhook.OrgWebhookModel
	tflog.Info(ctx, "Starting Webhook Update")

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	webhook, diags := resource_org_webhook.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_webhook\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	webhookId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_webhook\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Webhook Update for Webhook "+state.Id.ValueString())
	data, err := r.client.OrgsWebhooks().
		UpdateOrgWebhook(ctx, orgId, webhookId, &webhook)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Webhook",
			"Could not update Webhook, unexpected error: "+err.Error(),
		)
		return
	}

	state, diags = resource_org_webhook.SdkToTerraform(ctx, &data.Data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgWebhookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_webhook.OrgWebhookModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_webhook\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}
	webhookId, err := uuid.Parse(state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_webhook\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.Id.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting Webhook Delete: webhook_id "+state.Id.ValueString())
	httpr, err := r.client.OrgsWebhooks().DeleteOrgWebhook(ctx, orgId, webhookId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Webhook",
			"Could not delete Webhook, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *orgWebhookResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_webhook\" resource",
			"import \"id\" format must be \"{org_id}.{webhook_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_webhook\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{webhook_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("org_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"org_webhook\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s. Import \"id\" format must be \"{org_id}.{webhook_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
