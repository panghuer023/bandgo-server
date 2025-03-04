package bandgo

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/model/bandgo"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"

	"gorm.io/gorm"
)

const initOrderBandUser = system.InitOrderSystem + 1

type initBandUser struct{}

// auto run
func init() {
	system.RegisterInit(initOrderBandUser, &initBandUser{})
}

func (i *initBandUser) InitializerName() string {
	return bandgo.BandUser{}.TableName()
}

func (i *initBandUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&bandgo.BandUser{})
}

func (i *initBandUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&bandgo.BandUser{})
}

func (i *initBandUser) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil // 暂时不需要初始化数据
}

func (i *initBandUser) DataInserted(ctx context.Context) bool {
	return true // 暂时不需要检查数据
}
