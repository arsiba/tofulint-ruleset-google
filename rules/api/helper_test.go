package api

import (
	"testing"

	"github.com/SoeldnerConsult/tofulint-plugin-sdk/helper"
	"github.com/arsiba/tofulint-ruleset-google/google"
)

func NewTestRunner(t *testing.T, files map[string]string) *google.Runner {
	return &google.Runner{
		Runner:  helper.TestRunner(t, files),
		Client:  &google.Client{},
		Project: "foo-bar-baz",
	}
}
