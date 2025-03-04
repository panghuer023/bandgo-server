package bandgo

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bandgo"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"

	"gorm.io/gorm"
)

const initOrderBandPost = initOrderBandEvent + 1

type initBandPost struct{}

func init() {
	system.RegisterInit(initOrderBandPost, &initBandPost{})
}

func (i *initBandPost) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&bandgo.BandPost{},
		// &bandgo.PostComment{},
		// &bandgo.PostLike{},
	)
}

func (i *initBandPost) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&bandgo.BandPost{})
}

func (i *initBandPost) InitializerName() string {
	return bandgo.BandPost{}.TableName()
}

func (i *initBandPost) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (i *initBandPost) DataInserted(ctx context.Context) bool {
	return true
} 