package dbsvc

import (
	"testing"

	g "github.com/sandrolain/permissions/pkg/grpc"
	"github.com/sandrolain/permissions/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestEscapePattern(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "escape underscore",
			input:    "test_pattern",
			expected: "test\\_pattern",
		},
		{
			name:     "escape percent",
			input:    "test%pattern",
			expected: "test\\%pattern",
		},
		{
			name:     "convert asterisk to percent",
			input:    "test*pattern",
			expected: "test%pattern",
		},
		{
			name:     "multiple special characters",
			input:    "test_*%pattern",
			expected: "test\\_%\\%pattern",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EscapePattern(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetScopes(t *testing.T) {
	prefix := "R$"
	permissions := []models.Permission{
		{Scope: "R$admin", Allowed: true},
		{Scope: "R$user", Allowed: false},
		{Scope: "R$guest", Allowed: true},
	}

	expected := []string{"admin", "user", "guest"}
	result := GetScopes(permissions, prefix)

	assert.Equal(t, expected, result)
}

func TestGetScopeItems(t *testing.T) {
	permissions := []models.Permission{
		{Scope: "R$admin", Allowed: true},
		{Scope: "R$user", Allowed: false},
	}

	expected := []*g.ScopeItem{
		{Scope: "R$admin", Allowed: true},
		{Scope: "R$user", Allowed: false},
	}

	result := GetScopeItems(permissions)

	assert.Equal(t, len(expected), len(result))
	for i, v := range expected {
		assert.Equal(t, v.Scope, result[i].Scope)
		assert.Equal(t, v.Allowed, result[i].Allowed)
	}
}

func TestFormatRoleScope(t *testing.T) {
	tests := []struct {
		name     string
		role     string
		expected string
	}{
		{
			name:     "normal role",
			role:     "admin",
			expected: "R$admin",
		},
		{
			name:     "empty role",
			role:     "",
			expected: "R$",
		},
		{
			name:     "role with special characters",
			role:     "super-admin",
			expected: "R$super-admin",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatRoleScope(tt.role)
			assert.Equal(t, tt.expected, result)
		})
	}
}
