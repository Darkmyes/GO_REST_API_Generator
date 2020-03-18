package db_reader

import (
	"strings"

	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type MSSQLColumnDescription struct {
	TABLE_QUALIFIER   string
	TABLE_OWNER       string
	TABLE_NAME        string
	COLUMN_NAME       string
	KEY               string
	DATA_TYPE         string
	TYPE_NAME         string
	PRECISION         float32
	LENGTH            float32
	SCALE             string
	RADIX             string
	NULLABLE          bool
	REMARKS           string
	COLUMN_DEF        string
	SQL_DATA_TYPE     int
	SQL_DATETIME_SUB  string
	CHAR_OCTET_LENGTH string
	ORDINAL_POSITION  int
	IS_NULLABLE       string
	SS_DATA_TYPE      int
}

type MSSQLTable struct {
	Name    string
	Columns []MSSQLColumnDescription
}

func listTablesMSSQL(db *gorm.DB) []string {
	var ret []string
	err := db.Raw("SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES").Pluck("Tables_in_mssql", &ret).Error
	if err != nil {
		panic(err)
	}
	return ret
}

func listPKTableMSSQL(db *gorm.DB, table string) []string {
	var ret []string
	err := db.Raw("SELECT Col.Column_Name from INFORMATION_SCHEMA.TABLE_CONSTRAINTS Tab, "+
		"INFORMATION_SCHEMA.CONSTRAINT_COLUMN_USAGE Col WHERE Col.Constraint_Name = Tab.Constraint_Name "+
		"AND Col.Table_Name = Tab.Table_Name AND Constraint_Type = 'PRIMARY KEY' "+
		"AND Col.Table_Name = '"+table+"'").Pluck("PKs_in_mssql", &ret).Error
	if err != nil {
		panic(err)
	}
	return ret
}

func describeTableMSSQL(db *gorm.DB, name string) []MSSQLColumnDescription {
	var columnList []MSSQLColumnDescription

	pks := listPKTableMSSQL(db, name)

	rows, err := db.Raw("exec sp_columns " + name).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var TABLE_QUALIFIER, TABLE_OWNER, TABLE_NAME, COLUMN_NAME, DATA_TYPE, TYPE_NAME, SCALE, RADIX string
		var REMARKS, COLUMN_DEF, SQL_DATETIME_SUB, CHAR_OCTET_LENGTH, IS_NULLABLE string
		var NULLABLE bool
		var PRECISION, LENGTH float32
		var SQL_DATA_TYPE, ORDINAL_POSITION, SS_DATA_TYPE int

		rows.Scan(
			&TABLE_QUALIFIER,
			&TABLE_OWNER,
			&TABLE_NAME,
			&COLUMN_NAME,
			&DATA_TYPE,
			&TYPE_NAME,
			&PRECISION,
			&LENGTH,
			&SCALE,
			&RADIX,
			&NULLABLE,
			&REMARKS,
			&COLUMN_DEF,
			&SQL_DATA_TYPE,
			&SQL_DATETIME_SUB,
			&CHAR_OCTET_LENGTH,
			&ORDINAL_POSITION,
			&IS_NULLABLE,
			&SS_DATA_TYPE,
		)

		tableDescr := MSSQLColumnDescription{
			TABLE_QUALIFIER:   TABLE_QUALIFIER,
			TABLE_OWNER:       TABLE_OWNER,
			TABLE_NAME:        TABLE_NAME,
			COLUMN_NAME:       COLUMN_NAME,
			DATA_TYPE:         DATA_TYPE,
			TYPE_NAME:         TYPE_NAME,
			PRECISION:         PRECISION,
			LENGTH:            LENGTH,
			SCALE:             SCALE,
			RADIX:             RADIX,
			NULLABLE:          NULLABLE,
			REMARKS:           REMARKS,
			COLUMN_DEF:        COLUMN_DEF,
			SQL_DATA_TYPE:     SQL_DATA_TYPE,
			SQL_DATETIME_SUB:  SQL_DATETIME_SUB,
			CHAR_OCTET_LENGTH: CHAR_OCTET_LENGTH,
			ORDINAL_POSITION:  ORDINAL_POSITION,
			IS_NULLABLE:       IS_NULLABLE,
			SS_DATA_TYPE:      SS_DATA_TYPE,
		}

		for _, e := range pks {
			if e == COLUMN_NAME {
				tableDescr.KEY = "PRI"
			}
		}

		columnList = append(columnList, tableDescr)
	}
	return columnList
}

func MSSQLTableToJSON(table MSSQLTable) models.Object {
	var cols []models.Column

	for _, MSSQLCol := range table.Columns {
		col := models.Column{
			Name:        strings.ToLower(MSSQLCol.COLUMN_NAME),
			Col_Name:    strings.ToLower(MSSQLCol.COLUMN_NAME),
			Col_Type:    strings.ToUpper(MSSQLCol.TYPE_NAME),
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
		if MSSQLCol.NULLABLE {
			col.Is_Null = true
		} else {
			col.Is_Null = false
		}
		if MSSQLCol.KEY == "PRI" {
			col.Primary_Key = true
		}
		cols = append(cols, col)
	}

	obj := models.Object{
		Proyect_Id: "0",
		Name:       strings.ToLower(table.Name),
		Tbl_Name:   "",
		Is_Login:   false,
		LoginId:    "",
		LoginPass:  "",
		Columns:    cols,
	}
	return obj
}

func readObjectsMSSQL(path, bd, user, pass, port string) []models.Object {
	if port == " " {
		port = "1433"
	}
	db, err := gorm.Open("mssql", "sqlserver://"+user+":"+pass+"@"+path+":"+port+"?database="+bd)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	tables := listTablesMSSQL(db)

	var objs []models.Object

	for _, tab := range tables {
		cols := describeTableMSSQL(db, tab)
		table := MSSQLTable{
			Name:    tab,
			Columns: cols,
		}
		obj := MSSQLTableToJSON(table)
		objs = append(objs, obj)
	}

	return objs
}

func readObjectsMSSQL2(db_data models.DbData) []models.Object {
	if db_data.Db_Port == " " {
		db_data.Db_Port = "1433"
	}
	//(db_data.Db_Url, db_data.Db_Name, db_data.Db_User, db_data.Db_Pass, db_data.Db_Port)
	db, err := gorm.Open("mssql", "sqlserver://"+db_data.Db_User+":"+db_data.Db_Pass+"@"+db_data.Db_Url+":"+db_data.Db_Port+"?database="+db_data.Db_Name)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	tables := listTablesMSSQL(db)

	var objs []models.Object

	for _, tab := range tables {
		cols := describeTableMSSQL(db, tab)
		table := MSSQLTable{
			Name:    tab,
			Columns: cols,
		}
		obj := MSSQLTableToJSON(table)
		objs = append(objs, obj)
	}

	return objs
}
