package proyect_maker

import (
	"fmt"
	"os"
	"strings"
)

func AndroidMakeCode(P Proyect, Obs []Object) {
	mainfolder := "android_build"

	err := os.MkdirAll("./"+mainfolder+"/Models", os.ModePerm)
	check(err)

	err = os.MkdirAll("./"+mainfolder+"/Controllers", os.ModePerm)
	check(err)

	buildAndroidConnection(mainfolder, P)
	buildAndroidModels(mainfolder, Obs, P.Db)
	buildAndroidControllers(mainfolder, Obs, P.Db)
}

func buildAndroidConnection(mainfolder string, P Proyect) {
	f, err := os.Create("./" + mainfolder + "/Controllers/Connection.java")
	check(err)
	defer f.Close()

	_, err = f.WriteString("package Controllers;\n\n")

	_, err = f.WriteString("public class Database{\n")
	//_, err = f.WriteString("\t\tprivate conn;\n\n")
	_, err = f.WriteString("\t\tprivate String url = '" + P.Db_Data.Db_Url + "';\n")
	_, err = f.WriteString("\t\tprivate String database = '" + P.Db_Data.Db_Name + "';\n")
	_, err = f.WriteString("\t\tprivate String user = '" + P.Db_Data.Db_User + "';\n")
	_, err = f.WriteString("\t\tprivate String password = '" + P.Db_Data.Db_Pass + "';\n")
	_, err = f.WriteString("\t\tprivate String port = '" + P.Db_Data.Db_Port + "';\n")
	_, err = f.WriteString("\t\tprivate String charset = '" + P.Db_Data.DB_Charset + "';\n\n")

	_, err = f.WriteString("\t\tpublic Database(){\n")
	_, err = f.WriteString("\t\t\t\n")
	/*

	 */
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t\tpublic Database(){\n")
	_, err = f.WriteString("\t\t\t\n")
	_, err = f.WriteString("\t\t}\n\n")
	/*
		connstr := " "

		switch P.Db {
		case "MYSQL":
			//
			connstr = "\t\t\t\t$this->conn = new PDO('mysql:host='.$url.';dbname='.$database.';port:'.$port.';charset='.$charset,$user,$password);\n"
			break
		case "PSQL":
			// "pgsql:host=%s;port=%d;dbname=%s;user=%s;password=%s",
			connstr = "\t\t\t\t$this->conn = new PDO('pgsql:host=$url;port:$port;dbname=$_base;charset=$charset',$user,$password);\n"
			break
		case "SQLSERVER":
			//new PDO("sqlsrv:Server=localhost,1521;_base=testdb", "NombreUsuario", "Contraseña");
			connstr = "\t\t\t\t$this->conn = new PDO('sqlsrv:host=$url;port:$port;dbname=$_base;charset=$charset',$user,$password);\n"
			break
		}*/

	_, err = f.WriteString("\t\t\ttry{\n")
	//_, err = f.WriteString(connstr)
	_, err = f.WriteString("\t\t\t} catch(Exception $e) {\n")
	_, err = f.WriteString("\t\t\t\terror_log($e->getMessage());\n")
	_, err = f.WriteString("\t\t\t\t//exit('Something weird happened');\n")
	_, err = f.WriteString("\t\t\t}\n")

	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t\tpublic function connect(){\n")
	_, err = f.WriteString("\t\t\treturn $this->conn;\n")
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t\tpublic function disconnect(){\n")
	_, err = f.WriteString("\t\t\t$this->conn = null;\n")
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t}\n")
	_, err = f.WriteString("?>")
	fmt.Printf("Android: Connection has been created\n")
}

func buildAndroidModels(mainfolder string, Objs []Object, db string) {
	for _, O := range Objs {
		buildAndroidModel(mainfolder, O, db)
	}
}

func buildAndroidControllers(mainfolder string, Objs []Object, db string) {
	for _, O := range Objs {
		buildAndroidController(mainfolder, O, db)
	}
}

func buildAndroidModel(mainfolder string, O Object, db string) {
	strpath := "./" + mainfolder + "/Models/" + strings.Title(O.Name) + ".java"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	_, err = f.WriteString("package Models;\n\n")

	_, err = f.WriteString("import com.google.gson.annotations.SerializedName;\n\n")

	_, err = f.WriteString("public class " + strings.Title(O.Name) + " {\n")

	for _, C := range O.Columns {
		_, err = f.WriteString("\t@SerializedName(" + `"` + C.Name + `"` + ")\n")
		_, err = f.WriteString("\tprivate ")
		if C.Col_Type == "char" || C.Col_Type == "varchar" || C.Col_Type == "CHAR" || C.Col_Type == "VARCHAR" || C.Col_Type == "enum" || C.Col_Type == "ENUM" {
			_, err = f.WriteString("String ")
		} else if C.Col_Type == "real" || C.Col_Type == "REAL" {
			_, err = f.WriteString("double ")
		} else if C.Col_Type == "int" || C.Col_Type == "INT" {
			_, err = f.WriteString("int ")
		} else if C.Col_Type == "bool" || C.Col_Type == "BOOL" {
			_, err = f.WriteString("boolean ")
		}
		_, err = f.WriteString(strings.Title(C.Name) + ";\n")
	}

	_, err = f.WriteString("\n")
	_, err = f.WriteString("\tpublic " + strings.Title(O.Name) + "(){\n")

	for _, C := range O.Columns {
		_, err = f.WriteString("\t\tthis.")
		_, err = f.WriteString(strings.Title(C.Name) + " = ")
		if C.Col_Type == "char" || C.Col_Type == "varchar" || C.Col_Type == "CHAR" || C.Col_Type == "VARCHAR" || C.Col_Type == "enum" || C.Col_Type == "ENUM" {
			_, err = f.WriteString(`""`)
		} else if C.Col_Type == "real" || C.Col_Type == "REAL" || C.Col_Type == "int" || C.Col_Type == "INT" {
			_, err = f.WriteString("0")
		} else if C.Col_Type == "bool" || C.Col_Type == "BOOL" {
			_, err = f.WriteString("false")
		}
		_, err = f.WriteString(";\n")
	}

	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\tpublic " + strings.Title(O.Name) + "(")

	for i, C := range O.Columns {
		if C.Col_Type == "char" || C.Col_Type == "varchar" || C.Col_Type == "CHAR" || C.Col_Type == "VARCHAR" || C.Col_Type == "enum" || C.Col_Type == "ENUM" {
			_, err = f.WriteString("String ")
		} else if C.Col_Type == "real" || C.Col_Type == "REAL" {
			_, err = f.WriteString("double ")
		} else if C.Col_Type == "int" || C.Col_Type == "INT" {
			_, err = f.WriteString("int ")
		} else if C.Col_Type == "bool" || C.Col_Type == "BOOL" {
			_, err = f.WriteString("boolean ")
		}
		if i != (len(O.Columns) - 1) {
			_, err = f.WriteString(" " + strings.Title(C.Name) + ", ")
		} else {
			_, err = f.WriteString(" " + strings.Title(C.Name))
		}
	}

	_, err = f.WriteString("){\n")

	for _, C := range O.Columns {
		_, err = f.WriteString("\t\tthis.")
		_, err = f.WriteString(strings.Title(C.Name) + " = " + strings.Title(C.Name))
		_, err = f.WriteString(";\n")
	}

	_, err = f.WriteString("\t}\n")

	_, err = f.WriteString("\n")

	strP := ""

	for _, C := range O.Columns {

		if C.Col_Type == "char" || C.Col_Type == "varchar" || C.Col_Type == "CHAR" || C.Col_Type == "VARCHAR" || C.Col_Type == "enum" || C.Col_Type == "ENUM" {
			strP = "String "
		} else if C.Col_Type == "real" || C.Col_Type == "REAL" {
			strP = "double "
		} else if C.Col_Type == "int" || C.Col_Type == "INT" {
			strP = "int "
		} else if C.Col_Type == "bool" || C.Col_Type == "BOOL" {
			strP = "boolean "
		} else if C.Col_Type == "date" || C.Col_Type == "DATE" || C.Col_Type == "datetime" || C.Col_Type == "DATETIME" {
			strP = "Date "
		}

		_, err = f.WriteString("\tpublic " + strP + "get" + strings.Title(C.Name) + "(){\n")
		_, err = f.WriteString("\t\treturn this." + strings.Title(C.Name) + ";\n")
		_, err = f.WriteString("\t}\n\n")

		_, err = f.WriteString("\tpublic void set" + strings.Title(C.Name) + "(" + strP + strings.Title(C.Name) + "){\n")
		_, err = f.WriteString("\t\tthis." + strings.Title(C.Name) + " = " + strings.Title(C.Name) + ";\n")
		_, err = f.WriteString("\t}\n\n")
	}

	_, err = f.WriteString("}\n")

	fmt.Printf("Android: Model " + O.Name + " has been created\n")
}

func buildAndroidController(mainfolder string, O Object, db string) {
	strpath := "./" + mainfolder + "/Controllers/" + strings.Title(O.Name) + "DAO.java"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	_, err = f.WriteString("package Controllers;\n\n")

	//imports;
	_, err = f.WriteString("import java.util.ArrayList;\n\n")

	_, err = f.WriteString("public class " + strings.Title(O.Name) + "DAO {\n")

	_, err = f.WriteString("\tprivate final String strURL = " + `"` + `"` + ";\n\n")

	_, err = f.WriteString("\tpublic static boolean Insert(" + strings.Title(O.Name) + " O) {\n")
	_, err = f.WriteString("\t\treturn false;\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\tpublic static boolean Update(" + strings.Title(O.Name) + " O) {\n")
	_, err = f.WriteString("\t\treturn false;\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\tpublic static boolean Delete(" + strings.Title(O.Name) + " O) {\n")
	_, err = f.WriteString("\t\treturn false;\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\tpublic static ArrayList<" + strings.Title(O.Name) + "> FindAll() {\n")
	_, err = f.WriteString("\t\treturn null;\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\tpublic static " + strings.Title(O.Name) + " FindOne(String busq) {\n")
	_, err = f.WriteString("\t\treturn null;\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("}\n")

	fmt.Printf("Android: Controller " + strings.Title(O.Name) + "DAO has been created\n")
}

func strAndroidParams2(Cols []Column) string {
	strP := ""
	strP += "\t\t\t$data = [\n"

	for _, C := range Cols {
		strP += "\t\t\t\t'" + C.Col_Name + "' => $" + C.Col_Name + ",\n"
	}

	strP += "\t\t\t];\n"

	return strP
}

func strAndroidParams(Cols []Column) string {
	strP := "\t\t\t$data = array(\n"

	for i, C := range Cols {
		if i == (len(Cols) - 1) {
			strP += "\t\t\t\t'" + C.Col_Name + "' => $this->" + C.Name + "\n"
		} else {
			strP += "\t\t\t\t'" + C.Col_Name + "' => $this->" + C.Name + ",\n"
		}
	}

	strP += "\t\t\t\t);\n"

	return strP
}

func strAndroidLogin(O Object) string {
	strP := ""

	if O.Is_Login {
		strP += "\t\tfunction login(){\n"
		strP += "\t\t\t$sql = 'SELECT " + O.LoginId + ", " + O.LoginPass
		strP += " FROM " + O.Tbl_Name + " WHERE "
		strP += O.LoginId + " = ? AND " + O.LoginPass + " = ?';\n"

		strP += "\t\t\t$stmt = $this->connection->prepare($sql);\n"
		strP += "\t\t\tif($stmt->execute( array("
		strP += "$this->" + O.LoginId + ", $this->" + O.LoginPass + ")"
		strP += ") ) { return true; }\n"

		strP += "\t\t\treturn false;\n"

		strP += "\t\t}\n\n"
	}

	return strP
}

func strAndroidInsert(O Object) string {
	strP := ""
	strP += "\t\t\t$sql = 'INSERT INTO " + O.Tbl_Name + " ("

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
		} else {
			strP += ", " + O.Columns[i].Col_Name
		}
	}

	strP += ") VALUES ("

	for i, _ := range O.Columns {
		if i == 0 {
			strP += ":" + O.Columns[i].Col_Name
		} else {
			strP += ", :" + O.Columns[i].Col_Name
		}
	}

	strP += ")';\n"

	strQ := ""
	strQ += "\t\t\t$stmt= $this->connection->prepare($sql);\n"
	strQ += "\t\t\tif( $stmt->execute($data) ){ return true; }\n"
	strP += strQ

	strP += "\t\t\treturn false;\n"

	return strP
}

func strAndroidUpdate(O Object) string {
	strP := ""
	strQ := ""
	strP += "\t\t\t$sql = 'UPDATE " + O.Tbl_Name + " SET "
	strQ += "\t\t\t$stmt = $this->connection->prepare($sql);\n"
	strQ += "\t\t\tif( $stmt->execute( array("

	count := 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key != true {
			if count == 0 {
				strP += O.Columns[i].Col_Name + " = ?"
				strQ += "$this->" + O.Columns[i].Name
				count = 1
			} else {
				strP += ", " + O.Columns[i].Col_Name + " = ?"
				strQ += ",$this->" + O.Columns[i].Name
			}
		}
	}

	strP += " WHERE "
	count = 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += " " + O.Columns[i].Col_Name + " = ?"
				strQ += ",$this->" + O.Columns[i].Name
				count++
			} else {
				strP += " AND " + O.Columns[i].Col_Name + " = ?"
				strQ += ",$this->" + O.Columns[i].Name
			}
		}
	}

	strQ += ") ) ){ return true; }\n"
	strP += "';\n"
	strP += strQ
	strP += "\t\t\treturn false;\n"

	return strP
}

func strAndroidDelete(O Object) string {
	strP := ""
	strQ := ""
	strP += "\t\t\t$sql = 'DELETE FROM " + O.Tbl_Name + " WHERE "
	strQ += "\t\t\t$stmt = $this->connection->prepare($sql);\n"
	strQ += "\t\t\tif($stmt->execute( array("

	count := 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += O.Columns[i].Col_Name + " = ?"
				strQ += "$this->" + O.Columns[i].Name
				count++
			} else {
				strP += " AND " + O.Columns[i].Col_Name + " = ?"
				strQ += ",$this->" + O.Columns[i].Name
			}
		}
	}

	strQ += ") ) ){ return true; }\n"
	strP += "';\n"
	strP += strQ
	strP += "\t\t\treturn false;\n"

	return strP
}

func strAndroidFindAll(O Object) string {
	strP := ""
	strQ := ""
	strP += "\t\t\t$sql = 'SELECT "
	strQ += "\t\t\t$stmt = $this->connection->prepare($sql);\n"
	strQ += "\t\t\tif($stmt->execute()) { return $stmt; }\n"

	count := 0

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
			count++
		} else {
			strP += ", " + O.Columns[i].Col_Name
		}
	}

	strP += " FROM " + O.Tbl_Name + "';\n"
	strP += strQ

	strP += "\t\t\treturn null;\n"

	return strP
}

func strAndroidFindAllPag(O Object) string {
	strP := ""
	strQ := ""
	strP += "\t\t\t$sql = 'SELECT "
	strQ += "\t\t\t$stmt = $this->connection->prepare($sql);\n"
	strQ += "\t\t\tif($stmt->execute( array($cant,($index*$cant)) ) ) { return $stmt; }\n"

	count := 0

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
			count++
		} else {
			strP += ", " + O.Columns[i].Col_Name
		}
	}

	strP += " FROM " + O.Tbl_Name + " LIMIT ? OFFSET ?';\n"
	strP += strQ

	strP += "\t\t\treturn null;\n"

	return strP
}

func strAndroidFindOne(O Object) string {
	strP := ""
	strQ := ""
	strP += "\t\t\t$sql = 'SELECT "
	strQ += "\t\t\t$stmt = $this->connection->prepare($sql);\n"
	strQ += "\t\t\tif($stmt->execute( array("

	count := 0

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
			count = 1
		} else {
			strP += ", " + O.Columns[i].Col_Name
		}
	}

	strP += " FROM " + O.Tbl_Name + " WHERE "
	count = 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += O.Columns[i].Col_Name + " = ?"
				strQ += "$this->" + O.Columns[i].Name
				count++
			} else {
				strP += " AND " + O.Columns[i].Col_Name + " = ?"
				strQ += ",$this->" + O.Columns[i].Name
			}
		}
	}

	strQ += ") ) ) { return $stmt; }\n"
	strP += "';\n"
	strP += strQ

	strP += "\t\t\treturn null;\n"

	return strP
}
