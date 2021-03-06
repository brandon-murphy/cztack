package test

import (
	"testing"

	"github.com/chanzuckerberg/go-misc/tftest"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestIAMRoleBless(t *testing.T) {
	test := tftest.Test{
		Setup: func(t *testing.T) *terraform.Options {
			region := tftest.IAMRegion
			curAcct := tftest.AWSCurrentAccountID(t)

			project := tftest.UniqueID()
			env := tftest.UniqueID()
			service := tftest.UniqueID()
			owner := tftest.UniqueID()

			return &terraform.Options{
				TerraformDir: ".",

				Vars: map[string]interface{}{
					"role_name":         random.UniqueId(),
					"source_account_id": curAcct,
					"project":           project,
					"env":               env,
					"service":           service,
					"owner":             owner,
					"bless_lambda_arns": []string{"arn:aws:lambda:us-west-2:111111111111:function:test"},
				},
				EnvVars: map[string]string{
					"AWS_DEFAULT_REGION": region,
				},
			}
		},
		Validate: func(t *testing.T, options *terraform.Options) {},
	}

	test.Run(t)
}
