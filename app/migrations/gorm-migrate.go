package migrations

import (
	"myTaskApp/app/configs"
	"myTaskApp/app/databases"
	_projectData "myTaskApp/features/project/data"
	_userData "myTaskApp/features/user/data"
)

func InitialMigration() {
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&_userData.User{})
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&_projectData.Project{})
}
