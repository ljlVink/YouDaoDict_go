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
//json导出尚未实现
func ExportAllBooks(mode int,IsJson bool)(msg string,status bool){
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
			var book constant.Dao_wordbook
			for rows.Next() {
				err=rows.Scan(&book.Word,&book.Translate, &book.Dict_type, &book.Src_lang, &book.Dst_lang,&book.Wordgroup_type,&book.Lang_type,&book.Item_state,&book.Sync_state,&book.Timestamp,&book.Example)
				if err!=nil{
					log.Println(err)
					return  "",false
				}
				if mode == constant.Export_wordbooks_sentenses{
					if book.Wordgroup_type==4{//如果是单词，跳过
						continue
					}
				}
				if mode == constant.Export_wordbooks_word{
					if book.Wordgroup_type==0{//如果是句子 跳过
						continue
					}
				}
				result+=strconv.Itoa(cnt)+", "+book.Word+"\n"
				var translate_struct constant.Translate_struct
				errs := json.Unmarshal([]byte(book.Translate), &translate_struct)
				if errs != nil {
					log.Println(errs)
					return "",false
				}
				if translate_struct.Pos != ""{//有词性
					result+=translate_struct.Pos+" "
				}
				result+=translate_struct.Tran+"\n\n"//导出词性
				cnt++
			}
		}
	}
	return result,true
}

/*
func Insertwordbok(table string,book constant.Dao_wordbook) error {
	db, err := sql.Open(constant.DB_driver_name,constant.DB_wordbook_filepath)
	if err!=nil{
		return err
	}
	sql := "insert into "+table+" (word,translate,dict_type,src_lang,dst_lang,wordgroup_type,lang_type,item_state,sync_state,timestamp,example) values(?,?,?,?,?,?,?,?,?,?,?)"
	var transbase constant.Translate_struct
	transbase.Pos=""
	transbase.Tran=book.Translate
	bytes,_:=json.Marshal(&transbase)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(book.Word, string(bytes),406,"en","zh-CHS", book.Wordgroup_type,1,1,1,time.Now().UnixNano() / 1e6,book.Example)
	return err
}*/