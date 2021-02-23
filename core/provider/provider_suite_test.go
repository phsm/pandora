package provider

import (
	"testing"

	"github.com/phsm/pandora/core"
	"github.com/phsm/pandora/lib/ginkgoutil"
)

func TestProvider(t *testing.T) {
	ginkgoutil.RunSuite(t, "AmmoQueue Suite")
}

func testDeps() core.ProviderDeps {
	return core.ProviderDeps{ginkgoutil.NewLogger()}
}
