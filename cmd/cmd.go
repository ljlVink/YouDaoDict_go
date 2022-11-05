package cmd
import(
	"log"
	"fmt"
	"bytes"
	"errors"
	"os/exec"
)
func RunCommand(path, name string, arg ...string) (msg string, err error) {
	log.SetPrefix("[RUNCMD]")
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Dir = path
	err = cmd.Run()
	log.Println(cmd.Args)
	if err != nil {
		msg = fmt.Sprint(err) + ": " + stderr.String()
		err = errors.New(msg)
		log.Println("err", err.Error(), "cmd", cmd.Args)
	}
	log.Println(out.String())
	return
}