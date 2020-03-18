package api_maker

import (
	"fmt"
	"os"
	"strings"

	"./models"
)

func JWTGoMakeCode(P models.Project) {
	buildPath := "go_build"
	err := os.MkdirAll("./"+P.Folder_Path+"/"+buildPath+"/models", os.ModePerm)
	check(err)

	err = os.MkdirAll("./"+P.Folder_Path+"/"+buildPath+"/api", os.ModePerm)
	check(err)

	buildGoMain(buildPath, P, P.Objects)
	buildGoConnection(buildPath, P)
	buildGoStuffs(P.Folder_Path, buildPath)
	buildGoModels(P.Folder_Path, buildPath, P.Objects, P.Db)
	//buildJWTGoConfigFile()
	buildGoAPIEndPoints(P.Folder_Path, buildPath, P.Objects)
}

func buildGoMain(buildPath string, P models.Project, Objs []models.Object) {
	f, err := os.Create("./" + P.Folder_Path + "/" + buildPath + "/server.go")
	check(err)
	defer f.Close()

	_, err = f.WriteString("package main\n\n")

	_, err = f.WriteString("import (\n")
	_, err = f.WriteString("\t" + `"fmt"` + "\n")
	_, err = f.WriteString("\t" + `"log"` + "\n")
	//_, err = f.WriteString("\t" + `"database/sql"` + "\n")
	_, err = f.WriteString("\t" + `"net/http"` + "\n")
	_, err = f.WriteString("\t" + `"./api"` + "\n")
	_, err = f.WriteString("\t" + `"./models"` + "\n")
	_, err = f.WriteString("\t" + `"github.com/gorilla/mux"` + "\n")
	_, err = f.WriteString(")\n\n")

	_, err = f.WriteString("func main() {\n")

	//_, err = f.WriteString("\tdb, err := models.NewDB(" + `"root:Qweasdzxc1234@tcp(127.0.0.1:3307)/restaurantdb2?parseTime=true"` + ")\n")
	_, err = f.WriteString("\tdb, err := models.NewDB(" + `"` + P.Db_Data.Db_User + ":" + P.Db_Data.Db_Pass + "@" + "tcp(" + P.Db_Data.Db_Url + ":" + P.Db_Data.Db_Port + ")/" + P.Db_Data.Db_Name + "?allowOldPasswords=true&parseTime=true" + `"` + ")\n")
	_, err = f.WriteString("\tif err != nil {\n")
	_, err = f.WriteString("\t\tlog.Panic(err)\n")
	_, err = f.WriteString("\t}\n")

	_, err = f.WriteString("\tr := mux.NewRouter()\n\n")
	//	api.Find
	for _, O := range Objs {

		if O.Is_Login {
			_, err = f.WriteString("\tr.Handle(" + `"/login"` + ", api.Login(db)).Methods(" + `"POST"` + ")  \n")
		}

		_, err = f.WriteString("\tr.Handle(" + `"/` + O.Name + `"` + ", api.Find" + strings.Title(O.Name) + "(db)).Methods(" + `"GET"` + ")  \n")
		_, err = f.WriteString("\tr.Handle(" + `"/` + O.Name + `"` + ", api.Insert" + strings.Title(O.Name) + "(db)).Methods(" + `"POST"` + ")  \n")
		_, err = f.WriteString("\tr.Handle(" + `"/` + O.Name + `"` + ", api.Update" + strings.Title(O.Name) + "(db)).Methods(" + `"PUT"` + ")  \n")
		_, err = f.WriteString("\tr.Handle(" + `"/` + O.Name + `"` + ", api.Delete" + strings.Title(O.Name) + "(db)).Methods(" + `"DELETE"` + ")  \n")
		_, err = f.WriteString("\n")
	}

	_, err = f.WriteString("\tfmt.Println(" + `"EL SERVIDOR ESTA CORRIENDO EN EL PUERTO 8000"` + ")\n")
	_, err = f.WriteString("\tfmt.Println(http.ListenAndServe(" + `":8000"` + ", r))\n")

	_, err = f.WriteString("}\n")
}

func buildGoConnection(buildPath string, P models.Project) {
	f, err := os.Create("./" + P.Folder_Path + "/" + buildPath + "/models/db.go")
	check(err)
	defer f.Close()

	switch P.Db {
	case "MYSQL":
		_, err = f.WriteString(strGoMysqlConnection())
	case "POSTGRESQL":
		_, err = f.WriteString(strGoMongoDBConnection(P.Folder_Path, P.Db_Data.Db_Name, P.Db_Data.DB_Collate))
	case "SQLSERVER":
		_, err = f.WriteString(" -- ")
	}

	fmt.Printf("Model: DB has been created\n")
}

func strGoMysqlConnection() string {
	strP := ""

	strP += "package models\n\n"

	strP += "import (\n"
	strP += "\t" + `"database/sql"` + "\n"
	strP += "\t" + `_ "github.com/go-sql-driver/mysql"` + "\n"
	strP += ")\n\n"

	strP += "func NewDB(dataSourceName string) (*sql.DB, error) {\n"
	strP += "\tdb, err := sql.Open(" + `"mysql"` + ", dataSourceName)\n"
	strP += "\tif err != nil {\n"
	strP += "\t\treturn nil, err\n"
	strP += "\t}\n"
	strP += "\tif err = db.Ping(); err != nil {\n"
	strP += "\t\treturn nil, err\n"
	strP += "\t}\n"
	strP += "\treturn db, nil\n"
	strP += "}\n"

	return strP
}

func strGoMongoDBConnection(Folder_Path, DBname, DBcoll string) string {
	strP := ""

	strP += "package models\n\n"

	strP += "import (\n"
	strP += "\t" + `"context"` + "\n"
	strP += "\t" + `"go.mongodb.org/mongo-driver/mongo"` + "\n"
	strP += "\t" + `"go.mongodb.org/mongo-driver/mongo/options"` + "\n"
	strP += "\t" + `"go.mongodb.org/mongo-driver/mongo/readpref"` + "\n"
	strP += ")\n\n"

	strP += "func GetDatabaseName string {\n"
	strP += "\treturn " + `"` + DBname + `"` + "\n"
	strP += "}\n\n"

	strP += "func GetCollectionName string {\n"
	strP += "\treturn " + `"` + DBcoll + `"` + "\n"
	strP += "}\n\n"

	strP += "func GetClient(dataSourceName string) *mongo.Client, error {\n"
	strP += "\tclientOptions := options.Client().ApplyURI(dataSourceName)\n"
	strP += "\tclient, err := mongo.NewClient(clientOptions)\n"
	strP += "\tif err != nil {\n"
	strP += "\t\treturn nil, err\n"
	strP += "\t}\n"
	strP += "\terr = client.Connect(context.Background())\n"
	strP += "\tif err != nil {\n"
	strP += "\t\treturn nil, err\n"
	strP += "\t}\n"
	strP += "\t	err = client.Ping(context.Background(), readpref.Primary())\n"
	strP += "\tif err != nil {\n"
	strP += "\t\treturn nil, err\n"
	strP += "\t}\n\n"

	strP += "\treturn client, nil\n"
	strP += "}\n"

	return strP
}

func buildGoStuffs(Folder_Path, buildPath string) {
	f, err := os.Create("./" + Folder_Path + "/" + buildPath + "/models/stuffs.go")
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

func buildGoModels(Folder_Path, buildPath string, Objs []models.Object, db string) {
	for _, O := range Objs {
		buildGoModel(Folder_Path+"/"+buildPath, O, db)
	}
}

func buildGoModel(buildPath string, O models.Object, db string) {
	strpath := "./" + buildPath + "/models/" + O.Name + ".go"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	_, err = f.WriteString("package models\n\n")
	_, err = f.WriteString("import (\n")
	if strings.ToUpper(db) == "MONGODB" {
		_, err = f.WriteString("\t" + `"go.mongodb.org/mongo-driver/bson"` + "\n")
		_, err = f.WriteString("\t" + `"go.mongodb.org/mongo-driver/mongo"` + "\n")
		_, err = f.WriteString("\t" + `"go.mongodb.org/mongo-driver/mongo/options"` + "\n")
		_, err = f.WriteString("\t" + `"go.mongodb.org/mongo-driver/mongo/readpref"` + "\n")
	} else {
		_, err = f.WriteString("\t" + `"database/sql"` + "\n")
	}
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

		if strings.ToUpper(db) == "MONGODB" {
			_, err = f.WriteString(` bson:"` + C.Name + `"`)
		}

		_, err = f.Write([]byte{'\u0060'})
		_, err = f.WriteString("\n")
	}

	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString(strGoLogin(O))

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") Insert (db *sql.DB) error {\n")
	if strings.ToUpper(db) == "MONGODB" {
		_, err = f.WriteString(strGoMONGODBInsert(O))
	} else {
		_, err = f.WriteString(strGoInsert(O))
	}
	_, err = f.WriteString("\treturn nil\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") Update (db *sql.DB) error {\n")
	if strings.ToUpper(db) == "MONGODB" {
		_, err = f.WriteString(strGoMONGODBUpdate(O))
	} else {
		_, err = f.WriteString(strGoUpdate(O))
	}
	_, err = f.WriteString("\treturn nil\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") Delete (db *sql.DB) error {\n")
	if strings.ToUpper(db) == "MONGODB" {
		_, err = f.WriteString(strGoMONGODBDelete(O))
	} else {
		_, err = f.WriteString(strGoDelete(O))
	}
	_, err = f.WriteString("\treturn nil\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") FindOne (db *sql.DB) (*" + strings.Title(O.Name) + ", error) {\n")
	_, err = f.WriteString("\tobj := new(" + strings.Title(O.Name) + ")\n")
	if strings.ToUpper(db) == "MONGODB" {
		_, err = f.WriteString(strGoMONGODBFindOne(O))
	} else {
		_, err = f.WriteString(strGoFindOne(O))
	}
	_, err = f.WriteString("\treturn nil, nil\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") FindAll (db *sql.DB) ([]*" + strings.Title(O.Name) + ", error) {\n")
	_, err = f.WriteString("\tobjs := make([]*" + strings.Title(O.Name) + ", 0)\n")
	if strings.ToUpper(db) == "MONGODB" {
		_, err = f.WriteString(strGoMONGODBFindAll(O))
	} else {
		_, err = f.WriteString(strGoFindAll(O))
	}
	_, err = f.WriteString("\treturn objs, nil\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func (o " + strings.Title(O.Name) + ") FindAllPaginated (db *sql.DB, index int, nums int) ([]*" + strings.Title(O.Name) + ", error) {\n")
	_, err = f.WriteString("\tobjs := make([]*" + strings.Title(O.Name) + ", 0)\n")
	if strings.ToUpper(db) == "MONGODB" {
		_, err = f.WriteString(strGoMONGODBFindAllPag(O))
	} else {
		_, err = f.WriteString(strGoFindAllPag(O))
	}
	_, err = f.WriteString("\treturn objs, nil\n")
	_, err = f.WriteString("}\n\n")

	fmt.Printf("Model: " + O.Name + " has been created\n")
}

func strGoLogin(O models.Object) string {
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

func strGoInsert(O models.Object) string {
	strP := "\tstrSQL := " + `"` + "INSERT INTO " + O.Tbl_Name + " ("
	strQ := ""
	strW := ""

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
			strQ += "?"
			strW += "o." + strings.Title(O.Columns[i].Col_Name)
		} else {
			strP += ", " + O.Columns[i].Col_Name
			strQ += ", ?"
			strW += ", o." + strings.Title(O.Columns[i].Col_Name)
		}
	}

	strP += ") VALUES (" + strQ + ")" + `"` + ";\n"

	strP += "\tstmtInsert, err := db.Prepare(strSQL)\n\n"
	strP += "\tif err != nil { return err }\n"
	strP += "\tdefer stmtInsert.Close()\n\n"

	strP += "\t_, err = stmtInsert.Exec(" + strW + ")\n"
	strP += "\tif err != nil { return err }\n\n"

	return strP
}

func strGoMONGODBInsert(O models.Object) string {
	strP := ""

	strP += "\tcollection := client.Database(GetDatabaseName()).Collection(GetCollectionName())\n"
	strP += "\t/*insertResult*/_, err := collection.InsertOne(context.TODO(), o)\n"
	strP += "\tif err != nil {\n"
	strP += "\t\treturn err\n"
	strP += "\t}\n\n"

	strP += "\t//return insertResult.InsertedID\n"

	return strP
}

func strGoUpdate(O models.Object) string {
	strP := "\tstrSQL := " + `"` + "UPDATE " + O.Tbl_Name + " SET "
	strQ := ""

	count := 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key != true {
			if count == 0 {
				strP += O.Columns[i].Col_Name + " = ?"
				strQ += "o." + strings.Title(O.Columns[i].Col_Name)
				count = 1
			} else {
				strP += ", " + O.Columns[i].Col_Name + " = ?"
				strQ += ", o." + strings.Title(O.Columns[i].Col_Name)
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
	strP += "\tstmtInsert, err := db.Prepare(strSQL)\n\n"
	strP += "\tif err != nil { return err }\n"
	strP += "\tdefer stmtInsert.Close()\n\n"

	strP += "\t_, err = stmtInsert.Exec(" + strQ + ")\n"
	strP += "\tif err != nil { return err }\n\n"

	return strP
}

func strGoMONGODBUpdate(O models.Object) string {
	strP := ""

	strP += "\tcollection := client.Database(GetDatabaseName()).Collection(GetCollectionName())\n"

	strP += "\tupdatedData := bson.D{ {Key: " + `"$set"` + ", Value: o} }\n"
	strP += "\t/*updatedResult*/, err := collection.UpdateOne(context.TODO(), bson.M{"

	count := 0
	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += ` "` + O.Columns[i].Col_Name + `" : ` + "o." + strings.Title(O.Columns[i].Name)
				count++
			} else {
				strP += `, "` + O.Columns[i].Col_Name + `" : ` + "o." + strings.Title(O.Columns[i].Name)
			}
		}
	}

	strP += "} , updatedData)\n"
	strP += "\tif err != nil {\n"
	strP += "\t\treturn err\n"
	strP += "\t}\n\n"

	strP += "\t//return updatedResult.ModifiedCount"

	return strP
}

func strGoDelete(O models.Object) string {
	strP := "\tstrSQL := " + `"` + "DELETE FROM " + O.Tbl_Name + " WHERE "
	strQ := ""

	count := 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += O.Columns[i].Col_Name + " = ?"
				strQ += "o." + strings.Title(O.Columns[i].Col_Name)
				count++
			} else {
				strP += " AND " + O.Columns[i].Col_Name + " = ?"
				strQ += ", o." + strings.Title(O.Columns[i].Col_Name)
			}
		}
	}

	strP += `"` + ";\n"
	strP += "\tstmtInsert, err := db.Prepare(strSQL)\n\n"
	strP += "\tif err != nil { return err }\n"
	strP += "\tdefer stmtInsert.Close()\n\n"

	strP += "\t_, err = stmtInsert.Exec(" + strQ + ")\n"
	strP += "\tif err != nil { return err }\n\n"

	return strP
}

func strGoMONGODBDelete(O models.Object) string {
	strP := ""

	strP += "\tcollection := client.Database(GetDatabaseName()).Collection(GetCollectionName())\n"

	strP += "\t/*deleteResult*/, err := collection.DeleteOne(context.TODO(), bson.M{"

	count := 0
	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += ` "` + O.Columns[i].Col_Name + `" : ` + "o." + strings.Title(O.Columns[i].Name)
				count++
			} else {
				strP += `, "` + O.Columns[i].Col_Name + `" : ` + "o." + strings.Title(O.Columns[i].Name)
			}
		}
	}

	strP += "} )\n"
	strP += "\tif err != nil {\n"
	strP += "\t\treturn err\n"
	strP += "\t}\n\n"

	strP += "\t//return deleteResult.DeletedCount"

	return strP
}

func strGoFindOne(O models.Object) string {
	strP := "\tstrSQL := " + `"` + "SELECT "
	strQ := ""
	strW := ""

	count := 0

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
			strW += "&obj." + strings.Title(O.Columns[i].Col_Name)
			count = 1
		} else {
			strP += ", " + O.Columns[i].Col_Name
			strW += ", &obj." + strings.Title(O.Columns[i].Col_Name)
		}
	}

	strP += " FROM " + O.Tbl_Name + " WHERE "
	count = 0

	for i, _ := range O.Columns {
		if O.Columns[i].Primary_Key == true {
			if count == 0 {
				strP += O.Columns[i].Col_Name + " = ?"
				strQ += "o." + strings.Title(O.Columns[i].Col_Name)
				count++
			} else {
				strP += " AND " + O.Columns[i].Col_Name + " = ?"
				strQ += ", o." + strings.Title(O.Columns[i].Col_Name)
			}
		}
	}

	strP += `"` + ";\n"

	strP += "\trows, err := db.Query(strSQL, " + strQ + ")\n\n"
	strP += "\tif err != nil { return nil, err }\n"
	strP += "\tdefer rows.Close()\n\n"

	strP += "\tif rows.Next() {\n"
	strP += "\t\terr := rows.Scan(" + strW + ")\n"
	strP += "\t\tif err != nil { return nil, err }\n"
	strP += "\t\treturn obj, nil\n"
	strP += "\t}\n"
	strP += "\tif err = rows.Err(); err != nil { return nil, err }\n"

	return strP
}

func strGoMONGODBFindOne(O models.Object) string {
	strP := ""

	strP += "\tcollection := client.Database(GetDatabaseName()).Collection(GetCollectionName())\n"

	strP += "\tcur := collection.FindOne(context.TODO(), bson.M{})\n"
	strP += "\tif err != nil { return nil, err }\n"

	strP += "\tobj := new(" + strings.Title(O.Name) + ")\n"
	strP += "\terr = cur.Decode(&obj)\n"
	strP += "\tif err != nil { return nil, err }\n"

	strP += "\treturn obj, nil"

	return strP
}

func strGoFindAll(O models.Object) string {
	strP := "\tstrSQL := " + `"` + "SELECT "
	strW := ""

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
			strW += "&obj." + strings.Title(O.Columns[i].Col_Name)
		} else {
			strP += ", " + O.Columns[i].Col_Name
			strW += ", &obj." + strings.Title(O.Columns[i].Col_Name)
		}
	}

	strP += " FROM " + O.Tbl_Name + `"` + ";\n"

	strP += "\trows, err := db.Query(strSQL)\n\n"
	strP += "\tif err != nil { return nil, err }\n"
	strP += "\tdefer rows.Close()\n\n"

	strP += "\tfor rows.Next() {\n"
	strP += "\t\tobj := new(" + strings.Title(O.Name) + ")\n"
	strP += "\t\terr := rows.Scan(" + strW + ")\n"
	strP += "\t\tif err != nil { return nil, err }\n"
	strP += "\t\tobjs = append(objs, obj)\n"
	strP += "\t}\n"
	strP += "\tif err = rows.Err(); err != nil { return nil, err }\n"

	return strP
}

func strGoMONGODBFindAll(O models.Object) string {
	strP := ""

	strP += "\tcollection := client.Database(GetDatabaseName()).Collection(GetCollectionName())\n"

	strP += "\tcur, err := collection.Find(context.TODO(), bson.M{})\n"
	strP += "\tif err != nil { return nil, err }\n"

	strP += "\tfor cur.Next(context.TODO()) {"
	strP += "\t\tobj := new(" + strings.Title(O.Name) + ")\n"
	strP += "\t\terr = cur.Decode(&obj)\n"
	strP += "\t\tif err != nil { return nil, err }\n"
	strP += "\t\tobjs = append(objs, obj)\n"
	strP += "\t}\n\n"

	strP += "\treturn objs, nil"

	return strP
}

func strGoFindAllPag(O models.Object) string {
	strP := "\tstrSQL := " + `"` + "SELECT "
	strW := ""

	for i, _ := range O.Columns {
		if i == 0 {
			strP += O.Columns[i].Col_Name
			strW += "&obj." + strings.Title(O.Columns[i].Col_Name)
		} else {
			strP += ", " + O.Columns[i].Col_Name
			strW += ", &obj." + strings.Title(O.Columns[i].Col_Name)
		}
	}

	strP += " FROM " + O.Tbl_Name + " LIMIT ? OFFSET ?" + `"` + ";\n"

	strP += "\trows, err := db.Query(strSQL, index, nums)\n\n"
	strP += "\tif err != nil { return nil, err }\n"
	strP += "\tdefer rows.Close()\n\n"

	strP += "\tfor rows.Next() {\n"
	strP += "\t\tobj := new(" + strings.Title(O.Name) + ")\n"
	strP += "\t\terr := rows.Scan(" + strW + ")\n"
	strP += "\t\tif err != nil { return nil, err }\n"
	strP += "\t\tobjs = append(objs, obj)\n"
	strP += "\t}\n"
	strP += "\tif err = rows.Err(); err != nil { return nil, err }\n"

	return strP
}

func strGoMONGODBFindAllPag(O models.Object) string {
	strP := ""

	strP += "\tcollection := client.Database(GetDatabaseName()).Collection(GetCollectionName())\n"

	strP += "\tcur, err := collection.Find(context.TODO(), bson.M{})\n"
	strP += "\tif err != nil { return nil, err }\n"

	strP += "\tfor cur.Next(context.TODO()) {"
	strP += "\t\tobj := new(" + strings.Title(O.Name) + ")\n"
	strP += "\t\terr = cur.Decode(&obj)\n"
	strP += "\t\tif err != nil { return nil, err }\n"
	strP += "\t\tobjs = append(objs, obj)\n"
	strP += "\t}\n\n"

	strP += "\treturn objs, nil"

	return strP
}

func buildGoAPIEndPoints(Folder_Path, buildPath string, Objs []models.Object) {
	for _, O := range Objs {
		buildGoAPIEndPoint(Folder_Path, buildPath, O)
		if O.Is_Login {
			buildGoAPILogin(Folder_Path, buildPath, O)
		}
	}
}

func buildGoAPILogin(Folder_Path, buildPath string, O models.Object) {
	f, err := os.Create("./" + Folder_Path + "/" + buildPath + "/api/" + O.Name + "Login.go")
	check(err)
	defer f.Close()

	_, err = f.WriteString("package api\n\n")

	_, err = f.WriteString("import (\n")
	_, err = f.WriteString("\t" + `"encoding/json"` + "\n")
	_, err = f.WriteString("\t" + `"time"` + "\n")
	_, err = f.WriteString("\t" + `"net/http"` + "\n")
	_, err = f.WriteString("\t" + `"github.com/dgrijalva/jwt-go"` + "\n")
	_, err = f.WriteString("\t" + `"database/sql"` + "\n")
	_, err = f.WriteString(")\n\n")

	_, err = f.WriteString("var jwtKey = []byte(" + `"` + "my_secret_key" + `"` + ")\n\n")

	_, err = f.WriteString("type Credentials struct {\n")
	_, err = f.WriteString("\tPassword string ")
	_, err = f.Write([]byte{'\u0060'})
	_, err = f.WriteString(`json:"password"`)
	_, err = f.Write([]byte{'\u0060'})
	_, err = f.WriteString("\n")
	_, err = f.WriteString("\tUsername string ")
	_, err = f.Write([]byte{'\u0060'})
	_, err = f.WriteString(`json:"username"`)
	_, err = f.Write([]byte{'\u0060'})
	_, err = f.WriteString("\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("type Claims struct {")
	_, err = f.WriteString("\tUsername string ")
	_, err = f.Write([]byte{'\u0060'})
	_, err = f.WriteString(`json:"username"`)
	_, err = f.Write([]byte{'\u0060'})
	_, err = f.WriteString("\n")
	_, err = f.WriteString("\tjwt.StandardClaims\n")
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func Login (db *sql.DB) http.Handler {\n")
	_, err = f.WriteString("\treturn http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {\n")

	_, err = f.WriteString("\t\tvar creds Credentials\n")
	_, err = f.WriteString("\t\terr := json.NewDecoder(r.Body).Decode(&creds)\n")
	_, err = f.WriteString("\t\tif err != nil {\n")
	_, err = f.WriteString("\t\t\tw.WriteHeader(http.StatusBadRequest)\n")
	_, err = f.WriteString("\t\t\treturn\n")
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t\texpirationTime := time.Now().Add(2 * time.Hour)\n")
	_, err = f.WriteString("\t\tclaims := &Claims{\n")
	_, err = f.WriteString("\t\t\tUsername: creds.Username,\n")
	_, err = f.WriteString("\t\t\tStandardClaims: jwt.StandardClaims{\n")
	_, err = f.WriteString("\t\t\t\tExpiresAt: expirationTime.Unix(),\n")
	_, err = f.WriteString("\t\t\t},\n")
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t\ttoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)\n")
	_, err = f.WriteString("\t\ttokenString, err := token.SignedString(jwtKey)\n\n")

	_, err = f.WriteString("\t\tif err != nil {\n")
	_, err = f.WriteString("\t\t\tw.WriteHeader(http.StatusInternalServerError)\n")
	_, err = f.WriteString("\t\t\treturn\n")
	_, err = f.WriteString("\t\t}\n\n")

	_, err = f.WriteString("\t\thttp.SetCookie(w, &http.Cookie{\n")
	_, err = f.WriteString("\t\t\tName: " + `"token"` + ",\n")
	_, err = f.WriteString("\t\t\tValue: tokenString,\n")
	_, err = f.WriteString("\t\t\tExpires: expirationTime,\n")
	_, err = f.WriteString("\t\t})\n")

	_, err = f.WriteString("\t})\n")
	_, err = f.WriteString("}\n\n")

	// Verify Token

	_, err = f.WriteString("func VerifyToken(w http.ResponseWriter, r *http.Request) (*Claims, error) {\n")

	_, err = f.WriteString("\tc, err := r.Cookie(" + `"token"` + ")\n")

	_, err = f.WriteString("\tif err != nil {\n")
	_, err = f.WriteString("\t\tif err == http.ErrNoCookie {\n")
	_, err = f.WriteString("\t\t\tw.WriteHeader(http.StatusUnauthorized)\n")
	_, err = f.WriteString("\t\t\treturn nil, err\n")
	_, err = f.WriteString("\t\t}\n")
	_, err = f.WriteString("\t\tw.WriteHeader(http.StatusBadRequest)\n")
	_, err = f.WriteString("\t\treturn nil, err\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\ttknStr := c.Value\n")
	_, err = f.WriteString("\tclaims := &Claims{}\n\n")

	_, err = f.WriteString("\ttkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {\n")
	_, err = f.WriteString("\t\treturn jwtKey, nil\n")
	_, err = f.WriteString("\t})\n\n")

	_, err = f.WriteString("\tif err != nil {\n")
	_, err = f.WriteString("\t\tif err == jwt.ErrSignatureInvalid {\n")
	_, err = f.WriteString("\t\t\tw.WriteHeader(http.StatusUnauthorized)\n")
	_, err = f.WriteString("\t\t\treturn nil, err\n")
	_, err = f.WriteString("\t\t}\n")
	_, err = f.WriteString("\t\tw.WriteHeader(http.StatusBadRequest)\n")
	_, err = f.WriteString("\t\treturn nil, err\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\tif !tkn.Valid {\n")
	_, err = f.WriteString("\t\tw.WriteHeader(http.StatusUnauthorized)\n")
	_, err = f.WriteString("\t\treturn nil, err\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\treturn claims, nil\n")

	_, err = f.WriteString("}\n\n")

	// Refresh Token

	_, err = f.WriteString("func RefreshToken(w http.ResponseWriter, r *http.Request) {\n")

	_, err = f.WriteString("\tclaims, er := VerifyToken(w, r)\n\n")
	_, err = f.WriteString("\tif er != nil {\n")
	_, err = f.WriteString("\t\treturn\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\tif time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {\n")
	_, err = f.WriteString("\t\tw.WriteHeader(http.StatusBadRequest)\n")
	_, err = f.WriteString("\t\treturn\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\texpirationTime := time.Now().Add(2 * time.Hour)\n")
	_, err = f.WriteString("\tclaims.ExpiresAt = expirationTime.Unix()\n")
	_, err = f.WriteString("\ttoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)\n")
	_, err = f.WriteString("\ttokenString, err := token.SignedString(jwtKey)\n\n")

	_, err = f.WriteString("\tif err != nil {\n")
	_, err = f.WriteString("\t\tw.WriteHeader(http.StatusInternalServerError)\n")
	_, err = f.WriteString("\t\treturn\n")
	_, err = f.WriteString("\t}\n\n")

	_, err = f.WriteString("\thttp.SetCookie(w, &http.Cookie{\n")
	_, err = f.WriteString("\t\tName: " + `"token"` + ",\n")
	_, err = f.WriteString("\t\tValue: tokenString,\n")
	_, err = f.WriteString("\t\tExpires: expirationTime,\n")
	_, err = f.WriteString("\t})\n")

	_, err = f.WriteString("}\n\n")
}

func buildGoAPIEndPoint(Folder_Path, buildPath string, O models.Object) {
	f, err := os.Create("./" + Folder_Path + "/" + buildPath + "/api/" + O.Name + "EP.go")
	check(err)
	defer f.Close()

	_, err = f.WriteString("package api\n\n")

	_, err = f.WriteString("import (\n")
	_, err = f.WriteString("\t" + `"encoding/json"` + "\n")
	//_, err = f.WriteString("\t" + `"fmt"` + "\n")
	_, err = f.WriteString("\t" + `"net/http"` + "\n")
	_, err = f.WriteString("\t" + `"../models"` + "\n")
	//_, err = f.WriteString("\t" + `"github.com/gorilla/mux"` + "\n")
	_, err = f.WriteString("\t" + `"database/sql"` + "\n")
	_, err = f.WriteString(")\n\n")

	_, err = f.WriteString("func Find" + strings.Title(O.Name) + " (db *sql.DB) http.Handler {\n")
	_, err = f.WriteString(strGoAPI_GET(O))
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func Insert" + strings.Title(O.Name) + " (db *sql.DB) http.Handler {\n")
	_, err = f.WriteString(strGoAPI_POST(O))
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func Update" + strings.Title(O.Name) + " (db *sql.DB) http.Handler {\n")
	_, err = f.WriteString(strGoAPI_PUT(O))
	_, err = f.WriteString("}\n\n")

	_, err = f.WriteString("func Delete" + strings.Title(O.Name) + " (db *sql.DB) http.Handler {\n")
	_, err = f.WriteString(strGoAPI_DELETE(O))
	_, err = f.WriteString("}\n\n")

	fmt.Printf("API EndPoint: " + O.Name + " has been created\n")
}

func strGoAPI_GETPaginated(O models.Object) string {
	return ""
}

func strGoAPI_GET(O models.Object) string {
	strP := "\treturn http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {\n"

	strP += "\t\t_, er := VerifyToken(w, r)\n\n"
	strP += "\t\tif er != nil {\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n\n"

	strP += "\t\tdecoder := json.NewDecoder(r.Body)\n"
	strP += "\t\tvar obj models." + strings.Title(O.Name) + "\n"
	strP += "\t\terr := decoder.Decode(&obj)\n"

	/*strP += "\t\tif err != nil {\n"
	strP += "\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n"*/

	strP += "\t\tif(err == nil){\n"

	strP += "\t\t\tobj, err := obj.FindOne(db)\n"
	strP += "\t\t\tif err != nil {\n"
	strP += "\t\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\t\treturn\n"
	strP += "\t\t\t}\n"

	strP += "\t\t\tw.Header().Set(" + `"Content-Type"` + ", " + `"application/json"` + ")\n"
	strP += "\t\t\tjson.NewEncoder(w).Encode(obj)\n"

	strP += "\t\t}else {\n"

	//strP += "\t\t\tobjs, err := models." + strings.Title(O.Name) + ".FindAll(db)\n"
	strP += "\t\t\tobj = models." + strings.Title(O.Name) + "{}\n"
	strP += "\t\t\tobjs, err := obj.FindAll(db)\n"
	strP += "\t\t\tif err != nil {\n"
	strP += "\t\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\t\treturn\n"
	strP += "\t\t\t}\n"

	strP += "\t\t\tw.Header().Set(" + `"Content-Type"` + ", " + `"application/json"` + ")\n"
	strP += "\t\t\tjson.NewEncoder(w).Encode(objs)\n"

	strP += "\t\t}\n"
	strP += "\t})\n"

	return strP
}

func strGoAPI_POST(O models.Object) string {
	strP := "\treturn http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {\n"

	strP += "\t\t_, er := VerifyToken(w, r)\n\n"
	strP += "\t\tif er != nil {\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n\n"

	strP += "\t\tdecoder := json.NewDecoder(r.Body)\n"
	strP += "\t\tvar obj models." + strings.Title(O.Name) + "\n"
	strP += "\t\terr := decoder.Decode(&obj)\n"

	strP += "\t\tif err != nil {\n"
	strP += "\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n"

	//strP += "\t\terr = models.InsertarCliente(env.DB, obj)\n"
	strP += "\t\terr = obj.Insert(db)\n"
	strP += "\t\tif err != nil {\n"
	strP += "\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n"

	strP += "\t\tinfo := models.Info{Info: " + `"` + strings.Title(O.Name) + ` Guardado con Exito"` + "}\n"
	strP += "\t\tw.Header().Set(" + `"Content-Type"` + ", " + `"application/json"` + ")\n"
	strP += "\t\tjson.NewEncoder(w).Encode(info)\n"
	strP += "\t})\n"

	return strP
}

func strGoAPI_PUT(O models.Object) string {
	strP := "\treturn http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {\n"

	/*if O.Is_Login {
		strP += "\t\tclaims, er := VerifyToken(w, r)\n\n"
	} else {
		strP += "\t\t_, er := VerifyToken(w, r)\n\n"
	}*/

	strP += "\t\t_, er := VerifyToken(w, r)\n\n"
	strP += "\t\tif er != nil {\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n\n"

	strP += "\t\tdecoder := json.NewDecoder(r.Body)\n"
	strP += "\t\tvar obj models." + strings.Title(O.Name) + "\n"
	strP += "\t\terr := decoder.Decode(&obj)\n"

	strP += "\t\tif err != nil {\n"
	strP += "\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n"

	strP += "\t\tfobj,err2 := obj.FindOne(db)\n"
	strP += "\t\tif err2 != nil || fobj == nil {\n"
	strP += "\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\tinfo := models.Info{Info: " + `"` + strings.Title(O.Name) + ` No Encontrado"` + "}\n"
	strP += "\t\t\tw.Header().Set(" + `"Content-Type"` + ", " + `"application/json"` + ")\n"
	strP += "\t\t\tjson.NewEncoder(w).Encode(info)\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n"

	strP += "\t\terr = obj.Update(db)\n"
	strP += "\t\tif err != nil {\n"
	strP += "\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n"

	strP += "\t\tinfo := models.Info{Info: " + `"` + strings.Title(O.Name) + ` Modificado con Exito"` + "}\n"
	strP += "\t\tw.Header().Set(" + `"Content-Type"` + ", " + `"application/json"` + ")\n"
	strP += "\t\tjson.NewEncoder(w).Encode(info)\n"
	strP += "\t})\n"

	return strP
}

func strGoAPI_DELETE(O models.Object) string {
	strP := "\treturn http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {\n"

	strP += "\t\t_, er := VerifyToken(w, r)\n\n"
	strP += "\t\tif er != nil {\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n\n"

	strP += "\t\tdecoder := json.NewDecoder(r.Body)\n"
	strP += "\t\tvar obj models." + strings.Title(O.Name) + "\n"
	strP += "\t\terr := decoder.Decode(&obj)\n"

	strP += "\t\tif err != nil {\n"
	strP += "\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n"

	strP += "\t\tfobj,err2 := obj.FindOne(db)\n"
	strP += "\t\tif err2 != nil  || fobj == nil {\n"
	strP += "\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\tinfo := models.Info{Info: " + `"` + strings.Title(O.Name) + ` No Encontrado"` + "}\n"
	strP += "\t\t\tw.Header().Set(" + `"Content-Type"` + ", " + `"application/json"` + ")\n"
	strP += "\t\t\tjson.NewEncoder(w).Encode(info)\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n"

	strP += "\t\terr = obj.Delete(db)\n"
	strP += "\t\tif err != nil {\n"
	strP += "\t\t\thttp.Error(w, http.StatusText(500), 500)\n"
	strP += "\t\t\treturn\n"
	strP += "\t\t}\n"

	strP += "\t\tinfo := models.Info{Info: " + `"` + strings.Title(O.Name) + ` Eliminado con Exito"` + "}\n"
	strP += "\t\tw.Header().Set(" + `"Content-Type"` + ", " + `"application/json"` + ")\n"
	strP += "\t\tjson.NewEncoder(w).Encode(info)\n"
	strP += "\t})\n"

	return strP
}
