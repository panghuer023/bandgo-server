package bandgo

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bandgo"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"

	"gorm.io/gorm"
)

const initOrderBandEvent = initOrderBandTeam + 1

type initBandEvent struct{}

func init() {
	system.RegisterInit(initOrderBandEvent, &initBandEvent{})
}

func (i *initBandEvent) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&bandgo.BandEvent{},
		// &bandgo.EventParticipant{},
	)
}

func (i *initBandEvent) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&bandgo.BandEvent{})
}

func (i *initBandEvent) InitializerName() string {
	return bandgo.BandEvent{}.TableName()
}

func (i *initBandEvent) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (i *initBandEvent) DataInserted(ctx context.Context) bool {
	return true
} 