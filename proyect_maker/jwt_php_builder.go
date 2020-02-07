package proyect_maker

import (
	"fmt"
	"os"
)

func JWTPHPMakeCode(P Proyect, Obs []Object) {
	err := os.MkdirAll("./build/models", os.ModePerm)
	check(err)

	err = os.MkdirAll("./build/api", os.ModePerm)
	check(err)

	buildConnection(P)
	buildModels(Obs, P.Db)
	copyJWTPackage()
	buildJWTConfigFile()
	buildAPIEndPoints(Obs)
}

func buildConnection(P Proyect) {
	f, err := os.Create("./build/connection.php")
	check(err)
	defer f.Close()

	_, err = f.WriteString("<?php\n")

	_, err = f.WriteString("\tclass Database{\n")
	_, err = f.WriteString("\t\tprivate $conn;\n\n")

	_, err = f.WriteString("\t\tpublic function __construct(){\n")
	_, err = f.WriteString("\t\t\t$url = '" + P.Db_Data.Db_Url + "';\n")
	_, err = f.WriteString("\t\t\t$database = '" + P.Db_Data.Db_Name + "';\n")
	_, err = f.WriteString("\t\t\t$user = '" + P.Db_Data.Db_User + "';\n")
	_, err = f.WriteString("\t\t\t$password = '" + P.Db_Data.Db_Pass + "';\n")
	_, err = f.WriteString("\t\t\t$port = '" + P.Db_Data.Db_Port + "';\n")
	_, err = f.WriteString("\t\t\t$charset = '" + P.Db_Data.DB_Charset + "';\n")

	connstr := " "

	switch P.Db {
	case "MYSQL":
		//
		connstr = "\t\t\t\t$this->conn = new PDO('mysql:host='.$url.';dbname='.$database.';port:'.$port.';charset='.$charset,$user,$password);\n"
		break
	case "POSTGRESQL":
		// "pgsql:host=%s;port=%d;dbname=%s;user=%s;password=%s",
		connstr = "\t\t\t\t$this->conn = new PDO('pgsql:host=$url;port:$port;dbname=$_base;charset=$charset',$user,$password);\n"
		break
	case "SQLSERVER":
		//new PDO("sqlsrv:Server=localhost,1521;_base=testdb", "NombreUsuario", "Contraseña");
		connstr = "\t\t\t\t$this->conn = new PDO('sqlsrv:host=$url;port:$port;dbname=$_base;charset=$charset',$user,$password);\n"
		break
	}

	_, err = f.WriteString("\t\t\ttry{\n")
	_, err = f.WriteString(connstr)
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
	fmt.Printf("Connection has been created\n")
}

func buildModels(Objs []Object, db string) {
	for _, O := range Objs {
		buildModel(O, db)
	}
}

func buildModel(O Object, db string) {
	strpath := "./build/models/" + O.Name + ".php"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	_, err = f.WriteString("<?php\n")
	_, err = f.WriteString("\tclass " + O.Name + " {\n")

	_, err = f.WriteString("\t\tpublic $connection = null;\n")

	for _, C := range O.Columns {
		_, err = f.WriteString("\t\tpublic $" + C.Name + " = 0;\n")
	}

	_, err = f.WriteString("\n")

	_, err = f.WriteString("\t\tfunction insert(){\n")
	_, err = f.WriteString(strParams(O.Columns))
	_, err = f.WriteString(strInsert(O))
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t\tfunction update(){\n")
	_, err = f.WriteString(strUpdate(O))
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t\tfunction delete(){\n")
	_, err = f.WriteString(strDelete(O))
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString(strLogin(O))

	_, err = f.WriteString("\t\tfunction findOne(){\n")
	_, err = f.WriteString(strFindOne(O))
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t\tfunction findAll(){\n")
	_, err = f.WriteString(strFindAll(O))
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t\tfunction findAllPaginated($index, $cant){\n")
	_, err = f.WriteString(strFindAllPag(O))
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t}\n")
	_, err = f.WriteString("?>")
	fmt.Printf("Model: " + O.Name + " has been created\n")
}

func copyJWTPackage() {
	err := os.MkdirAll("./build/libs", os.ModePerm)
	check(err)

	err = CopyDir("./libraries/php/php-jwt-master", "./build/libs/php-jwt-master")
	check(err)
}

func buildJWTConfigFile() {
	err := os.MkdirAll("./build/config", os.ModePerm)
	check(err)

	f, err := os.Create("./build/config/jwt_config.php")
	check(err)
	defer f.Close()

	_, err = f.WriteString("<?php\n")
	_, err = f.WriteString("\t$key = " + `"example_key"` + ";\n")
	_, err = f.WriteString("\t$iss = " + `"example_key"` + ";\n")
	_, err = f.WriteString("\t$aud = " + `"http://example.com"` + ";\n")
	_, err = f.WriteString("\t$iat = 1356999524;")
	_, err = f.WriteString("\t$nbf = 1357000000;\n")
	_, err = f.WriteString("?>")

	fmt.Printf("REST API: JWT Config File has been created\n")
}

func strParams2(Cols []Column) string {
	strP := ""
	strP += "\t\t\t$data = [\n"

	for _, C := range Cols {
		strP += "\t\t\t\t'" + C.Col_Name + "' => $" + C.Col_Name + ",\n"
	}

	strP += "\t\t\t];\n"

	return strP
}

func strParams(Cols []Column) string {
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

func strLogin(O Object) string {
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

func strInsert(O Object) string {
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

func strUpdate(O Object) string {
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

func strDelete(O Object) string {
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

func strFindAll(O Object) string {
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

func strFindAllPag(O Object) string {
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

func strFindOne(O Object) string {
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

func strJWTImport() string {
	strP := "\tinclude_once '../libs/php-jwt-master/src/BeforeValidException.php';\n"
	strP += "\tinclude_once '../libs/php-jwt-master/src/ExpiredException.php';\n"
	strP += "\tinclude_once '../libs/php-jwt-master/src/SignatureInvalidException.php';\n"
	strP += "\tinclude_once '../libs/php-jwt-master/src/JWT.php';\n"
	strP += "\tinclude_once '../config/jwt_config.php';\n"
	strP += "\t" + `use \Firebase\JWT\JWT;` + "\n\n"
	return strP
}

func strJWTCheck() string {
	strP := "\t$jwt = isset($jsondata['token']) ? $jsondata['token'] : '';\n"
	strP += "\tif($jwt){\n"
	strP += "\t\ttry {\n"
	strP += "\t\t\t$decoded = JWT::decode($jwt, $key, array('HS256'));\n"
	strP += "\t\t}	catch (Exception $e){\n"
	strP += "\t\t\theader('HTTP/1.1 401 Unauthorized', true, 401);\n"
	strP += "\t\t\techo json_encode(array(\n"
	strP += "\t\t\t\t'message' => 'Access denied.',\n"
	strP += "\t\t\t\t'error' => $e->getMessage()\n"
	strP += "\t\t\t));\n"
	strP += "\t\t\treturn;\n"
	strP += "\t\t}\n\n"
	strP += "\t}\n\n"

	return strP
}

func strJWTGenerate(O Object) string {
	strP := "\t\t$token = array(\n"
	strP += "\t\t\t'iss' => $iss\n"
	strP += "\t\t\t'aud' => $aud\n"
	strP += "\t\t\t'iat' => $iat\n"
	strP += "\t\t\t'nbf' => $nbf\n"
	strP += "\t\t\t'data' => array(\n"

	for i, C := range O.Columns {
		if C.Col_Name == O.LoginId || C.Col_Name == O.LoginPass {
			if C.Col_Type == "CHAR" || C.Col_Type == "VARCHAR" {
				if i == (len(O.Columns) - 1) {
					strP += "\t\t\t'" + C.Name + "' => utf8_encode($jsondata['" + C.Name + "'])\n"
				} else {
					strP += "\t\t\t'" + C.Name + "' => utf8_encode($jsondata['" + C.Name + "']),\n"
				}
			} else {
				if i == (len(O.Columns) - 1) {
					strP += "\t\t\t'" + C.Name + "' => $row['" + C.Name + "']\n"
				} else {
					strP += "\t\t\t'" + C.Name + "' => $row['" + C.Name + "'],\n"
				}
			}
		}
	}

	strP += "\t\t\t)\n"
	strP += "\t\t);\n\n"

	strP += "\t\t$jwt = JWT::encode($token, $key);\n\n"

	return strP
}

func strJWTReGenerate() string {
	return ""
}

func buildAPIEndPoints(Objs []Object) {
	for _, O := range Objs {
		buildAPIEndPoint(O)
		if O.Is_Login {
			buildAPILogin(O)
		}
	}
}

func strLoginData(O Object) string {
	strP := ""
	for _, C := range O.Columns {
		if C.Col_Name == O.LoginId || C.Col_Name == O.LoginPass {
			if C.Col_Type == "CHAR" || C.Col_Type == "VARCHAR" {
				strP += "\t\t$" + O.Name + "->" + C.Name + " = utf8_decode($jsondata['" + C.Name + "']);\n"
			} else {
				strP += "\t\t$" + O.Name + "->" + C.Name + " = $jsondata['" + C.Name + "'];\n"
			}
		}
	}

	strP += "\n"

	return strP
}

func buildAPILogin(O Object) {
	strpath := "./build/api/" + O.Name + "_login.php"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	_, err = f.WriteString("<?php\n")

	_, err = f.WriteString("\theader('Access-Control-Allow-Origin: *');\n")
	_, err = f.WriteString("\theader('Content-Type: application/json');\n\n")

	_, err = f.WriteString("\t$contentType = isset($_SERVER['CONTENT_TYPE']) ? trim($_SERVER['CONTENT_TYPE']) : '';\n")
	_, err = f.WriteString("\tif( (strcasecmp($contentType, 'application/json') != 0)\n")
	_, err = f.WriteString("\t\t&& ($_SERVER['REQUEST_METHOD'] === 'POST') ){\n")
	_, err = f.WriteString("\t\theader('HTTP/1.1 404 Bad Request', true, 404);\n")
	_, err = f.WriteString("\t\treturn;\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString(strJWTImport())
	_, err = f.WriteString(strJSONGET())

	_, err = f.WriteString("\tinclude_once '../connection.php';\n")
	_, err = f.WriteString("\tinclude_once '../models/" + O.Name + ".php';\n\n")

	_, err = f.WriteString("\t$db = new Database();\n")
	_, err = f.WriteString("\t$conn = $db->connect();\n\n")

	_, err = f.WriteString("\t$" + O.Name + " = new " + O.Name + "();\n\n")
	_, err = f.WriteString("\t$" + O.Name + "->connection = $conn;\n\n")

	_, err = f.WriteString(strJSONtoParams(O.Name, O.Columns))

	_, err = f.WriteString("\tif( $" + O.Name + "->login() ){\n")
	_, err = f.WriteString(strJWTGenerate(O))
	_, err = f.WriteString("\t\theader('HTTP/1.1 200 Ok', true, 200);\n")
	_, err = f.WriteString("\t\techo json_encode(\n")
	_, err = f.WriteString("\t\t\tarray(\n")
	_, err = f.WriteString("\t\t\t\t'message' => 'Successful login.',\n")
	_, err = f.WriteString("\t\t\t\t'token' => $jwt\n")
	_, err = f.WriteString("\t\t\t)\n")
	_, err = f.WriteString("\t\t);\n")
	_, err = f.WriteString("\t}else{\n")
	_, err = f.WriteString("\t\theader('HTTP/1.1 401 Unauthorized', true, 401);\n")
	_, err = f.WriteString("\t\techo json_encode( array('message' => 'LOGIN FAILED, BAD CREDENTIALS') );\n")
	_, err = f.WriteString("\t}\n")

	_, err = f.WriteString("?>")
}

func buildAPIEndPoint(O Object) {
	strpath := "./build/api/" + O.Name + ".php"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	_, err = f.WriteString("<?php\n")

	_, err = f.WriteString("\theader('Access-Control-Allow-Origin: *');\n")
	_, err = f.WriteString("\theader('Content-Type: application/json');\n\n")

	_, err = f.WriteString("\t$contentType = isset($_SERVER['CONTENT_TYPE']) ? trim($_SERVER['CONTENT_TYPE']) : '';\n")
	_, err = f.WriteString("\tif(strcasecmp($contentType, 'application/json') != 0){\n")
	_, err = f.WriteString("\t\theader('HTTP/1.1 404 Bad Request', true, 404);//throw new Exception('Content type must be: application/json');\n")
	_, err = f.WriteString("\t\treturn;\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString(strJWTImport())
	_, err = f.WriteString(strJSONGET())
	_, err = f.WriteString(strJWTCheck())

	_, err = f.WriteString("\tinclude_once '../connection.php';\n")
	_, err = f.WriteString("\tinclude_once '../models/" + O.Name + ".php';\n\n")

	_, err = f.WriteString("\t$db = new Database();\n")
	_, err = f.WriteString("\t$conn = $db->connect();\n\n")

	_, err = f.WriteString("\t$" + O.Name + " = new " + O.Name + "();\n\n")
	_, err = f.WriteString("\t$" + O.Name + "->connection = $conn;\n\n")

	_, err = f.WriteString(strJSONtoParams(O.Name, O.Columns))

	_, err = f.WriteString("\tif ($_SERVER['REQUEST_METHOD'] === 'GET') {\n")
	_, err = f.WriteString(strAPI_GET(O))
	_, err = f.WriteString("\n\t}else if ($_SERVER['REQUEST_METHOD'] === 'DELETE') {\n")
	_, err = f.WriteString(strAPI_DELETE(O))
	_, err = f.WriteString("\n\t}else if ($_SERVER['REQUEST_METHOD'] === 'POST') {\n")
	_, err = f.WriteString(strAPI_POST(O))
	_, err = f.WriteString("\n\t}else if ($_SERVER['REQUEST_METHOD'] === 'PUT') {\n")
	_, err = f.WriteString(strAPI_PUT(O))
	_, err = f.WriteString("\n\t}\n\n")

	_, err = f.WriteString("?>")
	fmt.Printf("API EndPoint: " + O.Name + " has been created\n")
}

func strUrlParams(Cols []Column) string {
	strP := ""
	for i, C := range Cols {
		if C.Col_Type == "CHAR" || C.Col_Type == "VARCHAR" {
			if i == (len(Cols) - 1) {
				strP += "\t\t\t\t\t\t'" + C.Name + "' => utf8_encode($row['" + C.Name + "'])\n"
			} else {
				strP += "\t\t\t\t\t\t'" + C.Name + "' => utf8_encode($row['" + C.Name + "']),\n"
			}
		} else {
			if i == (len(Cols) - 1) {
				strP += "\t\t\t\t\t\t'" + C.Name + "' => $row['" + C.Name + "']\n"
			} else {
				strP += "\t\t\t\t\t\t'" + C.Name + "' => $row['" + C.Name + "'],\n"
			}
		}
	}
	return strP
}

func strRowtoJSON(Cols []Column) string {
	strP := ""
	strP += "array(\n"

	for i, C := range Cols {
		if C.Col_Type == "CHAR" || C.Col_Type == "VARCHAR" || C.Col_Type == "char" || C.Col_Type == "varchar" {
			if i == (len(Cols) - 1) {
				strP += "\t\t\t\t\t\t'" + C.Name + "' => utf8_encode($row['" + C.Name + "'])\n"
			} else {
				strP += "\t\t\t\t\t\t'" + C.Name + "' => utf8_encode($row['" + C.Name + "']),\n"
			}
		} else {
			if i == (len(Cols) - 1) {
				strP += "\t\t\t\t\t\t'" + C.Name + "' => $row['" + C.Name + "']\n"
			} else {
				strP += "\t\t\t\t\t\t'" + C.Name + "' => $row['" + C.Name + "'],\n"
			}
		}
	}

	strP += "\t\t\t\t\t);\n"

	return strP
}

func strAPI_GET(O Object) string {
	strP := "\t\tif( "
	count := 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += "isset($jsondata->" + O.Columns[i].Name + ")"
				count++
			} else {
				strP += " AND isset($jsondata->" + O.Columns[i].Name + ")"
			}
		}
	}

	strP += " ){\n"

	strP += "\t\t\t$stmt = $" + O.Name + "->findOne();\n\n"

	strP += "\t\t\tif($stmt->columnCount() > 0){\n"

	strP += "\t\t\t\tif($row = $stmt->fetch()){\n"
	strP += "\t\t\t\t\t$item = " + strRowtoJSON(O.Columns)
	strP += "\t\t\t\t\theader('HTTP/1.1 200 OK', true, 200);\n"
	strP += "\t\t\t\t\techo json_encode($item);\n"
	strP += "\t\t\t\t\treturn;\n"
	strP += "\t\t\t\t}\n"
	strP += "\t\t\t}else{\n"
	strP += "\t\t\t\theader('HTTP/1.1 204 No Content', true, 204);\n"
	strP += "\t\t\t\techo json_encode( array('message' => 'NO ROWS FOUND') );\n"
	strP += "\t\t\t\treturn;\n"
	strP += "\t\t\t}\n"

	strP += "\t\t}else{\n"

	strP += "\t\t\t$stmt = $" + O.Name + "->findAll();\n\n"

	strP += "\t\t\tif($stmt->columnCount() > 0){\n"
	strP += "\t\t\t\t$list = array();\n\n"

	strP += "\t\t\t\twhile($row = $stmt->fetch()){\n"
	strP += "\t\t\t\t\t$item = " + strRowtoJSON(O.Columns)
	strP += "\t\t\t\t\tarray_push($list, $item);\n"
	strP += "\t\t\t\t}\n"
	strP += "\t\t\t\theader('HTTP/1.1 200 OK', true, 200);\n"
	strP += "\t\t\t\techo json_encode($list);\n"
	strP += "\t\t\t\treturn;\n"
	strP += "\t\t\t}else{\n"
	strP += "\t\t\t\theader('HTTP/1.1 204 No Content', true, 204);\n"
	strP += "\t\t\t\techo json_encode( array('message' => 'NO ROWS FOUND') );\n"
	strP += "\t\t\t\treturn;\n"
	strP += "\t\t\t}\n"

	strP += "\t\t}\n"

	return strP
}

func strJSONGET() string {
	strP := "\t$bodydata = trim(file_get_contents('php://input'));\n"
	strP += "\t$jsondata = json_decode($bodydata, true);\n\n"
	strP += "\tif(!is_array($jsondata)){\n"
	strP += "\t\theader('HTTP/1.1 404 Bad Request');\n"
	strP += "\t}\n\n"

	return strP
}

func strJSONtoParams(tblName string, Cols []Column) string {
	strP := ""
	for _, C := range Cols {
		if C.Col_Type == "CHAR" || C.Col_Type == "VARCHAR" || C.Col_Type == "char" || C.Col_Type == "varchar" {
			strP += "\t$" + tblName + "->" + C.Name + " = isset($jsondata['" + C.Name + "']) ? utf8_decode($jsondata['" + C.Name + "']) : '';\n"
		} else {
			strP += "\t$" + tblName + "->" + C.Name + " = isset($jsondata['" + C.Name + "']) ? $jsondata['" + C.Name + "'] : '';\n"
		}
	}
	strP += "\n"

	return strP
}

func strAPI_POST(O Object) string {
	strP := "\t\tif( $" + O.Name + "->insert() ){\n"
	strP += "\t\t\theader('HTTP/1.1 201 Created', true, 201);\n"
	strP += "\t\t\techo json_encode( array('message' => 'SUCCESSFUL') );\n"
	strP += "\t\t\treturn;\n"
	strP += "\t\t}else{\n"
	strP += "\t\t\theader('HTTP/1.1 202 Accepted', true, 202);\n"
	strP += "\t\t\techo json_encode( array('message' => 'ERROR WHILE INSERT') );\n"
	strP += "\t\t}\n"

	return strP
}

func strAPI_PUT(O Object) string {
	strP := "\t\tif( $" + O.Name + "->update() ){\n"
	strP += "\t\t\theader('HTTP/1.1 200 OK', true, 200);\n"
	strP += "\t\t\techo json_encode( array('message' => 'SUCCESSFUL') );\n"
	strP += "\t\t\treturn;\n"
	strP += "\t\t}else{\n"
	strP += "\t\t\theader('HTTP/1.1 202 Accepted', true, 202);\n"
	strP += "\t\t\techo json_encode( array('message' => 'ERROR WHILE UPDATE') );\n"
	strP += "\t\t}\n"
	return strP
}

func strAPI_DELETE(O Object) string {
	strP := "\t\tif( $" + O.Name + "->delete() ){\n"
	strP += "\t\t\theader('HTTP/1.1 200 OK', true, 200);\n"
	strP += "\t\t\techo json_encode( array('message' => 'SUCCESSFUL') );\n"
	strP += "\t\t\treturn;\n"
	strP += "\t\t}else{\n"
	strP += "\t\t\theader('HTTP/1.1 202 Accepted', true, 202);\n"
	strP += "\t\t\techo json_encode( array('message' => 'ERROR WHILE DELETE') );\n"
	strP += "\t\t}\n"

	return strP
}
