# YouDaoManager

有道词典笔无线管理工具，后端

### 支持系统
|设备|版本|
|-|-|
|YDP021|2.0.4以上|
其余设备由于本人没有无法测试

### 安装到开机自启

```
sh -c "$(wget https://ghproxy.com/https://raw.githubusercontent.com/ljlVink/YouDaoDict_go/main/install -O -)"
```

### 功能

成功安装后会在内网ip:6588端口启动

|功能|api|方法|备注|
|-|-|-|-|
|获取单词本|/YouDaoManager/tool/GetWordBook|GET|动态获取，实时更新|
|开启SSH|/YouDaoManager/system/startssh|GET||
|关闭SSH|/YouDaoManager/system/stopssh|GET||
|开启adb|/YouDaoManager/system/openadb|GET|仍然需要在客户端打开，默认授权auth|
|关闭adb|/YouDaoManager/system/openadb|GET||
|重启|/YouDaoManager/system/reboot|GET||
|重启launcher|/YouDaoManager/system/restartlauncher|GET|重启有道桌面|
|关闭程序|/YouDaoManager/apitest/stop|**POST**|服务停止|
|链接测试|/YouDaoManager/apitest/conntest|GET|测试是否连接成功|
|获取音乐文件夹|/YouDaoManager/tool/getmusicFolder|GET|返回格式["aaa.mp3","bbb.mp3"]|

需要带参数的
|功能|api|方法|备注|
|-|-|-|-|
|上传文件到音乐文件夹|/YouDaoManager/tool/MusicUpload|**POST**|直接上传音乐文件即可|



### 构建

```
GOARM=7 CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o YouDaoManager
```

可以在actions下载最新构建产物

### 