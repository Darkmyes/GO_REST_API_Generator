package models

type Project struct {
	Id              int32    `json:"id"`
	Name            string   `json:"name"`
	Desc            string   `json:"desc"`
	Db              string   `json:"db"`
	Db_Data         DbData   `json:"db_data"`
	Lang            string   `json:"lang"`
	Folder_Path     string   `json:"folder_path"`
	Init_Date       string   `json:"init_date"`
	Upd_Date        string   `json:"upd_date"`
	GenerateSQL     bool     `json:"gen_sql"`
	GenerateAndroid bool     `json:"gen_android"`
	AuthMode        string   `json:"auth_mode"`
	Objects         []Object `json:"objects"`
}
