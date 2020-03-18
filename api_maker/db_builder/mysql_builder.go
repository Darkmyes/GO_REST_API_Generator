package db_builder

import (
	"fmt"
	"os"

	"../models"
	"../tools"
)

func MySQLMakeCode(P models.Project) {
	err := os.MkdirAll("./db-build/sql/mysql", os.ModePerm)
	tools.Check(err)

	f, err := os.Create("./db-build/sql/mysql/" + P.Name + ".sql")
	tools.Check(err)
	defer f.Close()

	_, err = f.WriteString(buildDataBase(P))
	_, err = f.WriteString(buildTables(P.Objects))

}

func buildDataBase(P models.Project) string {
	strP := tools.ProcessFile("templates/mysql/Database.tmpl", P)
	return strP
}

func buildTables(Objs []models.Object) string {
	strP := ""

	for _, O := range Objs {
		strP += tools.ProcessFile("templates/mysql/Table.tmpl", O)
		fmt.Println("SQL: " + O.Name + " has been written")
	}

	return strP
}
