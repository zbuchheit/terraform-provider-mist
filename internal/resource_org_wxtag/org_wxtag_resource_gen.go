// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_org_wxtag

import (
	"context"
	"fmt"
	"github.com/Juniper/terraform-provider-mist/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OrgWxtagResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"mac": schema.StringAttribute{
				Optional:            true,
				Description:         "if `type`==`client`, Client MAC Address",
				MarkdownDescription: "if `type`==`client`, Client MAC Address",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("client")),
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("client")),
				},
			},
			"match": schema.StringAttribute{
				Optional:            true,
				Description:         "required if `type`==`match`. enum: `ap_id`, `app`, `asset_mac`, `client_mac`, `hostname`, `ip_range_subnet`, `port`, `psk_name`, `psk_role`, `radius_attr`, `radius_class`, `radius_group`, `radius_username`, `sdkclient_uuid`, `wlan_id`",
				MarkdownDescription: "required if `type`==`match`. enum: `ap_id`, `app`, `asset_mac`, `client_mac`, `hostname`, `ip_range_subnet`, `port`, `psk_name`, `psk_role`, `radius_attr`, `radius_class`, `radius_group`, `radius_username`, `sdkclient_uuid`, `wlan_id`",
				Validators: []validator.String{
					stringvalidator.OneOf("ap_id", "app", "asset_mac", "client_mac", "hostname", "ip_range_subnet", "port", "psk_name", "psk_role", "radius_attr", "radius_class", "radius_group", "radius_username", "sdkclient_uuid", "wlan_id"),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("match")),
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("match")),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "The name",
				MarkdownDescription: "The name",
			},
			"op": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "required if `type`==`match`, type of tag (inclusive/exclusive). enum: `in`, `not_in`",
				MarkdownDescription: "required if `type`==`match`, type of tag (inclusive/exclusive). enum: `in`, `not_in`",
				Validators: []validator.String{
					stringvalidator.OneOf("in", "not_in"),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("match")),
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("match")),
				},
				Default: stringdefault.StaticString("in"),
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"specs": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_range": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "matched destination port, \"0\" means any",
							MarkdownDescription: "matched destination port, \"0\" means any",
							Validators: []validator.String{
								stringvalidator.Any(
									mistvalidator.ParseInt(0, 65535),
									mistvalidator.ParseRangeOfInt(0, 65535, true),
								),
							},
							Default: stringdefault.StaticString("0"),
						},
						"protocol": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Description:         "tcp / udp / icmp / gre / any / \":protocol_number\", `protocol_number` is between 1-254",
							MarkdownDescription: "tcp / udp / icmp / gre / any / \":protocol_number\", `protocol_number` is between 1-254",
							Validators: []validator.String{
								stringvalidator.Any(stringvalidator.OneOf("tcp", "udp", "icmp", "gre", "any"), mistvalidator.ParseInt(1, 254)),
							},
							Default: stringdefault.StaticString("any"),
						},
						"subnets": schema.ListAttribute{
							ElementType:         types.StringType,
							Required:            true,
							Description:         "matched destination subnets and/or IP Addresses",
							MarkdownDescription: "matched destination subnets and/or IP Addresses",
							Validators: []validator.List{
								listvalidator.SizeAtLeast(1),
								listvalidator.ValueStringsAre(
									stringvalidator.Any(
										mistvalidator.ParseCidr(true, true),
										mistvalidator.ParseIp(true, true),
										mistvalidator.ParseVar(),
									),
								),
							},
						},
					},
					CustomType: SpecsType{
						ObjectType: types.ObjectType{
							AttrTypes: SpecsValue{}.AttributeTypes(ctx),
						},
					},
					Validators: []validator.Object{
						mistvalidator.AtLeastNOf(
							1,
							path.MatchRelative().AtName("port_range"),
							path.MatchRelative().AtName("protocol"),
							path.MatchRelative().AtName("subnets"),
						),
					},
				},
				Optional:            true,
				Description:         "if `type`==`spec`",
				MarkdownDescription: "if `type`==`spec`",
				Validators: []validator.List{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("spec")),
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("spec")),
				},
			},
			"type": schema.StringAttribute{
				Required:            true,
				Description:         "enum: `client`, `match`, `resource`, `spec`, `subnet`, `vlan`",
				MarkdownDescription: "enum: `client`, `match`, `resource`, `spec`, `subnet`, `vlan`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"",
						"client",
						"match",
						"resource",
						"spec",
						"subnet",
						"vlan",
					),
				},
			},
			"values": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "required if `type`==`match` and\n  * `match`==`ap_id`: list of AP IDs\n  * `match`==`app`: list of Application Names\n  * `match`==`asset_mac`: list of Asset MAC Addresses\n  * `match`==`client_mac`: list of Client MAC Addresses\n  * `match`==`hostname`: list of Resources Hostnames\n  * `match`==`ip_range_subnet`: list of IP Addresses and/or CIDRs\n  * `match`==`psk_name`: list of PSK Names\n  * `match`==`psk_role`: list of PSK Roles\n  * `match`==`port`: list of Ports or Port Ranges\n  * `match`==`radius_attr`: list of RADIUS Attributes. The values are [ “6=1”, “26=10.2.3.4” ], this support other RADIUS attributes where we know the type\n  * `match`==`radius_class`: list of RADIUS Classes. This matches the ATTR-Class(25)\n  * `match`==`radius_group`: list of RADIUS Groups. This is a smart tag that matches RADIUS-Filter-ID, Airespace-ACL-Name (VendorID=14179, VendorType=6) / Aruba-User-Role (VendorID=14823, VendorType=1)\n  * `match`==`radius_username`: list of RADIUS Usernames. This matches the ATTR-User-Name(1)\n  * `match`==`sdkclient_uuid`: list of SDK UUIDs\n  * `match`==`wlan_id`: list of WLAN IDs\n\n**Notes**:\nVariables are not allowed",
				MarkdownDescription: "required if `type`==`match` and\n  * `match`==`ap_id`: list of AP IDs\n  * `match`==`app`: list of Application Names\n  * `match`==`asset_mac`: list of Asset MAC Addresses\n  * `match`==`client_mac`: list of Client MAC Addresses\n  * `match`==`hostname`: list of Resources Hostnames\n  * `match`==`ip_range_subnet`: list of IP Addresses and/or CIDRs\n  * `match`==`psk_name`: list of PSK Names\n  * `match`==`psk_role`: list of PSK Roles\n  * `match`==`port`: list of Ports or Port Ranges\n  * `match`==`radius_attr`: list of RADIUS Attributes. The values are [ “6=1”, “26=10.2.3.4” ], this support other RADIUS attributes where we know the type\n  * `match`==`radius_class`: list of RADIUS Classes. This matches the ATTR-Class(25)\n  * `match`==`radius_group`: list of RADIUS Groups. This is a smart tag that matches RADIUS-Filter-ID, Airespace-ACL-Name (VendorID=14179, VendorType=6) / Aruba-User-Role (VendorID=14823, VendorType=1)\n  * `match`==`radius_username`: list of RADIUS Usernames. This matches the ATTR-User-Name(1)\n  * `match`==`sdkclient_uuid`: list of SDK UUIDs\n  * `match`==`wlan_id`: list of WLAN IDs\n\n**Notes**:\nVariables are not allowed",
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("ap_id")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("app")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("asset_mac")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("client_mac")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("hostname")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("ip_range_subnet")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("psk_name")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("psk_role")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("port")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("radius_attr")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("radius_class")),

					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("radius_group")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("radius_username")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("sdkclient_uuid")),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("match"), types.StringValue("wlan_id")),
				},
			},
			"vlan_id": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.Any(mistvalidator.ParseInt(1, 4094), mistvalidator.ParseVar()),
				},
			},
		},
	}
}

type OrgWxtagModel struct {
	Id     types.String `tfsdk:"id"`
	Mac    types.String `tfsdk:"mac"`
	Match  types.String `tfsdk:"match"`
	Name   types.String `tfsdk:"name"`
	Op     types.String `tfsdk:"op"`
	OrgId  types.String `tfsdk:"org_id"`
	Specs  types.List   `tfsdk:"specs"`
	Type   types.String `tfsdk:"type"`
	Values types.List   `tfsdk:"values"`
	VlanId types.String `tfsdk:"vlan_id"`
}

var _ basetypes.ObjectTypable = SpecsType{}

type SpecsType struct {
	basetypes.ObjectType
}

func (t SpecsType) Equal(o attr.Type) bool {
	other, ok := o.(SpecsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t SpecsType) String() string {
	return "SpecsType"
}

func (t SpecsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	portRangeAttribute, ok := attributes["port_range"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`port_range is missing from object`)

		return nil, diags
	}

	portRangeVal, ok := portRangeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`port_range expected to be basetypes.StringValue, was: %T`, portRangeAttribute))
	}

	protocolAttribute, ok := attributes["protocol"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`protocol is missing from object`)

		return nil, diags
	}

	protocolVal, ok := protocolAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`protocol expected to be basetypes.StringValue, was: %T`, protocolAttribute))
	}

	subnetsAttribute, ok := attributes["subnets"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`subnets is missing from object`)

		return nil, diags
	}

	subnetsVal, ok := subnetsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`subnets expected to be basetypes.ListValue, was: %T`, subnetsAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return SpecsValue{
		PortRange: portRangeVal,
		Protocol:  protocolVal,
		Subnets:   subnetsVal,
		state:     attr.ValueStateKnown,
	}, diags
}

func NewSpecsValueNull() SpecsValue {
	return SpecsValue{
		state: attr.ValueStateNull,
	}
}

func NewSpecsValueUnknown() SpecsValue {
	return SpecsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewSpecsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (SpecsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing SpecsValue Attribute Value",
				"While creating a SpecsValue value, a missing attribute value was detected. "+
					"A SpecsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("SpecsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid SpecsValue Attribute Type",
				"While creating a SpecsValue value, an invalid attribute value was detected. "+
					"A SpecsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("SpecsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("SpecsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra SpecsValue Attribute Value",
				"While creating a SpecsValue value, an extra attribute value was detected. "+
					"A SpecsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra SpecsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewSpecsValueUnknown(), diags
	}

	portRangeAttribute, ok := attributes["port_range"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`port_range is missing from object`)

		return NewSpecsValueUnknown(), diags
	}

	portRangeVal, ok := portRangeAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`port_range expected to be basetypes.StringValue, was: %T`, portRangeAttribute))
	}

	protocolAttribute, ok := attributes["protocol"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`protocol is missing from object`)

		return NewSpecsValueUnknown(), diags
	}

	protocolVal, ok := protocolAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`protocol expected to be basetypes.StringValue, was: %T`, protocolAttribute))
	}

	subnetsAttribute, ok := attributes["subnets"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`subnets is missing from object`)

		return NewSpecsValueUnknown(), diags
	}

	subnetsVal, ok := subnetsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`subnets expected to be basetypes.ListValue, was: %T`, subnetsAttribute))
	}

	if diags.HasError() {
		return NewSpecsValueUnknown(), diags
	}

	return SpecsValue{
		PortRange: portRangeVal,
		Protocol:  protocolVal,
		Subnets:   subnetsVal,
		state:     attr.ValueStateKnown,
	}, diags
}

func NewSpecsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) SpecsValue {
	object, diags := NewSpecsValue(attributeTypes, attributes)

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

		panic("NewSpecsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t SpecsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewSpecsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewSpecsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewSpecsValueNull(), nil
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

	return NewSpecsValueMust(SpecsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t SpecsType) ValueType(ctx context.Context) attr.Value {
	return SpecsValue{}
}

var _ basetypes.ObjectValuable = SpecsValue{}

type SpecsValue struct {
	PortRange basetypes.StringValue `tfsdk:"port_range"`
	Protocol  basetypes.StringValue `tfsdk:"protocol"`
	Subnets   basetypes.ListValue   `tfsdk:"subnets"`
	state     attr.ValueState
}

func (v SpecsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["port_range"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["protocol"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["subnets"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.PortRange.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["port_range"] = val

		val, err = v.Protocol.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["protocol"] = val

		val, err = v.Subnets.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["subnets"] = val

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

func (v SpecsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v SpecsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v SpecsValue) String() string {
	return "SpecsValue"
}

func (v SpecsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	subnetsVal, d := types.ListValue(types.StringType, v.Subnets.Elements())

	diags.Append(d...)

	if d.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"port_range": basetypes.StringType{},
			"protocol":   basetypes.StringType{},
			"subnets": basetypes.ListType{
				ElemType: types.StringType,
			},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"port_range": basetypes.StringType{},
		"protocol":   basetypes.StringType{},
		"subnets": basetypes.ListType{
			ElemType: types.StringType,
		},
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
			"port_range": v.PortRange,
			"protocol":   v.Protocol,
			"subnets":    subnetsVal,
		})

	return objVal, diags
}

func (v SpecsValue) Equal(o attr.Value) bool {
	other, ok := o.(SpecsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.PortRange.Equal(other.PortRange) {
		return false
	}

	if !v.Protocol.Equal(other.Protocol) {
		return false
	}

	if !v.Subnets.Equal(other.Subnets) {
		return false
	}

	return true
}

func (v SpecsValue) Type(ctx context.Context) attr.Type {
	return SpecsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v SpecsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"port_range": basetypes.StringType{},
		"protocol":   basetypes.StringType{},
		"subnets": basetypes.ListType{
			ElemType: types.StringType,
		},
	}
}
