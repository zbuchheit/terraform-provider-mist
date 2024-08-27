// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_org_inventory

import (
	"context"
	"fmt"
	"strings"

	mistvalidator "github.com/Juniper/terraform-provider-mist/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func OrgInventoryResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"devices": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mac": schema.StringAttribute{
							Computed:            true,
							Optional:            true,
							Description:         "Device MAC address. Required to assign adopted devices to site. Cannot be specified when `claim_code` is used",
							MarkdownDescription: "Device MAC address. Required to assign adopted devices to site. Cannot be specified when `claim_code` is used",
							Validators: []validator.String{
								stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("claim_code")),
								mistvalidator.ParseMac(),
								mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("claim_code"), types.StringNull()),
							},
						},
						"model": schema.StringAttribute{
							Computed:            true,
							Description:         "Device model",
							MarkdownDescription: "Device model",
						},
						"org_id": schema.StringAttribute{
							Computed: true,
						},
						"claim_code": schema.StringAttribute{
							Computed:            true,
							Optional:            true,
							Description:         "Device Claim Code. Required for claimed devices",
							MarkdownDescription: "Device Claim Code. Required for claimed devices",
							Validators: []validator.String{
								stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("claim_code")),
								mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("mac"), types.StringNull()),
							},
						},
						"serial": schema.StringAttribute{
							Computed:            true,
							Description:         "Device serial",
							MarkdownDescription: "Device serial",
						},
						"site_id": schema.StringAttribute{
							Optional:            true,
							Description:         "Site ID. Used to assign device to a Site",
							MarkdownDescription: "Site ID. Used to assign device to a Site",
						},
						"type": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"",
									"ap",
									"switch",
									"gateway",
								),
							},
						},
						"vc_mac": schema.StringAttribute{
							Computed:            true,
							Description:         "Virtual Chassis MAC Address",
							MarkdownDescription: "Virtual Chassis MAC Address",
						},
						"hostname": schema.StringAttribute{
							Computed:            true,
							Description:         "Device Hostname",
							MarkdownDescription: "Device Hostname",
						},
						"id": schema.StringAttribute{
							Computed:            true,
							Description:         "Mist Device ID",
							MarkdownDescription: "Mist Device ID",
						},
					},
					CustomType: DevicesType{
						ObjectType: types.ObjectType{
							AttrTypes: DevicesValue{}.AttributeTypes(ctx),
						},
					},
				},
				Optional: true,
				Computed: true,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
				},
			},
		},
	}
}

type OrgInventoryModel struct {
	OrgId   types.String `tfsdk:"org_id"`
	Devices types.List   `tfsdk:"devices"`
}

var _ basetypes.ObjectTypable = DevicesType{}

type DevicesType struct {
	basetypes.ObjectType
}

func (t DevicesType) Equal(o attr.Type) bool {
	other, ok := o.(DevicesType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t DevicesType) String() string {
	return "DevicesType"
}

func (t DevicesType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	macAttribute, ok := attributes["mac"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`mac is missing from object`)

		return nil, diags
	}

	macVal, ok := macAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`mac expected to be basetypes.StringValue, was: %T`, macAttribute))
	}

	claim_codeAttribute, ok := attributes["claim_code"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`claim_code is missing from object`)

		return nil, diags
	}

	claim_codeVal, ok := claim_codeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`claim_code expected to be basetypes.StringValue, was: %T`, claim_codeAttribute))
	}

	modelAttribute, ok := attributes["model"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`model is missing from object`)

		return nil, diags
	}

	modelVal, ok := modelAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`model expected to be basetypes.StringValue, was: %T`, modelAttribute))
	}

	orgIdAttribute, ok := attributes["org_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`org_id is missing from object`)

		return nil, diags
	}

	orgIdVal, ok := orgIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`org_id expected to be basetypes.StringValue, was: %T`, orgIdAttribute))
	}

	serialAttribute, ok := attributes["serial"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`serial is missing from object`)

		return nil, diags
	}

	serialVal, ok := serialAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`serial expected to be basetypes.StringValue, was: %T`, serialAttribute))
	}

	siteIdAttribute, ok := attributes["site_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`site_id is missing from object`)

		return nil, diags
	}

	siteIdVal, ok := siteIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`site_id expected to be basetypes.StringValue, was: %T`, siteIdAttribute))
	}

	deviceTypeAttribute, ok := attributes["type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`type is missing from object`)

		return nil, diags
	}

	deviceTypeVal, ok := deviceTypeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`type expected to be basetypes.StringValue, was: %T`, deviceTypeAttribute))
	}

	vcMacAttribute, ok := attributes["vc_mac"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`vc_mac is missing from object`)

		return nil, diags
	}

	vcMacVal, ok := vcMacAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`vc_mac expected to be basetypes.StringValue, was: %T`, vcMacAttribute))
	}

	hostnameAttribute, ok := attributes["hostname"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`hostname is missing from object`)

		return nil, diags
	}

	hostnameVal, ok := hostnameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`hostname expected to be basetypes.StringValue, was: %T`, hostnameAttribute))
	}

	deviceIdAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return nil, diags
	}

	deviceIdVal, ok := deviceIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.StringValue, was: %T`, deviceIdAttribute))
	}

	return DevicesValue{
		Magic:      claim_codeVal,
		Mac:        macVal,
		Model:      modelVal,
		OrgId:      orgIdVal,
		Serial:     serialVal,
		SiteId:     siteIdVal,
		DeviceType: deviceTypeVal,
		VcMac:      vcMacVal,
		Hostname:   hostnameVal,
		Id:         deviceIdVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewDevicesValueNull() DevicesValue {
	return DevicesValue{
		state: attr.ValueStateNull,
	}
}

func NewDevicesValueUnknown() DevicesValue {
	return DevicesValue{
		state: attr.ValueStateUnknown,
	}
}

func NewDevicesValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (DevicesValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing DevicesValue Attribute Value",
				"While creating a DevicesValue value, a missing attribute value was detected. "+
					"A DevicesValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("DevicesValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid DevicesValue Attribute Type",
				"While creating a DevicesValue value, an invalid attribute value was detected. "+
					"A DevicesValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("DevicesValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("DevicesValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra DevicesValue Attribute Value",
				"While creating a DevicesValue value, an extra attribute value was detected. "+
					"A DevicesValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra DevicesValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewDevicesValueUnknown(), diags
	}

	claim_codeAttribute, ok := attributes["claim_code"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`claim_code is missing from object`)

		return NewDevicesValueUnknown(), diags
	}

	claim_codeVal, ok := claim_codeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`claim_code expected to be basetypes.StringValue, was: %T`, claim_codeAttribute))
	}

	macAttribute, ok := attributes["mac"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`mac is missing from object`)

		return NewDevicesValueUnknown(), diags
	}

	macVal, ok := macAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`mac expected to be basetypes.StringValue, was: %T`, macAttribute))
	}

	modelAttribute, ok := attributes["model"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`model is missing from object`)

		return NewDevicesValueUnknown(), diags
	}

	modelVal, ok := modelAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`model expected to be basetypes.StringValue, was: %T`, modelAttribute))
	}

	orgIdAttribute, ok := attributes["org_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`org_id is missing from object`)

		return NewDevicesValueUnknown(), diags
	}

	orgIdVal, ok := orgIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`org_id expected to be basetypes.StringValue, was: %T`, orgIdAttribute))
	}

	serialAttribute, ok := attributes["serial"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`serial is missing from object`)

		return NewDevicesValueUnknown(), diags
	}

	serialVal, ok := serialAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`serial expected to be basetypes.StringValue, was: %T`, serialAttribute))
	}

	siteIdAttribute, ok := attributes["site_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`site_id is missing from object`)

		return NewDevicesValueUnknown(), diags
	}

	siteIdVal, ok := siteIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`site_id expected to be basetypes.StringValue, was: %T`, siteIdAttribute))
	}

	deviceTypeAttribute, ok := attributes["type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`type is missing from object`)

		return NewDevicesValueUnknown(), diags
	}

	deviceTypeVal, ok := deviceTypeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`type expected to be basetypes.StringValue, was: %T`, deviceTypeAttribute))
	}

	vcMacAttribute, ok := attributes["vc_mac"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`vc_mac is missing from object`)

		return NewDevicesValueUnknown(), diags
	}

	vcMacVal, ok := vcMacAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`vc_mac expected to be basetypes.StringValue, was: %T`, vcMacAttribute))
	}

	hostnameAttribute, ok := attributes["hostname"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`hostname is missing from object`)

		return NewDevicesValueUnknown(), diags
	}

	hostnameVal, ok := hostnameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`hostname expected to be basetypes.StringValue, was: %T`, hostnameAttribute))
	}

	deviceIdAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewDevicesValueUnknown(), diags
	}

	deviceIdVal, ok := deviceIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.StringValue, was: %T`, deviceIdAttribute))
	}

	if diags.HasError() {
		return NewDevicesValueUnknown(), diags
	}

	return DevicesValue{
		Magic:      claim_codeVal,
		Mac:        macVal,
		Model:      modelVal,
		OrgId:      orgIdVal,
		Serial:     serialVal,
		SiteId:     siteIdVal,
		DeviceType: deviceTypeVal,
		VcMac:      vcMacVal,
		Hostname:   hostnameVal,
		Id:         deviceIdVal,
		state:      attr.ValueStateKnown,
	}, diags
}

func NewDevicesValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) DevicesValue {
	object, diags := NewDevicesValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewDevicesValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t DevicesType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewDevicesValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewDevicesValueUnknown(), nil
	}

	if in.IsNull() {
		return NewDevicesValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewDevicesValueMust(DevicesValue{}.AttributeTypes(ctx), attributes), nil
}

func (t DevicesType) ValueType(ctx context.Context) attr.Value {
	return DevicesValue{}
}

var _ basetypes.ObjectValuable = DevicesValue{}

type DevicesValue struct {
	Magic      basetypes.StringValue `tfsdk:"claim_code"`
	Mac        basetypes.StringValue `tfsdk:"mac"`
	Model      basetypes.StringValue `tfsdk:"model"`
	OrgId      basetypes.StringValue `tfsdk:"org_id"`
	Serial     basetypes.StringValue `tfsdk:"serial"`
	SiteId     basetypes.StringValue `tfsdk:"site_id"`
	DeviceType basetypes.StringValue `tfsdk:"type"`
	VcMac      basetypes.StringValue `tfsdk:"vc_mac"`
	Hostname   basetypes.StringValue `tfsdk:"hostname"`
	Id         basetypes.StringValue `tfsdk:"id"`
	state      attr.ValueState
}

func (v DevicesValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["claim_code"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["mac"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["model"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["org_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["serial"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["site_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["type"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["vc_mac"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["hostname"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["id"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 8)

		val, err = v.Magic.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["claim_code"] = val

		val, err = v.Mac.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["mac"] = val

		val, err = v.Model.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["model"] = val

		val, err = v.OrgId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["org_id"] = val

		val, err = v.Serial.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["serial"] = val

		val, err = v.SiteId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["site_id"] = val

		val, err = v.DeviceType.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["type"] = val

		val, err = v.VcMac.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["vc_mac"] = val

		val, err = v.Hostname.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["hostname"] = val

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v DevicesValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v DevicesValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v DevicesValue) String() string {
	return "DevicesValue"
}

func (v DevicesValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"claim_code": basetypes.StringType{},
		"mac":        basetypes.StringType{},
		"model":      basetypes.StringType{},
		"org_id":     basetypes.StringType{},
		"serial":     basetypes.StringType{},
		"site_id":    basetypes.StringType{},
		"type":       basetypes.StringType{},
		"vc_mac":     basetypes.StringType{},
		"hostname":   basetypes.StringType{},
		"id":         basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"claim_code": v.Magic,
			"mac":        v.Mac,
			"model":      v.Model,
			"org_id":     v.OrgId,
			"serial":     v.Serial,
			"site_id":    v.SiteId,
			"type":       v.DeviceType,
			"vc_mac":     v.VcMac,
			"hostname":   v.Hostname,
			"id":         v.Id,
		})

	return objVal, diags
}

func (v DevicesValue) Equal(o attr.Value) bool {
	other, ok := o.(DevicesValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Magic.Equal(other.Magic) {
		return false
	}

	if !v.Mac.Equal(other.Mac) {
		return false
	}

	if !v.Model.Equal(other.Model) {
		return false
	}

	if !v.OrgId.Equal(other.OrgId) {
		return false
	}

	if !v.Serial.Equal(other.Serial) {
		return false
	}

	if !v.SiteId.Equal(other.SiteId) {
		return false
	}

	if !v.DeviceType.Equal(other.DeviceType) {
		return false
	}

	if !v.Hostname.Equal(other.Hostname) {
		return false
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.VcMac.Equal(other.VcMac) {
		return false
	}

	return true
}

func (v DevicesValue) Type(ctx context.Context) attr.Type {
	return DevicesType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v DevicesValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"claim_code": basetypes.StringType{},
		"mac":        basetypes.StringType{},
		"model":      basetypes.StringType{},
		"org_id":     basetypes.StringType{},
		"serial":     basetypes.StringType{},
		"site_id":    basetypes.StringType{},
		"type":       basetypes.StringType{},
		"vc_mac":     basetypes.StringType{},
		"hostname":   basetypes.StringType{},
		"id":         basetypes.StringType{},
	}
}
