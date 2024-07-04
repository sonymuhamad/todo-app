package providerwire

import (
	"github.com/sonymuhamad/todo-app/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func provideGormSQLDatabase(cfg pkg.EnvConfig) *gorm.DB {
	//var logLevel logger.LogLevel
	//if cfg.Env == "development" {
	//	logLevel = logger.Info
	//} else {
	//	logLevel = logger.Warn
	//}
	//
	//customLogger := &logger2.GormCustomLogger{
	//	LogLevel: logLevel,
	//}
	//
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  cfg.DatabaseUrl,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		FullSaveAssociations:   false,
	})

	if err != nil {
		panic(err)
	}

	return db
}
