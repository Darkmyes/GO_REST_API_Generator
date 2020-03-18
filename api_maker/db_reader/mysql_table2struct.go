package db_reader

import (
	"strings"

	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySQLColumnDescription struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

type MySQLTable struct {
	Name    string
	Columns []MySQLColumnDescription
}

func listTablesMySQL(db *gorm.DB) []string {
	var ret []string
	err := db.Raw("SHOW TABLES").Pluck("Tables_in_mysql", &ret).Error
	if err != nil {
		panic(err)
	}
	return ret
}

func describeTableMySQL(db *gorm.DB, name string) []MySQLColumnDescription {
	var columnList []MySQLColumnDescription

	rows, err := db.Raw("DESCRIBE " + name).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var Field, Type, Null, Key, Default, Extra string
		rows.Scan(&Field, &Type, &Null, &Key, &Default, &Extra)

		tableDescr := MySQLColumnDescription{
			Field:   Field,
			Type:    Type,
			Null:    Null,
			Key:     Key,
			Default: Default,
			Extra:   Extra,
		}

		columnList = append(columnList, tableDescr)
	}
	return columnList
}

func MySQLTableToJSON(table MySQLTable) models.Object {
	var cols []models.Column

	for _, MySQLCol := range table.Columns {
		col := models.Column{
			Name:        MySQLCol.Field,
			Col_Name:    MySQLCol.Field,
			Col_Type:    strings.ToUpper(MySQLCol.Type),
			Col_Lenght:  0,
			Enum_List:   nil,
			Is_Null:     false,
			Is_Unique:   false,
			Is_Index:    false,
			Primary_Key: false,
			Foreign_Key: false,
			Tbl_Ref:     "",
			Col_Refs:    nil,
			On_Delete:   "",
			On_Update:   "",
		}
		if MySQLCol.Null == "YES" {
			col.Is_Null = true
		} else {
			col.Is_Null = false
		}
		if MySQLCol.Key == "PRI" {
			col.Primary_Key = true
		}
		if MySQLCol.Key == "PRI" {
			col.Is_Null = false
		}
		if MySQLCol.Key == "PRI" {
			col.Is_Null = false
		}
		cols = append(cols, col)
	}

	obj := models.Object{
		Proyect_Id: "0",
		Name:       table.Name,
		Tbl_Name:   "",
		Is_Login:   false,
		LoginId:    "",
		LoginPass:  "",
		Columns:    cols,
	}
	return obj
}

func readObjectsMySQL(path, bd, user, pass, port string) []models.Object {
	if port == " " {
		port = "3306"
	}
	//(db_data.Db_Url, db_data.Db_Name, db_data.Db_User, db_data.Db_Pass, db_data.Db_Port)
	db, err := gorm.Open("mysql", user+":"+pass+"@("+path+":"+port+")/"+bd+"?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", user+":"+pass+"@("+path+")/"+bd+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	tables := listTablesMySQL(db)

	var objs []models.Object

	for _, tab := range tables {
		cols := describeTableMySQL(db, tab)
		table := MySQLTable{
			Name:    tab,
			Columns: cols,
		}
		obj := MySQLTableToJSON(table)
		objs = append(objs, obj)
	}

	return objs
}

func readObjectsMySQL2(db_data models.DbData) []models.Object {
	if db_data.Db_Port == " " {
		db_data.Db_Port = "3306"
	}
	//(db_data.Db_Url, db_data.Db_Name, db_data.Db_User, db_data.Db_Pass, db_data.Db_Port)
	db, err := gorm.Open("mysql", db_data.Db_User+":"+db_data.Db_Pass+"@("+db_data.Db_Url+":"+db_data.Db_Port+")/"+db_data.Db_Name+"?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", user+":"+pass+"@("+path+")/"+bd+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	tables := listTablesMySQL(db)

	var objs []models.Object

	for _, tab := range tables {
		cols := describeTableMySQL(db, tab)
		table := MySQLTable{
			Name:    tab,
			Columns: cols,
		}
		obj := MySQLTableToJSON(table)
		objs = append(objs, obj)
	}

	return objs
}
