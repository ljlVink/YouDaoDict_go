package adb

import (
	cmd "YouDaoManager/cmd"
	"log"
	"os"
)
func OpenAdb(open bool){
	log.SetPrefix("[ADButils] ")
	if open{
		writefile("/tmp/.usb_config","usb_mtp_en\nusb_adb_en\n")
		//追加
		cmd.RunCommand("./","/bin/touch","/tmp/.adb_auth_verified")
		//重启系统服务
		cmd.RunCommand("./","/etc/init.d/S98usbdevice","restart")
	}else{
		removeFile("/tmp/.usb_config")
		writefile("/tmp/.usb_config","usb_mtp_en\n")
		removeFile("/tmp/.adb_auth_verified")
		cmd.RunCommand("./","/etc/init.d/S98usbdevice","restart","mtp")
	}
}
func writefile  (file string,content string){
	ffile,err :=os.OpenFile(file, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil{
		log.Println("Open file err =", err)
		return
	}
	defer ffile.Close()
	n, err := ffile.Write([]byte(content))
	if err != nil{
		log.Println("Write file error =", err)
		return
	}
	log.Println("WriteTo file success, n =", n)
}
func removeFile(filename string)(err error){
	e:=os.Remove(filename)
	if e!=nil{
		log.Println(e)
		return e
	}
	return
}