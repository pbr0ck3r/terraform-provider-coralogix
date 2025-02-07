package coralogix

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProvider *schema.Provider
var testAccProviderFactories map[string]func() (*schema.Provider, error)
var testAccProtoV6ProviderFactories map[string]func() (tfprotov6.ProviderServer, error)

func init() {
	testAccProvider = OldProvider()
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"coralogix": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"coralogix": providerserver.NewProtocol6WithError(NewCoralogixProvider()),
	}
}

func TestProvider(t *testing.T) {
	provider := OldProvider()
	if err := provider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = OldProvider()
}

func testAccPreCheck(t *testing.T) {
	//ctx := context.TODO()

	if os.Getenv("CORALOGIX_API_KEY") == "" {
		t.Fatalf("CORALOGIX_API_KEY must be set for acceptance tests")
	}

	if os.Getenv("CORALOGIX_ENV") == "" {
		t.Fatalf("CORALOGIX_ENV must be set for acceptance tests")
	}

	//diags := testAccProvider.Configure(ctx, terraform.NewResourceConfigRaw(nil))
	//if diags.HasError() {
	//	t.Fatal(diags[0].Summary)
	//}
}
