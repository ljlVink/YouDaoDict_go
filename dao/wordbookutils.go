package dao

import (
	"YouDaoManager/constant"
	"strings"
	"encoding/json"
	"database/sql"
	"log"
	"strconv"
	_ "github.com/logoove/sqlite"
)

//status true 成功
//status false 失败
func GetWordBook() (msg string,status bool){
	log.SetPrefix("[DAO_wordbook]")
	db, err := sql.Open(constant.DB_driver_name,constant.DB_wordbook_filepath)
	if err!=nil{
		log.Println(err)
		return  "",false
	}
	tables,err:=db.Query(constant.DB_query_tables_cmd)
	if err!=nil{
		log.Println(err)
		return  "",false
	}
	var tableName string
	result:=""
	cnt:=1
	for tables.Next(){
		err=tables.Scan(&tableName)
		if err!=nil{
			log.Println(err)
			return "",false
		}
		if strings.Contains(tableName,"table_wordbook"){
			rows, err := db.Query("SELECT * FROM "+tableName)
			if err!=nil{
				log.Println(err)
				return  "",false
			}
			var word string
			var translate string
			var dict_type int
			var src_lang string
			var dst_lang string
			var wordgroup_type int
			var lang_type int
			var item_state int
			var sync_state int
			var timestamp int
			var example sql.NullString
			for rows.Next() {
				err=rows.Scan(&word,&translate, &dict_type, &src_lang, &dst_lang,&wordgroup_type,&lang_type,&item_state,&sync_state,&timestamp,&example)
				if err!=nil{
					log.Println(err)
					return  "",false
				}
				result+=strconv.Itoa(cnt)+", "+word+"\n"
				var translate_struct constant.Translate_struct
				translate_struct.Pos=""
				translate_struct.Tran=""
		
				errs := json.Unmarshal([]byte(translate), &translate_struct)
				if errs != nil {
					log.Println(errs)
					return "",false
				}
				result+=translate_struct.Tran+"\n\n"
				cnt++
			}
		}
	}
	return result,true
}
