package main

import (
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/plugin"
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/tflint"
	"github.com/SoeldnerConsult/tofulint-ruleset-google/google"
	"github.com/SoeldnerConsult/tofulint-ruleset-google/project"
	"github.com/SoeldnerConsult/tofulint-ruleset-google/rules"
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
