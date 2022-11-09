package main

import (
	adbutils "YouDaoManager/adb"
	constant "YouDaoManager/constant"
	_ "YouDaoManager/daemon"
	wordbookutils "YouDaoManager/dao"
	musicfolder "YouDaoManager/musicfolder"
	sshutils "YouDaoManager/ssh"
	systemutils "YouDaoManager/utils"
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)
func applog(logg string){
	log.SetPrefix("[APP] ")
	log.Println(logg)
}
func main(){
	applog("YouDaoManager Version:"+constant.Version)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	//链接测试
	r.GET(constant.Api_connect_test, func(context *gin.Context) {
		context.String(200, "connect OK!")
	})
	//服务终止信号
	r.POST(constant.Api_stop_self,func(context *gin.Context) {
		context.String(200,"done")
		os.Exit(0)
	})
	//获取 单词本
	r.GET(constant.Tool_get_WordBook,func (context *gin.Context)  {
		bookcontent,status:=wordbookutils.GetWordBook()
		if status{
			context.String(200,bookcontent)
		}else{
			context.String(400,"ERROR")
		}
	})
	//开启 ssh
	r.GET(constant.Sys_open_ssh,func(context *gin.Context) {
		if sshutils.OpenSSH(true){
			context.String(200,"OK")
			
		}else {
			context.String(400,"ERROR")
		}
	})
	//关闭 ssh
	r.GET(constant.Sys_stop_ssh,func (context *gin.Context)  {
		if sshutils.OpenSSH(false){
			context.String(200,"OK,default user=root,password=CherryYoudao")
		}else{
			context.String(400,"ERROR")
		}
	})
	//开启adb 自动授权
	r.GET(constant.Sys_open_adb,func(context *gin.Context) {
		adbutils.OpenAdb(true)
		context.String(200,"OK")
	})
	//关闭adb 自动授权
	r.GET(constant.Sys_stop_adb,func(context *gin.Context) {
		adbutils.OpenAdb(false)
		context.String(200,"OK")
	})
	//设备重启
	r.GET(constant.Sys_reboot,func (context *gin.Context)  {
		systemutils.System_reboot()
		context.String(200,"OK")
	})
	//重启 launcher
	r.GET(constant.Sys_restart_ydlauncher,func(context *gin.Context) {
		systemutils.Restart_launcher()
		context.String(200,"OK")
	})

	//音乐上传
	r.POST(constant.Tool_upload_music, func(c *gin.Context) {
        file, _ := c.FormFile("file")
        applog("upload music:"+file.Filename)
        dst := "/userdisk/Music/" + file.Filename
        c.SaveUploadedFile(file,dst)
        c.String(200, fmt.Sprintf("'%s' uploaded", file.Filename))
    })
	//获取音乐文件夹内文件
	r.GET(constant.Tool_get_musicFolder,func(context *gin.Context) {
		result:=musicfolder.GetFolder()
		context.String(200,result)
	})
	r.POST(constant.Tool_remove_music,func(context *gin.Context) {
		json := constant.Music_remove_struct{}
		context.BindJSON(&json)
		filename:=json.Filename
		applog(filename)
		err:=musicfolder.RemoveMusic(string(filename))
		if err!=nil{
			context.String(400,"ERROR")
		}else{
			context.String(200,"OK")
		}
	})
	r.Run(":6588") 
}