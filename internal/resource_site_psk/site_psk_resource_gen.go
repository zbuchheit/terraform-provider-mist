// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_site_psk

import (
	"context"
	"github.com/Juniper/terraform-provider-mist/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func SitePskResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				Optional:            true,
				Description:         "email to send psk expiring notifications to",
				MarkdownDescription: "email to send psk expiring notifications to",
			},
			"expire_time": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Description:         "Expire time for this PSK key (epoch time in seconds). Default `null` (as no expiration)",
				MarkdownDescription: "Expire time for this PSK key (epoch time in seconds). Default `null` (as no expiration)",
				Default:             int64default.StaticInt64(0),
			},
			"expiry_notification_time": schema.Int64Attribute{
				Optional:            true,
				Description:         "Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire",
				MarkdownDescription: "Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire",
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "PSK ID",
				MarkdownDescription: "PSK ID",
			},
			"mac": schema.StringAttribute{
				Optional:            true,
				Description:         "if `usage`==`single`, the mac that this PSK ties to, empty if `auto-binding`",
				MarkdownDescription: "if `usage`==`single`, the mac that this PSK ties to, empty if `auto-binding`",
				Validators: []validator.String{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("usage"), types.StringValue("single")), mistvalidator.ParseMac(),
				},
			},
			"macs": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "if `usage`==`macs`, this list contains N number of client mac addresses or mac patterns(11:22:*) or both. This list is capped at 5000",
				MarkdownDescription: "if `usage`==`macs`, this list contains N number of client mac addresses or mac patterns(11:22:*) or both. This list is capped at 5000",
				Validators: []validator.List{
					listvalidator.SizeAtMost(5000),
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("usage"),
						types.StringValue("macs")),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"note": schema.StringAttribute{
				Optional: true,
			},
			"notify_expiry": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "If set to true, reminder notification will be sent when psk is about to expire",
				MarkdownDescription: "If set to true, reminder notification will be sent when psk is about to expire",
				Default:             booldefault.StaticBool(false),
			},
			"notify_on_create_or_edit": schema.BoolAttribute{
				Optional:            true,
				Description:         "If set to true, notification will be sent when psk is created or edited",
				MarkdownDescription: "If set to true, notification will be sent when psk is created or edited",
			},
			"old_passphrase": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				Description:         "previous passphrase of the PSK if it has been rotated",
				MarkdownDescription: "previous passphrase of the PSK if it has been rotated",
			},
			"org_id": schema.StringAttribute{
				Computed: true,
			},
			"passphrase": schema.StringAttribute{
				Required:            true,
				Sensitive:           true,
				Description:         "passphrase of the PSK (8-63 character or 64 in hex)",
				MarkdownDescription: "passphrase of the PSK (8-63 character or 64 in hex)",
				Validators: []validator.String{
					stringvalidator.LengthBetween(8, 64),
				},
			},
			"role": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
				},
			},
			"site_id": schema.StringAttribute{
				Required: true,
			},
			"ssid": schema.StringAttribute{
				Required:            true,
				Description:         "SSID this PSK should be applicable to",
				MarkdownDescription: "SSID this PSK should be applicable to",
			},
			"usage": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "enum: `macs`, `multi`, `single`",
				MarkdownDescription: "enum: `macs`, `multi`, `single`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"macs",
						"multi",
						"single",
					),
				},
				Default: stringdefault.StaticString("multi"),
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

type SitePskModel struct {
	Email                  types.String `tfsdk:"email"`
	ExpireTime             types.Int64  `tfsdk:"expire_time"`
	ExpiryNotificationTime types.Int64  `tfsdk:"expiry_notification_time"`
	Id                     types.String `tfsdk:"id"`
	Mac                    types.String `tfsdk:"mac"`
	Macs                   types.List   `tfsdk:"macs"`
	Name                   types.String `tfsdk:"name"`
	Note                   types.String `tfsdk:"note"`
	NotifyExpiry           types.Bool   `tfsdk:"notify_expiry"`
	NotifyOnCreateOrEdit   types.Bool   `tfsdk:"notify_on_create_or_edit"`
	OldPassphrase          types.String `tfsdk:"old_passphrase"`
	OrgId                  types.String `tfsdk:"org_id"`
	Passphrase             types.String `tfsdk:"passphrase"`
	Role                   types.String `tfsdk:"role"`
	SiteId                 types.String `tfsdk:"site_id"`
	Ssid                   types.String `tfsdk:"ssid"`
	Usage                  types.String `tfsdk:"usage"`
	VlanId                 types.String `tfsdk:"vlan_id"`
}
