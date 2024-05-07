package migrations

import (
	"myTaskApp/app/configs"
	"myTaskApp/app/databases"
	projectData "myTaskApp/features/project/data"
	userData "myTaskApp/features/user/data"
)

func InitialMigration() {
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&userData.User{}, &projectData.Project{})
}
