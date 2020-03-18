package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"./api_maker"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

/* func CreateServer() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public/"))))
	port := ":8080"
	openBrowser("http://localhost" + port)
	http.ListenAndServe(port, nil)
} */

func unquoteCodePoint(s string) (string, error) {
	r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	return string(r), err
}

func ShowInfo() {
	alien_emoji, err := unquoteCodePoint("\\U0001F47E")
	if err != nil {
		alien_emoji = ""
	}
	check_emoji, err := unquoteCodePoint("\\U00002714")
	if err != nil {
		check_emoji = "-"
	}
	sparkle_emoji, err := unquoteCodePoint("\\U00002728")
	if err != nil {
		check_emoji = "-"
	}
	fmt.Println("\t" + alien_emoji + " GO_REST_API_Generator " + alien_emoji)
	fmt.Println(sparkle_emoji + " This is a REST API generator based in GO language that can generate code in NodeJS, PHP and GO.")
	fmt.Println(sparkle_emoji + " Can connect to MySQL and SQLServer Databases and generate a json file with the basic information of the DB and tables Schema")
	fmt.Println(sparkle_emoji + " Supported Languages with DB Connection:")
	fmt.Println("\t" + check_emoji + " PHP: MySQL")
	fmt.Println("\t" + check_emoji + " NodeJS: MySQL - MSSQL")
	fmt.Println("\t" + check_emoji + " Go: MySQL")
	fmt.Println(sparkle_emoji + " Flags Examples:")
	fmt.Println("\t" + check_emoji + " go run main.go -mode=gen -folder=public -lang=nodejs")
	fmt.Println("\t" + check_emoji + " go run main.go -mode=db_read -db=mysql -server=localhost -user=root -pass=Password -db_name=DB -folder=project_folder")
	fmt.Println("\t" + check_emoji + " go run main.go -mode=db_read_gen -db=mysql -server=localhost -user=root -pass=Password -db_name=DB -folder=project_folder")
}

func validateDBVars(db_name, user, db_server string) error {
	valid := true
	message := ""
	if db_name == " " {
		valid = false
		message += "- Error: No Database Name provided\n"
	}
	if user == " " {
		valid = false
		message += "- Error: No Database User provided\n"
	}
	if db_server == " " {
		valid = false
		message += "- Error: No Database URL provided\n"
	}

	if !valid {
		return errors.New(message)
	}

	return nil
}

func main() {
	var mode, folderPath, db_type, db_name, user, pass, port, lang, db_server string

	flag.StringVar(&mode, "mode", "info",
		"Generator mode :(gen, db_read, db_read_gen)\n"+
			"gen : Generate the REST API from the project.json file in the given folder\n"+
			"db_read : Read the given DB and generate the project.json file in the given folder\n"+
			"db_read_gen : Read the given DB and generate the project.json file and the REST API in the given folder\n",
	)
	flag.StringVar(&folderPath, "folder", "project", "Folder Name of the Project to generate\n")

	flag.StringVar(&db_type, "db", "MYSQL", "DB Type (mysql, mssql)\n")
	flag.StringVar(&db_server, "server", "localhost", "DB Server URL\n")
	flag.StringVar(&user, "user", " ", "DB User\n")
	flag.StringVar(&pass, "pass", " ", "DB Password\n")
	flag.StringVar(&db_name, "db_name", " ", "DB Name\n")
	flag.StringVar(&port, "port", " ", "DB Port\n")
	flag.StringVar(&lang, "lang", "php", "Server Language (go, nodejs, php)\n")

	flag.Parse()

	switch mode {
	/* case "server":
	CreateServer() */
	case "info":
		ShowInfo()
		break
	case "gen":
		api_maker.GenerateFromJSON(folderPath, strings.ToUpper(lang))
		break
	case "db_read":
		valErr := validateDBVars(db_name, user, db_server)
		if valErr != nil {
			fmt.Println(valErr)
			break
		}
		api_maker.ReadFromDB(folderPath, strings.ToUpper(db_type), db_server, db_name, user, pass, port)
		break
	case "db_read_gen":
		valErr := validateDBVars(db_name, user, db_server)
		if valErr != nil {
			fmt.Println(valErr)
			break
		}
		api_maker.ReadGenerateFromBD(folderPath, strings.ToUpper(db_type), db_server, db_name, user, pass, port, lang)
		break
	default:
		fmt.Println("- Error: Bad mode given")
		break
	}

}
