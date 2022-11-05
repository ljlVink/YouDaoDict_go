package main

import (
	_ "YouDaoManager/daemon"
	adbutils "YouDaoManager/adb"
	constant "YouDaoManager/constant"
	sshutils "YouDaoManager/ssh"
	wordbookutils "YouDaoManager/dao"
	"YouDaoManager/utils"
	"log"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	log.Println("hihi")
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
			context.String(200,"ERROR")
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
		}
		context.String(400,"ERROR")
	})
	//开启adb 自动授权
	r.GET(constant.Sys_open_adb,func(context *gin.Context) {
		adbutils.OpenAdb(true)
		context.String(200,"OK")
	})
	r.GET(constant.Sys_stop_adb,func(context *gin.Context) {
		adbutils.OpenAdb(false)
		context.String(200,"OK")
	})
	r.GET(constant.Sys_reboot,func (context *gin.Context)  {
		utils.System_reboot()
	})
	r.POST(constant.Tool_upload_music, func(c *gin.Context) {
        file, _ := c.FormFile("file")
        log.Println(file.Filename)
        dst := "/userdisk/Music/" + file.Filename
        c.SaveUploadedFile(file,dst)
        c.String(200, fmt.Sprintf("'%s' uploaded", file.Filename))
    })
	r.Run(":6588") 
}