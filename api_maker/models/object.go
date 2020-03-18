package models

type Object struct {
	Proyect_Id string   `json:"proyect_id"`
	Name       string   `json:"name"`
	Tbl_Name   string   `json:"tbl_name"`
	Is_Login   bool     `json:"is_login"`
	LoginId    string   `json:"login_id"`
	LoginPass  string   `json:"login_pass"`
	Columns    []Column `json:"columns"`
}
