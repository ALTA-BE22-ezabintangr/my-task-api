package migrations

import (
	"myTaskApp/app/configs"
	"myTaskApp/app/databases"
	projectData "myTaskApp/features/project/data"
	taskData "myTaskApp/features/task/data"
	userData "myTaskApp/features/user/data"
)

func InitialMigration() {
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&userData.User{}, &projectData.Project{}, &taskData.Task{})
}
