// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_site_wxrule

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func SiteWxruleResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"action": schema.StringAttribute{
				Required:            true,
				Description:         "type of action, allow / block",
				MarkdownDescription: "type of action, allow / block",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"",
						"allow",
						"block",
					),
				},
			},
			"apply_tags": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
			"blocked_apps": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "blocked apps (always blocking, ignoring action), the key of Get Application List",
				MarkdownDescription: "blocked apps (always blocking, ignoring action), the key of Get Application List",
			},
			"dst_allow_wxtags": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "tag list to indicate these tags are allowed access",
				MarkdownDescription: "tag list to indicate these tags are allowed access",
			},
			"dst_deny_wxtags": schema.ListAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				Description:         "tag list to indicate these tags are blocked access",
				MarkdownDescription: "tag list to indicate these tags are blocked access",
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"order": schema.Int64Attribute{
				Required:            true,
				Description:         "the order how rules would be looked up, > 0 and bigger order got matched first, -1 means LAST, uniqueness not checked",
				MarkdownDescription: "the order how rules would be looked up, > 0 and bigger order got matched first, -1 means LAST, uniqueness not checked",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"org_id": schema.StringAttribute{
				Computed: true,
			},
			"site_id": schema.StringAttribute{
				Required: true,
			},
			"src_wxtags": schema.ListAttribute{
				ElementType:         types.StringType,
				Required:            true,
				Description:         "tag list to determine if this rule would match",
				MarkdownDescription: "tag list to determine if this rule would match",
			},
			"template_id": schema.StringAttribute{
				Optional:            true,
				Description:         "Only for Org Level WxRule",
				MarkdownDescription: "Only for Org Level WxRule",
			},
		},
	}
}

type SiteWxruleModel struct {
	Action         types.String `tfsdk:"action"`
	ApplyTags      types.List   `tfsdk:"apply_tags"`
	BlockedApps    types.List   `tfsdk:"blocked_apps"`
	DstAllowWxtags types.List   `tfsdk:"dst_allow_wxtags"`
	DstDenyWxtags  types.List   `tfsdk:"dst_deny_wxtags"`
	Enabled        types.Bool   `tfsdk:"enabled"`
	Id             types.String `tfsdk:"id"`
	Order          types.Int64  `tfsdk:"order"`
	OrgId          types.String `tfsdk:"org_id"`
	SiteId         types.String `tfsdk:"site_id"`
	SrcWxtags      types.List   `tfsdk:"src_wxtags"`
	TemplateId     types.String `tfsdk:"template_id"`
}
