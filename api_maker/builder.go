package api_maker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"./db_builder"
	"./db_reader"
	"./models"
	"./tools"
)

//func MakeRestAPI(P Project) {
func MakeRestAPI(P models.Project) {
	switch strings.ToUpper(P.Lang) {
	case "PHP":
		//PHPMakeCode(P, Objs)
		JWTPHPMakeCode(P)
		break
	case "NODEJS":
		JWTNodeMakeCode(P)
		break
	case "GO":
		JWTGoMakeCode(P)
		break
	case "JAVA":
		//JAVAMakeCode(getObjects(P.Id))
		break
	default:
		//Some Stuff
		break
	}

	// Generate SQL Files
	if P.GenerateSQL {
		db_builder.GenerateDB(P)
	}

	// Generate Android
	if P.GenerateAndroid {
		AndroidMakeCode(P)
	}
}

func GenerateFromJSON(folderPath, lang string) {
	data, err := ioutil.ReadFile("./" + folderPath + "/project.json")
	if err != nil {
		fmt.Print(err)
	}

	var p models.Project

	err = json.Unmarshal(data, &p)
	if err != nil {
		fmt.Println("error:", err)
	}

	p.Folder_Path = folderPath
	p.Lang = strings.ToUpper(lang)

	MakeRestAPI(p)
}

func ReadFromDB(folderPath, dbtype, server, dbname, user, pass, port string) {
	db_reader.ImportTablesToJSON(folderPath, dbtype, server, dbname, user, pass, port)
}

func ReadGenerateFromBD(folderPath, dbtype, server, dbname, user, pass, port, lang string) {
	db_reader.ImportTablesToJSON(folderPath, dbtype, server, dbname, user, pass, port)
	GenerateFromJSON(folderPath, lang)
}

func ReadGenerateFromBD2(folderPath, dbtype, server, dbname, user, pass, port string) {
	db_data := models.DbData{
		Db_Type: dbtype,
		Db_Url:  server,
		Db_Port: port,
		Db_Name: dbname,
		Db_User: user,
		Db_Pass: pass,
	}
	objs := db_reader.ReadObjects2(db_data)
	p := models.Project{
		Id:              0,
		Name:            dbname,
		Desc:            "",
		Db:              dbtype,
		Db_Data:         db_data,
		Lang:            "NODEJS",
		Folder_Path:     folderPath,
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

	MakeRestAPI(p)
}
