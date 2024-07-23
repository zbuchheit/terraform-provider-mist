// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_org_service

import (
	"context"
	"fmt"
	"github.com/Juniper/terraform-provider-mist/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OrgServiceResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"addresses": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "if `type`==`custom`, ip subnets",
				MarkdownDescription: "if `type`==`custom`, ip subnets",
				Validators: []validator.List{
					listvalidator.ValueStringsAre(stringvalidator.Any(mistvalidator.ParseCidr(true, false), mistvalidator.ParseVar())),
				},
				Default: listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
			},
			"app_categories": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "when `type`==`app_categories`\nlist of application categories are available through /api/v1/const/app_categories",
				MarkdownDescription: "when `type`==`app_categories`\nlist of application categories are available through /api/v1/const/app_categories",
				Default:             listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
			},
			"app_subcategories": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "when `type`==`app_categories`\nlist of application categories are available through /api/v1/const/app_subcategories",
				MarkdownDescription: "when `type`==`app_categories`\nlist of application categories are available through /api/v1/const/app_subcategories",
				Default:             listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
			},
			"apps": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "when `type`==`apps`\nlist of applications are available through:\n  - /api/v1/const/applications,\n  - /api/v1/const/gateway_applications\n  - /insight/top_app_by-bytes?wired=true",
				MarkdownDescription: "when `type`==`apps`\nlist of applications are available through:\n  - /api/v1/const/applications,\n  - /api/v1/const/gateway_applications\n  - /insight/top_app_by-bytes?wired=true",
				Default:             listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"dscp": schema.Int64Attribute{
				Optional:            true,
				Description:         "when `traffic_type`==`custom`",
				MarkdownDescription: "when `traffic_type`==`custom`",
			},
			"failover_policy": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"",
						"revertable",
						"non_revertable",
						"none",
					),
				},
				Default: stringdefault.StaticString("revertable"),
			},
			"hostnames": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "if `type`==`custom`, web filtering",
				MarkdownDescription: "if `type`==`custom`, web filtering",
				Default:             listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"max_jitter": schema.Int64Attribute{
				Optional:            true,
				Description:         "when `traffic_type`==`custom`, for uplink selection",
				MarkdownDescription: "when `traffic_type`==`custom`, for uplink selection",
			},
			"max_latency": schema.Int64Attribute{
				Optional:            true,
				Description:         "when `traffic_type`==`custom`, for uplink selection",
				MarkdownDescription: "when `traffic_type`==`custom`, for uplink selection",
			},
			"max_loss": schema.Int64Attribute{
				Optional:            true,
				Description:         "when `traffic_type`==`custom`, for uplink selection",
				MarkdownDescription: "when `traffic_type`==`custom`, for uplink selection",
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.All(stringvalidator.LengthBetween(2, 32), mistvalidator.ParseName()),
				},
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"sle_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "whether to enable measure SLE",
				MarkdownDescription: "whether to enable measure SLE",
				Default:             booldefault.StaticBool(false),
			},
			"specs": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port_range": schema.StringAttribute{
							Optional: true,
						},
						"protocol": schema.StringAttribute{
							Optional:            true,
							Description:         "`https`/ `tcp` / `udp` / `icmp` / `gre` / `any` / `:protocol_number`.\n`protocol_number` is between 1-254",
							MarkdownDescription: "`https`/ `tcp` / `udp` / `icmp` / `gre` / `any` / `:protocol_number`.\n`protocol_number` is between 1-254",
							Validators: []validator.String{
								stringvalidator.Any(stringvalidator.OneOf("https", "tcp", "udp", "icmp", "gre", "any"), mistvalidator.ParseInt(1, 254)),
							},
							Default: stringdefault.StaticString("any"),
						},
					},
					CustomType: SpecsType{
						ObjectType: types.ObjectType{
							AttrTypes: SpecsValue{}.AttributeTypes(ctx),
						},
					},
				},
				Optional: true,
			},
			"ssr_relaxed_tcp_state_enforcement": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(false),
			},
			"traffic_class": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "when `traffic_type`==`custom`",
				MarkdownDescription: "when `traffic_type`==`custom`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"",
						"best_effort",
						"high",
						"medium",
						"low",
					),
				},
				Default: stringdefault.StaticString("best_effort"),
			},
			"traffic_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "values from `/api/v1/consts/traffic_types`\n* when `type`==`apps`, we'll choose traffic_type automatically\n* when `type`==`addresses` or `type`==`hostnames`, you can provide your own settings (optional)",
				MarkdownDescription: "values from `/api/v1/consts/traffic_types`\n* when `type`==`apps`, we'll choose traffic_type automatically\n* when `type`==`addresses` or `type`==`hostnames`, you can provide your own settings (optional)",
				Default:             stringdefault.StaticString("data_best_effort"),
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"",
						"apps",
						"app_categories",
						"custom",
						"urls",
					),
				},
				Default: stringdefault.StaticString("custom"),
			},
			"urls": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "when `type`==`urls\nno need for spec as URL can encode the ports being used`",
				MarkdownDescription: "when `type`==`urls\nno need for spec as URL can encode the ports being used`",
				Default:             listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
			},
		},
	}
}

type OrgServiceModel struct {
	Addresses                     types.List   `tfsdk:"addresses"`
	AppCategories                 types.List   `tfsdk:"app_categories"`
	AppSubcategories              types.List   `tfsdk:"app_subcategories"`
	Apps                          types.List   `tfsdk:"apps"`
	Description                   types.String `tfsdk:"description"`
	Dscp                          types.Int64  `tfsdk:"dscp"`
	FailoverPolicy                types.String `tfsdk:"failover_policy"`
	Hostnames                     types.List   `tfsdk:"hostnames"`
	Id                            types.String `tfsdk:"id"`
	MaxJitter                     types.Int64  `tfsdk:"max_jitter"`
	MaxLatency                    types.Int64  `tfsdk:"max_latency"`
	MaxLoss                       types.Int64  `tfsdk:"max_loss"`
	Name                          types.String `tfsdk:"name"`
	OrgId                         types.String `tfsdk:"org_id"`
	SleEnabled                    types.Bool   `tfsdk:"sle_enabled"`
	Specs                         types.List   `tfsdk:"specs"`
	SsrRelaxedTcpStateEnforcement types.Bool   `tfsdk:"ssr_relaxed_tcp_state_enforcement"`
	TrafficClass                  types.String `tfsdk:"traffic_class"`
	TrafficType                   types.String `tfsdk:"traffic_type"`
	Type                          types.String `tfsdk:"type"`
	Urls                          types.List   `tfsdk:"urls"`
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

	if diags.HasError() {
		return nil, diags
	}

	return SpecsValue{
		PortRange: portRangeVal,
		Protocol:  protocolVal,
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

	if diags.HasError() {
		return NewSpecsValueUnknown(), diags
	}

	return SpecsValue{
		PortRange: portRangeVal,
		Protocol:  protocolVal,
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
	state     attr.ValueState
}

func (v SpecsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["port_range"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["protocol"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

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

	attributeTypes := map[string]attr.Type{
		"port_range": basetypes.StringType{},
		"protocol":   basetypes.StringType{},
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
	}
}
