package dao

import (
	"fmt"
	"grpc-todolist/app/task/internal/repository/db/model"
	"os"
)

func migration() {
	//自动迁移模式
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.Task{},
		)
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
