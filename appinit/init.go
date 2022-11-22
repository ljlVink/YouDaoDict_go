package daemon

import (
	"YouDaoManager/cmd"
	"YouDaoManager/constant"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sevlyar/go-daemon"
)
func init() {
	go updatecheck()
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
	log.Println("update check!")
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
	body, err := ioutil.ReadAll(resp.Body)
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
	if versioncode!=constant.Version_code{
		log.Println("need to update")
		cmd.RunCommand("./","/bin/sh","-c","$(wget https://ghproxy.com/https://raw.githubusercontent.com/ljlVink/YouDaoDict_go/main/install -O -)")
	}

}