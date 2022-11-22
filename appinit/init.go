package daemon

import (
	"YouDaoManager/constant"
	"io/ioutil"
	"log"
	"crypto/tls"
	"net/http"
	"strconv"
	"github.com/sevlyar/go-daemon"
)
func init() {
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
	versioncode,err :=strconv.Atoi(string(body))
	if err!=nil{
		log.Println("update check error",err)
		return
	}
	log.Println("version:",versioncode)
	if versioncode!=constant.Version_code{
		log.Println("update")

	}

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