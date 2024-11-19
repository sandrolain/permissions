package dbsvc

import (
	"context"
	"testing"

	"github.com/sandrolain/permissions/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Permission{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestNewDBService(t *testing.T) {
	db := setupTestDB(t)
	svc := NewDBService(db)
	assert.NotNil(t, svc)
	assert.NotNil(t, svc.db)
}

func TestDBService_GetScopes(t *testing.T) {
	db := setupTestDB(t)
	svc := NewDBService(db)
	ctx := context.Background()

	// Create test data
	testData := []models.Permission{
		{Entity: "user1", Scope: "read:docs", Allowed: true},
		{Entity: "user1", Scope: "write:docs", Allowed: false},
		{Entity: "user1", Scope: "R$admin", Allowed: true}, // This should be filtered out
		{Entity: "user2", Scope: "read:docs", Allowed: true},
	}
	for _, p := range testData {
		err := db.Create(&p).Error
		assert.NoError(t, err)
	}

	// Test cases
	tests := []struct {
		name         string
		entity       string
		scopePattern string
		wantCount    int
		wantError    bool
	}{
		{
			name:         "get all scopes for user1",
			entity:       "user1",
			scopePattern: "",
			wantCount:    2, // Should not include R$admin
		},
		{
			name:         "get scopes with pattern",
			entity:       "user1",
			scopePattern: "read:*",
			wantCount:    1,
		},
		{
			name:         "get scopes for non-existent user",
			entity:       "user3",
			scopePattern: "",
			wantCount:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			perms, err := svc.GetScopes(ctx, tt.entity, tt.scopePattern)
			if tt.wantError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, perms, tt.wantCount)
		})
	}
}

func TestDBService_SetScope(t *testing.T) {
	db := setupTestDB(t)
	svc := NewDBService(db)
	ctx := context.Background()

	tests := []struct {
		name        string
		entity      string
		scope       string
		allow       bool
		wantChanged bool
		wantError   bool
	}{
		{
			name:        "create new permission",
			entity:      "user1",
			scope:       "read:docs",
			allow:       true,
			wantChanged: true,
		},
		{
			name:        "update existing permission",
			entity:      "user1",
			scope:       "read:docs",
			allow:       false,
			wantChanged: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			affected, err := svc.SetScope(ctx, tt.entity, tt.scope, tt.allow)
			if tt.wantError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantChanged, affected)

			// Verify the permission was set correctly
			var perm models.Permission
			err = db.Where("entity = ? AND scope = ?", tt.entity, tt.scope).First(&perm).Error
			assert.NoError(t, err)
			assert.Equal(t, tt.allow, perm.Allowed)
		})
	}
}

func TestDBService_UnsetScope(t *testing.T) {
	db := setupTestDB(t)
	svc := NewDBService(db)
	ctx := context.Background()

	// Create test data
	perm := models.Permission{
		Entity:  "user1",
		Scope:   "read:docs",
		Allowed: true,
	}
	err := db.Create(&perm).Error
	assert.NoError(t, err)

	tests := []struct {
		name        string
		entity      string
		scope       string
		wantChanged bool
		wantError   bool
	}{
		{
			name:        "unset existing permission",
			entity:      "user1",
			scope:       "read:docs",
			wantChanged: true,
		},
		{
			name:        "unset non-existent permission",
			entity:      "user1",
			scope:       "write:docs",
			wantChanged: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			affected, err := svc.UnsetScope(ctx, tt.entity, tt.scope)
			if tt.wantError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantChanged, affected)

			// Verify the permission was unset
			var count int64
			db.Model(&models.Permission{}).Where("entity = ? AND scope = ?", tt.entity, tt.scope).Count(&count)
			assert.Equal(t, int64(0), count)
		})
	}
}

func TestDBService_IsAllowed(t *testing.T) {
	db := setupTestDB(t)
	svc := NewDBService(db)
	ctx := context.Background()

	// Create test data
	testData := []models.Permission{
		{Entity: "user1", Scope: "read:docs", Allowed: true},
		{Entity: "user1", Scope: "write:docs", Allowed: false},
	}
	for _, p := range testData {
		err := db.Create(&p).Error
		assert.NoError(t, err)
	}

	tests := []struct {
		name      string
		entity    string
		scope     string
		wantFound bool
		wantRes   bool
		wantError bool
	}{
		{
			name:      "allowed permission",
			entity:    "user1",
			scope:     "read:docs",
			wantFound: true,
			wantRes:   true,
		},
		{
			name:      "denied permission",
			entity:    "user1",
			scope:     "write:docs",
			wantFound: true,
			wantRes:   false,
		},
		{
			name:      "non-existent permission",
			entity:    "user1",
			scope:     "delete:docs",
			wantFound: false,
			wantRes:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			found, res, err := svc.IsAllowed(ctx, tt.entity, tt.scope)
			if tt.wantError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantFound, found)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

func TestDBService_GetUserRolesAndScopes(t *testing.T) {
	db := setupTestDB(t)
	svc := NewDBService(db)
	ctx := context.Background()

	// Create test data
	testData := []models.Permission{
		{Entity: "user1", Scope: "R$admin", Allowed: true},
		{Entity: "user1", Scope: "R$user", Allowed: true},
		{Entity: "user1", Scope: "read:docs", Allowed: true}, // Non-role permission
	}
	for _, p := range testData {
		err := db.Create(&p).Error
		assert.NoError(t, err)
	}

	t.Run("GetUserRolesPermissions", func(t *testing.T) {
		perms, err := svc.GetUserRolesPermissions(ctx, "user1")
		assert.NoError(t, err)
		assert.Len(t, perms, 2) // Should only get role permissions
		scopes := []string{"R$admin", "R$user"}
		for _, p := range perms {
			assert.Contains(t, scopes, p.Scope)
		}
	})

	t.Run("GetUserRolesScopes", func(t *testing.T) {
		scopes, err := svc.GetUserRolesScopes(ctx, "user1")
		assert.NoError(t, err)
		assert.Len(t, scopes, 2)
		assert.ElementsMatch(t, []string{"R$admin", "R$user"}, scopes)
	})

	t.Run("GetUserRoles", func(t *testing.T) {
		roles, err := svc.GetUserRoles(ctx, "user1")
		assert.NoError(t, err)
		assert.Len(t, roles, 2)
		assert.ElementsMatch(t, []string{"admin", "user"}, roles)
	})
}
