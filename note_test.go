package nope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const yx = `
permissions:
  - user:read
  - user:write
  - user:delete
  - user:impersonate
  - user:admin:
      - user:read
      - user:write
      - user:delete
roles:
  admin:
    description: Administrators
    permissions:
      - user:admin
`

func TestFromYAML(t *testing.T) {
	nope, err := FromYAML([]byte(yx))
	assert.Nil(t, err)
	assert.Equal(t, "user:admin", nope.Roles["admin"].Permissions[0])
	assert.Equal(t, "user:read", nope.Permissions[0].Key())
	/*assert.Empty(t, nope.ResolvePermissions("foo"))

	assert.True(t, nope.HasAtLeastOnePermission([]string{"user:read"}, "admin"))
	assert.False(t, nope.HasAtLeastOnePermission([]string{"user:authenticate"}, "admin"))
	*/
}

func TestNope_HasAtLeastOnePermission(t *testing.T) {
	nope, _ := FromYAML([]byte(yx))
	assert.True(t, nope.HasAtLeastOnePermission([]string{"user:read"}, "admin"))
	assert.True(t, nope.HasAtLeastOnePermission([]string{"user:read", "user:impersonate"}, "admin"))
	assert.False(t, nope.HasAtLeastOnePermission([]string{"user:impersonate"}, "admin"))
}

func TestNope_HasAllPermissions(t *testing.T) {
	nope, _ := FromYAML([]byte(yx))
	assert.True(t, nope.HasAllPermissions([]string{"user:foo", "user:read", "user:write", "user:delete"}, "admin"))
	assert.False(t, nope.HasAllPermissions([]string{"user:read"}, "admin"))
}
