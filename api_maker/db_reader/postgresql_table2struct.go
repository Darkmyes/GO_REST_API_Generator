package db_reader

import (
	"strings"

	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgreSQLColumnDescription struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

type PostgreSQLTable struct {
	Name    string
	Columns []PostgreSQLColumnDescription
}

func listTablesPostgreSQL(db *gorm.DB) []string {
	var ret []string
	err := db.Raw("SHOW TABLES").Pluck("Tables_in_PostgreSQL", &ret).Error
	if err != nil {
		panic(err)
	}
	return ret
}

func describeTablePostgreSQL(db *gorm.DB, name string) []PostgreSQLColumnDescription {
	var columnList []PostgreSQLColumnDescription

	rows, err := db.Raw("DESCRIBE " + name).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var Field, Type, Null, Key, Default, Extra string
		rows.Scan(&Field, &Type, &Null, &Key, &Default, &Extra)

		tableDescr := PostgreSQLColumnDescription{
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

func PostgreSQLTableToJSON(table PostgreSQLTable) models.Object {
	var cols []models.Column

	for _, PostgreSQLCol := range table.Columns {
		col := models.Column{
			Name:        PostgreSQLCol.Field,
			Col_Name:    PostgreSQLCol.Field,
			Col_Type:    strings.ToUpper(PostgreSQLCol.Type),
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
		if PostgreSQLCol.Null == "YES" {
			col.Is_Null = true
		} else {
			col.Is_Null = false
		}
		if PostgreSQLCol.Key == "PRI" {
			col.Primary_Key = true
		}
		if PostgreSQLCol.Key == "PRI" {
			col.Is_Null = false
		}
		if PostgreSQLCol.Key == "PRI" {
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

func readObjectsPostgreSQL(path, bd, user, pass, port string) []models.Object {
	if port == "" {
		port = "3306"
	}
	db, err := gorm.Open("postgres", "host="+path+" port="+port+" user="+user+" dbname="+bd+" password="+pass)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	tables := listTablesPostgreSQL(db)

	var objs []models.Object

	for _, tab := range tables {
		cols := describeTablePostgreSQL(db, tab)
		table := PostgreSQLTable{
			Name:    tab,
			Columns: cols,
		}
		obj := PostgreSQLTableToJSON(table)
		objs = append(objs, obj)
	}

	return objs
}
