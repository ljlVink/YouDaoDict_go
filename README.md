# YouDaoManager

有道词典笔无线管理工具，后端

### 支持系统
|设备|版本|
|-|-|
|YDP021|2.0.4以上|
其余设备由于本人没有无法测试

### 安装到开机自启

```
目前还不完善，建议吧主程序挪到/userdata中，拷贝项目`S75YouDaoMgr`到/etc/init.d中
```

### 功能
|功能|api|方法|备注|
|-|-|-|-|
|获取单词本|/YouDaoManager/tool/GetWordBook|GET|动态获取，实时更新|
|开启SSH|/YouDaoManager/system/startssh|GET||
|关闭SSH|/YouDaoManager/system/stopssh|GET||
|开启adb|/YouDaoManager/system/openadb|GET|仍然需要在客户端打开，默认授权auth|
|关闭adb|/YouDaoManager/system/openadb|GET||
|重启|/YouDaoManager/system/reboot|GET||
|上传文件到音乐文件夹|/YouDaoManager/tool/MusicUpload|**POST**|
|关闭程序|/YouDaoManager/apitest/stop|**POST**|服务停止|
|链接测试|/YouDaoManager/apitest/conntest|GET|测试是否连接成功|


### 构建

```
GOARM=7 CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o YouDaoManager
```

可以在actions下载最新构建产物

### 