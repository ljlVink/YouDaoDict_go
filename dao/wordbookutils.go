package dao

import (
	"YouDaoManager/constant"
	"database/sql"
	"log"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

//status true 成功
//status false 失败
func GetWordBook() (msg string,status bool){
	log.SetPrefix("[SQL_wordbook]")
	log.Println("wordbook!")
	db, err := sql.Open(constant.DB_driver_name,constant.DB_wordbook_filepath)
	if err!=nil{
		log.Println(err)
		return  "",false
	}
    rows, err := db.Query("SELECT * FROM table_wordbook_InvalidId")
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
	cnt:=1
	result:=""
	for rows.Next() {
		err=rows.Scan(&word,&translate, &dict_type, &src_lang, &dst_lang,&wordgroup_type,&lang_type,&item_state,&sync_state,&timestamp,&example)
		if err!=nil{
			log.Println(err)
			return  "",false
		}
		result+=strconv.Itoa(cnt)+", "+word+"\n"+translate+"\n"
	}
	return result,true
}
