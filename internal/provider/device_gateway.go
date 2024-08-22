package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/Juniper/terraform-provider-mist/internal/resource_device_gateway"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ resource.Resource                = &deviceGatewayResource{}
	_ resource.ResourceWithConfigure   = &deviceGatewayResource{}
	_ resource.ResourceWithImportState = &deviceGatewayResource{}
)

func NewDeviceGatewayResource() resource.Resource {
	return &deviceGatewayResource{}
}

type deviceGatewayResource struct {
	client mistapi.ClientInterface
}

func (r *deviceGatewayResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist DeviceGateway client")
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
func (r *deviceGatewayResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_gateway"
}

func (r *deviceGatewayResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This resource manages the Gateway configuration." +
			"It can be used to define specific configuration at the device level or to override Org Gateway template settings.",
		Attributes: resource_device_gateway.DeviceGatewayResourceSchema(ctx).Attributes,
	}
}

func (r *deviceGatewayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Starting DeviceGateway Create")
	var plan, state resource_device_gateway.DeviceGatewayModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_gateway, diags := resource_device_gateway.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"device_gateway\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"device_gateway\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.DeviceId.ValueString(), err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Starting DeviceGateway Create on Site "+plan.SiteId.ValueString()+" for device "+plan.DeviceId.ValueString())
	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_gateway)

	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error creating device_gateway",
			"Could not create device_gateway, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_gateway := models.DeviceGateway{}
	json.Unmarshal(body, &mist_gateway)
	state, diags = resource_device_gateway.SdkToTerraform(ctx, &mist_gateway)
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

func (r *deviceGatewayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resource_device_gateway.DeviceGatewayModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGateway Read: device_gateway_id "+state.DeviceId.ValueString())
	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"device_gateway\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"device_gateway\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.DeviceId.ValueString(), err.Error()),
		)
		return
	}

	httpr, err := r.client.SitesDevices().GetSiteDevice(ctx, siteId, deviceId)
	if httpr.Response.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if httpr.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error getting device_gateway",
			"Could not get device_gateway, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(httpr.Response.Body)
	mist_gateway := models.DeviceGateway{}
	json.Unmarshal(body, &mist_gateway)

	state, diags = resource_device_gateway.SdkToTerraform(ctx, &mist_gateway)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceGatewayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state, plan resource_device_gateway.DeviceGatewayModel
	tflog.Info(ctx, "Starting DeviceGateway Update")

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

	device_gateway, diags := resource_device_gateway.TerraformToSdk(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGateway Update for DeviceGateway "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(plan.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"device_gateway\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(plan.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"device_gateway\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", plan.DeviceId.ValueString(), err.Error()),
		)
		return
	}

	data, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_gateway)

	if data.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error updating device_gateway",
			"Could not update device_gateway, unexpected error: "+err.Error(),
		)
		return
	}

	body, _ := io.ReadAll(data.Response.Body)
	mist_gateway := models.DeviceGateway{}
	json.Unmarshal(body, &mist_gateway)
	state, diags = resource_device_gateway.SdkToTerraform(ctx, &mist_gateway)
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

func (r *deviceGatewayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resource_device_gateway.DeviceGatewayModel

	diags := resp.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	device_gateway, e := resource_device_gateway.DeleteTerraformToSdk(ctx)
	resp.Diagnostics.Append(e...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Starting DeviceGateway Delete: device_gateway_id "+state.DeviceId.ValueString())

	siteId, err := uuid.Parse(state.SiteId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"device_gateway\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.SiteId.ValueString(), err.Error()),
		)
		return
	}

	deviceId, err := uuid.Parse(state.DeviceId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"device_id\" value for \"device_gateway\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s", state.DeviceId.ValueString(), err.Error()),
		)
		return
	}

	httpr, err := r.client.SitesDevices().UpdateSiteDevice(ctx, siteId, deviceId, &device_gateway)
	if httpr.Response.StatusCode != 404 && httpr.Response.StatusCode != 200 && err != nil {
		resp.Diagnostics.AddError(
			"Error deleting device_gateway",
			"Could not delete device_gateway, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *deviceGatewayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	importIds := strings.Split(req.ID, ".")
	if len(importIds) != 2 {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"device_gateway\" resource",
			"Import \"id\" format must be \"{site_id}.{device_id}\"",
		)
		return
	}
	_, err := uuid.Parse(importIds[0])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"site_id\" value for \"device_gateway\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s. Import \"id\" format must be \"{site_id}.{device_id}\"", importIds[0], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("site_id"), importIds[0])...)

	_, err = uuid.Parse(importIds[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"id\" value for \"device_gateway\" resource",
			fmt.Sprintf("Could not parse the UUID \"%s\": %s. Import \"id\" format must be \"{site_id}.{device_id}\"", importIds[1], err.Error()),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), importIds[1])...)
}
