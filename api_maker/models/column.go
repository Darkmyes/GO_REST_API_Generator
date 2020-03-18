package models

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
