package model

// 执行数据迁移
func migration() {
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).AutoMigrate(&Task{})
	/*
		Model()函数传入那个表需要的外键
		AddForeignKey()
		函数需要四个参数：
		1.外键键名（迁移数据库之后生成的字段名，不是结构体中的字段名）
		2.关联的外键：关联的表名(字段名)，同样都是生成的数据库的表名和字段名
		3.删除时的状态：set null ,no action,cascade,restrict，
		其中restrict和no action相同，如果子表有关联数据，父表对应数据不能进行删除，
		cascade模式：如果删除父表数据，则对应的子表数据也相应删除，
		set null模式：如果父表删除数据，对应的子表的数据的外键被设置为null，但是相应的这个模式要求子表外键字段可以为null
		4.更新时表的状态，与删除时的四种状态相同
	*/
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")

}
