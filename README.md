# Nope ( WORK IN PROGRESS!! )

A lightweight role and permission management library for Go applications.

## Overview

Nope provides a simple way to define and manage role-based access control (RBAC) in your Go applications.
Define your permissions and roles in YAML, then use the library to check access rights with minimal overhead.

## Features

- **YAML-Ready** - Define permissions and roles in a human-readable format (or not)
- **Extended permissions** - Support for hierarchical permissions with parent-child relationships
- **Flexible permission checks** - Validate if roles have at least one or all required permissions

## Installation

```bash
go get github.com/theirish81/nope
```

## Usage

### Define Your Permissions and Roles

Create a YAML file or a data structure as the following example:

```yaml
permissions:
  - user:read
  - user:write
  - user:admin:
      - user:read
      - user:write
      - user:delete

roles:
  viewer:
    description: Can view user data
    permissions:
      - user:read
  
  editor:
    description: Can view and edit user data
    permissions:
      - user:read
      - user:write

  user:admin:
    description: Full system access
    permissions:
      - admin
```

### Load and Use

```go
package main

import (
    "fmt"
    "os"
    "github.com/theirish81/nope"
)

func main() {
    // Load configuration from YAML
    data, _ := os.ReadFile("permissions.yaml")
    config, err := nope.FromYAML(data)
    if err != nil {
        panic(err)
    }

    // Get all permissions for a role
    permissions := config.GetPermissions("editor")
    fmt.Println(permissions) // [user:read user:write]

    // Check if a role has at least one permission
    hasAccess := config.HasAtLeastOnePermission(
        []string{"user:read", "user:delete"},
        "editor",
    )
    fmt.Println(hasAccess) // true (has user:read)

    // Check if a role has all permissions
    hasAll := config.HasAllPermissions(
        []string{"user:read", "user:write"},
        "editor",
    )
    fmt.Println(hasAll) // true

    // Check multiple roles
    adminPerms := config.GetPermissions("admin")
    fmt.Println(adminPerms) // [admin user:read user:write user:delete system:manage]
}
```

## API Reference

### Types

#### `Nope`
Main configuration structure containing permissions and roles.

#### `Role`
Represents a role with a description and list of associated permissions.

#### `PermissionUnion`
Allows permissions to be defined as either:
- A simple string (e.g., `user:read`)
- An extended map with sub-permissions (e.g., `user:admin [user:read, user:write, ...]`)

### Methods

#### `FromYAML(data []byte) (Nope, error)`
Parses YAML configuration data into a Nope instance.

#### `GetPermissions(roles ...string) []string`
Returns all permissions for the given roles, expanding extended permissions into their constituent parts.

#### `HasAtLeastOnePermission(permissions []string, roles ...string) bool`
Checks if the given roles have at least one of the specified permissions.

#### `HasAllPermissions(permissions []string, roles ...string) bool`
Checks if the given roles have all of the specified permissions.

## Permission Patterns

### Simple Permissions
```yaml
permissions:
  - user:read
  - user:write
  - blog:read
  - blog:write
```

### Extended Permissions
Extended permissions allow you to define a parent permission that includes multiple sub-permissions:

```yaml
permissions:
  - user:read
  - user:write
  - blog:read
  - blog:write
  - user:admin:
      - user:read
      - user:write
      - user:delete
```

When a role is granted the `admin` permission, it automatically receives all listed sub-permissions.

## Use Cases

- **API authorization** - Check if a user's roles allow access to specific endpoints
- **Feature flags** - Control feature access based on user roles
- **Resource access control** - Manage permissions for different resource types
- **Multi-tenant applications** - Define role hierarchies per tenant

## Dependencies

- [samber/lo](https://github.com/samber/lo) - Utility functions
- [gopkg.in/yaml.v3](https://gopkg.in/yaml.v3) - YAML parsing

## License

MIT License - See LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.