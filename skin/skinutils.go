package skin

import(
	"YouDaoManager/cmd"
	"YouDaoManager/utils"
)
func ApplySkin(filepath string){
 	utils.Remount_system_write()
	cmd.RunCommand("./","/usr/bin/unzip","-o",filepath,"-d","/oem/YoudaoDictPen/output/images")
	utils.Restart_launcher()
}