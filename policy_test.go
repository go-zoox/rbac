package rbac

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPolicyTemplate(t *testing.T) {
	// t.Skip("TODO: read/write to/from a buffer")

	p := NewPolicyTemplate("Admin")
	p.AddPermission("glob", "*", "grid:*:userID:*")
	p.AddPermission("glob", "read:*", "*")

	bytes, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	if err := ioutil.WriteFile("admin.json", bytes, 0644); err != nil {
		t.Fatal(err)
	}

	bytes, err = ioutil.ReadFile("admin.json")
	if err != nil {
		t.Fatal(err)
	}

	var policy PolicyTemplate
	if err := json.Unmarshal(bytes, &policy); err != nil {
		t.Fatal(err)
	}

	role, err := policy.Role(*strings.NewReplacer("$userID", "u123"))
	if err != nil {
		t.Fatal(err)
	}

	can, err := role.Can("read:comment", "c123")
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, can)

	os.Remove("admin.json")
}
