package provider

import (
	"context"
	"fmt"

	"github.com/Juniper/terraform-provider-mist/internal/resource_org_wlan_portal_image"

	"github.com/tmunzer/mistapi-go/mistapi"
	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource              = &orgWlanPortalImageResource{}
	_ resource.ResourceWithConfigure = &orgWlanPortalImageResource{}
)

func NewOrgWlanPortalImage() resource.Resource {
	return &orgWlanPortalImageResource{}
}

type orgWlanPortalImageResource struct {
	client mistapi.ClientInterface
}

func (r *orgWlanPortalImageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org WLAN Portal Image client")
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
func (r *orgWlanPortalImageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_wlan_portal_image"
}

func (r *orgWlanPortalImageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryWlan + "This resource is used to upload a WLAN Captive Web Portal background image." +
			"The WLAN object contains all the required configuration to broadcast an SSID (Authentication, VLAN, ...)",
		Attributes: resource_org_wlan_portal_image.OrgWlanPortalImageResourceSchema(ctx).Attributes,
	}
}

func (r *orgWlanPortalImageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting Org WLAN Portal Image Create")
	var plan, state resource_org_wlan_portal_image.OrgWlanPortalImageModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(plan.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"org_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	file := getFile(plan.File.ValueString(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	var json string = ""
	_, err = r.client.OrgsWlans().UploadOrgWlanPortalImage(ctx, orgId, wlanId, file, &json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error uploadting Wlan Portal Image",
			"Could not create Wlan, unexpected error: "+err.Error(),
		)
		return
	}

	state.File = plan.File
	state.OrgId = plan.OrgId
	state.WlanId = plan.WlanId

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgWlanPortalImageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (r *orgWlanPortalImageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_org_wlan_portal_image.OrgWlanPortalImageModel
	tflog.Info(ctx, "Starting Org WLAN Portal Image Update")

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(plan.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(plan.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"org_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.OrgId.ValueString(), err.Error()),
		)
		return
	}

	var json string
	var file models.FileWrapper
	file.File = []byte(plan.File.ValueString())

	_, err = r.client.OrgsWlans().UploadOrgWlanPortalImage(ctx, orgId, wlanId, file, &json)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error uploadting Wlan Portal Image",
			"Could not create Wlan, unexpected error: "+err.Error(),
		)
		return
	}

	state.File = plan.File
	state.OrgId = plan.OrgId
	state.WlanId = plan.WlanId

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *orgWlanPortalImageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_org_wlan_portal_image.OrgWlanPortalImageModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(state.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	wlanId, err := uuid.Parse(state.WlanId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"wlan_id\" value for \"org_wlan_portal_image\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.OrgId.ValueString(), err.Error()),
		)
		return
	}

	httpr, err := r.client.OrgsWlans().DeleteOrgWlanPortalImage(ctx, orgId, wlanId)
	if httpr.StatusCode != 404 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting Wlan Portal Image",
			"Could not delete Wlan, unexpected error: "+err.Error(),
		)
		return
	}
}

func getFile(fileData string, diags *diag.Diagnostics) models.FileWrapper {
	file, err := models.GetFile(fileData)
	if err != nil {
		diags.AddError(
			"Invalid \"file\" value for \"org_wlan_portal_image\" resource",
			fmt.Sprintf("Could not open file \"%s\": %s", fileData, err.Error()),
		)
	}
	return file
}
