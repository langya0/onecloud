package db

import (
	"context"
	"time"

	"github.com/yunionio/onecloud/pkg/mcclient"
	"github.com/yunionio/pkg/util/timeutils"
	"github.com/yunionio/sqlchemy"
)

type SResourceBase struct {
	SModelBase

	CreatedAt     time.Time `nullable:"false" created_at:"true" get:"user"`
	UpdatedAt     time.Time `nullable:"false" updated_at:"true"`
	UpdateVersion int       `default:"0" nullable:"false" auto_version:"true"`
	DeletedAt     time.Time ``
	Deleted       bool      `nullable:"false" default:"false"`
}

type SResourceBaseManager struct {
	SModelBaseManager
}

func NewResourceBaseManager(dt interface{}, tableName string, keyword string, keywordPlural string) SResourceBaseManager {
	return SResourceBaseManager{NewModelBaseManager(dt, tableName, keyword, keywordPlural)}
}

func (manager *SResourceBaseManager) Query(fields ...string) *sqlchemy.SQuery {
	return manager.SModelBaseManager.Query(fields...).IsFalse("deleted")
}

func CanDelete(model IModel, ctx context.Context) bool {
	err := model.ValidateDeleteCondition(ctx)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (model *SResourceBase) ResourceModelManager() IResourceModelManager {
	return model.GetModelManager().(IResourceModelManager)
}

/*func (model *SResourceBase) GetCustomizeColumns(ctx context.Context, userCred mcclient.TokenCredential, query jsonutils.JSONObject) *jsonutils.JSONDict {
	extra := model.SModelBase.GetCustomizeColumns(ctx, userCred, query)
	canDelete := CanDelete(model, ctx)
	if canDelete {
		extra.Add(jsonutils.JSONTrue, "can_delete")
	} else {
		extra.Add(jsonutils.JSONFalse, "can_delete")
	}
	return extra
}

func (model *SResourceBase) GetExtraDetails(ctx context.Context, userCred mcclient.TokenCredential, query jsonutils.JSONObject) *jsonutils.JSONDict {
	extra := model.SModelBase.GetExtraDetails(ctx, userCred, query)
	canDelete := CanDelete(model, ctx)
	if canDelete {
		extra.Add(jsonutils.JSONTrue, "can_delete")
	} else {
		extra.Add(jsonutils.JSONFalse, "can_delete")
	}
	return extra
}*/

func (model *SResourceBase) MarkDelete() error {
	model.Deleted = true
	model.DeletedAt = timeutils.UtcNow()
	return nil
}

func (model *SResourceBase) Delete(ctx context.Context, userCred mcclient.TokenCredential) error {
	return DeleteModel(ctx, userCred, model)
}
