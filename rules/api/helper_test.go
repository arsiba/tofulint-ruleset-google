package api

import (
	"testing"

	"github.com/arsiba/tofulint-plugin-sdk/helper"
	"github.com/arsiba/tofulint-ruleset-google/google"
)

func NewTestRunner(t *testing.T, files map[string]string) *google.Runner {
	return &google.Runner{
		Runner:  helper.TestRunner(t, files),
		Client:  &google.Client{},
		Project: "foo-bar-baz",
	}
}
