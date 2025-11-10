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

// Relation defines a directed edge in an entity graph, flowing from source to target.
// OutboundType names the relation type (e.g., "works_for") flowing toward the target.
// InboundType optionally names the reverse relation type (e.g., "employs") flowing back.
// Each direction can specify default roles to be applied.
type Relation struct {
	OutboundType         string    `yaml:"outboundType"`
	InboundType          *string   `yaml:"inboundType"`
	Description          string    `yaml:"description"`
	DefaultOutboundRoles []string  `yaml:"defaultOutboundRoles"`
	DefaultInboundRoles  *[]string `yaml:"defaultInboundRoles"`
}
