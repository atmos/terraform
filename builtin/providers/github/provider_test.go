package github

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

const testRepo string = "test-repo"

var testUser string = os.Getenv("GITHUB_TEST_USER")
var testCollaborator string = os.Getenv("GITHUB_TEST_COLLABORATOR")

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"github": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("GITHUB_TOKEN"); v == "" {
		t.Fatal("GITHUB_TOKEN must be set for acceptance tests.")
	}
	if v := os.Getenv("GITHUB_ORGANIZATION"); v == "" {
		t.Fatalf("GITHUB_ORGANIZATION must be set for acceptance tests. The organization must be owned by GITHUB_TEST_USER and have a repository named '%s' in it.", testRepo)
	}
	if v := os.Getenv("GITHUB_TEST_USER"); v == "" {
		t.Fatal("GITHUB_TEST_USER must be set for acceptance tests. The token set as GITHUB_TOKEN must be a token for this user.")
	}
	if v := os.Getenv("GITHUB_TEST_COLLABORATOR"); v == "" {
		t.Fatal("GITHUB_TEST_COLLABORATOR must be set for acceptance tests.")
	}
	if os.Getenv("GITHUB_TEST_COLLABORATOR") == os.Getenv("GITHUB_TEST_USER") {
		t.Fatal("GITHUB_TEST_COLLABORATOR must be a different user than GITHUB_TEST_USER.")
	}
}
