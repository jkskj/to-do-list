package service

import (
	"3/model"
	"3/pkg/e"
	"3/serializer"
	"fmt"
	"strconv"
	"time"
)

// CreateTaskService 创建任务的服务
type CreateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
	EndTime int64  `form:"end_time" json:"end_time"`
}

type ShowTaskService struct {
}

type ListTaskService struct {
	Status   string `form:"status" json:"status"` //0 待办   1已完成
	PageNum  int    `form:"page_num" json:"page_num"`
	PageSize int    `form:"page_size" json:"page_size"`
}

type UpdateTaskService struct {
	Status int `form:"status" json:"status"` //0 待办   1已完成
}

type UpdateALLTaskService struct {
	Status   int `form:"status" json:"status"` //0 待办   1已完成
	PageNum  int `form:"page_num" json:"page_num"`
	PageSize int `form:"page_size" json:"page_size"`
}

type SearchTaskService struct {
	Info     string `form:"info" json:"info"`
	PageNum  int    `form:"page_num" json:"page_num"`
	PageSize int    `form:"page_size" json:"page_size"`
}
type DeleteTaskService struct {
}
type DeleteAllTaskService struct {
	Status string `form:"status" json:"status"` //0 待办   1已完成
}

// Create 新增一条备忘录
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    service.Status,
		StartTime: time.Now().Unix(),
		EndTime:   service.EndTime,
	}
	code := e.SUCCESS
	//数据库储存任务
	err := model.DB.Create(&task).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}
}

// Show 展示一条备忘录
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//task.AddView() //增加点击数
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}
}

func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if service.Status == "" {
		model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	} else {
		status, _ := strconv.Atoi(service.Status)
		model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Where("status=?", status).Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	}

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}

func (service *UpdateTaskService) Update(uid uint, tid string) serializer.Response {
	code := e.SUCCESS
	var task model.Task
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).First(&task, tid).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	task.Status = service.Status
	model.DB.Save(&task)
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}

}

func (service *UpdateALLTaskService) UpdateAll(uid uint) serializer.Response {
	code := e.SUCCESS
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Find(&tasks).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	for _, task := range tasks {
		task.Status = service.Status
		model.DB.Save(&task)
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}
func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)

	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}
func (service *DeleteTaskService) Delete(uid uint, tid string) serializer.Response {
	code := e.SUCCESS
	var task model.Task
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).First(&task, tid).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	model.DB.Delete(&task)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
func (service *DeleteAllTaskService) DeleteAll(uid uint) serializer.Response {
	code := e.SUCCESS
	var tasks []model.Task
	count := 0
	var err error
	if service.Status == "" {
		err = model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).Find(&tasks).Error
	} else {
		status, _ := strconv.Atoi(service.Status)
		err = model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Where("status=?", status).Count(&count).Find(&tasks).Error
	}
	fmt.Println(len(tasks))
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if len(tasks) == 0 {
		code = e.ErrorNotExistData
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	for _, task := range tasks {
		model.DB.Delete(&task)
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
