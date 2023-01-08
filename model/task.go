package model

import (
	"3/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Task struct {
	gorm.Model
	User      User   ` gorm:"ForeignKey:Uid"` //FOREIGN KEY：指定外表字段，作为外键关联主表主键
	Uid       uint   `gorm:"not null"`        //id
	Title     string `gorm:"index;not null"`  //内容
	Status    int    `gorm:"default:'0'"`     //完成状态:0 未完成 1 已完成
	Content   string `gorm:"type:longtext"`
	StartTime int64  //添加时间
	EndTime   int64  //截止时间
}

func (Task *Task) View() uint64 {
	//增加点击数
	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(Task.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (Task *Task) AddView() {
	cache.RedisClient.Incr(cache.TaskViewKey(Task.ID))                      //增加视频点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.ID))) //增加排行点击数
}
