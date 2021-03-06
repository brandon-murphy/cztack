<!-- START -->
## Requirements

| Name | Version |
|------|---------|
| snowflake | >= 0.20.0 |

## Providers

| Name | Version |
|------|---------|
| snowflake | >= 0.20.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| monitor\_name | Identifier for the resource monitor; must be unique for your account. | `string` | `null` | no |
| per\_privilege\_grants | A map of privileges to authorized roles and shares. Privileges must be UPPER case.<br>  This allows you to authorize extra roles/shares for specific privileges. | `map(object({roles = list(string)}))` | `{}` | no |
| roles | Grants privilege to these roles. | `set(string)` | `null` | no |
| with\_grant\_option | When this is set to true, allows the recipient role to grant the privileges to other roles. | `bool` | `null` | no |

## Outputs

No output.

<!-- END -->
