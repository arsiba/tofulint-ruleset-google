package api

import "github.com/arsiba/tofulint-plugin-sdk/tflint"

// Rules is a list of rules with invoking APIs
var Rules = []tflint.Rule{
	NewGoogleDisabledAPIRule(),
}
