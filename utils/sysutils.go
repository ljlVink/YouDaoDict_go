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
