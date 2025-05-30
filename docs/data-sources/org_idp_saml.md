---
page_title: "zitadel_org_idp_saml Data Source - terraform-provider-zitadel"
subcategory: ""
description: |-
  Datasource representing a SAML IdP of the organization.
---

# zitadel_org_idp_saml (Data Source)

Datasource representing a SAML IdP of the organization.

## Example Usage

```terraform
data "zitadel_org_idp_saml" "default" {
  org_id = data.zitadel_org.default.id
  id     = "123456789012345678"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The ID of this resource.

### Optional

- `org_id` (String) ID of the organization

### Read-Only

- `auto_linking` (String) Enable if users should get prompted to link an existing ZITADEL user to an external account if the selected attribute matches, supported values: AUTO_LINKING_OPTION_UNSPECIFIED, AUTO_LINKING_OPTION_USERNAME, AUTO_LINKING_OPTION_EMAIL
- `binding` (String) The binding
- `is_auto_creation` (Boolean) enabled if a new account in ZITADEL are created automatically on login with an external account
- `is_auto_update` (Boolean) enabled if a the ZITADEL account fields are updated automatically on each login
- `is_creation_allowed` (Boolean) enabled if users are able to create a new account in ZITADEL when using an external account
- `is_linking_allowed` (Boolean) enabled if users are able to link an existing ZITADEL user with an external account
- `metadata_xml` (String) The metadata XML as plain string
- `name` (String) Name of the IDP
- `with_signed_request` (String) Whether the SAML IDP requires signed requests
