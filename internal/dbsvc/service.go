package dbsvc

import (
	"context"
	"fmt"

	"github.com/sandrolain/gomsvc/pkg/dblib"
	"github.com/sandrolain/permissions/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewDBService(db *gorm.DB) *DBService {
	return &DBService{
		db: db,
	}
}

type DBService struct {
	db *gorm.DB
}

func (m *DBService) GetScopes(ctx context.Context, entity string, scopePattern string) (res []models.Permission, err error) {
	q := m.db.WithContext(ctx).
		Where("entity = ?", entity).
		Where("scope NOT LIKE ?", models.RolePrefix+"%")

	if scopePattern != "" {
		scopePattern = EscapePattern(scopePattern)
		q = q.Where("scope LIKE ?", scopePattern)
	}

	tx := q.Select("scope", "allowed").Find(&res)

	if dblib.IsQueryError(tx.Error) {
		err = tx.Error
		return
	}
	return
}

func (m *DBService) SetScope(ctx context.Context, entity string, scope string, allow bool) (affected bool, err error) {
	permission := models.Permission{
		Entity:  entity,
		Scope:   scope,
		Allowed: allow,
	}

	tx := m.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "entity"}, {Name: "scope"}},
		DoUpdates: clause.AssignmentColumns([]string{"allowed", "updated_at"}),
	}).Create(&permission)

	if dblib.IsQueryError(tx.Error) {
		err = tx.Error
		return
	}
	if tx.RowsAffected > 0 {
		affected = true
	}
	return
}

func (m *DBService) UnsetScope(ctx context.Context, entity string, scope string) (affected bool, err error) {
	q := m.db.WithContext(ctx).Where("entity = ? AND scope = ?", entity, scope)
	tx := q.Unscoped().Delete(&models.Permission{})
	if dblib.IsQueryError(tx.Error) {
		err = tx.Error
		return
	}
	if tx.RowsAffected > 0 {
		affected = true
	}
	return
}

func (m *DBService) IsAllowed(ctx context.Context, entity string, scope string) (found bool, res bool, err error) {
	res = false
	data := models.Permission{}

	tx := m.db.WithContext(ctx).
		Where(&models.Permission{
			Entity: entity,
			Scope:  scope,
		}).
		Select("allowed").
		First(&data)

	if dblib.IsQueryError(tx.Error) {
		err = tx.Error
		return
	}
	if tx.RowsAffected > 0 {
		found = true
		res = data.Allowed
	}
	return
}

func (m *DBService) IsAllowedNegated(ctx context.Context, scopes []string, scope string, negate bool) (found bool, res bool, err error) {
	res = false
	data := models.Permission{}

	tx := m.db.WithContext(ctx).
		Where("entity IN ?", scopes).
		Where("scope = ?", scope).
		Where("allow <> ?", negate).
		Select("allowed").
		First(&data)

	if dblib.IsQueryError(tx.Error) {
		err = tx.Error
		return
	}
	if tx.RowsAffected > 0 {
		found = true
		res = data.Allowed
	}
	return
}

func (m *DBService) GetUserRolesPermissions(ctx context.Context, user string) (res []models.Permission, err error) {
	q := m.db.WithContext(ctx).Select("scope").Where("entity = ?", user)
	q = q.Where(fmt.Sprintf("scope LIKE '%s%%'", models.RolePrefix))
	err = q.Select("scope").Find(&res).Error
	return
}

func (m *DBService) GetUserRolesScopes(ctx context.Context, user string) (res []string, err error) {
	perm, err := m.GetUserRolesPermissions(ctx, user)
	if err == nil {
		res = GetScopes(perm, "")
	}
	return
}

func (m *DBService) GetUserRoles(ctx context.Context, user string) (res []string, err error) {
	perm, err := m.GetUserRolesPermissions(ctx, user)
	if err == nil {
		res = GetScopes(perm, models.RolePrefix)
	}
	return
}
