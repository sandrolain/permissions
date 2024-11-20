package svcgrpc

import (
	context "context"
	"strings"
	"testing"

	"github.com/sandrolain/permissions/internal/dbsvc"
	"github.com/sandrolain/permissions/internal/models"
	g "github.com/sandrolain/permissions/pkg/grpc"
)

type TestDBService struct {
	dbsvc.DBServiceInterface
	db map[string]map[string]bool
}

func NewTestDBService() *TestDBService {
	return &TestDBService{
		db: make(map[string]map[string]bool),
	}
}

func (m *TestDBService) GetScopes(ctx context.Context, entity string, scopePattern string) (res []models.Permission, err error) {
	if _, ok := m.db[entity]; !ok {
		return
	}

	for scope, allow := range m.db[entity] {
		if scopePattern != "" && !match(scope, scopePattern) {
			continue
		}
		res = append(res, models.Permission{
			Entity:  entity,
			Scope:   scope,
			Allowed: allow,
		})
	}

	return
}

func (m *TestDBService) SetScope(ctx context.Context, entity string, scope string, allow bool) (affected bool, err error) {
	if _, ok := m.db[entity]; !ok {
		m.db[entity] = make(map[string]bool)
	}

	if _, ok := m.db[entity][scope]; ok && m.db[entity][scope] == allow {
		return
	}

	m.db[entity][scope] = allow
	affected = true
	return
}

func (m *TestDBService) UnsetScope(ctx context.Context, entity string, scope string) (affected bool, err error) {
	if _, ok := m.db[entity]; !ok {
		return
	}

	if _, ok := m.db[entity][scope]; !ok {
		return
	}

	delete(m.db[entity], scope)
	affected = true
	return
}

func (m *TestDBService) IsAllowed(ctx context.Context, entity string, scope string) (found bool, res bool, err error) {
	if _, ok := m.db[entity]; !ok {
		return
	}

	if _, ok := m.db[entity][scope]; !ok {
		return
	}

	found = true
	res = m.db[entity][scope]
	return
}

func (m *TestDBService) IsAllowedNegated(ctx context.Context, scopes []string, scope string, negate bool) (found bool, res bool, err error) {
	for _, entity := range scopes {
		if found, res, err = m.IsAllowed(ctx, entity, scope); err != nil {
			return
		}

		if found && res != negate {
			return
		}
	}

	if len(scopes) > 0 {
		return
	}

	found = true
	res = negate
	return
}

func (m *TestDBService) GetUserRolesPermissions(ctx context.Context, user string) (res []models.Permission, err error) {
	if _, ok := m.db[user]; !ok {
		return
	}

	for scope, allow := range m.db[user] {
		if strings.HasPrefix(scope, models.RolePrefix) {
			continue
		}

		res = append(res, models.Permission{
			Entity:  user,
			Scope:   scope,
			Allowed: allow,
		})
	}

	return
}

func (m *TestDBService) GetUserRolesScopes(ctx context.Context, user string) (res []string, err error) {
	if _, ok := m.db[user]; !ok {
		return
	}

	for scope := range m.db[user] {
		if strings.HasPrefix(scope, models.RolePrefix) {
			res = append(res, scope)
		}
	}

	return
}

func (m *TestDBService) GetUserRoles(ctx context.Context, user string) (res []string, err error) {
	scopes, err := m.GetUserRolesScopes(ctx, user)
	if err != nil {
		return
	}

	res = make([]string, len(scopes))
	for i, scope := range scopes {
		res[i] = scope[len(models.RolePrefix):]
	}
	return
}

func match(scope, pattern string) bool {
	if pattern == "" {
		return true
	}
	if strings.HasSuffix(pattern, "*") {
		return strings.HasPrefix(scope, strings.TrimSuffix(pattern, "*"))
	}
	return scope == pattern
}

func TestGrpcServiceSetUserScope(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	req := &g.SetUserScopeRequest{
		User:    "user",
		Scope:   "scope1",
		Allowed: true,
	}

	_, err := srv.SetUserScope(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	perm, err := db.GetUserRolesPermissions(context.Background(), "user")
	if err != nil {
		t.Fatal(err)
	}

	if len(perm) != 1 {
		t.Fatal("expected 1 permission")
	}

	if !perm[0].Allowed {
		t.Fatal("expected permission 0 to be allowed")
	}
}

func TestGrpcServiceGetUserRoles(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), "user1", models.RolePrefix+"admin", true)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.SetScope(context.Background(), "user1", models.RolePrefix+"viewer", true)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.GetUserRolesRequest{
		User: "user1",
	}

	res, err := srv.GetUserRoles(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Roles) != 2 {
		t.Fatalf("expected 2 roles, got %d", len(res.Roles))
	}
}

func TestGrpcServiceSetUserRole(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	req := &g.SetUserRoleRequest{
		User: "user1",
		Role: "admin",
	}

	res, err := srv.SetUserRole(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Affected {
		t.Fatal("expected role to be set")
	}

	if len(res.Roles) != 1 {
		t.Fatalf("expected 1 role, got %d", len(res.Roles))
	}
}

func TestGrpcServiceUnsetUserRole(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), "user1", models.RolePrefix+"admin", true)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.UnsetUserRoleRequest{
		User: "user1",
		Role: "admin",
	}

	res, err := srv.UnsetUserRole(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Affected {
		t.Fatal("expected role to be unset")
	}

	if len(res.Roles) != 0 {
		t.Fatalf("expected 0 roles, got %d", len(res.Roles))
	}
}

func TestGrpcServiceGetGlobalScopes(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), models.GlobalEntity, "read", true)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.SetScope(context.Background(), models.GlobalEntity, "write", false)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.GetGlobalScopesRequest{
		ScopePattern: "*",
	}

	res, err := srv.GetGlobalScopes(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Scopes) != 2 {
		t.Fatalf("expected 2 scopes, got %d", len(res.Scopes))
	}
}

func TestGrpcServiceGetRoleScopes(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), dbsvc.FormatRoleScope("admin"), "read", true)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.SetScope(context.Background(), dbsvc.FormatRoleScope("admin"), "write", true)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.GetRoleScopesRequest{
		Role:         "admin",
		ScopePattern: "*",
	}

	res, err := srv.GetRoleScopes(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Scopes) != 2 {
		t.Fatalf("expected 2 scopes, got %d", len(res.Scopes))
	}
}

func TestGrpcServiceGetUserScopes(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), "user1", "read", true)
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.SetScope(context.Background(), "user1", "write", false)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.GetUserScopesRequest{
		User:         "user1",
		ScopePattern: "*",
	}

	res, err := srv.GetUserScopes(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if len(res.Scopes) != 2 {
		t.Fatalf("expected 2 scopes, got %d", len(res.Scopes))
	}
}

func TestGrpcServiceSetGlobalScope(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	req := &g.SetGlobalScopeRequest{
		Scope:   "read",
		Allowed: true,
	}

	res, err := srv.SetGlobalScope(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Affected {
		t.Fatal("expected scope to be set")
	}

	scopes, err := db.GetScopes(context.Background(), models.GlobalEntity, "*")
	if err != nil {
		t.Fatal(err)
	}

	if len(scopes) != 1 {
		t.Fatalf("expected 1 scope, got %d", len(scopes))
	}
}

func TestGrpcServiceSetRoleScope(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	req := &g.SetRoleScopeRequest{
		Role:    "admin",
		Scope:   "read",
		Allowed: true,
	}

	res, err := srv.SetRoleScope(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Affected {
		t.Fatal("expected scope to be set")
	}

	scopes, err := db.GetScopes(context.Background(), dbsvc.FormatRoleScope("admin"), "*")
	if err != nil {
		t.Fatal(err)
	}

	if len(scopes) != 1 {
		t.Fatalf("expected 1 scope, got %d", len(scopes))
	}
}

func TestGrpcServiceUnsetGlobalScope(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), models.GlobalEntity, "read", true)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.UnsetGlobalScopeRequest{
		Scope: "read",
	}

	res, err := srv.UnsetGlobalScope(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Affected {
		t.Fatal("expected scope to be unset")
	}

	scopes, err := db.GetScopes(context.Background(), models.GlobalEntity, "*")
	if err != nil {
		t.Fatal(err)
	}

	if len(scopes) != 0 {
		t.Fatalf("expected 0 scopes, got %d", len(scopes))
	}
}

func TestGrpcServiceUnsetRoleScope(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), dbsvc.FormatRoleScope("admin"), "read", true)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.UnsetRoleScopeRequest{
		Role:  "admin",
		Scope: "read",
	}

	res, err := srv.UnsetRoleScope(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Affected {
		t.Fatal("expected scope to be unset")
	}

	scopes, err := db.GetScopes(context.Background(), dbsvc.FormatRoleScope("admin"), "*")
	if err != nil {
		t.Fatal(err)
	}

	if len(scopes) != 0 {
		t.Fatalf("expected 0 scopes, got %d", len(scopes))
	}
}

func TestGrpcServiceUnsetUserScope(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), "user1", "read", true)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.UnsetUserScopeRequest{
		User:  "user1",
		Scope: "read",
	}

	res, err := srv.UnsetUserScope(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Affected {
		t.Fatal("expected scope to be unset")
	}

	scopes, err := db.GetScopes(context.Background(), "user1", "*")
	if err != nil {
		t.Fatal(err)
	}

	if len(scopes) != 0 {
		t.Fatalf("expected 0 scopes, got %d", len(scopes))
	}
}

func TestGrpcServiceGlobalAllowed(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), models.GlobalEntity, "read", true)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.GlobalAllowedRequest{
		Scope: "read",
	}

	res, err := srv.GlobalAllowed(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Allowed {
		t.Fatal("expected scope to be allowed")
	}
}

func TestGrpcServiceRoleAllowed(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), dbsvc.FormatRoleScope("admin"), "read", true)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.RoleAllowedRequest{
		Role:  "admin",
		Scope: "read",
	}

	res, err := srv.RoleAllowed(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Allowed {
		t.Fatal("expected scope to be allowed")
	}
}

func TestGrpcServiceUserAllowed(t *testing.T) {
	db := NewTestDBService()
	srv := NewGrpcServer(db)

	// Set up test data
	_, err := db.SetScope(context.Background(), "user1", "read", true)
	if err != nil {
		t.Fatal(err)
	}

	req := &g.UserAllowedRequest{
		User:  "user1",
		Scope: "read",
	}

	res, err := srv.UserAllowed(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if !res.Allowed {
		t.Fatal("expected scope to be allowed")
	}
}
