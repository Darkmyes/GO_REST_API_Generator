package proyect_maker

import (
	"fmt"
	"os"
	"strings"
)

/*
req.params mejorado {for de variables para cada pk y en las rutas tambien o dejarla de las dos formas}
*/

func JWTNodeMakeCode(P Proyect, Obs []Object) {
	buildPath := "node_build"
	err := os.MkdirAll("./"+buildPath+"/src/controllers", os.ModePerm)
	check(err)

	//err = os.MkdirAll("./"+buildPath+"/src/models", os.ModePerm)
	//check(err)

	err = os.MkdirAll("./"+buildPath+"/src/routes", os.ModePerm)
	check(err)

	buildNodeMain(buildPath, P, Obs)
	buildNodeConnection(buildPath, P)
	//buildNodeStuffs(buildPath)
	//buildNodeModels(buildPath, Obs, P.Db)
	buildNodeControllers(buildPath, Obs, P.Db)
	buildNodeRoutes(buildPath, Obs, P.Db)
	//buildJWTNodeConfigFile()
	//buildNodeAPIEndPoints(buildPath, Obs)
}

func buildNodeMain(buildPath string, P Proyect, Objs []Object) {
	//index.js File
	f, err := os.Create("./" + buildPath + "/src/index.js")
	check(err)
	defer f.Close()

	_, err = f.WriteString("require('dotenv').config();\n")
	_, err = f.WriteString("const app = require('./app');\n")
	_, err = f.WriteString("require('./database');\n\n")

	_, err = f.WriteString("async function main() {\n")
	_, err = f.WriteString("\tawait app.listen(app.get('port'));\n")
	_, err = f.WriteString("\tconsole.log('Server on port', app.get('port'));\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("main();\n")

	fmt.Printf("Main: index.js has been created\n")

	//app.js File
	f, err = os.Create("./" + buildPath + "/src/app.js")
	check(err)
	defer f.Close()

	_, err = f.WriteString("const express = require('express');\n")
	_, err = f.WriteString("const cors = require('cors');\n\n")

	_, err = f.WriteString("const app = express();\n\n")

	_, err = f.WriteString("// settings\n")
	_, err = f.WriteString("app.set('port', process.env.PORT || 4000);\n\n")

	_, err = f.WriteString("// middlewares\n")
	_, err = f.WriteString("app.use(cors());\n")
	_, err = f.WriteString("app.use(express.json());\n\n")

	_, err = f.WriteString("// routes\n")

	for _, O := range Objs {
		if O.Is_Login {
			_, err = f.WriteString("app.use('/api/login', require('./routes/login'));\n")
		}

		_, err = f.WriteString("app.use('/api/" + O.Name + "', require('./routes/" + O.Name + "'));\n")
		//_, err = f.WriteString("\n")
	}

	_, err = f.WriteString("\nmodule.exports = app;\n")

	fmt.Printf("Main: app.js has been created\n")
}

func buildNodeConnection(buildPath string, P Proyect) {
	//database.js File
	f, err := os.Create("./" + buildPath + "/src/database.js")
	check(err)
	defer f.Close()

	_, err = f.WriteString("const mysql = require('mysql');\n")
	_, err = f.WriteString("const { promisify }= require('util');\n\n")

	_, err = f.WriteString("pool = mysql.createPool({\n")
	_, err = f.WriteString("\thost: '" + P.Db_Data.Db_Url + "',\n")
	_, err = f.WriteString("\tuser: '" + P.Db_Data.Db_User + "',\n")
	_, err = f.WriteString("\tpassword: '" + P.Db_Data.Db_Pass + "',\n")
	_, err = f.WriteString("\tdatabase: '" + P.Db_Data.Db_Name + "',\n")
	_, err = f.WriteString("\tport: " + P.Db_Data.Db_Port + "\n")
	_, err = f.WriteString("});\n\n")

	_, err = f.WriteString("pool.getConnection((err, connection) => {\n")
	_, err = f.WriteString("\tif (err) {\n")
	_, err = f.WriteString("\t\tif (err.code === 'PROTOCOL_CONNECTION_LOST') {\n")
	_, err = f.WriteString("\t\tconsole.error('Database connection was closed.');\n")
	_, err = f.WriteString("\t\t}\n")
	_, err = f.WriteString("\t\tif (err.code === 'ER_CON_COUNT_ERROR') {\n")
	_, err = f.WriteString("\t\tconsole.error('Database has to many connections.');\n")
	_, err = f.WriteString("\t\t}\n")
	_, err = f.WriteString("\t\tif (err.code === 'ECONNREFUSED') {\n")
	_, err = f.WriteString("\t\tconsole.error('Database connection was refused.');\n")
	_, err = f.WriteString("\t\t}\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\tif (connection) connection.release();\n")
	_, err = f.WriteString("\tconsole.log('DB is Connected');\n\n")

	_, err = f.WriteString("\treturn;\n")
	_, err = f.WriteString("});\n")

	_, err = f.WriteString("// Promisify Pool Querys\n")
	_, err = f.WriteString("pool.query = promisify(pool.query);\n\n")

	_, err = f.WriteString("module.exports = pool;\n")
	/*
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
	*/
	fmt.Printf("Model: DB has been created\n")
}

func buildNodeStuffs(buildPath string) {
	f, err := os.Create("./" + buildPath + "/models/stuffs.Node")
	check(err)
	defer f.Close()

	_, err = f.WriteString("package models\n\n")

	_, err = f.WriteString("type Info struct {\n")
	_, err = f.WriteString("\tInfo string ")
	_, err = f.Write([]byte{'\u0060'})
	_, err = f.WriteString(`"json:info"`)
	_, err = f.Write([]byte{'\u0060'})
	_, err = f.WriteString("\n}\n")

	fmt.Printf("Model: Stuff has been created\n")
}

func buildNodeModels(buildPath string, Objs []Object, db string) {
	for _, O := range Objs {
		buildNodeModel(buildPath, O, db)
	}
}

func buildNodeRoutes(buildPath string, Objs []Object, db string) {
	for _, O := range Objs {
		buildNodeRoute(buildPath+"/src", O, db)
	}
}

func buildNodeControllers(buildPath string, Objs []Object, db string) {
	for _, O := range Objs {
		buildNodeController(buildPath+"/src", O, db)
		if O.Is_Login {
			buildNodeAPILogin(buildPath, O)
		}
	}
}

func buildNodeModel(buildPath string, O Object, db string) {
	strpath := "./" + buildPath + "/models/" + O.Name + ".Node"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	_, err = f.WriteString("package models\n\n")
	_, err = f.WriteString("import (\n")
	_, err = f.WriteString("\t" + `"database/sql"` + "\n")
	_, err = f.WriteString(")\n\n")

	_, err = f.WriteString("type " + strings.Title(O.Name) + " struct  {\n")

	for _, C := range O.Columns {
		_, err = f.WriteString("\t" + strings.Title(C.Name) + " ")
		if C.Col_Type == "char" || C.Col_Type == "varchar" || C.Col_Type == "CHAR" || C.Col_Type == "VARCHAR" || C.Col_Type == "enum" || C.Col_Type == "ENUM" {
			_, err = f.WriteString("string ")
		} else if C.Col_Type == "real" || C.Col_Type == "REAL" {
			_, err = f.WriteString("float64 ")
		} else if C.Col_Type == "int" || C.Col_Type == "INT" {
			_, err = f.WriteString("int32 ")
		} else if C.Col_Type == "bool" || C.Col_Type == "BOOL" {
			_, err = f.WriteString("bool ")
		}
		_, err = f.Write([]byte{'\u0060'})
		_, err = f.WriteString(`json:"` + C.Name + `"`)
		_, err = f.Write([]byte{'\u0060'})
		_, err = f.WriteString("\n")
	}

	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString(strNodeLogin(O))

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") Insert (db *sql.DB) error {\n")
	_, err = f.WriteString(strNodeInsert(O))
	_, err = f.WriteString("\treturn nil\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") Update (db *sql.DB) error {\n")
	_, err = f.WriteString(strNodeUpdate(O))
	_, err = f.WriteString("\treturn nil\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") Delete (db *sql.DB) error {\n")
	_, err = f.WriteString(strNodeDelete(O))
	_, err = f.WriteString("\treturn nil\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") FindOne (db *sql.DB) (*" + strings.Title(O.Name) + ", error) {\n")
	_, err = f.WriteString("\tobj := new(" + strings.Title(O.Name) + ")\n")
	_, err = f.WriteString(strNodeFindOne(O))
	_, err = f.WriteString("\treturn nil, nil\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") FindAll (db *sql.DB) ([]*" + strings.Title(O.Name) + ", error) {\n")
	_, err = f.WriteString("\tobjs := make([]*" + strings.Title(O.Name) + ", 0)\n")
	_, err = f.WriteString(strNodeFindAll(O))
	_, err = f.WriteString("\treturn objs, nil\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") FindAllPaginated (db *sql.DB, index int, nums int) ([]*" + strings.Title(O.Name) + ", error) {\n")
	_, err = f.WriteString("\tobjs := make([]*" + strings.Title(O.Name) + ", 0)\n")
	_, err = f.WriteString(strNodeFindAllPag(O))
	_, err = f.WriteString("\treturn objs, nil\n")
	_, err = f.WriteString("}\n\n")

	fmt.Printf("Model: " + O.Name + " has been created\n")
}

func buildNodeRoute(buildPath string, O Object, db string) {
	strpath := "./" + buildPath + "/routes/" + O.Name + ".js"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	_, err = f.WriteString("const { Router } = require('express');\n")
	_, err = f.WriteString("const router = Router();\n")
	_, err = f.WriteString("const verifyToken = require('./verifyToken')\n\n")

	_, err = f.WriteString("const " + O.Name + " = require('../controllers/" + O.Name + ".controller');\n\n")

	_, err = f.WriteString("router.route('/')\n")
	_, err = f.WriteString("\t.get(verifyToken, " + O.Name + ".FindAll)\n")
	_, err = f.WriteString("\t.post(verifyToken, " + O.Name + ".Insert)\n")
	_, err = f.WriteString("\t.put(verifyToken, " + O.Name + ".Update)\n")
	_, err = f.WriteString("\t.delete(verifyToken, " + O.Name + ".Delete);\n\n")

	_, err = f.WriteString("router.route('/:id')\n")
	_, err = f.WriteString("\t.get(verifyToken, " + O.Name + ".FindOne)\n\n")

	_, err = f.WriteString("router.route('/:index/:count')\n")
	_, err = f.WriteString("\t.get(verifyToken, " + O.Name + ".FindAllPaginated);\n\n")

	_, err = f.WriteString("module.exports = router;")

	fmt.Printf("Route: " + O.Name + " has been created\n")
}

func buildNodeController(buildPath string, O Object, db string) {
	strpath := "./" + buildPath + "/controllers/" + O.Name + ".controller.js"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	_, err = f.WriteString("const " + O.Name + "Ctrl = {};\n\n")

	_, err = f.WriteString("const pool = require('../database.js');\n\n")

	_, err = f.WriteString(strNodeInsert(O))
	_, err = f.WriteString(strNodeUpdate(O))
	_, err = f.WriteString(strNodeDelete(O))

	_, err = f.WriteString(strNodeFindOne(O))
	_, err = f.WriteString(strNodeFindAll(O))
	_, err = f.WriteString(strNodeFindAllPag(O))

	_, err = f.WriteString("module.exports = " + O.Name + "Ctrl;")

	fmt.Printf("Controller: " + O.Name + " has been created\n")
}

func strNodeLogin(O Object) string {
	strP := ""

	if O.Is_Login {

		strP += "func (o " + strings.Title(O.Name) + ") Login (db *sql.DB) (bool, error) {\n"
		strP += "\tobj := new(" + strings.Title(O.Name) + ")\n"

		strP += "\tstrSQL := " + `"` + "SELECT " + O.LoginId + ", " + O.LoginPass
		strP += " FROM " + O.Tbl_Name + " WHERE "
		strP += O.LoginId + " = ? AND " + O.LoginPass + " = ?" + `"` + "\n"

		strP += "\trows, err := db.Query(strSQL, o." + strings.Title(O.LoginId) + " , o." + strings.Title(O.LoginPass) + ")\n\n"
		strP += "\tif err != nil { return false, err }\n"
		strP += "\tdefer rows.Close()\n\n"

		strP += "\tif rows.Next() {\n"
		strP += "\t\terr := rows.Scan( &obj." + strings.Title(O.LoginId) + ", &obj." + strings.Title(O.LoginPass) + ")\n"
		strP += "\t\tif err != nil { return false, err }\n"

		strP += "\t\tif " + "(obj." + strings.Title(O.LoginId) + " == o." + strings.Title(O.LoginId) + " && "
		strP += "obj." + strings.Title(O.LoginPass) + " == o." + strings.Title(O.LoginId) + ") {\n"
		strP += "\t\t\treturn true, nil\n"
		strP += "\t\t}\n"

		strP += "\t}\n"
		strP += "\tif err = rows.Err(); err != nil { return false, err }\n"

		strP += "\treturn false, err\n"

		strP += "}\n\n"
	}

	return strP
}

func strNodeInsert(O Object) string {
	strP := O.Name + "Ctrl.Insert = async (req, res) => {\n"
	strP += "\ttry {\n"

	strP += "\t\tconst obj = req.body;\n"

	strP += "\t\tconst strSQL = " + `"` + "INSERT INTO " + O.Tbl_Name + " ("
	strQ := ""
	strW := "\t\tconst objdata = ["

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
			strQ += "?"
			strW += "obj." + strings.Title(O.Columns[i].Col_Name)
		} else {
			strP += ", " + O.Columns[i].Col_Name
			strQ += ", ?"
			strW += ", obj." + strings.Title(O.Columns[i].Col_Name)
		}
	}

	strW += "];\n\n"
	strP += ") VALUES (" + strQ + ")" + `"` + ";\n"
	strP += strW

	strP += "\t\tconst objs = await pool.query(strSQL, objdata);\n"
	strP += "\t\tres.json('" + strings.Title(O.Name) + " created');\n"

	strP += "\t} catch(err) {\n"

	strP += "\t\tres.status(400).json({\n"
	strP += "\t\t\terror: err\n"
	strP += "\t\t});\n"

	strP += "\t}\n"
	strP += "}\n\n"

	return strP
}

func strNodeUpdate(O Object) string {
	strP := O.Name + "Ctrl.Update = async (req, res) => {\n"
	strP += "\ttry {\n"

	strP += "\t\tconst obj = req.body;\n"

	strP += "\t\tconst strSQL = " + `"` + "UPDATE " + O.Tbl_Name + " SET "
	strQ := "\t\tconst objdata = ["

	count := 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key != true {
			if count == 0 {
				strP += O.Columns[i].Col_Name + " = ?"
				strQ += "obj." + strings.Title(O.Columns[i].Col_Name)
				count = 1
			} else {
				strP += ", " + O.Columns[i].Col_Name + " = ?"
				strQ += ", obj." + strings.Title(O.Columns[i].Col_Name)
			}
		}
	}

	strP += " WHERE "
	count = 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += " " + O.Columns[i].Col_Name + " = ?"
				strQ += ", o." + strings.Title(O.Columns[i].Col_Name)
				count++
			} else {
				strP += " AND " + O.Columns[i].Col_Name + " = ?"
				strQ += ", o." + strings.Title(O.Columns[i].Col_Name)
			}
		}
	}

	strP += `"` + ";\n"
	strQ += "];\n\n"

	strP += strQ

	strP += "\t\tconst objs = await pool.query(strSQL, objdata);\n"
	strP += "\t\tres.json('" + strings.Title(O.Name) + " updated');\n"

	strP += "\t} catch(err) {\n"

	strP += "\t\tres.status(400).json({\n"
	strP += "\t\t\terror: err\n"
	strP += "\t\t});\n"

	strP += "\t}\n"
	strP += "}\n\n"

	return strP
}

func strNodeDelete(O Object) string {
	strP := O.Name + "Ctrl.Delete = async (req, res) => {\n"
	strP += "\ttry {\n"

	strP += "\t\tconst obj = req.body;\n"

	strP += "\t\tconst strSQL = " + `"` + "DELETE FROM " + O.Tbl_Name + " WHERE "
	strQ := "\t\tconst objdata = ["

	count := 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += O.Columns[i].Col_Name + " = ?"
				strQ += "obj." + strings.Title(O.Columns[i].Col_Name)
				count++
			} else {
				strP += " AND " + O.Columns[i].Col_Name + " = ?"
				strQ += ", obj." + strings.Title(O.Columns[i].Col_Name)
			}
		}
	}

	strP += `"` + ";\n"
	strQ += "];\n\n"

	strP += strQ

	strP += "\t\tconst objs = await pool.query(strSQL, objdata);\n"
	strP += "\t\tres.json('" + strings.Title(O.Name) + " deleted');\n"

	strP += "\t} catch(err) {\n"

	strP += "\t\tres.status(400).json({\n"
	strP += "\t\t\terror: err\n"
	strP += "\t\t});\n"

	strP += "\t}\n"
	strP += "}\n\n"

	return strP
}

func strNodeFindOne(O Object) string {
	strP := O.Name + "Ctrl.FindOne = async (req, res) => {\n"
	strP += "\ttry {\n"

	strP += "\t\tconst strSQL = " + `"` + "SELECT "

	//strQ := ""
	strW := "\t\tconst objdata = ["

	count := 0

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
			//strW += "obj." + strings.Title(O.Columns[i].Col_Name)
			count = 1
		} else {
			strP += ", " + O.Columns[i].Col_Name
			//strW += ", obj." + strings.Title(O.Columns[i].Col_Name)
		}
	}

	strP += " FROM " + O.Tbl_Name + " WHERE "
	count = 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += O.Columns[i].Col_Name + " = ?"
				strW += "obj." + strings.Title(O.Columns[i].Col_Name)
				count++
			} else {
				strP += " AND " + O.Columns[i].Col_Name + " = ?"
				strW += ", obj." + strings.Title(O.Columns[i].Col_Name)
			}
		}
	}

	strP += `"` + ";\n"

	strW += "];\n\n"

	strP += strW

	strP += "\t\tconst objs = await pool.query(strSQL, objdata);\n"
	strP += "\t\tres.json(objs);\n"

	strP += "\t} catch(err) {\n"

	strP += "\t\tres.status(400).json({\n"
	strP += "\t\t\terror: err\n"
	strP += "\t\t});\n"

	strP += "\t}\n"
	strP += "}\n\n"

	return strP
}

func strNodeFindAll(O Object) string {
	strP := O.Name + "Ctrl.FindAll = async (req, res) => {\n"
	strP += "\ttry {\n"

	strP += "\t\tconst strSQL = " + `"` + "SELECT "

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
		} else {
			strP += ", " + O.Columns[i].Col_Name
		}
	}

	strP += " FROM " + O.Tbl_Name + `"` + ";\n"

	strP += "\t\tconst objs = await pool.query(strSQL);\n"
	strP += "\t\tres.json(objs);\n"

	strP += "\t} catch(err) {\n"

	strP += "\t\tres.status(400).json({\n"
	strP += "\t\t\terror: err\n"
	strP += "\t\t});\n"

	strP += "\t}\n"
	strP += "}\n\n"

	return strP
}

func strNodeFindAllPag(O Object) string {
	strP := O.Name + "Ctrl.FindAllPaginated = async (req, res) => {\n"
	strP += "\ttry {\n"

	strP += "\t\tconst strSQL = " + `"` + "SELECT "

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
		} else {
			strP += ", " + O.Columns[i].Col_Name
		}
	}

	strP += " FROM " + O.Tbl_Name + " LIMIT ? OFFSET ?" + `"` + ";\n"

	strP += "\t\tconst { index, count } = req.params;\n"
	strP += "\t\tconst objdata = [index, count];\n"

	strP += "\t\tconst objs = await pool.query(strSQL, objdata);\n"
	strP += "\t\tres.json(objs);\n"

	strP += "\t} catch(err) {\n"

	strP += "\t\tres.status(400).json({\n"
	strP += "\t\t\terror: err\n"
	strP += "\t\t});\n"

	strP += "\t}\n"
	strP += "}\n\n"

	return strP
}

func buildNodeAPILogin(buildPath string, O Object) {
	f, err := os.Create("./" + buildPath + "/src/config.js")
	check(err)
	defer f.Close()

	_, err = f.WriteString("module.exports = {\n")
	_, err = f.WriteString("\tsecret: 'mysecretkey'\n")
	_, err = f.WriteString("}\n")

	fmt.Printf("Login: Secret has been created\n")

	f, err = os.Create("./" + buildPath + "/src/controllers/verifyToken.js")
	check(err)
	defer f.Close()

	_, err = f.WriteString("const jwt = require('jsonwebtoken');\n")
	_, err = f.WriteString("const config = require('../config');\n\n")

	_, err = f.WriteString("async function verifyToken(req, res, next) {\n")
	_, err = f.WriteString("\tconst token = req.headers['x-access-token'];\n")
	_, err = f.WriteString("\tif (!token) {\n")
	_, err = f.WriteString("\t\treturn res.status(401).send({ auth: false, message: 'No token provided' });\n")
	_, err = f.WriteString("\t}\n")
	_, err = f.WriteString("\tconst decoded = await jwt.verify(token, config.secret);\n")
	_, err = f.WriteString("\treq.userId = decoded.id;\n")
	_, err = f.WriteString("\tnext();\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("module.exports = verifyToken;\n")
	fmt.Printf("Login: verifyToken has been created\n")

	//Login EndPoint
	f, err = os.Create("./" + buildPath + "/src/routes/login.js")
	check(err)
	defer f.Close()

	_, err = f.WriteString("const { Router } = require('express');\n")
	_, err = f.WriteString("const router = Router();\n\n")

	//_, err = f.WriteString("const verifyToken = require('./verifyToken')\n\n")

	_, err = f.WriteString("const jwt = require('jsonwebtoken');\n")
	_, err = f.WriteString("const config = require('../config');\n\n")

	_, err = f.WriteString("const pool = require('../database.js');\n\n")

	_, err = f.WriteString("router.route('/')\n")
	_, err = f.WriteString("\t.get(")

	//strP := "Login = async (req, res) => {\n"
	strP := "async (req, res) => {\n"
	strP += "\t\ttry {\n"

	strP += "\t\t\tconst obj = req.body;\n"

	strP += "\t\t\tconst strSQL = " + `"` + "SELECT " + O.LoginId + ", " + O.LoginPass
	strP += " FROM " + O.Tbl_Name + " WHERE "
	strP += O.LoginId + " = ? AND " + O.LoginPass + " = ?" + `"` + ";\n"

	strP += "\t\t\tconst objdata = [ obj." + O.LoginId + ", obj." + O.LoginPass + "];\n\n"

	strP += "\t\t\tconst objs = await pool.query(strSQL, objdata);\n"
	strP += "\t\t\tconst token = jwt.sign({id: obj." + O.LoginId + "}, config.secret, {\n"
	strP += "\t\t\t\texpiresIn: 60 * 60 * 24\n"
	strP += "\t\t\t});\n"

	strP += "\t\t\tres.cookie('jwt-token', token);\n"
	strP += "\t\t\t//res.status(200).json({auth: true, token});\n"

	strP += "\t\t} catch(err) {\n"

	strP += "\t\t\tres.status(400).json({\n"
	strP += "\t\t\t\terror: err\n"
	strP += "\t\t\t});\n"

	strP += "\t\t}\n"
	strP += "\t});\n\n"

	_, err = f.WriteString(strP)

	_, err = f.WriteString("module.exports = router;\n")
	fmt.Printf("Login: Login has been created\n")
}
