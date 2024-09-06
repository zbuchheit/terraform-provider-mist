package resource_device_switch

import (
	"context"

	"github.com/tmunzer/mistapi-go/mistapi/models"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	mist_transform "github.com/Juniper/terraform-provider-mist/internal/commons/utils"
)

func portMirroringTerraformToSdk(ctx context.Context, diags *diag.Diagnostics, d basetypes.MapValue) map[string]models.SwitchPortMirroringProperty {
	data := make(map[string]models.SwitchPortMirroringProperty)
	for item_name, item_value := range d.Elements() {
		var item_interface interface{} = item_value
		item_obj := item_interface.(PortMirroringValue)

		data_item := models.SwitchPortMirroringProperty{}
		if !item_obj.InputNetworksIngress.IsNull() && !item_obj.InputNetworksIngress.IsUnknown() {
			data_item.InputNetworksIngress = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.InputNetworksIngress)
		}
		if data_item.InputNetworksIngress == nil {
			data_item.InputNetworksIngress = make([]string, 0)
		}

		if !item_obj.InputPortIdsEgress.IsNull() && !item_obj.InputPortIdsEgress.IsUnknown() {
			data_item.InputPortIdsEgress = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.InputPortIdsEgress)
		}
		if data_item.InputPortIdsEgress == nil {
			data_item.InputPortIdsEgress = make([]string, 0)
		}

		if !item_obj.InputPortIdsIngress.IsNull() && !item_obj.InputPortIdsIngress.IsUnknown() {
			data_item.InputPortIdsIngress = mist_transform.ListOfStringTerraformToSdk(ctx, item_obj.InputPortIdsIngress)
		}
		if data_item.InputPortIdsIngress == nil {
			data_item.InputPortIdsIngress = make([]string, 0)
		}

		if item_obj.OutputNetwork.ValueStringPointer() != nil {
			data_item.OutputNetwork = models.ToPointer(item_obj.OutputNetwork.ValueString())
		}
		if item_obj.OutputPortId.ValueStringPointer() != nil {
			data_item.OutputPortId = models.ToPointer(item_obj.OutputPortId.ValueString())
		}

		data[item_name] = data_item
	}
	return data
}
