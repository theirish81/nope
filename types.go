package nope

// Nope represents the main configuration structure containing permissions and roles
type Nope struct {
	Permissions []PermissionUnion `yaml:"permissions"`
	Roles       RolesMap          `yaml:"roles"`
}

// RolesMap maps role names to their definitions
type RolesMap map[string]Role

// Role defines a role with its description and associated permissions
type Role struct {
	Description string   `yaml:"description"`
	Permissions []string `yaml:"permissions"`
}

// PermissionsExt represents extended permissions as a map of permission to sub-permissions
type PermissionsExt map[string][]string

// PermissionUnion allows a permission to be either a simple string or an extended map
type PermissionUnion struct {
	string
	PermissionsExt
}
