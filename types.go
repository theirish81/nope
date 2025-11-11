package nope

// Nope represents the main configuration structure containing permissions and roles
type Nope struct {
	Permissions []PermissionUnion `yaml:"permissions" json:"permissions"`
	Roles       RolesMap          `yaml:"roles" json:"roles"`
	Relations   RelationsMap      `yaml:"relations" json:"relations,omitempty"`
}

// RolesMap maps role names to their definitions
type RolesMap map[string]Role

// RelationsMap maps relation names to their definitions
type RelationsMap map[string]Relation

// Role defines a role with its description and associated permissions
type Role struct {
	Description string   `yaml:"description" json:"description"`
	Permissions []string `yaml:"permissions" json:"permissions"`
}

// PermissionsExt represents extended permissions as a map of permission to sub-permissions
type PermissionsExt struct {
	Alias       string   `yaml:"alias" json:"alias"`
	Permissions []string `yaml:"permissions" json:"permissions"`
}

// PermissionUnion allows a permission to be either a simple string or an extended map
type PermissionUnion struct {
	string
	*PermissionsExt
}

// Relation defines an edge type in the authorization graph connecting identities.
// It specifies the semantic meaning of relationships between entities and the roles
// that can be assigned across those relationships.
// Description provides human-readable documentation for this relation type.
// DefaultRoles are automatically assigned to identities connected by this relation.
// AllowedRoles enumerates all roles that may be assigned across this relation type,
// acting as a whitelist that constrains which roles are valid for this relationship
// BackRef an optional reference to the inverse edge, if applicable.
type Relation struct {
	Description  string   `yaml:"description" json:"description"`
	DefaultRoles []string `yaml:"defaultRoles" json:"default_roles"`
	AllowedRoles []string `yaml:"allowedRoles" json:"allowed_roles"`
	BackRef      *string  `yaml:"backRef" json:"back_ref"`
}
