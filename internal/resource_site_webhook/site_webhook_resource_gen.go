// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_site_webhook

import (
	"context"
	"github.com/Juniper/terraform-provider-mist/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func SiteWebhookResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "whether webhook is enabled",
				MarkdownDescription: "whether webhook is enabled",
				Default:             booldefault.StaticBool(true),
			},
			"headers": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "if `type`=`http-post`, additional custom HTTP headers to add\nthe headers name and value must be string, total bytes of headers name and value must be less than 1000",
				MarkdownDescription: "if `type`=`http-post`, additional custom HTTP headers to add\nthe headers name and value must be string, total bytes of headers name and value must be less than 1000",
				Validators: []validator.Map{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("http-post")),
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("oauth2")),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "name of the webhook",
				MarkdownDescription: "name of the webhook",
			},
			"oauth2_client_id": schema.StringAttribute{
				Optional:            true,
				Description:         "required when `oauth2_grant_type`==`client_credentials`",
				MarkdownDescription: "required when `oauth2_grant_type`==`client_credentials`",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("oauth2_grant_type"), types.StringValue("client_credentials")),
				},
			},
			"oauth2_client_secret": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				Description:         "required when `oauth2_grant_type`==`client_credentials`",
				MarkdownDescription: "required when `oauth2_grant_type`==`client_credentials`",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("oauth2_grant_type"), types.StringValue("client_credentials")),
				},
			},
			"oauth2_grant_type": schema.StringAttribute{
				Optional:            true,
				Description:         "required when `type`==`oauth2`. enum: `client_credentials`, `password`",
				MarkdownDescription: "required when `type`==`oauth2`. enum: `client_credentials`, `password`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"client_credentials",
						"password",
					),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("oauth2")),
				},
			},
			"oauth2_password": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				Description:         "required when `oauth2_grant_type`==`password`",
				MarkdownDescription: "required when `oauth2_grant_type`==`password`",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("oauth2_grant_type"), types.StringValue("password")),
				},
			},
			"oauth2_scopes": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "required when `type`==`oauth2`, if provided, will be used in the token request",
				MarkdownDescription: "required when `type`==`oauth2`, if provided, will be used in the token request",
			},
			"oauth2_token_url": schema.StringAttribute{
				Optional:            true,
				Description:         "required when `type`==`oauth2`",
				MarkdownDescription: "required when `type`==`oauth2`",
				Validators: []validator.String{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("oauth2")),
				},
			},
			"oauth2_username": schema.StringAttribute{
				Optional:            true,
				Description:         "required when `oauth2_grant_type`==`password`",
				MarkdownDescription: "required when `oauth2_grant_type`==`password`",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("oauth2_grant_type"), types.StringValue("password")),
				},
			},
			"org_id": schema.StringAttribute{
				Computed: true,
			},
			"secret": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				Description:         "only if `type`=`http-post` \n\nwhen `secret` is provided, two  HTTP headers will be added: \n  * X-Mist-Signature-v2: HMAC_SHA256(secret, body)\n  * X-Mist-Signature: HMAC_SHA1(secret, body)",
				MarkdownDescription: "only if `type`=`http-post` \n\nwhen `secret` is provided, two  HTTP headers will be added: \n  * X-Mist-Signature-v2: HMAC_SHA256(secret, body)\n  * X-Mist-Signature: HMAC_SHA1(secret, body)",
				Validators: []validator.String{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("http-post")),
				},
			},
			"site_id": schema.StringAttribute{
				Required: true,
			},
			"splunk_token": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				Description:         "required if `type`=`splunk`\nIf splunk_token is not defined for a type Splunk webhook, it will not send, regardless if the webhook receiver is configured to accept it.'",
				MarkdownDescription: "required if `type`=`splunk`\nIf splunk_token is not defined for a type Splunk webhook, it will not send, regardless if the webhook receiver is configured to accept it.'",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("type"), types.StringValue("splunk")),
				},
			},
			"topics": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "enum: `alarms`, `asset-raw`, `asset-raw-rssi`, `audits`, `client-info`, `client-join`, `client-latency`, `client-sessions`, `device-updowns`, `device-events`, `discovered-raw-rssi`, `location`, `location_asset`, `location_centrak`, `location_client`, `location_sdk`, `location_unclient`, `mxedge-events`, `nac-accounting`, `nac_events`, `occupancy-alerts`, `rssizone`, `sdkclient_scan_data`, `site_sle`, `vbeacon`, `wifi-conn-raw`, `wifi-unconn-raw`, `zone`",
				MarkdownDescription: "enum: `alarms`, `asset-raw`, `asset-raw-rssi`, `audits`, `client-info`, `client-join`, `client-latency`, `client-sessions`, `device-updowns`, `device-events`, `discovered-raw-rssi`, `location`, `location_asset`, `location_centrak`, `location_client`, `location_sdk`, `location_unclient`, `mxedge-events`, `nac-accounting`, `nac_events`, `occupancy-alerts`, `rssizone`, `sdkclient_scan_data`, `site_sle`, `vbeacon`, `wifi-conn-raw`, `wifi-unconn-raw`, `zone`",
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
					listvalidator.ValueStringsAre(
						stringvalidator.OneOf(
							"alarms",
							"asset-raw",
							"asset-raw-rssi",
							"audits",
							"client-info",
							"client-join",
							"client-latency",
							"client-sessions",
							"device-updowns",
							"device-events",
							"discovered-raw-rssi",
							"location",
							"location_asset",
							"location_centrak",
							"location_client",
							"location_sdk",
							"location_unclient",
							"mxedge-events",
							"nac-accounting",
							"nac_events",
							"occupancy-alerts",
							"rssizone",
							"sdkclient_scan_data",
							"site_sle",
							"vbeacon",
							"wifi-conn-raw",
							"wifi-unconn-raw",
							"zone",
						),
					),
				},
			},
			"type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "enum: `aws-sns`, `google-pubsub`, `http-post`, `oauth2`, `splunk`",
				MarkdownDescription: "enum: `aws-sns`, `google-pubsub`, `http-post`, `oauth2`, `splunk`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"aws-sns",
						"google-pubsub",
						"http-post",
						"oauth2",
						"splunk",
					),
				},
				Default: stringdefault.StaticString("http-post"),
			},
			"url": schema.StringAttribute{
				Required: true,
			},
			"verify_cert": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "when url uses HTTPS, whether to verify the certificate",
				MarkdownDescription: "when url uses HTTPS, whether to verify the certificate",
				Default:             booldefault.StaticBool(true),
			},
		},
	}
}

type SiteWebhookModel struct {
	Enabled            types.Bool   `tfsdk:"enabled"`
	Headers            types.Map    `tfsdk:"headers"`
	Id                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	Oauth2ClientId     types.String `tfsdk:"oauth2_client_id"`
	Oauth2ClientSecret types.String `tfsdk:"oauth2_client_secret"`
	Oauth2GrantType    types.String `tfsdk:"oauth2_grant_type"`
	Oauth2Password     types.String `tfsdk:"oauth2_password"`
	Oauth2Scopes       types.List   `tfsdk:"oauth2_scopes"`
	Oauth2TokenUrl     types.String `tfsdk:"oauth2_token_url"`
	Oauth2Username     types.String `tfsdk:"oauth2_username"`
	OrgId              types.String `tfsdk:"org_id"`
	Secret             types.String `tfsdk:"secret"`
	SiteId             types.String `tfsdk:"site_id"`
	SplunkToken        types.String `tfsdk:"splunk_token"`
	Topics             types.List   `tfsdk:"topics"`
	Type               types.String `tfsdk:"type"`
	Url                types.String `tfsdk:"url"`
	VerifyCert         types.Bool   `tfsdk:"verify_cert"`
}