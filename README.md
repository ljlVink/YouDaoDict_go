# YouDaoDict_go

只是突发奇想 想到了golang一份代码可以编译出在不同系统和cpu架构运行的二进制文件

set_arm_linux_env.bat 是设置环境为arm64编译环境，即在首次编译前运行下这个bat(pwsh自己研究)

欢迎PR

编译
```shell
go build xxx.go
```
|项目|作用|可执行文件|用法|
|--|--|--|--|
YdFileServer|文件服务器,默认/|YdServer|直接运行