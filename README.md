# rds-agent
rds agent

基于小米运维开源的open-falcon，阿里云RDS监控专用agent。

  
##配置说明
配置文件请参照cfg.example.json，修改该文件名为cfg.json，然后将阿里云平台生成的accessKeyId，accessKeySecret以及需要监控的RDS实例ID填入配置文件即可。
```
{
  "debug": true,
  "hostname": "zsf-test-rds",
  "ip": "",
  "smartAPI": {
    "enabled": true,            // 这是我们自行扩展的功能，需要有对应的http接口接收。不用的可以置为false
    "url": "http://127.0.0.1/api/service/version"
  },
  "plugin": {
    "enabled": false,
    "dir": "./plugin",
    "git": "https://github.com/open-falcon/plugin.git",
    "logs": "./logs"
  },
  "heartbeat": {
    "enabled": true,
    "addr": "127.0.0.1:6030",
    "interval": 60,
    "timeout": 1000
  },
  "transfer": {
    "enabled": true,
    "addrs": [
      "127.0.0.1:8433",
      "127.0.0.1:8433"
    ],
    "interval": 60,
    "timeout": 1000
  },
  "http": {
    "enabled": true,
    "listen": ":1988",
    "backdoor": false
  },
  "mysqlMetric": {
    "MySQL_NetworkTraffic": true,
    "MySQL_QPSTPS": true,
    "MySQL_Sessions": true,
    "MySQL_InnoDBBufferRatio": true,
    "MySQL_InnoDBDataReadWriten": true,
    "MySQL_InnoDBLogRequests": true,
    "MySQL_InnoDBLogWrites": true,
    "MySQL_TempDiskTableCreates": true,
    "MySQL_MyISAMKeyBufferRatio": true,
    "MySQL_MyISAMKeyReadWrites": true,
    "MySQL_COMDML": true,
    "MySQL_RowDML": true,
    "MySQL_MemCpuUsage": true,
    "MySQL_IOPS": true,
    "MySQL_DetailedSpaceUsage": true,
    "slavestat": true
  },
  "sqlserverMetric": {
    "SQLServer_Transactions": true,
    "SQLServer_Sessions": true,
    "SQLServer_BufferHit": true,
    "SQLServer_FullScans": true,
    "SQLServer_SQLCompilations": true,
    "SQLServer_CheckPoint": true,
    "SQLServer_Logins": true,
    "SQLServer_LockTimeout": true,
    "SQLServer_Deadlock": true,
    "SQLServer_LockWaits": true,
    "SQLServer_NetworkTrffic": true,
    "SQLServer_QPS": true,
    "SQLServer_InstanceCPUUsage": true,
    "SQLServer_IOPS": true,
    "SQLServer_SpaceUsage": true
  },
  "dbType": "",        //rds_mysql or rds_sqlserver
  "accessKeyId": "",
  "accessKeySecret": "",
  "dbInstanceId": ""
}

```

