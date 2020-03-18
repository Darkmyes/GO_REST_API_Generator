package models

type DbData struct {
	Db_Type    string `json:"db_type"`
	Db_Url     string `json:"db_url"`
	Db_Port    string `json:"db_port"`
	Db_Name    string `json:"db_name"`
	Db_User    string `json:"db_user"`
	Db_Pass    string `json:"db_pass"`
	DB_Charset string `json:"db_charset"`
	DB_Collate string `json:"db_collate"`
}
