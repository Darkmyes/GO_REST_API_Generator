package proyect_maker

type DbData struct {
	Db_Url     string `json:"db_url"`
	Db_Port    string `json:"db_port"`
	Db_Name    string `json:"db_name"`
	Db_User    string `json:"db_user"`
	Db_Pass    string `json:"db_pass"`
	DB_Charset string `json:"db_charset"`
	DB_Collate string `json:"db_collate"`
}

type Proyect struct {
	Id              int32  `json:"id"`
	Name            string `json:"name"`
	Desc            string `json:"desc"`
	Db              string `json:"db"`
	Db_Data         DbData `json:"db_data"`
	Lang            string `json:"lang"`
	Folder_Path     string `json:"folder_path"`
	Init_Date       string `json:"init_date"`
	Upd_Date        string `json:"upd_date"`
	GenerateSQL     bool   `json:"gen_sql"`
	GenerateAndroid bool   `json:"gen_android"`
	AuthMode        string `json:"auth_mode"`
}

type Column struct {
	Name        string   `json:"name"`
	Col_Name    string   `json:"col_name"`
	Col_Type    string   `json:"col_type"`
	Col_Lenght  int      `json:"col_lenght"`
	Enum_List   []string `json:"enum_list"`
	Is_Null     bool     `json:"is_null"`
	Is_Unique   bool     `json:"is_unique"`
	Is_Index    bool     `json:"is_index"`
	Primary_Key bool     `json:"primary_key"`
	Foreign_Key bool     `json:"foreign_key"`
	Tbl_Ref     string   `json:"tbl_ref"`
	Col_Refs    []string `json:"col_refs"`
	On_Delete   string   `json:"on_delete"`
	On_Update   string   `json:"on_update"`
}

type Object struct {
	Proyect_Id string   `json:"proyect_id"`
	Name       string   `json:"name"`
	Tbl_Name   string   `json:"tbl_name"`
	Is_Login   bool     `json:"is_login"`
	LoginId    string   `json:"login_id"`
	LoginPass  string   `json:"login_pass"`
	Columns    []Column `json:"columns"`
}
