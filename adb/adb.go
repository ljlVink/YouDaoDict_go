package adb

import (
	cmd "YouDaoManager/cmd"
)
func OpenAdb(open bool){
	if open{
		cmd.RunCommand("./","/bin/bash","echo","usb_mtp_on",">","/tmp/.usb_config")
		//追加
		cmd.RunCommand("./","/bin/bash","echo","usb_adb_on",">>","/tmp/.usb_config")
		//有道是不是打错了...反正都写进去吧，在2.0.4至少这样的
		cmd.RunCommand("./","/bin/bash","echo","usb_adb_en",">>","/tmp/.usb_config")
		//放权
		cmd.RunCommand("./","/bin/touch","/tmp/.adb_auth_verified")
		//?不太好使，还是要手动确定
		cmd.RunCommand("./","/bin/bash","mkdir","/dev/usb-ffs/adb","-m","0770")
		//mount -o uid=2000,gid=2000 -t functionfs adb /dev/usb-ffs/adb
		cmd.RunCommand("./","/bin/mount","-o","uid=2000,gid=2000","-t","functionfs","adb","/dev/usb-ffs/adb")
		//start-stop-daemon --start --quiet --background --exec /usr/bin/adbd
		cmd.RunCommand("./","/sbin/start-stop-daemon","--start","--quiet","--background","--exec","/usr/bin/adbd")
	}else{
		cmd.RunCommand("./","/bin/bash","echo","usb_mtp_on",">","/tmp/.usb_config")
		cmd.RunCommand("./","/bin/bash","rm","/tmp/.adb_auth_verified")
	}
}
