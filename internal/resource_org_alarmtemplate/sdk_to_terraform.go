package resource_org_alarmtemplate

import (
	"context"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func SdkToTerraform(ctx context.Context, data models.AlarmTemplate) (OrgAlarmtemplateModel, diag.Diagnostics) {
	var state OrgAlarmtemplateModel
	var diags diag.Diagnostics

	var delivery DeliveryValue
	var id types.String
	var name types.String
	var org_id types.String
	var rules types.Map = basetypes.NewMapNull(RulesValue{}.Type(ctx))

	delivery = deliverySdkToTerraform(ctx, &diags, &data.Delivery)
	id = types.StringValue(data.Id.String())
	if data.Name != nil {
		name = types.StringValue(*data.Name)
	}
	org_id = types.StringValue(data.OrgId.String())
	rules = rulesSdkToTerraform(ctx, &diags, data.Rules)

	state.Delivery = delivery
	state.Id = id
	state.Name = name
	state.OrgId = org_id
	state.Rules = rules

	return state, diags
}

func deliverySdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data *models.Delivery) DeliveryValue {
	var additional_emails types.List = mist_transform.ListOfStringSdkToTerraformEmpty(ctx)
	var enabled types.Bool
	var to_org_admins types.Bool
	var to_site_admins types.Bool

	if data != nil {
		if data.AdditionalEmails != nil && len(data.AdditionalEmails) > 0 {
			additional_emails = mist_transform.ListOfStringSdkToTerraform(ctx, data.AdditionalEmails)
		}
		enabled = types.BoolValue(data.Enabled)
		if data.ToOrgAdmins != nil {
			to_org_admins = types.BoolValue(*data.ToOrgAdmins)
		}
		if data.ToSiteAdmins != nil {
			to_site_admins = types.BoolValue(*data.ToSiteAdmins)
		}
	}

	data_map_attr_type := DeliveryValue{}.AttributeTypes(ctx)
	data_map_value := map[string]attr.Value{
		"additional_emails": additional_emails,
		"enabled":           enabled,
		"to_org_admins":     to_org_admins,
		"to_site_admins":    to_site_admins,
	}
	r, e := NewDeliveryValue(data_map_attr_type, data_map_value)
	diags.Append(e...)
	return r
}

func rulesSdkToTerraform(ctx context.Context, diags *diag.Diagnostics, data map[string]models.AlarmTemplateRule) basetypes.MapValue {
	rules_map := make(map[string]attr.Value)
	for k, v := range data {
		var delivery types.Object = basetypes.NewObjectNull(DeliveryValue{}.AttributeTypes(ctx))
		var enabled types.Bool

		if v.Delivery != nil {
			tmp, e := deliverySdkToTerraform(ctx, diags, v.Delivery).ToObjectValue(ctx)
			if e != nil {
				diags.Append(e...)
			} else {
				delivery = tmp
			}
		}
		if v.Enabled != nil {
			enabled = types.BoolValue(*v.Enabled)
		}

		data_map_attr_type := RulesValue{}.AttributeTypes(ctx)
		data_map_value := map[string]attr.Value{
			"delivery": delivery,
			"enabled":  enabled,
		}
		data, e := NewRulesValue(data_map_attr_type, data_map_value)
		diags.Append(e...)

		rules_map[k] = data
	}
	r, e := basetypes.NewMapValueFrom(ctx, RulesValue{}.Type(ctx), rules_map)
	if e != nil {
		diags.Append(e...)
	}
	return r
}