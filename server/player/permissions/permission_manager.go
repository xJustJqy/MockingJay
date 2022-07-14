package permissions

type PermissionManager struct {
	permissions map[string]bool
}

func New(permissions []string) *PermissionManager {
	perms := map[string]bool{}
	for _, perm := range permissions {
		perms[perm] = true
	}
	return &PermissionManager{
		permissions: perms,
	}
}

func (p *PermissionManager) Allow(permission string) {
	p.permissions[permission] = true
}

func (p *PermissionManager) Deny(permission string) {
	p.permissions[permission] = false
}

func (p *PermissionManager) HasPermission(permission string) bool {
	if len(permission) < 1 {
		return true
	}
	for perm, allowed := range p.permissions {
		if permission == perm && allowed {
			return true
		}
	}
	return false
}

func (p *PermissionManager) GetPermissions() []string {
	perms := []string{}
	for perm, allowed := range p.permissions {
		if allowed {
			perms = append(perms, perm)
		}
	}
	return perms
}