package test

/*
 NOTE: automatically generated by scripts/snowflake_generate_grant_all.
			 Changes made directly to this file will be overwritten.
			 Make improvements there so everyone can benefit.
*/

import (
	"encoding/json"
	"testing"

	"github.com/chanzuckerberg/go-misc/tftest"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

const vars string = "{\"database_name\":\"database\",\"roles\":[\"role_a\",\"role_b\",\"role_c\"],\"schema_name\":\"schema\",\"sequence_name\":\"sequence\"}"

func TestModule(t *testing.T) {
	test := tftest.Test{
		// just run init, swtich to Plan or Apply when you can
		Mode: tftest.Init,

		Setup: func(t *testing.T) *terraform.Options {
			opts := &terraform.Options{
				TerraformDir: ".",
				EnvVars: map[string]string{
					"AWS_DEFAULT_REGION": tftest.DefaultRegion,
				},
				Vars: mustJson(vars),
			}

			return opts
		},
		Validate: func(t *testing.T, options *terraform.Options) {},
	}

	test.Run(t)
}

func mustJson(jsonData string) map[string]interface{} {
	unmarshalled := map[string]interface{}{}

	err := json.Unmarshal([]byte(jsonData), &unmarshalled)
	if err != nil {
		panic(err) // should never happen since generated code
	}
	return unmarshalled
}
