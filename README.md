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
成功安装后会在内网ip:6588端口启动

### 功能
api测试功能
|功能|api|方法|备注|返回|status code|
|-|-|-|-|-|-|
|关闭程序|/YouDaoManager/apitest/stop|**POST**|服务停止|无返回或者done|200|
|链接测试|/YouDaoManager/apitest/conntest|GET|测试是否连接成功|connect OK!|200|



系统功能

|功能|api|方法|返回|status code|
|-|-|-|-|-|
|开启SSH|/YouDaoManager/system/startssh|GET|成功OK,错误ERROR|200,400|
|关闭SSH|/YouDaoManager/system/stopssh|GET|成功OK,错误ERROR|200,400|
|开启adb|/YouDaoManager/system/openadb|GET|OK|200|
|关闭adb|/YouDaoManager/system/openadb|GET|OK|200|
|重启|/YouDaoManager/system/reboot|GET|OK或无|200或无响应|
|重启launcher|/YouDaoManager/system/restartlauncher|GET|OK|200|


单词本相关
|功能|api|方法|返回|status code|
|-|-|-|-|-|
|导出全部单词本|/YouDaoManager/tool/ExportWordBook_all|GET|类似词典笔导出单词本|200|
|导出单词|/YouDaoManager/tool/ExportWordBook_word|GET|仅导出单词|200|
|导出句子|/YouDaoManager/tool/ExportWordBook_sentence|GET|仅导出句子|200|

历史记录相关

|功能|api|方法|返回|status code|
|-|-|-|-|-|
|获取历史记录|/YouDaoManager/tool/GetLastHistory|GET|获取最近第一条历史记录|200|


皮肤相关

|功能|api|方法|返回|status code|
|-|-|-|-|-|
|更换皮肤|/YouDaoManager/tool/ApplySkin|**POST**|上传images文件夹的压缩包即可|200|

关于压缩包注意事项

1.词典笔的一些图片资源在`/oem/YoudaoDictPen/output/images`文件夹

打包时类似zip的文件格式

```
images.zip
│  history_clear_close.png
│  home-audioplayer-bg.png
│  home-audioplayer.png
│  home-dict-bg.png
│  home-dict.png
│  home-fav-bg.png
```

而不是嵌套文件夹，错误示例

```
images.zip
├─images
|   history_clear_close.png
|   home-audioplayer-bg.png
|   home-audioplayer.png
```



音乐相关
|功能|api|方法|返回|status code|
|-|-|-|-|-|
|获取音乐文件夹|/YouDaoManager/tool/getmusicFolder|GET|返回格式["aaa.mp3","bbb.mp3"]|200|
|上传文件到音乐文件夹|/YouDaoManager/tool/MusicUpload|**POST**|直接上传音乐文件即可|200|
|移除文件|/YouDaoManager/tool/MusicRemove|**POST**|json格式:{"filename":"xxx"}|成功200,错误400|


### 编译

```
GOARM=7 CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o YouDaoManager
```

可以在actions下载最新构建产物

### 