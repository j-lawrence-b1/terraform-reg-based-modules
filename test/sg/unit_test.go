// h/t: https://github.com/gruntwork-io/terratest/blob/master/test/terraform_aws_example_test.go

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the Terraform module in examples/terraform-aws-example using Terratest.
func Test_sg(t *testing.T) {
	t.Parallel()

    module := "sg"

    // Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: "../../examples/" + module,
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	security_group_id := terraform.Output(t, terraformOptions, "security_group_id")
	assert.NotNil(t, security_group_id)
	vpc_id := terraform.Output(t, terraformOptions, "vpc_id")
	assert.NotNil(t, vpc_id)
}
