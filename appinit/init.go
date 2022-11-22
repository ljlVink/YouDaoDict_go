package daemon

import (
	"log"
	"github.com/sevlyar/go-daemon"
)
func init() {
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