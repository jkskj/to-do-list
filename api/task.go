package api

import (
	"3/pkg/utils"
	"3/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"net/http"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&createTask)
	if err == nil {
		res := createTask.Create(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}

}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	err := c.ShouldBind(&showTask)
	if err == nil {
		res := showTask.Show(c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}

func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	err := c.ShouldBind(&listTask)
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err == nil {
		res := listTask.List(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
func UpdateAllTask(c *gin.Context) {
	var updateAllTask service.UpdateALLTaskService
	err := c.ShouldBind(&updateAllTask)
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err == nil {
		res := updateAllTask.UpdateAll(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	err := c.ShouldBind(&updateTask)
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err == nil {
		res := updateTask.Update(chaim.Id, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	err := c.ShouldBind(&searchTask)
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err == nil {
		res := searchTask.Search(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	err := c.ShouldBind(&deleteTask)
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err == nil {
		res := deleteTask.Delete(chaim.Id, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
func DeleteAllTask(c *gin.Context) {
	var deleteAllTask service.DeleteAllTaskService
	err := c.ShouldBind(&deleteAllTask)
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err == nil {
		res := deleteAllTask.DeleteAll(chaim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Error(err)
	}
}
