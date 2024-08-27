// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_site_psks

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func SitePsksDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"limit": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"page": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"role": schema.StringAttribute{
				Optional: true,
			},
			"site_id": schema.StringAttribute{
				Required: true,
			},
			"site_psks": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"admin_sso_id": schema.StringAttribute{
							Computed:            true,
							Description:         "sso id for psk created from psk portal",
							MarkdownDescription: "sso id for psk created from psk portal",
						},
						"created_time": schema.NumberAttribute{
							Computed: true,
						},
						"email": schema.StringAttribute{
							Computed:            true,
							Description:         "email to send psk expiring notifications to",
							MarkdownDescription: "email to send psk expiring notifications to",
						},
						"expire_time": schema.Int64Attribute{
							Computed:            true,
							Description:         "Expire time for this PSK key (epoch time in seconds). Default `null` (as no expiration)",
							MarkdownDescription: "Expire time for this PSK key (epoch time in seconds). Default `null` (as no expiration)",
						},
						"expiry_notification_time": schema.Int64Attribute{
							Computed:            true,
							Description:         "Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire",
							MarkdownDescription: "Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire",
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"mac": schema.StringAttribute{
							Computed:            true,
							Description:         "if `usage`==`single`, the mac that this PSK ties to, empty if `auto-binding`",
							MarkdownDescription: "if `usage`==`single`, the mac that this PSK ties to, empty if `auto-binding`",
						},
						"modified_time": schema.NumberAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"note": schema.StringAttribute{
							Computed: true,
						},
						"notify_expiry": schema.BoolAttribute{
							Computed:            true,
							Description:         "If set to true, reminder notification will be sent when psk is about to expire",
							MarkdownDescription: "If set to true, reminder notification will be sent when psk is about to expire",
						},
						"notify_on_create_or_edit": schema.BoolAttribute{
							Computed:            true,
							Description:         "If set to true, notification will be sent when psk is created or edited",
							MarkdownDescription: "If set to true, notification will be sent when psk is created or edited",
						},
						"old_passphrase": schema.StringAttribute{
							Computed:            true,
							Description:         "previous passphrase of the PSK if it has been rotated",
							MarkdownDescription: "previous passphrase of the PSK if it has been rotated",
						},
						"org_id": schema.StringAttribute{
							Computed: true,
						},
						"passphrase": schema.StringAttribute{
							Computed:            true,
							Description:         "passphrase of the PSK (8-63 character or 64 in hex)",
							MarkdownDescription: "passphrase of the PSK (8-63 character or 64 in hex)",
						},
						"role": schema.StringAttribute{
							Computed: true,
						},
						"site_id": schema.StringAttribute{
							Computed: true,
						},
						"ssid": schema.StringAttribute{
							Computed:            true,
							Description:         "SSID this PSK should be applicable to",
							MarkdownDescription: "SSID this PSK should be applicable to",
						},
						"usage": schema.StringAttribute{
							Computed:            true,
							Description:         "enum: `macs`, `multi`, `single`",
							MarkdownDescription: "enum: `macs`, `multi`, `single`",
						},
						"vlan_id": schema.StringAttribute{
							Computed: true,
						},
					},
					CustomType: SitePsksType{
						ObjectType: types.ObjectType{
							AttrTypes: SitePsksValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
			"ssid": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

type SitePsksModel struct {
	Limit    types.Int64  `tfsdk:"limit"`
	Name     types.String `tfsdk:"name"`
	Page     types.Int64  `tfsdk:"page"`
	Role     types.String `tfsdk:"role"`
	SiteId   types.String `tfsdk:"site_id"`
	SitePsks types.Set    `tfsdk:"site_psks"`
	Ssid     types.String `tfsdk:"ssid"`
}

var _ basetypes.ObjectTypable = SitePsksType{}

type SitePsksType struct {
	basetypes.ObjectType
}

func (t SitePsksType) Equal(o attr.Type) bool {
	other, ok := o.(SitePsksType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t SitePsksType) String() string {
	return "SitePsksType"
}

func (t SitePsksType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	adminSsoIdAttribute, ok := attributes["admin_sso_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`admin_sso_id is missing from object`)

		return nil, diags
	}

	adminSsoIdVal, ok := adminSsoIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`admin_sso_id expected to be basetypes.StringValue, was: %T`, adminSsoIdAttribute))
	}

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

	emailAttribute, ok := attributes["email"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`email is missing from object`)

		return nil, diags
	}

	emailVal, ok := emailAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`email expected to be basetypes.StringValue, was: %T`, emailAttribute))
	}

	expireTimeAttribute, ok := attributes["expire_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`expire_time is missing from object`)

		return nil, diags
	}

	expireTimeVal, ok := expireTimeAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`expire_time expected to be basetypes.Int64Value, was: %T`, expireTimeAttribute))
	}

	expiryNotificationTimeAttribute, ok := attributes["expiry_notification_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`expiry_notification_time is missing from object`)

		return nil, diags
	}

	expiryNotificationTimeVal, ok := expiryNotificationTimeAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`expiry_notification_time expected to be basetypes.Int64Value, was: %T`, expiryNotificationTimeAttribute))
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

	noteAttribute, ok := attributes["note"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`note is missing from object`)

		return nil, diags
	}

	noteVal, ok := noteAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`note expected to be basetypes.StringValue, was: %T`, noteAttribute))
	}

	notifyExpiryAttribute, ok := attributes["notify_expiry"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`notify_expiry is missing from object`)

		return nil, diags
	}

	notifyExpiryVal, ok := notifyExpiryAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`notify_expiry expected to be basetypes.BoolValue, was: %T`, notifyExpiryAttribute))
	}

	notifyOnCreateOrEditAttribute, ok := attributes["notify_on_create_or_edit"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`notify_on_create_or_edit is missing from object`)

		return nil, diags
	}

	notifyOnCreateOrEditVal, ok := notifyOnCreateOrEditAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`notify_on_create_or_edit expected to be basetypes.BoolValue, was: %T`, notifyOnCreateOrEditAttribute))
	}

	oldPassphraseAttribute, ok := attributes["old_passphrase"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`old_passphrase is missing from object`)

		return nil, diags
	}

	oldPassphraseVal, ok := oldPassphraseAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`old_passphrase expected to be basetypes.StringValue, was: %T`, oldPassphraseAttribute))
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

	passphraseAttribute, ok := attributes["passphrase"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`passphrase is missing from object`)

		return nil, diags
	}

	passphraseVal, ok := passphraseAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`passphrase expected to be basetypes.StringValue, was: %T`, passphraseAttribute))
	}

	roleAttribute, ok := attributes["role"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`role is missing from object`)

		return nil, diags
	}

	roleVal, ok := roleAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`role expected to be basetypes.StringValue, was: %T`, roleAttribute))
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

	ssidAttribute, ok := attributes["ssid"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ssid is missing from object`)

		return nil, diags
	}

	ssidVal, ok := ssidAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ssid expected to be basetypes.StringValue, was: %T`, ssidAttribute))
	}

	usageAttribute, ok := attributes["usage"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`usage is missing from object`)

		return nil, diags
	}

	usageVal, ok := usageAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`usage expected to be basetypes.StringValue, was: %T`, usageAttribute))
	}

	vlanIdAttribute, ok := attributes["vlan_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`vlan_id is missing from object`)

		return nil, diags
	}

	vlanIdVal, ok := vlanIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`vlan_id expected to be basetypes.StringValue, was: %T`, vlanIdAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return SitePsksValue{
		AdminSsoId:             adminSsoIdVal,
		CreatedTime:            createdTimeVal,
		Email:                  emailVal,
		ExpireTime:             expireTimeVal,
		ExpiryNotificationTime: expiryNotificationTimeVal,
		Id:                     idVal,
		Mac:                    macVal,
		ModifiedTime:           modifiedTimeVal,
		Name:                   nameVal,
		Note:                   noteVal,
		NotifyExpiry:           notifyExpiryVal,
		NotifyOnCreateOrEdit:   notifyOnCreateOrEditVal,
		OldPassphrase:          oldPassphraseVal,
		OrgId:                  orgIdVal,
		Passphrase:             passphraseVal,
		Role:                   roleVal,
		SiteId:                 siteIdVal,
		Ssid:                   ssidVal,
		Usage:                  usageVal,
		VlanId:                 vlanIdVal,
		state:                  attr.ValueStateKnown,
	}, diags
}

func NewSitePsksValueNull() SitePsksValue {
	return SitePsksValue{
		state: attr.ValueStateNull,
	}
}

func NewSitePsksValueUnknown() SitePsksValue {
	return SitePsksValue{
		state: attr.ValueStateUnknown,
	}
}

func NewSitePsksValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (SitePsksValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing SitePsksValue Attribute Value",
				"While creating a SitePsksValue value, a missing attribute value was detected. "+
					"A SitePsksValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("SitePsksValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid SitePsksValue Attribute Type",
				"While creating a SitePsksValue value, an invalid attribute value was detected. "+
					"A SitePsksValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("SitePsksValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("SitePsksValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra SitePsksValue Attribute Value",
				"While creating a SitePsksValue value, an extra attribute value was detected. "+
					"A SitePsksValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra SitePsksValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewSitePsksValueUnknown(), diags
	}

	adminSsoIdAttribute, ok := attributes["admin_sso_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`admin_sso_id is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	adminSsoIdVal, ok := adminSsoIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`admin_sso_id expected to be basetypes.StringValue, was: %T`, adminSsoIdAttribute))
	}

	createdTimeAttribute, ok := attributes["created_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`created_time is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	createdTimeVal, ok := createdTimeAttribute.(basetypes.NumberValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`created_time expected to be basetypes.NumberValue, was: %T`, createdTimeAttribute))
	}

	emailAttribute, ok := attributes["email"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`email is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	emailVal, ok := emailAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`email expected to be basetypes.StringValue, was: %T`, emailAttribute))
	}

	expireTimeAttribute, ok := attributes["expire_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`expire_time is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	expireTimeVal, ok := expireTimeAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`expire_time expected to be basetypes.Int64Value, was: %T`, expireTimeAttribute))
	}

	expiryNotificationTimeAttribute, ok := attributes["expiry_notification_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`expiry_notification_time is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	expiryNotificationTimeVal, ok := expiryNotificationTimeAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`expiry_notification_time expected to be basetypes.Int64Value, was: %T`, expiryNotificationTimeAttribute))
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	idVal, ok := idAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be basetypes.StringValue, was: %T`, idAttribute))
	}

	macAttribute, ok := attributes["mac"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`mac is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	macVal, ok := macAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`mac expected to be basetypes.StringValue, was: %T`, macAttribute))
	}

	modifiedTimeAttribute, ok := attributes["modified_time"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`modified_time is missing from object`)

		return NewSitePsksValueUnknown(), diags
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

		return NewSitePsksValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be basetypes.StringValue, was: %T`, nameAttribute))
	}

	noteAttribute, ok := attributes["note"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`note is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	noteVal, ok := noteAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`note expected to be basetypes.StringValue, was: %T`, noteAttribute))
	}

	notifyExpiryAttribute, ok := attributes["notify_expiry"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`notify_expiry is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	notifyExpiryVal, ok := notifyExpiryAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`notify_expiry expected to be basetypes.BoolValue, was: %T`, notifyExpiryAttribute))
	}

	notifyOnCreateOrEditAttribute, ok := attributes["notify_on_create_or_edit"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`notify_on_create_or_edit is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	notifyOnCreateOrEditVal, ok := notifyOnCreateOrEditAttribute.(basetypes.BoolValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`notify_on_create_or_edit expected to be basetypes.BoolValue, was: %T`, notifyOnCreateOrEditAttribute))
	}

	oldPassphraseAttribute, ok := attributes["old_passphrase"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`old_passphrase is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	oldPassphraseVal, ok := oldPassphraseAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`old_passphrase expected to be basetypes.StringValue, was: %T`, oldPassphraseAttribute))
	}

	orgIdAttribute, ok := attributes["org_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`org_id is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	orgIdVal, ok := orgIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`org_id expected to be basetypes.StringValue, was: %T`, orgIdAttribute))
	}

	passphraseAttribute, ok := attributes["passphrase"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`passphrase is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	passphraseVal, ok := passphraseAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`passphrase expected to be basetypes.StringValue, was: %T`, passphraseAttribute))
	}

	roleAttribute, ok := attributes["role"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`role is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	roleVal, ok := roleAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`role expected to be basetypes.StringValue, was: %T`, roleAttribute))
	}

	siteIdAttribute, ok := attributes["site_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`site_id is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	siteIdVal, ok := siteIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`site_id expected to be basetypes.StringValue, was: %T`, siteIdAttribute))
	}

	ssidAttribute, ok := attributes["ssid"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`ssid is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	ssidVal, ok := ssidAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`ssid expected to be basetypes.StringValue, was: %T`, ssidAttribute))
	}

	usageAttribute, ok := attributes["usage"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`usage is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	usageVal, ok := usageAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`usage expected to be basetypes.StringValue, was: %T`, usageAttribute))
	}

	vlanIdAttribute, ok := attributes["vlan_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`vlan_id is missing from object`)

		return NewSitePsksValueUnknown(), diags
	}

	vlanIdVal, ok := vlanIdAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`vlan_id expected to be basetypes.StringValue, was: %T`, vlanIdAttribute))
	}

	if diags.HasError() {
		return NewSitePsksValueUnknown(), diags
	}

	return SitePsksValue{
		AdminSsoId:             adminSsoIdVal,
		CreatedTime:            createdTimeVal,
		Email:                  emailVal,
		ExpireTime:             expireTimeVal,
		ExpiryNotificationTime: expiryNotificationTimeVal,
		Id:                     idVal,
		Mac:                    macVal,
		ModifiedTime:           modifiedTimeVal,
		Name:                   nameVal,
		Note:                   noteVal,
		NotifyExpiry:           notifyExpiryVal,
		NotifyOnCreateOrEdit:   notifyOnCreateOrEditVal,
		OldPassphrase:          oldPassphraseVal,
		OrgId:                  orgIdVal,
		Passphrase:             passphraseVal,
		Role:                   roleVal,
		SiteId:                 siteIdVal,
		Ssid:                   ssidVal,
		Usage:                  usageVal,
		VlanId:                 vlanIdVal,
		state:                  attr.ValueStateKnown,
	}, diags
}

func NewSitePsksValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) SitePsksValue {
	object, diags := NewSitePsksValue(attributeTypes, attributes)

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

		panic("NewSitePsksValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t SitePsksType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewSitePsksValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewSitePsksValueUnknown(), nil
	}

	if in.IsNull() {
		return NewSitePsksValueNull(), nil
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

	return NewSitePsksValueMust(SitePsksValue{}.AttributeTypes(ctx), attributes), nil
}

func (t SitePsksType) ValueType(ctx context.Context) attr.Value {
	return SitePsksValue{}
}

var _ basetypes.ObjectValuable = SitePsksValue{}

type SitePsksValue struct {
	AdminSsoId             basetypes.StringValue `tfsdk:"admin_sso_id"`
	CreatedTime            basetypes.NumberValue `tfsdk:"created_time"`
	Email                  basetypes.StringValue `tfsdk:"email"`
	ExpireTime             basetypes.Int64Value  `tfsdk:"expire_time"`
	ExpiryNotificationTime basetypes.Int64Value  `tfsdk:"expiry_notification_time"`
	Id                     basetypes.StringValue `tfsdk:"id"`
	Mac                    basetypes.StringValue `tfsdk:"mac"`
	ModifiedTime           basetypes.NumberValue `tfsdk:"modified_time"`
	Name                   basetypes.StringValue `tfsdk:"name"`
	Note                   basetypes.StringValue `tfsdk:"note"`
	NotifyExpiry           basetypes.BoolValue   `tfsdk:"notify_expiry"`
	NotifyOnCreateOrEdit   basetypes.BoolValue   `tfsdk:"notify_on_create_or_edit"`
	OldPassphrase          basetypes.StringValue `tfsdk:"old_passphrase"`
	OrgId                  basetypes.StringValue `tfsdk:"org_id"`
	Passphrase             basetypes.StringValue `tfsdk:"passphrase"`
	Role                   basetypes.StringValue `tfsdk:"role"`
	SiteId                 basetypes.StringValue `tfsdk:"site_id"`
	Ssid                   basetypes.StringValue `tfsdk:"ssid"`
	Usage                  basetypes.StringValue `tfsdk:"usage"`
	VlanId                 basetypes.StringValue `tfsdk:"vlan_id"`
	state                  attr.ValueState
}

func (v SitePsksValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 20)

	var val tftypes.Value
	var err error

	attrTypes["admin_sso_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["created_time"] = basetypes.NumberType{}.TerraformType(ctx)
	attrTypes["email"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["expire_time"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["expiry_notification_time"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["mac"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["modified_time"] = basetypes.NumberType{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["note"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["notify_expiry"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["notify_on_create_or_edit"] = basetypes.BoolType{}.TerraformType(ctx)
	attrTypes["old_passphrase"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["org_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["passphrase"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["role"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["site_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["ssid"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["usage"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["vlan_id"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 20)

		val, err = v.AdminSsoId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["admin_sso_id"] = val

		val, err = v.CreatedTime.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["created_time"] = val

		val, err = v.Email.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["email"] = val

		val, err = v.ExpireTime.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["expire_time"] = val

		val, err = v.ExpiryNotificationTime.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["expiry_notification_time"] = val

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.Mac.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["mac"] = val

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

		val, err = v.Note.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["note"] = val

		val, err = v.NotifyExpiry.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["notify_expiry"] = val

		val, err = v.NotifyOnCreateOrEdit.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["notify_on_create_or_edit"] = val

		val, err = v.OldPassphrase.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["old_passphrase"] = val

		val, err = v.OrgId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["org_id"] = val

		val, err = v.Passphrase.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["passphrase"] = val

		val, err = v.Role.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["role"] = val

		val, err = v.SiteId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["site_id"] = val

		val, err = v.Ssid.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["ssid"] = val

		val, err = v.Usage.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["usage"] = val

		val, err = v.VlanId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["vlan_id"] = val

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

func (v SitePsksValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v SitePsksValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v SitePsksValue) String() string {
	return "SitePsksValue"
}

func (v SitePsksValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"admin_sso_id":             basetypes.StringType{},
		"created_time":             basetypes.NumberType{},
		"email":                    basetypes.StringType{},
		"expire_time":              basetypes.Int64Type{},
		"expiry_notification_time": basetypes.Int64Type{},
		"id":                       basetypes.StringType{},
		"mac":                      basetypes.StringType{},
		"modified_time":            basetypes.NumberType{},
		"name":                     basetypes.StringType{},
		"note":                     basetypes.StringType{},
		"notify_expiry":            basetypes.BoolType{},
		"notify_on_create_or_edit": basetypes.BoolType{},
		"old_passphrase":           basetypes.StringType{},
		"org_id":                   basetypes.StringType{},
		"passphrase":               basetypes.StringType{},
		"role":                     basetypes.StringType{},
		"site_id":                  basetypes.StringType{},
		"ssid":                     basetypes.StringType{},
		"usage":                    basetypes.StringType{},
		"vlan_id":                  basetypes.StringType{},
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
			"admin_sso_id":             v.AdminSsoId,
			"created_time":             v.CreatedTime,
			"email":                    v.Email,
			"expire_time":              v.ExpireTime,
			"expiry_notification_time": v.ExpiryNotificationTime,
			"id":                       v.Id,
			"mac":                      v.Mac,
			"modified_time":            v.ModifiedTime,
			"name":                     v.Name,
			"note":                     v.Note,
			"notify_expiry":            v.NotifyExpiry,
			"notify_on_create_or_edit": v.NotifyOnCreateOrEdit,
			"old_passphrase":           v.OldPassphrase,
			"org_id":                   v.OrgId,
			"passphrase":               v.Passphrase,
			"role":                     v.Role,
			"site_id":                  v.SiteId,
			"ssid":                     v.Ssid,
			"usage":                    v.Usage,
			"vlan_id":                  v.VlanId,
		})

	return objVal, diags
}

func (v SitePsksValue) Equal(o attr.Value) bool {
	other, ok := o.(SitePsksValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.AdminSsoId.Equal(other.AdminSsoId) {
		return false
	}

	if !v.CreatedTime.Equal(other.CreatedTime) {
		return false
	}

	if !v.Email.Equal(other.Email) {
		return false
	}

	if !v.ExpireTime.Equal(other.ExpireTime) {
		return false
	}

	if !v.ExpiryNotificationTime.Equal(other.ExpiryNotificationTime) {
		return false
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.Mac.Equal(other.Mac) {
		return false
	}

	if !v.ModifiedTime.Equal(other.ModifiedTime) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.Note.Equal(other.Note) {
		return false
	}

	if !v.NotifyExpiry.Equal(other.NotifyExpiry) {
		return false
	}

	if !v.NotifyOnCreateOrEdit.Equal(other.NotifyOnCreateOrEdit) {
		return false
	}

	if !v.OldPassphrase.Equal(other.OldPassphrase) {
		return false
	}

	if !v.OrgId.Equal(other.OrgId) {
		return false
	}

	if !v.Passphrase.Equal(other.Passphrase) {
		return false
	}

	if !v.Role.Equal(other.Role) {
		return false
	}

	if !v.SiteId.Equal(other.SiteId) {
		return false
	}

	if !v.Ssid.Equal(other.Ssid) {
		return false
	}

	if !v.Usage.Equal(other.Usage) {
		return false
	}

	if !v.VlanId.Equal(other.VlanId) {
		return false
	}

	return true
}

func (v SitePsksValue) Type(ctx context.Context) attr.Type {
	return SitePsksType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v SitePsksValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"admin_sso_id":             basetypes.StringType{},
		"created_time":             basetypes.NumberType{},
		"email":                    basetypes.StringType{},
		"expire_time":              basetypes.Int64Type{},
		"expiry_notification_time": basetypes.Int64Type{},
		"id":                       basetypes.StringType{},
		"mac":                      basetypes.StringType{},
		"modified_time":            basetypes.NumberType{},
		"name":                     basetypes.StringType{},
		"note":                     basetypes.StringType{},
		"notify_expiry":            basetypes.BoolType{},
		"notify_on_create_or_edit": basetypes.BoolType{},
		"old_passphrase":           basetypes.StringType{},
		"org_id":                   basetypes.StringType{},
		"passphrase":               basetypes.StringType{},
		"role":                     basetypes.StringType{},
		"site_id":                  basetypes.StringType{},
		"ssid":                     basetypes.StringType{},
		"usage":                    basetypes.StringType{},
		"vlan_id":                  basetypes.StringType{},
	}
}
