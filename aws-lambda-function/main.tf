locals {
  name = "${var.project}-${var.env}-${var.service}"

  tags = {
    managedBy = "terraform"
    Name      = local.name
    env       = var.env
    owner     = var.owner
    service   = var.service
    owner     = var.owner
  }
}

resource aws_lambda_function lambda {
  s3_bucket = var.source_s3_bucket
  s3_key    = var.source_s3_key

  filename         = var.filename
  source_code_hash = var.source_code_hash

  function_name = local.name
  handler       = var.handler

  runtime     = var.runtime
  role        = aws_iam_role.role.arn
  timeout     = var.timeout
  kms_key_arn = var.kms_key_arn

  dynamic environment {
    for_each = length(var.environment) > 0 ? [0] : []

    content {
      variables = var.environment
    }
  }

  tags = local.tags
}

data aws_iam_policy_document role {
  statement {
    actions   = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}


resource aws_iam_role role {
  name               = local.name
  assume_role_policy = data.aws_iam_policy_document.role.json
  tags               = local.tags
}

resource aws_cloudwatch_log_group log {
  name              = "/aws/lambda/${local.name}"
  retention_in_days = var.log_retention_in_days
}

data aws_iam_policy_document lambda_logging {
  statement {
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]
    resource = aws_cloudwatch_log_group.log.arn
  }
}

resource aws_iam_policy lambda_logging {
  name_prefix = "${local.name}-lambda-logging"
  path        = "/"
  description = "IAM policy for logging from the ${local.name} lambda."
  policy      = data.aws_iam_policy_document.lambda_logging.json
}

resource aws_iam_role_policy_attachment lambda_logs {
  role       = aws_iam_role.role.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}
