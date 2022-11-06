package ssh
import(
	cmd "YouDaoManager/cmd"
	"log"
)
func OpenSSH(open bool)bool{
	log.SetPrefix("[SSH] ")
	args:="start"
	if !open{
		args="stop"
	}
	_,err:=cmd.RunCommand("./","/bin/bash","sshd_sevice",args)
	if err!=nil{
		log.Println("SSH open failed!")
		log.Println(err)
		return false
	}else{
		log.Println("SSH open success!")
		return true
	}
}
