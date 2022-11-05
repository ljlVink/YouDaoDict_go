package ssh
import(
	cmd "YouDaoManager/cmd"
)
func OpenSSH(open bool)bool{
	args:="start"
	if !open{
		args="stop"
	}
	_,err:=cmd.RunCommand("./","/bin/bash","sshd_sevice",args)
	if err!=nil{
		return false
	}else{
		return true
	}
}
