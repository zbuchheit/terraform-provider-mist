package resource_org_wlantemplate

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func TerraformToSdk(ctx context.Context, plan *OrgWlantemplateModel) (*models.Template, diag.Diagnostics) {
	var diags diag.Diagnostics
	unset := make(map[string]interface{})

	data := models.Template{}

	data.Name = plan.Name.ValueString()

	if !plan.Applies.IsNull() && !plan.Applies.IsUnknown() {
		data.Applies = appliesTerraformToSdk(ctx, &diags, plan.Applies)
	} else {
		unset["-applies"] = ""
	}

	if !plan.DeviceprofileIds.IsNull() && !plan.DeviceprofileIds.IsUnknown() {
		data.DeviceprofileIds = mist_transform.ListOfUuidTerraformToSdk(ctx, plan.DeviceprofileIds)
	} else {
		unset["-deviceprofile_ids"] = ""
	}

	if !plan.Exceptions.IsNull() && !plan.Exceptions.IsUnknown() {
		data.Exceptions = exceptionsTerraformToSdk(ctx, &diags, plan.Exceptions)
	} else {
		unset["-exceptions"] = ""
	}

	if !plan.FilterByDeviceprofile.IsNull() && !plan.FilterByDeviceprofile.IsUnknown() {
		data.FilterByDeviceprofile = models.ToPointer(plan.FilterByDeviceprofile.ValueBool())
	} else {
		unset["-filter_by_deviceprofile"] = ""
	}

	data.AdditionalProperties = unset
	return &data, diags
}
