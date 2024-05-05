package migrations

import (
	"myTaskApp/app/configs"
	"myTaskApp/app/databases"
	"myTaskApp/features/user/data"
)

func InitialMigration() {
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&data.User{})
}
