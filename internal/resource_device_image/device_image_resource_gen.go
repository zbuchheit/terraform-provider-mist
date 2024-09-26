// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_device_image

import (
	"context"
	"github.com/Juniper/terraform-provider-mist/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func DeviceImageResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_id": schema.StringAttribute{
				Required: true,
			},
			"file": schema.StringAttribute{
				Required:            true,
				Description:         "path to the device image file to upload. File must be a `jpeg`, `jpg` or `png` image`",
				MarkdownDescription: "path to the device image file to upload. File must be a `jpeg`, `jpg` or `png` image`",
				Validators: []validator.String{
					stringvalidator.Any(
						mistvalidator.ParseImageType(true, true),
					),
				},
			},
			"image_number": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.Between(1, 3),
				},
			},
			"site_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

type DeviceImageModel struct {
	DeviceId    types.String `tfsdk:"device_id"`
	File        types.String `tfsdk:"file"`
	ImageNumber types.Int64  `tfsdk:"image_number"`
	SiteId      types.String `tfsdk:"site_id"`
}
