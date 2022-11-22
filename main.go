package main

import (
	adbutils "YouDaoManager/adb"
	constant "YouDaoManager/constant"
	_ "YouDaoManager/appinit"
	Daoutils "YouDaoManager/dao"
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
	r.GET(constant.Tool_export_WordBook_All,func (context *gin.Context)  {
		bookcontent,status:=Daoutils.ExportAllBooks(constant.Export_wordbooks_all,false)
		if status{
			context.String(200,bookcontent)
		}else{
			context.String(400,"ERROR")
		}
	})
	//获取 单词本单词部分
	r.GET(constant.Tool_export_WordBook_word,func (context *gin.Context)  {
		bookcontent,status:=Daoutils.ExportAllBooks(constant.Export_wordbooks_word,false)
		if status{
			context.String(200,bookcontent)
		}else{
			context.String(400,"ERROR")
		}
	})
	r.GET(constant.Tool_export_WordBook_sentence,func (context *gin.Context)  {
		bookcontent,status:=Daoutils.ExportAllBooks(constant.Export_wordbooks_sentenses,false)
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
			context.String(200,"OK")
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
	//删除音乐文件夹的文件
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
	//获取最先的历史记录 
	r.GET(constant.GetLastHistory,func(context *gin.Context) {
		result,stat:=Daoutils.ExportLastHistory(0,false)
		if !stat{
			context.String(400,"ERROR")
		}else{
			context.String(200,result)
		}
	})
	r.Run(":6588") 
}