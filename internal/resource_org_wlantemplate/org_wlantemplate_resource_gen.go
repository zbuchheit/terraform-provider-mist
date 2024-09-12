// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_org_wlantemplate

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OrgWlantemplateResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"applies": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"org_id": schema.StringAttribute{
						Optional: true,
					},
					"site_ids": schema.ListAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						Computed:            true,
						Description:         "list of site ids",
						MarkdownDescription: "list of site ids",
					},
					"sitegroup_ids": schema.ListAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						Computed:            true,
						Description:         "list of sitegroup ids",
						MarkdownDescription: "list of sitegroup ids",
					},
				},
				CustomType: AppliesType{
					ObjectType: types.ObjectType{
						AttrTypes: AppliesValue{}.AttributeTypes(ctx),
					},
				},
				Optional:            true,
				Computed:            true,
				Description:         "where this template should be applied to, can be org_id, site_ids, sitegroup_ids",
				MarkdownDescription: "where this template should be applied to, can be org_id, site_ids, sitegroup_ids",
			},
			"deviceprofile_ids": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Computed:            true,
				Description:         "list of Device Profile ids",
				MarkdownDescription: "list of Device Profile ids",
				Default:             listdefault.StaticValue(types.ListValueMust(types.StringType, []attr.Value{})),
			},
			"exceptions": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"site_ids": schema.ListAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						Computed:            true,
						Description:         "list of site ids",
						MarkdownDescription: "list of site ids",
					},
					"sitegroup_ids": schema.ListAttribute{
						ElementType:         types.StringType,
						Optional:            true,
						Computed:            true,
						Description:         "list of sitegroup ids",
						MarkdownDescription: "list of sitegroup ids",
					},
				},
				CustomType: ExceptionsType{
					ObjectType: types.ObjectType{
						AttrTypes: ExceptionsValue{}.AttributeTypes(ctx),
					},
				},
				Optional:            true,
				Computed:            true,
				Description:         "where this template should not be applied to (takes precedence)",
				MarkdownDescription: "where this template should not be applied to (takes precedence)",
			},
			"filter_by_deviceprofile": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "whether to further filter by Device Profile",
				MarkdownDescription: "whether to further filter by Device Profile",
				Default:             booldefault.StaticBool(false),
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

type OrgWlantemplateModel struct {
	Applies               AppliesValue    `tfsdk:"applies"`
	DeviceprofileIds      types.List      `tfsdk:"deviceprofile_ids"`
	Exceptions            ExceptionsValue `tfsdk:"exceptions"`
	FilterByDeviceprofile types.Bool      `tfsdk:"filter_by_deviceprofile"`
	Id                    types.String    `tfsdk:"id"`
	Name                  types.String    `tfsdk:"name"`
	OrgId                 types.String    `tfsdk:"org_id"`
}

var _ basetypes.ObjectTypable = AppliesType{}

type AppliesType struct {
	basetypes.ObjectType
}

func (t AppliesType) Equal(o attr.Type) bool {
	other, ok := o.(AppliesType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t AppliesType) String() string {
	return "AppliesType"
}

func (t AppliesType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

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

	siteIdsAttribute, ok := attributes["site_ids"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`site_ids is missing from object`)

		return nil, diags
	}

	siteIdsVal, ok := siteIdsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`site_ids expected to be basetypes.ListValue, was: %T`, siteIdsAttribute))
	}

	sitegroupIdsAttribute, ok := attributes["sitegroup_ids"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`sitegroup_ids is missing from object`)

		return nil, diags
	}

	sitegroupIdsVal, ok := sitegroupIdsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`sitegroup_ids expected to be basetypes.ListValue, was: %T`, sitegroupIdsAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return AppliesValue{
		OrgId:        orgIdVal,
		SiteIds:      siteIdsVal,
		SitegroupIds: sitegroupIdsVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewAppliesValueNull() AppliesValue {
	return AppliesValue{
		state: attr.ValueStateNull,
	}
}

func NewAppliesValueUnknown() AppliesValue {
	return AppliesValue{
		state: attr.ValueStateUnknown,
	}
}

func NewAppliesValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (AppliesValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing AppliesValue Attribute Value",
				"While creating a AppliesValue value, a missing attribute value was detected. "+
					"A AppliesValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("AppliesValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid AppliesValue Attribute Type",
				"While creating a AppliesValue value, an invalid attribute value was detected. "+
					"A AppliesValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("AppliesValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("AppliesValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra AppliesValue Attribute Value",
				"While creating a AppliesValue value, an extra attribute value was detected. "+
					"A AppliesValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra AppliesValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewAppliesValueUnknown(), diags
	}

	orgIdAttribute, ok := attributes["org_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`org_id is missing from object`)

		return NewAppliesValueUnknown(), diags
	}

	orgIdVal, ok := orgIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`org_id expected to be basetypes.StringValue, was: %T`, orgIdAttribute))
	}

	siteIdsAttribute, ok := attributes["site_ids"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`site_ids is missing from object`)

		return NewAppliesValueUnknown(), diags
	}

	siteIdsVal, ok := siteIdsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`site_ids expected to be basetypes.ListValue, was: %T`, siteIdsAttribute))
	}

	sitegroupIdsAttribute, ok := attributes["sitegroup_ids"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`sitegroup_ids is missing from object`)

		return NewAppliesValueUnknown(), diags
	}

	sitegroupIdsVal, ok := sitegroupIdsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`sitegroup_ids expected to be basetypes.ListValue, was: %T`, sitegroupIdsAttribute))
	}

	if diags.HasError() {
		return NewAppliesValueUnknown(), diags
	}

	return AppliesValue{
		OrgId:        orgIdVal,
		SiteIds:      siteIdsVal,
		SitegroupIds: sitegroupIdsVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewAppliesValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) AppliesValue {
	object, diags := NewAppliesValue(attributeTypes, attributes)

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

		panic("NewAppliesValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t AppliesType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewAppliesValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewAppliesValueUnknown(), nil
	}

	if in.IsNull() {
		return NewAppliesValueNull(), nil
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

	return NewAppliesValueMust(AppliesValue{}.AttributeTypes(ctx), attributes), nil
}

func (t AppliesType) ValueType(ctx context.Context) attr.Value {
	return AppliesValue{}
}

var _ basetypes.ObjectValuable = AppliesValue{}

type AppliesValue struct {
	OrgId        basetypes.StringValue `tfsdk:"org_id"`
	SiteIds      basetypes.ListValue   `tfsdk:"site_ids"`
	SitegroupIds basetypes.ListValue   `tfsdk:"sitegroup_ids"`
	state        attr.ValueState
}

func (v AppliesValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["org_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["site_ids"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["sitegroup_ids"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.OrgId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["org_id"] = val

		val, err = v.SiteIds.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["site_ids"] = val

		val, err = v.SitegroupIds.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["sitegroup_ids"] = val

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

func (v AppliesValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v AppliesValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v AppliesValue) String() string {
	return "AppliesValue"
}

func (v AppliesValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	siteIdsVal, d := types.ListValue(types.StringType, v.SiteIds.Elements())

	diags.Append(d...)

	if d.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"org_id": basetypes.StringType{},
			"site_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"sitegroup_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		}), diags
	}

	sitegroupIdsVal, d := types.ListValue(types.StringType, v.SitegroupIds.Elements())

	diags.Append(d...)

	if d.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"org_id": basetypes.StringType{},
			"site_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"sitegroup_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"org_id": basetypes.StringType{},
		"site_ids": basetypes.ListType{
			ElemType: types.StringType,
		},
		"sitegroup_ids": basetypes.ListType{
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
			"org_id":        v.OrgId,
			"site_ids":      siteIdsVal,
			"sitegroup_ids": sitegroupIdsVal,
		})

	return objVal, diags
}

func (v AppliesValue) Equal(o attr.Value) bool {
	other, ok := o.(AppliesValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.OrgId.Equal(other.OrgId) {
		return false
	}

	if !v.SiteIds.Equal(other.SiteIds) {
		return false
	}

	if !v.SitegroupIds.Equal(other.SitegroupIds) {
		return false
	}

	return true
}

func (v AppliesValue) Type(ctx context.Context) attr.Type {
	return AppliesType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v AppliesValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"org_id": basetypes.StringType{},
		"site_ids": basetypes.ListType{
			ElemType: types.StringType,
		},
		"sitegroup_ids": basetypes.ListType{
			ElemType: types.StringType,
		},
	}
}

var _ basetypes.ObjectTypable = ExceptionsType{}

type ExceptionsType struct {
	basetypes.ObjectType
}

func (t ExceptionsType) Equal(o attr.Type) bool {
	other, ok := o.(ExceptionsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ExceptionsType) String() string {
	return "ExceptionsType"
}

func (t ExceptionsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	siteIdsAttribute, ok := attributes["site_ids"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`site_ids is missing from object`)

		return nil, diags
	}

	siteIdsVal, ok := siteIdsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`site_ids expected to be basetypes.ListValue, was: %T`, siteIdsAttribute))
	}

	sitegroupIdsAttribute, ok := attributes["sitegroup_ids"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`sitegroup_ids is missing from object`)

		return nil, diags
	}

	sitegroupIdsVal, ok := sitegroupIdsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`sitegroup_ids expected to be basetypes.ListValue, was: %T`, sitegroupIdsAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ExceptionsValue{
		SiteIds:      siteIdsVal,
		SitegroupIds: sitegroupIdsVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewExceptionsValueNull() ExceptionsValue {
	return ExceptionsValue{
		state: attr.ValueStateNull,
	}
}

func NewExceptionsValueUnknown() ExceptionsValue {
	return ExceptionsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewExceptionsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ExceptionsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ExceptionsValue Attribute Value",
				"While creating a ExceptionsValue value, a missing attribute value was detected. "+
					"A ExceptionsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ExceptionsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ExceptionsValue Attribute Type",
				"While creating a ExceptionsValue value, an invalid attribute value was detected. "+
					"A ExceptionsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ExceptionsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ExceptionsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ExceptionsValue Attribute Value",
				"While creating a ExceptionsValue value, an extra attribute value was detected. "+
					"A ExceptionsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ExceptionsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewExceptionsValueUnknown(), diags
	}

	siteIdsAttribute, ok := attributes["site_ids"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`site_ids is missing from object`)

		return NewExceptionsValueUnknown(), diags
	}

	siteIdsVal, ok := siteIdsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`site_ids expected to be basetypes.ListValue, was: %T`, siteIdsAttribute))
	}

	sitegroupIdsAttribute, ok := attributes["sitegroup_ids"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`sitegroup_ids is missing from object`)

		return NewExceptionsValueUnknown(), diags
	}

	sitegroupIdsVal, ok := sitegroupIdsAttribute.(basetypes.ListValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`sitegroup_ids expected to be basetypes.ListValue, was: %T`, sitegroupIdsAttribute))
	}

	if diags.HasError() {
		return NewExceptionsValueUnknown(), diags
	}

	return ExceptionsValue{
		SiteIds:      siteIdsVal,
		SitegroupIds: sitegroupIdsVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewExceptionsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ExceptionsValue {
	object, diags := NewExceptionsValue(attributeTypes, attributes)

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

		panic("NewExceptionsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ExceptionsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewExceptionsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewExceptionsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewExceptionsValueNull(), nil
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

	return NewExceptionsValueMust(ExceptionsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ExceptionsType) ValueType(ctx context.Context) attr.Value {
	return ExceptionsValue{}
}

var _ basetypes.ObjectValuable = ExceptionsValue{}

type ExceptionsValue struct {
	SiteIds      basetypes.ListValue `tfsdk:"site_ids"`
	SitegroupIds basetypes.ListValue `tfsdk:"sitegroup_ids"`
	state        attr.ValueState
}

func (v ExceptionsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["site_ids"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)
	attrTypes["sitegroup_ids"] = basetypes.ListType{
		ElemType: types.StringType,
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.SiteIds.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["site_ids"] = val

		val, err = v.SitegroupIds.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["sitegroup_ids"] = val

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

func (v ExceptionsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ExceptionsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ExceptionsValue) String() string {
	return "ExceptionsValue"
}

func (v ExceptionsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	siteIdsVal, d := types.ListValue(types.StringType, v.SiteIds.Elements())

	diags.Append(d...)

	if d.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"site_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"sitegroup_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		}), diags
	}

	sitegroupIdsVal, d := types.ListValue(types.StringType, v.SitegroupIds.Elements())

	diags.Append(d...)

	if d.HasError() {
		return types.ObjectUnknown(map[string]attr.Type{
			"site_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
			"sitegroup_ids": basetypes.ListType{
				ElemType: types.StringType,
			},
		}), diags
	}

	attributeTypes := map[string]attr.Type{
		"site_ids": basetypes.ListType{
			ElemType: types.StringType,
		},
		"sitegroup_ids": basetypes.ListType{
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
			"site_ids":      siteIdsVal,
			"sitegroup_ids": sitegroupIdsVal,
		})

	return objVal, diags
}

func (v ExceptionsValue) Equal(o attr.Value) bool {
	other, ok := o.(ExceptionsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.SiteIds.Equal(other.SiteIds) {
		return false
	}

	if !v.SitegroupIds.Equal(other.SitegroupIds) {
		return false
	}

	return true
}

func (v ExceptionsValue) Type(ctx context.Context) attr.Type {
	return ExceptionsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ExceptionsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"site_ids": basetypes.ListType{
			ElemType: types.StringType,
		},
		"sitegroup_ids": basetypes.ListType{
			ElemType: types.StringType,
		},
	}
}
