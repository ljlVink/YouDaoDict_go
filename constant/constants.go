package constant

var Version ="1.1"

var Api_connect_test="/YouDaoManager/apitest/conntest"

var Api_stop_self="/YouDaoManager/apitest/stop"

var Sys_open_ssh="/YouDaoManager/system/startssh"

var Sys_stop_ssh="/YouDaoManager/system/stopssh"

var Sys_open_adb="/YouDaoManager/system/openadb"

var Sys_stop_adb="/YouDaoManager/system/stopadb"

var Sys_reboot="/YouDaoManager/system/reboot"

var Sys_restart_ydlauncher="/YouDaoManager/system/restartlauncher"

var Tool_get_WordBook="/YouDaoManager/tool/GetWordBook"

var Tool_upload_music="/YouDaoManager/tool/MusicUpload"

var DB_driver_name = "sqlite3"

var DB_wordbook_filepath = "/userdisk/database/wordbook.db"

var DB_query_tables_cmd = "SELECT name FROM sqlite_master WHERE type='table' order by name"


type Translate_struct struct {
	Pos    string    `json:"pos"`
	Tran string `json:"tran"`
}

