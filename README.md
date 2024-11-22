# 腾讯云消息队列CKafka版 连接器 状态 Exporter

## 项目背景
连接器为 CKafka 的重要功能模块之一，旨在搭建 CKafka 的上下游数据流转链路。
因业务需要，把数据库增量 使用 连接器流入到Ckafka ,提供增量数据订阅和消费。
但腾讯云没有现成的告警指标，状态异常时不能及时发现。所以目前只能通过腾讯云API获取连接器状态做告警通知。
- API文档: 
- https://console.cloud.tencent.com/api/explorer?Product=ckafka&Version=2019-08-19&Action=DescribeConnectResource
- https://console.cloud.tencent.com/api/explorer?Product=ckafka&Version=2019-08-19&Action=DescribeDatahubTasks

## 克隆到本地还后初始化
```
cd  ./datahub_exporter 
go mod tidy
```
## running
```
go run main.go
```
## Building


### Build

    Mac下编译Linux平台的64位可执行程序：
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o datahub_exporter main.go
    Linux下编译Linux平台的64位可执行程序
    go build -o datahub_exporter main.go

### Running

        ./datahub_exporter

可选参数

*   --web.listen-address  监听地址和端口,默认为  0.0.0.0:48080
*   --conf   配置文件,默认位置为 当前用户家目录下的 .datahub_exporter.conf&#x20;

自定义传参示例:

        ./datahub_exporter  --web.listen-address=:8080  --conf=/tmp/.datahub_exporter.conf

### 最小权限策略
    QcloudCkafkaReadOnlyAccess
    消息队列 CKafka 版（Ckafka）只读访问策略

配置文件内容格式
```
[common]
secret_id = xxx
secret_key = xxx
region = ap-guangzhou`
```


### &#x20;爬取路径

*   连接列表 (datahub-connect)    /api/v1/ckafka/datahub_tasks/metrics

```
curl http://127.0.0.1:48080/api/v1/ckafka/datahub_tasks/metrics
datahub_task_error_count{region="ap-guangzhou"} 0
datahub_task_status{region="ap-guangzhou",TaskId="task-xxx",TaskName="test-xxx"} 1
datahub_task_status{region="ap-guangzhou",TaskId="task-xxx",TaskName="test-xxx"} 1
datahub_task_status{region="ap-guangzhou",TaskId="task-xxx",TaskName="prod-xxx"} 1
datahub_task_status{region="ap-guangzhou",TaskId="task-xxx",TaskName="prod-xxx"} 1
datahub_task_status{region="ap-guangzhou",TaskId="task-xxx",TaskName="prod-xxx"} 1

```



*   任务列表 (datahub-task)    /api/v1/ckafka/connect_resources/metrics

```
curl http://127.0.0.1:48080/api/v1/ckafka/connect_resources/metrics
connect_resource_error_count{region="ap-guangzhou"} 0
connect_resource_status{region="ap-guangzhou",ResourceId="resource-xxx",ResourceName="test-xxx"} 1
connect_resource_status{region="ap-guangzhou",ResourceId="resource-xxx",ResourceName="xxxx"} 1
connect_resource_status{region="ap-guangzhou",ResourceId="resource-xxx",ResourceName="xxx"} 1
connect_resource_status{region="ap-guangzhou",ResourceId="resource-xxx",ResourceName="prod-xxx"} 1
connect_resource_status{region="ap-guangzhou",ResourceId="resource-xxx",ResourceName="prod-xxx"} 1

```

### prometheus爬取job配置
```
  - job_name: datahub-exporter-datahub-task
    honor_labels: true
    honor_timestamps: true
    scrape_interval: 1m
    metrics_path: /api/v1/ckafka/datahub_tasks/metrics
    scheme: http
    static_configs:
    - targets:
        - xx.xx.xx.xx:48080
        labels:
        instance: datahub-exporter-datahub-task



  - job_name: datahub-exporter-connect-resources
    honor_labels: true
    honor_timestamps: true
    scrape_interval: 1m
    metrics_path: /api/v1/ckafka/connect_resources/metrics
    scheme: http
    static_configs:
    - targets:
    - xx.xx.xx.xx:48080
    labels:
        instance: datahub-exporter-connect-resources
```

### prometheus告警rule配置
```
groups:
  - name: datahub-exporter
    rules:
      - alert: connect_resource_status
        expr: connect_resource_status{job="datahub-exporter-connect-resources"} !=1
        for: 2m
        annotations:
          description: "\nResourceId: {{ $labels.ResourceId }}\nResourceName: {{ $labels.ResourceName }}  "
          summary: ckafka连接器异常
      - alert: datahub_task_status
        expr: datahub_task_status{job="datahub-exporter-datahub-task"}  !=1
        for: 2m
        annotations:
          description: |-
              TaskId: {{ $labels.TaskId }}
              TaskName: {{ $labels.TaskName }}
          summary: ckafka连接器任务异常
      - alert: datahub_task_error_count
        expr: datahub_task_error_count{job="datahub-exporter-datahub-task"} > 10
        for: 2m
        annotations:
          description: datahub_task爬取异常
          summary: datahub_task爬取异常
      - alert: connect_resource_error_count
        expr: connect_resource_error_count{job="datahub-exporter-connect-resources"}  > 10
        for: 2m
        annotations:
          description: connect_resource 爬取异常
          summary: connect_resource 爬取异常
```