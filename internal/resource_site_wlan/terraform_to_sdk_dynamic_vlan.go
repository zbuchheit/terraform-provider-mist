package resource_site_wlan

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func dynamicVlanTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, plan DynamicVlanValue) *models.WlanDynamicVlan {

	var local_vlan_ids []models.WlanDynamicVlanLocalVlanIds
	for _, item := range plan.LocalVlanIds.Elements() {
		var item_interface interface{} = item
		i := item_interface.(basetypes.StringValue)
		j := models.WlanDynamicVlanLocalVlanIdsContainer.FromString(i.ValueString())
		local_vlan_ids = append(local_vlan_ids, j)
	}

	vlans := make(map[string]string)
	for k, v := range plan.Vlans.Elements() {
		var v_interface interface{} = v
		v_plan := v_interface.(basetypes.StringValue)
		vlans[k] = v_plan.ValueString()
	}

	data := models.WlanDynamicVlan{}
	v := models.ToPointer(models.WlanDynamicVlanDefaultVlanIdContainer.FromString(plan.DefaultVlanId.ValueString()))
	data.DefaultVlanId = models.NewOptional(v)
	data.Enabled = plan.Enabled.ValueBoolPointer()
	data.LocalVlanIds = local_vlan_ids
	data.Type = models.ToPointer(models.WlanDynamicVlanTypeEnum(string(plan.DynamicVlanType.ValueString())))
	data.Vlans = vlans

	return &data
}
