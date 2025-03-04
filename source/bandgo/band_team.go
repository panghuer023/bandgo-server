package bandgo

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bandgo"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"

	"gorm.io/gorm"
)

const initOrderBandTeam = initOrderBandUser + 1

type initBandTeam struct{}

func init() {
	system.RegisterInit(initOrderBandTeam, &initBandTeam{})
}

func (i *initBandTeam) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&bandgo.BandTeam{},
		&bandgo.BandMember{},
		&bandgo.BandJoinRequest{},
	)
}

func (i *initBandTeam) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&bandgo.BandTeam{})
}

func (i *initBandTeam) InitializerName() string {
	return bandgo.BandTeam{}.TableName()
}

func (i *initBandTeam) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (i *initBandTeam) DataInserted(ctx context.Context) bool {
	return true
} 