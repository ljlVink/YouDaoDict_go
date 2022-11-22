package constant

import ("database/sql")

var Version ="1.5"

var Version_code=1

var Api_connect_test="/YouDaoManager/apitest/conntest"

var Api_stop_self="/YouDaoManager/apitest/stop"

var Sys_open_ssh="/YouDaoManager/system/startssh"

var Sys_stop_ssh="/YouDaoManager/system/stopssh"

var Sys_open_adb="/YouDaoManager/system/openadb"

var Sys_stop_adb="/YouDaoManager/system/stopadb"

var Sys_reboot="/YouDaoManager/system/reboot"

var Sys_restart_ydlauncher="/YouDaoManager/system/restartlauncher"

var Tool_export_WordBook_word="/YouDaoManager/tool/ExportWordBook_word"

var Tool_export_WordBook_All="/YouDaoManager/tool/ExportWordBook_all"

var Tool_export_WordBook_sentence="/YouDaoManager/tool/ExportWordBook_sentence"


//var Tool_remove_WordBook="/YouDaoManager/tool/RemoveWordBookItem"

//var Tool_Add_WordBook="/YouDaoManager/tool/AddWordBookItem"

var Tool_apply_skins="/YouDaoManager/tool/ApplySkin"

var Tool_upload_music="/YouDaoManager/tool/MusicUpload"

var Tool_remove_music="/YouDaoManager/tool/MusicRemove"

var Tool_Get_Last_History="/YouDaoManager/tool/GetLastHistory"

var Tool_get_musicFolder="/YouDaoManager/tool/getmusicFolder"

var DB_driver_name = "sqlite3"

var DB_wordbook_filepath = "/userdisk/database/wordbook.db"

var DB_query_tables_cmd = "SELECT name FROM sqlite_master WHERE type='table' order by name"

var DB_history_filepath="/userdisk/database/history.db"

var Export_wordbooks_all=0

var Export_wordbooks_word=1

var Export_wordbooks_sentenses=2


type Translate_struct struct {
	Pos    string    `json:"pos"`
	Tran string `json:"tran"`
}
type History_struct struct {
	ScanResult    string    `json:"scan"`
	Timestamp int `json:"timestamp"`
}

type Music_remove_struct struct{
	Filename    string `json:"filename"`
}
type Dao_wordbook struct{
	Word string
	Translate string
	Dict_type int
	Src_lang string
	Dst_lang string
	Wordgroup_type int
	Lang_type int
	Item_state int
	Sync_state int
	Timestamp int
	Example sql.NullString
}
type Dao_history struct{
	Word string
	Translate string
	Dict_type int
	Src_lang string
	Dst_lang string
	Wordgroup_type int
	Timestamp int
	Uploaded int
}
