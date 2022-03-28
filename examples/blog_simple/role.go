package main

import "github.com/go-zoox/rbac"

func NewAdminRole() rbac.Role {
	return rbac.Role{
		RoleID: "Admin",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("*", "*"),
		},
	}
}

func NewGuestRole() rbac.Role {
	return rbac.Role{
		RoleID: "Guest",
		Permissions: []rbac.Permission{
			rbac.NewGlobPermission("ReadArticle", "*"),
			rbac.NewGlobPermission("RateArticle", "*"),
		},
	}
}
