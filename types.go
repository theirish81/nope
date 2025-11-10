package nope

// Nope represents the main configuration structure containing permissions and roles
type Nope struct {
	Permissions []PermissionUnion `yaml:"permissions"`
	Roles       RolesMap          `yaml:"roles"`
	Relations   RelationsMap      `yaml:"relations"`
}

// RolesMap maps role names to their definitions
type RolesMap map[string]Role

// RelationsMap maps relation names to their definitions
type RelationsMap map[string]Relation

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

// Relation defines an edge type in the authorization graph connecting identities.
// It specifies the semantic meaning of relationships between entities and the roles
// that can be assigned across those relationships.
// Type is the semantic verb describing this relationship (e.g., "owns", "is_member", "employs").
// Description provides human-readable documentation for this relation type.
// DefaultRoles are automatically assigned to identities connected by this relation.
// AllowedRoles enumerates all roles that may be assigned across this relation type,
// acting as a whitelist that constrains which roles are valid for this relationship.
type Relation struct {
	Type         string   `yaml:"type"`
	Description  string   `yaml:"description"`
	DefaultRoles []string `yaml:"defaultRoles"`
	AllowedRoles []string `yaml:"allowedRoles"`
}
