package db_reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"../models"
	"../tools"
)

func ObjectsPOSTGRESSQL(path, bd, user, pass, port string) []models.Object {
	/* db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	defer db.Close() */
	return nil
}

func ReadObjects(dbType, path, bd, user, pass, port string) []models.Object {
	switch dbType {
	case "MYSQL":
		return readObjectsMySQL(path, bd, user, pass, port)
	case "MSSQL":
		return readObjectsMSSQL(path, bd, user, pass, port)
		/* FALTA IMPLEMENTAR FUNCIONAMIENTO
		case "POSTGRESQL":
			return readmodels.ObjectsPOSTGRESQL(path, bd, user, pass, port)
		*/
	}
	return nil
}

func ReadObjects2(db_data models.DbData) []models.Object {
	switch db_data.Db_Type {
	case "MYSQL":
		return readObjectsMySQL2(db_data)
	case "MSSQL":
		return readObjectsMSSQL2(db_data)
		/* FALTA IMPLEMENTAR FUNCIONAMIENTO
		case "POSTGRESQL":
			return readmodels.ObjectsPOSTGRESQL(path, bd, user, pass, port)
		*/
	}
	return nil
}

func ImportTablesToJSON(folderPath, dbtype, server, dbname, user, pass, port string) {
	if port == " " {
		switch strings.ToUpper(dbtype) {
		case "MYSQL":
			port = "3306"
			break
		case "MSSQL":
			port = "1433"
			break
		}
	}

	db_data := models.DbData{
		Db_Type: dbtype,
		Db_Url:  server,
		Db_Port: port,
		Db_Name: dbname,
		Db_User: user,
		Db_Pass: pass,
	}
	objs := ReadObjects2(db_data)
	p := models.Project{
		Id:              0,
		Name:            dbname,
		Desc:            "",
		Db:              dbtype,
		Db_Data:         db_data,
		Lang:            "NODEJS",
		Folder_Path:     "",
		Init_Date:       "",
		Upd_Date:        "",
		GenerateSQL:     false,
		GenerateAndroid: false,
		AuthMode:        "none",
		Objects:         objs,
	}

	fmt.Println("Total Objects extracted: " + string(len(objs)))

	project_json, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.MkdirAll("./"+folderPath, os.ModePerm)
	tools.Check(err)

	err = ioutil.WriteFile("./"+folderPath+"/project.json", project_json, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

}
