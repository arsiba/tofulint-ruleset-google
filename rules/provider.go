package rules

import (
	"github.com/arsiba/tofulint-plugin-sdk/tflint"
	"github.com/arsiba/tofulint-ruleset-google/rules/api"
	"github.com/arsiba/tofulint-ruleset-google/rules/magicmodules"
)

var manualRules = []tflint.Rule{
	NewGoogleComposerEnvironmentInvalidMachineTypeRule(),
	NewGoogleComputeInstanceInvalidMachineTypeRule(),
	NewGoogleComputeReservationInvalidMachineTypeRule(),
	NewGoogleComputeInstanceTemplateInvalidMachineTypeRule(),
	NewGoogleContainerClusterInvalidMachineTypeRule(),
	NewGoogleContainerNodePoolInvalidMachineTypeRule(),
	NewGoogleDataflowJobInvalidMachineTypeRule(),
	NewGoogleComputeResourcePolicyInvalidNameRule(),
	NewGoogleProjectIamMemberInvalidMemberRule(),
	NewGoogleProjectIamAuditConfigInvalidMemberRule(),
	NewGoogleProjectIamBindingInvalidMemberRule(),
	NewGoogleProjectIamPolicyInvalidMemberRule(),
}

// Rules is a list of all rules
var Rules []tflint.Rule

func init() {
	Rules = append(Rules, manualRules...)
	Rules = append(Rules, magicmodules.Rules...)
	Rules = append(Rules, api.Rules...)
}
