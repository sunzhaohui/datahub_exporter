package tencentcloudapi

import (
	"fmt"

	ckafka "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ckafka/v20190819"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func getCkafkaClient(SecretId string, SecretKey string, region string) (*ckafka.Client, error) {
	credential := common.NewCredential(
		SecretId,
		SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ckafka.ap-guangzhou.tencentcloudapi.com"
	client, err := ckafka.NewClient(credential, region, cpf)
	return client, err
}

type ConnectResource struct {
	ResourceId   string
	ResourceName string
	Status       int64
}

type ConnectResources struct {
	TotalCount          int64
	ConnectResourceList []ConnectResource
}

type DatahubTask struct {
	TaskId   string
	TaskName string
	Status   int64
}
type DatahubTasks struct {
	TotalCount      int64
	DatahubTaskList []DatahubTask
}

// 查询Datahub连接源列表
func CkafkaDescribeConnectResources(SecretId string, SecretKey string, region string, page int64, page_size int64) (ConnectResources, error) {

	connect_resources := ConnectResources{}

	client, err := getCkafkaClient(SecretId, SecretKey, region)
	if err != nil {
		return connect_resources, err
	}
	offset := (page - 1) * page_size
	request := ckafka.NewDescribeConnectResourcesRequest()
	request.Offset = common.Int64Ptr(offset)
	request.Limit = common.Int64Ptr(page_size)
	response, err := client.DescribeConnectResources(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return connect_resources, err
	}

	ConnectResourceList := response.Response.Result.ConnectResourceList
	TotalCount := *response.Response.Result.TotalCount
	connect_resources.TotalCount = TotalCount

	for _, item := range ConnectResourceList {
		connect_resource := ConnectResource{}
		connect_resource.ResourceId = *item.ResourceId
		connect_resource.ResourceName = *item.ResourceName
		connect_resource.Status = *item.Status

		connect_resources.ConnectResourceList = append(connect_resources.ConnectResourceList, connect_resource)
	}

	return connect_resources, err

}

// 查询所有Datahub连接源列表
func AllCkafkaDescribeConnectResources(SecretId string, SecretKey string, region string) (ConnectResources, error) {
	var all_connect_resources = ConnectResources{}
	var page int64 = 1
	var page_size int64 = 100
	var Count int64 = 0
	var err error
	for {
		res, err := CkafkaDescribeConnectResources(SecretId, SecretKey, region, page, page_size)
		if err != nil {
			fmt.Println(err)

			return all_connect_resources, err
		}

		all_connect_resources.TotalCount = res.TotalCount
		all_connect_resources.ConnectResourceList = append(all_connect_resources.ConnectResourceList, res.ConnectResourceList...)

		Count += int64(len(res.ConnectResourceList))

		if Count == all_connect_resources.TotalCount {
			break
		}
		page += 1

	}
	return all_connect_resources, err

}

// 查询Datahub任务列表
func CkafkaDescribeDatahubTasks(SecretId string, SecretKey string, region string, page int64, page_size int64) (DatahubTasks, error) {

	datahub_tasks := DatahubTasks{}

	client, err := getCkafkaClient(SecretId, SecretKey, region)
	if err != nil {
		return datahub_tasks, err
	}
	offset := (page - 1) * page_size
	request := ckafka.NewDescribeDatahubTasksRequest()
	request.Offset = common.Int64Ptr(offset)
	request.Limit = common.Int64Ptr(page_size)
	response, err := client.DescribeDatahubTasks(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return datahub_tasks, err
	}

	TaskList := response.Response.Result.TaskList
	TotalCount := *response.Response.Result.TotalCount
	datahub_tasks.TotalCount = TotalCount

	for _, item := range TaskList {
		task := DatahubTask{}
		task.TaskId = *item.TaskId
		task.TaskName = *item.TaskName
		task.Status = *item.Status

		datahub_tasks.DatahubTaskList = append(datahub_tasks.DatahubTaskList, task)
	}

	return datahub_tasks, err

}

// 查询所有Datahub任务列表
func AllCkafkaDescribeDatahubTasks(SecretId string, SecretKey string, region string) (DatahubTasks, error) {
	var all_datahub_tasks = DatahubTasks{}
	var page int64 = 1
	var page_size int64 = 100
	var Count int64 = 0
	var err error
	for {
		res, err := CkafkaDescribeDatahubTasks(SecretId, SecretKey, region, page, page_size)
		if err != nil {
			fmt.Println(err)

			return all_datahub_tasks, err
		}

		all_datahub_tasks.TotalCount = res.TotalCount
		all_datahub_tasks.DatahubTaskList = append(all_datahub_tasks.DatahubTaskList, res.DatahubTaskList...)

		Count += int64(len(res.DatahubTaskList))

		if Count == all_datahub_tasks.TotalCount {
			break
		}
		page += 1

	}
	return all_datahub_tasks, err
}
