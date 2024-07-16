// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_org_networktemplates

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func OrgNetworktemplatesDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"org_id": schema.StringAttribute{
				Required: true,
			},
			"org_networktemplates": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_time": schema.NumberAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"modified_time": schema.NumberAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"org_id": schema.StringAttribute{
							Computed: true,
						},
					},
					CustomType: OrgNetworktemplatesType{
						ObjectType: types.ObjectType{
							AttrTypes: OrgNetworktemplatesValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
		},
	}
}

type OrgNetworktemplatesModel struct {
	OrgId               types.String `tfsdk:"org_id"`
	OrgNetworktemplates types.Set    `tfsdk:"org_networktemplates"`
}

var _ basetypes.ObjectTypable = OrgNetworktemplatesType{}

type OrgNetworktemplatesType struct {
	basetypes.ObjectType
}

func (t OrgNetworktemplatesType) Equal(o attr.Type) bool {
	other, ok := o.(OrgNetworktemplatesType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t OrgNetworktemplatesType) String() string {
	return "OrgNetworktemplatesType"
}

func (t OrgNetworktemplatesType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	createdTimeAttribute, ok := attributes["created_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`created_time is missing from object`)

		return nil, diags
	}

	createdTimeVal, ok := createdTimeAttribute.(basetypes.NumberValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`created_time expected to be basetypes.NumberValue, was: %T`, createdTimeAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return nil, diags
	}

	idVal, ok := idAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.StringValue, was: %T`, idAttribute))
	}

	modifiedTimeAttribute, ok := attributes["modified_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`modified_time is missing from object`)

		return nil, diags
	}

	modifiedTimeVal, ok := modifiedTimeAttribute.(basetypes.NumberValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`modified_time expected to be basetypes.NumberValue, was: %T`, modifiedTimeAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
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

	if diags.HasError() {
		return nil, diags
	}

	return OrgNetworktemplatesValue{
		CreatedTime:  createdTimeVal,
		Id:           idVal,
		ModifiedTime: modifiedTimeVal,
		Name:         nameVal,
		OrgId:        orgIdVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewOrgNetworktemplatesValueNull() OrgNetworktemplatesValue {
	return OrgNetworktemplatesValue{
		state: attr.ValueStateNull,
	}
}

func NewOrgNetworktemplatesValueUnknown() OrgNetworktemplatesValue {
	return OrgNetworktemplatesValue{
		state: attr.ValueStateUnknown,
	}
}

func NewOrgNetworktemplatesValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (OrgNetworktemplatesValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing OrgNetworktemplatesValue Attribute Value",
				"While creating a OrgNetworktemplatesValue value, a missing attribute value was detected. "+
					"A OrgNetworktemplatesValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OrgNetworktemplatesValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid OrgNetworktemplatesValue Attribute Type",
				"While creating a OrgNetworktemplatesValue value, an invalid attribute value was detected. "+
					"A OrgNetworktemplatesValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("OrgNetworktemplatesValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("OrgNetworktemplatesValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra OrgNetworktemplatesValue Attribute Value",
				"While creating a OrgNetworktemplatesValue value, an extra attribute value was detected. "+
					"A OrgNetworktemplatesValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra OrgNetworktemplatesValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewOrgNetworktemplatesValueUnknown(), diags
	}

	createdTimeAttribute, ok := attributes["created_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`created_time is missing from object`)

		return NewOrgNetworktemplatesValueUnknown(), diags
	}

	createdTimeVal, ok := createdTimeAttribute.(basetypes.NumberValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`created_time expected to be basetypes.NumberValue, was: %T`, createdTimeAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewOrgNetworktemplatesValueUnknown(), diags
	}

	idVal, ok := idAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.StringValue, was: %T`, idAttribute))
	}

	modifiedTimeAttribute, ok := attributes["modified_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`modified_time is missing from object`)

		return NewOrgNetworktemplatesValueUnknown(), diags
	}

	modifiedTimeVal, ok := modifiedTimeAttribute.(basetypes.NumberValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`modified_time expected to be basetypes.NumberValue, was: %T`, modifiedTimeAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewOrgNetworktemplatesValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	orgIdAttribute, ok := attributes["org_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`org_id is missing from object`)

		return NewOrgNetworktemplatesValueUnknown(), diags
	}

	orgIdVal, ok := orgIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`org_id expected to be basetypes.StringValue, was: %T`, orgIdAttribute))
	}

	if diags.HasError() {
		return NewOrgNetworktemplatesValueUnknown(), diags
	}

	return OrgNetworktemplatesValue{
		CreatedTime:  createdTimeVal,
		Id:           idVal,
		ModifiedTime: modifiedTimeVal,
		Name:         nameVal,
		OrgId:        orgIdVal,
		state:        attr.ValueStateKnown,
	}, diags
}

func NewOrgNetworktemplatesValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) OrgNetworktemplatesValue {
	object, diags := NewOrgNetworktemplatesValue(attributeTypes, attributes)

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

		panic("NewOrgNetworktemplatesValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t OrgNetworktemplatesType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewOrgNetworktemplatesValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewOrgNetworktemplatesValueUnknown(), nil
	}

	if in.IsNull() {
		return NewOrgNetworktemplatesValueNull(), nil
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

	return NewOrgNetworktemplatesValueMust(OrgNetworktemplatesValue{}.AttributeTypes(ctx), attributes), nil
}

func (t OrgNetworktemplatesType) ValueType(ctx context.Context) attr.Value {
	return OrgNetworktemplatesValue{}
}

var _ basetypes.ObjectValuable = OrgNetworktemplatesValue{}

type OrgNetworktemplatesValue struct {
	CreatedTime  basetypes.NumberValue `tfsdk:"created_time"`
	Id           basetypes.StringValue `tfsdk:"id"`
	ModifiedTime basetypes.NumberValue `tfsdk:"modified_time"`
	Name         basetypes.StringValue `tfsdk:"name"`
	OrgId        basetypes.StringValue `tfsdk:"org_id"`
	state        attr.ValueState
}

func (v OrgNetworktemplatesValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 6)

	var val tftypes.Value
	var err error

	attrTypes["created_time"] = basetypes.NumberType{}.TerraformType(ctx)
	attrTypes["id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["modified_time"] = basetypes.NumberType{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["org_id"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 6)

		val, err = v.CreatedTime.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["created_time"] = val

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.ModifiedTime.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["modified_time"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.OrgId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["org_id"] = val

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

func (v OrgNetworktemplatesValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v OrgNetworktemplatesValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v OrgNetworktemplatesValue) String() string {
	return "OrgNetworktemplatesValue"
}

func (v OrgNetworktemplatesValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"created_time":  basetypes.NumberType{},
		"id":            basetypes.StringType{},
		"modified_time": basetypes.NumberType{},
		"name":          basetypes.StringType{},
		"org_id":        basetypes.StringType{},
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
			"created_time":  v.CreatedTime,
			"id":            v.Id,
			"modified_time": v.ModifiedTime,
			"name":          v.Name,
			"org_id":        v.OrgId,
		})

	return objVal, diags
}

func (v OrgNetworktemplatesValue) Equal(o attr.Value) bool {
	other, ok := o.(OrgNetworktemplatesValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.CreatedTime.Equal(other.CreatedTime) {
		return false
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.ModifiedTime.Equal(other.ModifiedTime) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.OrgId.Equal(other.OrgId) {
		return false
	}

	return true
}

func (v OrgNetworktemplatesValue) Type(ctx context.Context) attr.Type {
	return OrgNetworktemplatesType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v OrgNetworktemplatesValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"created_time":  basetypes.NumberType{},
		"id":            basetypes.StringType{},
		"modified_time": basetypes.NumberType{},
		"name":          basetypes.StringType{},
		"org_id":        basetypes.StringType{},
	}
}
