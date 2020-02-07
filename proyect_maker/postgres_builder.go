package proyect_maker

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PostgresMakeCode(P Proyect, Obs []Object) {
	err := os.MkdirAll("./db-build/sql/postgres", os.ModePerm)
	check(err)

	f, err := os.Create("./db-build/sql/postgres/" + P.Name + ".sql")
	check(err)
	defer f.Close()

	_, err = f.WriteString(postgresBuildDataBase(P))
	_, err = f.WriteString(postgresBuildTables(Obs))

}

func postgresBuildDataBase(P Proyect) string {
	strP := "CREATE DATABASE " + P.Db_Data.Db_Name
	if P.Db_Data.DB_Charset != "none" {
		strP += " CHARACTER SET " + P.Db_Data.DB_Charset + " COLLATE " + P.Db_Data.DB_Collate
	}
	strP += ";\n\n"
	strP += "USE " + P.Db_Data.Db_Name + ";\n\n"
	return strP
}

func postgresBuildTables(Objs []Object) string {
	strP := ""

	for _, O := range Objs {
		strP += postgresBuildTable(O)
		fmt.Println("SQL: " + O.Name + " has been written")
	}

	return strP
}

func postgresBuildTable(O Object) string {
	strP := "CREATE TABLE " + O.Tbl_Name + " (\n"
	strP += postgresBuildColumns(O.Columns)
	strP += postgresBuildFKs(O.Columns)
	strP += postgresBuildPKs(O.Columns)
	strP += ") ENGINE = INNODB;\n\n"
	return strP
}

func postgresBuildColumns(Cols []Column) string {
	strP := ""
	for _, C := range Cols {
		strP += "\t" + C.Col_Name + " " + strings.ToUpper(C.Col_Type)
		if C.Col_Lenght > 0 {
			strP += "(" + strconv.Itoa(C.Col_Lenght) + ")"
		}
		if len(C.Enum_List) > 0 {
			strP += "("
			for i, E := range C.Enum_List {
				if i != (len(C.Enum_List) - 1) {
					strP += E + ","
				} else {
					strP += E
				}
			}
			strP += ")"
		}
		if C.Is_Unique {
			strP += " UNIQUE "
		}
		if !C.Is_Null {
			strP += " NOT NULL"
		}
		strP += ",\n"
	}
	return strP
}

func postgresBuildFKs(Cols []Column) string {
	strP := ""
	for _, C := range Cols {
		if C.Foreign_Key {
			strP += "\tFOREIGN KEY (" + C.Col_Name
			strP += ") REFERENCES " + C.Tbl_Ref + "("
			for j, Ref := range C.Col_Refs {
				if j < (len(C.Col_Refs) - 1) {
					strP += Ref + ", "
				} else {
					strP += Ref
				}
			}
			strP += ")"
			strP += " ON UPDATE " + strings.ToUpper(C.On_Update)
			strP += " ON DELETE " + strings.ToUpper(C.On_Delete)
			strP += ",\n"
		}
	}
	return strP
}

func postgresBuildPKs(Cols []Column) string {
	strP := "\tPRIMARY KEY ("

	var Cs = make([]Column, 0)

	for _, C := range Cols {
		if C.Primary_Key {
			Cs = append(Cs, C)
		}
	}

	for i, C := range Cs {
		if i < (len(Cs) - 1) {
			strP += C.Col_Name + ", "
		} else {
			strP += C.Col_Name
		}
	}

	strP += ")\n"

	return strP
}

func postgresBuildIndexes(Tbl_Name string, Cols []Column) string {
	strP := "\tINDEX index_" + Tbl_Name + " ("
	for i, C := range Cols {
		if C.Is_Index {
			if i < (len(Cols) - 1) {
				strP += C.Col_Name + ", "
			} else {
				strP += C.Col_Name
			}
		}
	}

	strP += ")\n"

	return strP
}
