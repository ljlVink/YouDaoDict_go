package musicfolder

import(
	"encoding/json"
	"path/filepath"
	"log"
	"os"
)

func GetFolder() string{
	folder:=make([]interface{},0)
	root := "/userdisk/Music/"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Name()!= "Music"{
			folder=append(folder,info.Name())
		}
        return nil
    })
	if err!=nil{
		log.Println(err)
	}
	bytes,_:=json.Marshal(folder)
	return string(bytes)
}
func RemoveMusic(filename string)(err error){
	abspath:="/userdisk/Music/"+filename
	log.Println("abs:",abspath)
	e:=os.Remove(abspath)
	if e!=nil{
		log.Println(e)
		return e
	}
	return
}