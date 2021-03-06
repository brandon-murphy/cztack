# AWS SNS to LAMBDA Integration
This module is for the integration of AWS SNS subscribing to lambda. It supports multiple SNS topics subscribing to lambda. The module references the [cztack/aws-lambda-function](https://github.com/chanzuckerberg/cztack/tree/main/aws-lambda-function) module for the lambda function.

## Example

```hcl
module aws-sns-lambda {
  source                  = "github.com/chanzuckerberg/cztack//aws-sns-lambda?ref=v0.36.0"
  env                     = local.env
  lambda_environment      = {
    "foob" : bar
  }
  lambda_function_name    = "lambda"
  lambda_handler          = "index.handler"
  lambda_runtime          = "nodejs12.x"
  lambda_source_code_hash = data.archive_file.lambda_archive.output_base64sha256
  lambda_timeout          = local.lambda_timeout
  lambda_zip_file         = data.archive_file.lambda_archive.output_path
  owner                   = var.owner
  project                 = var.project
  service                 = var.service
  #  If there is only one SNS config, and topic_name is not provided, then use [var.project]-[var.env]-[var.service] as default topic name, otherwise, the SNS config without the topic_name will be discarded.
  sns_configurations      = [
    {
      topic_name : "...",
      topic_display_name : "...",
      endpoint_auto_confirms          : true
    },
    {
      topic_name : "heroku-dyno",
      topic_display_name : null,
      endpoint_auto_confirms          : false
      confirmation_timeout_in_minutes : 10
    }
  ]
  publish_lambda          = true
  lambda_role_path        = "/service-role/"
  function_description    = "A lambda function..."

}

data "archive_file" "lambda_archive" {
  type        = "zip"
  source_file = "${path.module}/lambda_src/index.js"
  output_path = "${path.module}/lambda_src/lambda_build.zip"
}
```


<!-- START -->
## Requirements

| Name | Version |
|------|---------|
| aws | < 3.0.0 |

## Providers

| Name | Version |
|------|---------|
| aws | < 3.0.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| env | Env for tagging and naming. | `string` | n/a | yes |
| function\_description | Description for the lambda function. | `string` | `""` | no |
| lambda\_environment | Environment for lambda. | `map` | `{}` | no |
| lambda\_function\_name | Name for lambda function. If not set, function use default naming convention of $project-$env-$service. | `string` | `null` | no |
| lambda\_handler | Name of the lambda handler. | `string` | n/a | yes |
| lambda\_role\_path | Path of the IAM role for lambda. | `string` | `null` | no |
| lambda\_runtime | Lambda language runtime. | `string` | n/a | yes |
| lambda\_source\_code\_hash | Lambda source code, used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file. | `string` | `null` | no |
| lambda\_timeout | Execution timeout for the lambda. | `number` | `null` | no |
| lambda\_zip\_file | The path to the function's deployment package within the local filesystem. | `string` | n/a | yes |
| log\_retention\_in\_days | The number of days to retain log events in the log group for lambda. | `number` | `null` | no |
| owner | Owner for tagging and naming. | `string` | n/a | yes |
| project | Project for tagging and naming. | `string` | n/a | yes |
| publish\_lambda | Whether to publish creation/change as new Lambda Function Version. | `bool` | `false` | no |
| service | Service for tagging and naming | `string` | n/a | yes |
| sns\_configurations | List of objects containing the SNS configurations. | `list(any)` | `[]` | no |

## Outputs

| Name | Description |
|------|-------------|
| lambda\_role\_name | Name for the lambda execution role. |
| sns\_topic\_arns | SNS topic ARNs. |

<!-- END -->
