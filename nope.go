package nope

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

// MarshalJSON implements custom JSON marshaling
func (u *PermissionUnion) MarshalJSON() ([]byte, error) {
	// If Simple has a value, marshal it as a string
	if u.string != "" {
		return json.Marshal(u.string)
	}
	// Otherwise, marshal Extended as an object
	return json.Marshal(u.PermissionsExt)
}

// UnmarshalYAML implements custom YAML unmarshaling for PermissionUnion
// It attempts to decode as a string first, then as a map if that fails
func (u *PermissionUnion) UnmarshalYAML(value *yaml.Node) error {
	// Try to decode as a simple string
	var s string
	if value.Decode(&s) == nil {
		u.string = s
		u.PermissionsExt = nil
		return nil
	}

	// Try to decode as an extended permissions map
	var m *PermissionsExt
	if value.Decode(&m) == nil {
		u.PermissionsExt = m
		return nil
	}

	// Both attempts failed
	return fmt.Errorf("permission union must be a string or a map, got: %v", value.Tag)
}

func (u *PermissionUnion) IsComposite() bool {
	return u.PermissionsExt != nil
}

func (u *PermissionUnion) Key() string {
	if u.IsComposite() {
		return u.PermissionsExt.Alias
	}
	return u.string
}

// FromYAML parses YAML data into a Nope configuration
func FromYAML(data []byte) (Nope, error) {
	var nope Nope
	err := yaml.Unmarshal(data, &nope)
	return nope, err
}

// ResolvePermissions resolves all permissions for the given roles
// It expands extended permissions into their constituent parts
func (a Nope) ResolvePermissions(roles ...string) []string {
	var permissions []string

	// Iterate through each role
	for _, role := range roles {
		// Get permissions defined for this role
		for _, p := range a.Roles[role].Permissions {
			// Find the permission definition in the permissions list
			if p2, found := lo.Find(a.Permissions, func(item PermissionUnion) bool {
				return item.Key() == p
			}); found {
				// If it's an extended permission, all sub-permissions
				if p2.IsComposite() {
					permissions = append(permissions, p2.PermissionsExt.Permissions...)
				} else {
					// Simple permission, add it directly
					permissions = append(permissions, p)
				}
			}
		}
	}

	return permissions
}

// HasAtLeastOnePermission checks if the given roles have at least one of the specified permissions
func (a Nope) HasAtLeastOnePermission(permissions []string, roles ...string) bool {
	for _, px := range a.ResolvePermissions(roles...) {
		if slices.Contains(permissions, px) {
			return true
		}
	}
	return false
}

// HasAllPermissions checks if the given roles have all of the specified permissions
func (a Nope) HasAllPermissions(permissions []string, roles ...string) bool {
	for _, px := range a.ResolvePermissions(roles...) {
		if !slices.Contains(permissions, px) {
			return false
		}
	}
	return true
}
