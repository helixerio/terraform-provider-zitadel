package machine_key

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zitadel/zitadel-go/v3/pkg/client/zitadel/authn"

	"github.com/zitadel/terraform-provider-zitadel/v2/zitadel/helper"
)

func GetResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing a machine key",
		Schema: map[string]*schema.Schema{
			helper.OrgIDVar: helper.OrgIDResourceField,
			UserIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the user",
				ForceNew:    true,
			},
			keyTypeVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the machine key" + helper.DescriptionEnumValuesList(authn.KeyType_name),
				ForceNew:    true,
				ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
					return helper.EnumValueValidation(keyTypeVar, value, authn.KeyType_value)
				},
			},
			PublicKeyVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optionally provide a public key of your own generated RSA private key",
				ForceNew:    true,
			},
			ExpirationDateVar: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Expiration date of the machine key in the RFC3339 format",
				ForceNew:    true,
				Computed:    true,
			},
			KeyDetailsVar: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Value of the machine key",
				Sensitive:   true,
			},
		},
		DeleteContext: delete,
		CreateContext: create,
		ReadContext:   read,
		Importer: helper.ImportWithIDAndOptionalOrg(
			keyIDVar,
			helper.NewImportAttribute(UserIDVar, helper.ConvertID, false),
			helper.NewImportAttribute(KeyDetailsVar, helper.ConvertJSON, true),
			helper.NewImportAttribute(PublicKeyVar, helper.ConvertBase64, true),
		),
	}
}
