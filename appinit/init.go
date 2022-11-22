package daemon

import (
	"YouDaoManager/cmd"
	"YouDaoManager/constant"
	"crypto/tls"
	"io"
	"log"
	"os"
	"net/http"
	"strconv"
	"strings"
    "time"

	"github.com/sevlyar/go-daemon"
)
func init() {
	go updatecheck()
	go forkDeamon()
	
}
func forkDeamon(){
	log.SetPrefix("[DAEMON]")
	cntxt := &daemon.Context{
		PidFileName: "YoudaoMgr",
		PidFilePerm: 0644,
		LogFileName: "YoudaoMgr.log",
		LogFilePerm: 0640,
		WorkDir:     "/userdata",
		Umask:       027,
		Args:        []string{"[YoudaoMgr]"},
	}
	d, err := cntxt.Reborn()
	if err != nil {
		log.Println("Unable to run: ", err)
		log.Println("maybe it is windows env")	
	}
	log.Println("Daemon started")
	if d != nil {
		return
	}
	defer cntxt.Release()

}
func updatecheck(){
	log.SetPrefix("[UPDATE]")
	log.Println("update check! Wait 30s to Intetrnet")
	time.Sleep(30 * time.Second)
	tr := &http.Transport{
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
	resp, err := client.Get("http://ghproxy.com/https://raw.githubusercontent.com/ljlVink/YouDaoDict_go/main/update_tag") // url
	if err != nil {
		log.Println("update check error",err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("update check error",err)
		return
	}
	var versioncode int
	if strings.Contains(string(body),"\n"){
		versioncode,_ =strconv.Atoi(string(body[:1]))//去掉回车!
	}else {
		versioncode,_ =strconv.Atoi(string(body))
	}
	if err!=nil{
		log.Println("update check error",err)
		return
	}
	log.Println("version:",versioncode)
	if versioncode>constant.Version_code{
		log.Println("need to update")
		res2,err:=client.Get("https://ghproxy.com/https://raw.githubusercontent.com/ljlVink/YouDaoDict_go/main/install")
		if err!=nil{
			log.Println("Download script fail!")
			return
		}
		defer res2.Body.Close()
		body2, err := io.ReadAll(res2.Body)
		if err != nil {
			log.Println("Download script fail",err)
			return
		}
		script:=string(body2)
		log.Print(script)
		writefile("/tmp/update",script)
		cmd.RunCommand("/tmp","/bin/bash","./update")
	}else{
		log.Println("latest version,good!")
	}
}
func writefile(file string,content string){
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
