package db_builder

import (
	"strings"

	"../models"
)

func GenerateDB(P models.Project) {
	switch strings.ToUpper(P.Db_Data.Db_Type) {
	case "MYSQL":
		MySQLMakeCode(P)
		break
		/* case "MSSQL":
		buildMSSQLNodeConnection(buildPath, P)*/
	}
}
