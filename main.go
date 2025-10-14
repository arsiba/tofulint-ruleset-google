package main

import (
	"github.com/arsiba/tofulint-plugin-sdk/plugin"
	"github.com/arsiba/tofulint-plugin-sdk/tflint"
	"github.com/arsiba/tofulint-ruleset-google/google"
	"github.com/arsiba/tofulint-ruleset-google/project"
	"github.com/arsiba/tofulint-ruleset-google/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &google.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "google",
				Version: project.Version,
				Rules:   rules.Rules,
			},
		},
	})
}
