// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_org_nacidp

import (
	"context"
	"github.com/Juniper/terraform-provider-mist/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OrgNacidpResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"group_filter": schema.StringAttribute{
				Optional:            true,
				Description:         "Required if `ldap_type`==`custom`, LDAP filter that will identify the type of group",
				MarkdownDescription: "Required if `ldap_type`==`custom`, LDAP filter that will identify the type of group",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("ldap_type"), types.StringValue("custom")),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"idp_type": schema.StringAttribute{
				Required:            true,
				Description:         "enum: `ldap`, `mxedge_proxy`, `oauth`",
				MarkdownDescription: "enum: `ldap`, `mxedge_proxy`, `oauth`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ldap",
						"mxedge_proxy",
						"oauth",
					),
				},
			},
			"ldap_base_dn": schema.StringAttribute{
				Optional:            true,
				Description:         "Required if `idp_type`==`ldap`, whole domain or a specific organization unit (container) in Search base to specify where users and groups are found in the LDAP tree",
				MarkdownDescription: "Required if `idp_type`==`ldap`, whole domain or a specific organization unit (container) in Search base to specify where users and groups are found in the LDAP tree",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("ldap")),
				},
			},
			"ldap_bind_dn": schema.StringAttribute{
				Optional:            true,
				Description:         "Required if `idp_type`==`ldap`, the account used to authenticate against the LDAP",
				MarkdownDescription: "Required if `idp_type`==`ldap`, the account used to authenticate against the LDAP",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("ldap")),
				},
			},
			"ldap_bind_password": schema.StringAttribute{
				Optional:            true,
				Description:         "Required if `idp_type`==`ldap`, the password used to authenticate against the LDAP",
				MarkdownDescription: "Required if `idp_type`==`ldap`, the password used to authenticate against the LDAP",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("ldap")),
				},
			},
			"ldap_cacerts": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "Required if `idp_type`==`ldap`, list of CA certificates to validate the LDAP certificate",
				MarkdownDescription: "Required if `idp_type`==`ldap`, list of CA certificates to validate the LDAP certificate",
			},
			"ldap_client_cert": schema.StringAttribute{
				Optional:            true,
				Description:         "if `idp_type`==`ldap`, LDAPS Client certificate",
				MarkdownDescription: "if `idp_type`==`ldap`, LDAPS Client certificate",
				Validators: []validator.String{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("ldap")),
				},
			},
			"ldap_client_key": schema.StringAttribute{
				Optional:            true,
				Description:         "if `idp_type`==`ldap`, Key for the `ldap_client_cert`",
				MarkdownDescription: "if `idp_type`==`ldap`, Key for the `ldap_client_cert`",
				Validators: []validator.String{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("ldap")),
				},
			},
			"ldap_group_attr": schema.StringAttribute{
				Optional:            true,
				Description:         "if `ldap_type`==`custom`",
				MarkdownDescription: "if `ldap_type`==`custom`",
				Validators: []validator.String{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("ldap_type"), types.StringValue("custom")),
				},
			},
			"ldap_group_dn": schema.StringAttribute{
				Optional:            true,
				Description:         "if `ldap_type`==`custom`",
				MarkdownDescription: "if `ldap_type`==`custom`",
				Validators: []validator.String{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("ldap_type"), types.StringValue("custom")),
				},
			},
			"ldap_resolve_groups": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "if `idp_type`==`ldap`, whether to recursively resolve LDAP groups",
				MarkdownDescription: "if `idp_type`==`ldap`, whether to recursively resolve LDAP groups",
				Validators: []validator.Bool{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("ldap")),
				},
			},
			"ldap_server_hosts": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "if `idp_type`==`ldap`, list of LDAP/LDAPS server IP Addresses or Hostnames",
				MarkdownDescription: "if `idp_type`==`ldap`, list of LDAP/LDAPS server IP Addresses or Hostnames",
				Validators: []validator.List{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("ldap")),
				},
			},
			"ldap_type": schema.StringAttribute{
				Optional:            true,
				Description:         "if `idp_type`==`ldap`. enum: `azure`, `custom`, `google`, `okta`, `ping_identity`",
				MarkdownDescription: "if `idp_type`==`ldap`. enum: `azure`, `custom`, `google`, `okta`, `ping_identity`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"azure",
						"custom",
						"google",
						"okta",
						"ping_identity",
					),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("ldap")),
				},
			},
			"ldap_user_filter": schema.StringAttribute{
				Optional:            true,
				Description:         "Required if `ldap_type`==`custom`, LDAP filter that will identify the type of user",
				MarkdownDescription: "Required if `ldap_type`==`custom`, LDAP filter that will identify the type of user",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("ldap_type"), types.StringValue("custom")),
				},
			},
			"member_filter": schema.StringAttribute{
				Optional:            true,
				Description:         "Required if `ldap_type`==`custom`,LDAP filter that will identify the type of member",
				MarkdownDescription: "Required if `ldap_type`==`custom`,LDAP filter that will identify the type of member",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("ldap_type"), types.StringValue("custom")),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "name",
				MarkdownDescription: "name",
			},
			"oauth_cc_client_id": schema.StringAttribute{
				Optional:            true,
				Description:         "Required if `idp_type`==`oauth`, Client Credentials",
				MarkdownDescription: "Required if `idp_type`==`oauth`, Client Credentials",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("oauth")),
				},
			},
			"oauth_cc_client_secret": schema.StringAttribute{
				Optional:            true,
				Description:         "Required if `idp_type`==`oauth`, oauth_cc_client_secret is RSA private key, of the form \"-----BEGIN RSA PRIVATE KEY--....\"",
				MarkdownDescription: "Required if `idp_type`==`oauth`, oauth_cc_client_secret is RSA private key, of the form \"-----BEGIN RSA PRIVATE KEY--....\"",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("oauth")),
				},
			},
			"oauth_discovery_url": schema.StringAttribute{
				Optional:            true,
				Description:         "if `idp_type`==`oauth`",
				MarkdownDescription: "if `idp_type`==`oauth`",
				Validators: []validator.String{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("oauth")),
				},
			},
			"oauth_ropc_client_id": schema.StringAttribute{
				Optional:            true,
				Description:         "if `idp_type`==`oauth`, ropc = Resource Owner Password Credentials",
				MarkdownDescription: "if `idp_type`==`oauth`, ropc = Resource Owner Password Credentials",
				Validators: []validator.String{
					mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("oauth")),
				},
			},
			"oauth_ropc_client_secret": schema.StringAttribute{
				Optional:            true,
				Description:         "if `oauth_type`==`azure` or `oauth_type`==`azure-gov`. oauth_ropc_client_secret can be empty",
				MarkdownDescription: "if `oauth_type`==`azure` or `oauth_type`==`azure-gov`. oauth_ropc_client_secret can be empty",
				Validators: []validator.String{
					stringvalidator.Any(
						mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("oauth_type"), types.StringValue("azure")),
						mistvalidator.AllowedWhenValueIs(path.MatchRelative().AtParent().AtName("oauth_type"), types.StringValue("azure-gov")),
					),
				},
			},
			"oauth_tenant_id": schema.StringAttribute{
				Optional:            true,
				Description:         "Required if `idp_type`==`oauth`, oauth_tenant_id",
				MarkdownDescription: "Required if `idp_type`==`oauth`, oauth_tenant_id",
				Validators: []validator.String{
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("oauth")),
				},
			},
			"oauth_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "if `idp_type`==`oauth`. enum: `azure`, `azure-gov`, `okta`, `ping_identity`",
				MarkdownDescription: "if `idp_type`==`oauth`. enum: `azure`, `azure-gov`, `okta`, `ping_identity`",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"azure",
						"azure-gov",
						"okta",
						"ping_identity",
					),
					mistvalidator.RequiredWhenValueIs(path.MatchRelative().AtParent().AtName("idp_type"), types.StringValue("oauth")),
				},
			},
			"org_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

type OrgNacidpModel struct {
	GroupFilter           types.String `tfsdk:"group_filter"`
	Id                    types.String `tfsdk:"id"`
	IdpType               types.String `tfsdk:"idp_type"`
	LdapBaseDn            types.String `tfsdk:"ldap_base_dn"`
	LdapBindDn            types.String `tfsdk:"ldap_bind_dn"`
	LdapBindPassword      types.String `tfsdk:"ldap_bind_password"`
	LdapCacerts           types.List   `tfsdk:"ldap_cacerts"`
	LdapClientCert        types.String `tfsdk:"ldap_client_cert"`
	LdapClientKey         types.String `tfsdk:"ldap_client_key"`
	LdapGroupAttr         types.String `tfsdk:"ldap_group_attr"`
	LdapGroupDn           types.String `tfsdk:"ldap_group_dn"`
	LdapResolveGroups     types.Bool   `tfsdk:"ldap_resolve_groups"`
	LdapServerHosts       types.List   `tfsdk:"ldap_server_hosts"`
	LdapType              types.String `tfsdk:"ldap_type"`
	LdapUserFilter        types.String `tfsdk:"ldap_user_filter"`
	MemberFilter          types.String `tfsdk:"member_filter"`
	Name                  types.String `tfsdk:"name"`
	OauthCcClientId       types.String `tfsdk:"oauth_cc_client_id"`
	OauthCcClientSecret   types.String `tfsdk:"oauth_cc_client_secret"`
	OauthDiscoveryUrl     types.String `tfsdk:"oauth_discovery_url"`
	OauthRopcClientId     types.String `tfsdk:"oauth_ropc_client_id"`
	OauthRopcClientSecret types.String `tfsdk:"oauth_ropc_client_secret"`
	OauthTenantId         types.String `tfsdk:"oauth_tenant_id"`
	OauthType             types.String `tfsdk:"oauth_type"`
	OrgId                 types.String `tfsdk:"org_id"`
}
