package controllers

import (
	"datahub_exporter/config"
	tencentcloudapi "datahub_exporter/utils/tencentCloudAPI"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Data    any    `json:"data"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type CkafkaController struct{}

var ConnectResourcesErrorCount int64
var DatahubTasksErrorCount int64

// 查询Datahub连接源指标
func (g CkafkaController) CkafkaDescribeConnectResourcesMetrcis(c *gin.Context) {

	var all_res_metric string

	var SecretId = config.Config.SecretId
	var SecretKey = config.Config.SecretKey
	var region = config.Config.Region

	allRes, err := tencentcloudapi.AllCkafkaDescribeConnectResources(SecretId, SecretKey, region)
	if err != nil {
		ConnectResourcesErrorCount += 1
	}
	error_count_metric := fmt.Sprintf("connect_resource_error_count{region=\"%s\"} %d\n", region, ConnectResourcesErrorCount)
	all_res_metric += error_count_metric
	if err != nil {
		c.String(http.StatusOK, all_res_metric)
		return
	}

	for _, res := range allRes.ConnectResourceList {

		res_metric := fmt.Sprintf("connect_resource_status{region=\"%s\",ResourceId=\"%s\",ResourceName=\"%s\"} %d\n", region, res.ResourceId, res.ResourceName, res.Status)
		all_res_metric += res_metric
	}

	c.String(http.StatusOK, all_res_metric)

}

// 查询Datahub任务指标
func (g CkafkaController) CkafkaDescribeDatahubTasksMetrcis(c *gin.Context) {

	var all_res_metric string

	var SecretId = config.Config.SecretId
	var SecretKey = config.Config.SecretKey
	var region = config.Config.Region

	allRes, err := tencentcloudapi.AllCkafkaDescribeDatahubTasks(SecretId, SecretKey, region)
	if err != nil {
		DatahubTasksErrorCount += 1
	}
	error_count_metric := fmt.Sprintf("datahub_task_error_count{region=\"%s\"} %d\n", region, DatahubTasksErrorCount)
	all_res_metric += error_count_metric
	if err != nil {
		c.String(http.StatusOK, all_res_metric)
		return
	}

	for _, res := range allRes.DatahubTaskList {

		res_metric := fmt.Sprintf("datahub_task_status{region=\"%s\",TaskId=\"%s\",TaskName=\"%s\"} %d\n", region, res.TaskId, res.TaskName, res.Status)
		all_res_metric += res_metric
	}

	c.String(http.StatusOK, all_res_metric)

}
