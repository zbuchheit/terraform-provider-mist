package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/tmunzer/mistapi-go/mistapi"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/Juniper/terraform-provider-mist/internal/datasource_org_deviceprofiles_gateway"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ datasource.DataSource = (*orgDeviceprofilesGatewayDataSource)(nil)

func NewOrgDeviceprofilesGatewayDataSource() datasource.DataSource {
	return &orgDeviceprofilesGatewayDataSource{}
}

type orgDeviceprofilesGatewayDataSource struct {
	client mistapi.ClientInterface
}

func (d *orgDeviceprofilesGatewayDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Mist Org Device Profiles Gateway Datasource client")
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

	d.client = client
}
func (d *orgDeviceprofilesGatewayDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_deviceprofiles_gateway"
}

func (d *orgDeviceprofilesGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: docCategoryDevices + "This provides the list of Gateway Device Profiles (Hub Profile)." +
			"A HUB profile is a configuration profile that automates the creation of overlay networks and defines the attributes of a hub device in a network. " +
			"It includes settings for wan interfaces, lan interfaces, dns servers, traffic steering preferences, application policies, and routing options. " +
			"HUB profiles are used to create consistent configurations for hub devices and ensure efficient connectivity between hubs and spokes in a network.",
		Attributes: datasource_org_deviceprofiles_gateway.OrgDeviceprofilesGatewayDataSourceSchema(ctx).Attributes,
	}
}

func (d *orgDeviceprofilesGatewayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var ds datasource_org_deviceprofiles_gateway.OrgDeviceprofilesGatewayModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &ds)...)

	if resp.Diagnostics.HasError() {
		return
	}

	orgId, err := uuid.Parse(ds.OrgId.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid \"org_id\" value for \"org_deviceprofiles_gateway\" data_source",
			"Could parse the UUID: "+err.Error(),
		)
		return
	}

	var mType models.DeviceTypeEnum = models.DeviceTypeEnum("ap")

	var limit int = 1000
	var page int = 0
	var total int = 9999
	var elements []attr.Value
	var diags diag.Diagnostics

	for limit*page < total {
		page += 1
		tflog.Debug(ctx, "Pagination Info", map[string]interface{}{
			"page":  page,
			"limit": limit,
			"total": total,
		})
		data, err := d.client.OrgsDeviceProfiles().ListOrgDeviceProfiles(ctx, orgId, &mType, &page, &limit)
		if data.Response.StatusCode != 200 && err != nil {
			resp.Diagnostics.AddError(
				"Error getting Org Gateway Device Profiles",
				"Could not get the list of Org Gateway Device Profiles, unexpected error: "+err.Error(),
			)
			return
		}

		limit_string := data.Response.Header.Get("X-Page-Limit")
		if limit, err = strconv.Atoi(limit_string); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Could not convert X-Page-Limit value into int, unexcpected error: "+err.Error(),
			)
			return
		}

		total_string := data.Response.Header.Get("X-Page-Total")
		if total, err = strconv.Atoi(total_string); err != nil {
			resp.Diagnostics.AddError(
				"Error extracting HTTP Response Headers",
				"Could not convert X-Page-Total value into int, unexcpected error: "+err.Error(),
			)
			return
		}

		body, _ := io.ReadAll(data.Response.Body)
		mist_deviceprofiles := []models.DeviceprofileGateway{}
		json.Unmarshal(body, &mist_deviceprofiles)

		diags = datasource_org_deviceprofiles_gateway.SdkToTerraform(ctx, &mist_deviceprofiles, &elements)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}

	dataSet, diags := types.SetValue(datasource_org_deviceprofiles_gateway.OrgDeviceprofilesGatewayValue{}.Type(ctx), elements)
	if diags != nil {
		diags.Append(diags...)
	}

	if err := resp.State.SetAttribute(ctx, path.Root("deviceprofiles"), dataSet); err != nil {
		resp.Diagnostics.Append(err...)
		return
	}
}
