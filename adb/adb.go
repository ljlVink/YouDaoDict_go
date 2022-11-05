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
		cmd.RunCommand("./","/bin/bash","touch","/tmp/.adb_auth_verified")
	}else{
		cmd.RunCommand("./","/bin/bash","echo","usb_mtp_on",">","/tmp/.usb_config")
		cmd.RunCommand("./","/bin/bash","rm","/tmp/.adb_auth_verified")
	}
}