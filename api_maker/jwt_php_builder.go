package api_maker

import (
	"fmt"
	"os"

	"./models"
)

func JWTPHPMakeCode(P models.Project) {
	err := os.MkdirAll("./"+P.Folder_Path+"/php_build/models", os.ModePerm)
	check(err)

	err = os.MkdirAll("./"+P.Folder_Path+"/php_build/api", os.ModePerm)
	check(err)

	buildConnection(P)
	buildModels(P.Folder_Path, P.Objects)
	copyJWTPackage(P.Folder_Path)
	buildJWTConfigFile(P.Folder_Path)
	buildAPIEndPoints(P.Folder_Path, P.Objects)
}

func buildConnection(P models.Project) {
	f, err := os.Create("./" + P.Folder_Path + "/php_build/connection.php")
	check(err)
	defer f.Close()

	conTemplate := ProcessFile("templates/php/Conexion.tmpl", P)
	_, err = f.WriteString(conTemplate)

	fmt.Printf("Connection has been created\n")
}

func buildModels(folder string, Objs []models.Object) {
	for _, O := range Objs {
		buildModel(folder, O)
	}
}

func buildModel(folder string, O models.Object) {
	strpath := "./" + folder + "/php_build/models/" + O.Name + ".php"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	conTemplate := ProcessFile("templates/php/Model.tmpl", O)
	_, err = f.WriteString(conTemplate)

	fmt.Printf("Model: " + O.Name + " has been created\n")
}

func copyJWTPackage(folder string) {
	err := os.MkdirAll("./"+folder+"/php_build/libs", os.ModePerm)
	check(err)

	err = CopyDir("./libraries/php/php-jwt-master", "./"+folder+"/php_build/libs/php-jwt-master")
	check(err)
}

func buildJWTConfigFile(folder string) {
	err := os.MkdirAll("./"+folder+"/php_build/config", os.ModePerm)
	check(err)

	f, err := os.Create("./" + folder + "/php_build/config/jwt_config.php")
	check(err)
	defer f.Close()

	conTemplate := ProcessFile("templates/php/jwtConfig.tmpl", nil)
	_, err = f.WriteString(conTemplate)

	fmt.Printf("REST API: JWT Config File has been created\n")
}

func buildAPIEndPoints(folder string, Objs []models.Object) {
	for _, O := range Objs {
		buildAPIEndPoint(folder, O)
		if O.Is_Login {
			buildAPILogin(folder, O)
		}
	}
}

func buildAPILogin(folder string, O models.Object) {
	strpath := "./" + folder + "/php_build/api/" + O.Name + "_login.php"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	conTemplate := ProcessFile("templates/php/APIEndPointLogin.tmpl", O)
	_, err = f.WriteString(conTemplate)

	fmt.Printf("Login API EndPoint: " + O.Name + " has been created\n")
}

func buildAPIEndPoint(folder string, O models.Object) {
	strpath := "./" + folder + "/php_build/api/" + O.Name + ".php"
	f, err := os.Create(strpath)
	check(err)
	defer f.Close()

	conTemplate := ProcessFile("templates/php/APIEndPoint.tmpl", O)
	_, err = f.WriteString(conTemplate)

	fmt.Printf("API EndPoint: " + O.Name + " has been created\n")
}
