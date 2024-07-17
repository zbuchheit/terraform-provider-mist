package resource_org_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/tmunzer/mistapi-go/mistapi/models"
)

func jcloudTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d JcloudValue) *models.OrgSettingJcloud {
	data := models.OrgSettingJcloud{}

	if d.OrgApitoken.ValueStringPointer() != nil {
		data.OrgApitoken = d.OrgApitoken.ValueStringPointer()
	}

	if d.OrgApitokenName.ValueStringPointer() != nil {
		data.OrgApitokenName = d.OrgApitokenName.ValueStringPointer()
	}

	if d.OrgId.ValueStringPointer() != nil {
		data.OrgId = d.OrgId.ValueStringPointer()
	}

	return &data
}
