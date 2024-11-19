package dbsvc

import (
	"strings"

	g "github.com/sandrolain/permissions/pkg/grpc"
	"github.com/sandrolain/permissions/pkg/models"
)

func EscapePattern(pattern string) string {
	pattern = strings.ReplaceAll(pattern, "_", "\\_")
	pattern = strings.ReplaceAll(pattern, "%", "\\%")
	pattern = strings.ReplaceAll(pattern, "*", "%")
	return pattern
}

func GetScopes(data []models.Permission, prefix string) []string {
	from := len(prefix)
	res := make([]string, len(data))
	for i, v := range data {
		res[i] = v.Scope[from:]
	}
	return res
}

func GetScopeItems(data []models.Permission) []*g.ScopeItem {
	res := make([]*g.ScopeItem, len(data))

	for i, v := range data {
		res[i] = &g.ScopeItem{
			Scope:   v.Scope,
			Allowed: v.Allowed,
		}
	}
	return res
}

func FormatRoleScope(role string) string {
	return models.RolePrefix + role
}
