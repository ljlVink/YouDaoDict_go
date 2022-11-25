package skin

import(
	"YouDaoManager/cmd"
	"YouDaoManager/utils"
	"os"
	"log"
)
func ApplySkin(filepath string){
	log.SetPrefix("[Skinutils] ")
 	utils.Remount_system_write()
	cmd.RunCommand("./","/usr/bin/unzip","-o",filepath,"-d","/oem/YoudaoDictPen/output/images")
	utils.Restart_launcher()
	removeFile(filepath)
}
func removeFile(filename string)(err error){

	e:=os.Remove(filename)
	if e!=nil{
		log.Println(e)
		return e
	}
	return
}