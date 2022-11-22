package utils

import(
	"log"
	"os"
	"YouDaoManager/cmd"
)
func Readfile(file string) string{
	log.SetPrefix("[READFILE] ")
	content, err := os.ReadFile(file)
    if err != nil {
		log.Println("file not found")
		return ""
	}
    return string(content)
}
func System_reboot(){
	cmd.RunCommand("./","/sbin/reboot","")
}
func Restart_launcher(){
	//考虑到杀死之后会自动重启 于是不需要继续启动
	cmd.RunCommand("./","/usr/bin/killall","YoudaoDictPen")
}
//挂载系统分区可读写
func Remount_system_write(){
	//mount -o remount -rw /
	cmd.RunCommand("./","/bin/mount","-o","rw,remount","/")
}