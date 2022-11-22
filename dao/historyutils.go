package dao

import (
	"YouDaoManager/constant"
	"encoding/json"
	"database/sql"
	"log"
	_ "github.com/logoove/sqlite"
)	

//status true 成功
//status false 失败
func ExportLastHistory(mode int,IsJson bool)(msg string,status bool){
	log.SetPrefix("[DAO_history]")
	db, err := sql.Open(constant.DB_driver_name,constant.DB_history_filepath)
	if err!=nil{
		log.Println(err)
		return  "",false
	}
	rows, err := db.Query("SELECT * FROM table_history order by timestamp desc")
	if err!=nil{
		log.Println(err)
	}
	var book constant.Dao_history
	rows.Next()
	err=rows.Scan(&book.Word,&book.Translate, &book.Dict_type, &book.Src_lang, &book.Dst_lang,&book.Wordgroup_type,&book.Timestamp,&book.Uploaded)
	if err!=nil{
		log.Println(err)
	}
	defer rows.Close()
	defer db.Close()//一定要关闭db
	if (IsJson){
		var scanResult constant.History_struct
		scanResult.ScanResult=book.Word
		scanResult.Timestamp=book.Timestamp
		json,_:=json.Marshal(&scanResult)
		return string(json),true
	}else{
		return book.Word,true
	}
}