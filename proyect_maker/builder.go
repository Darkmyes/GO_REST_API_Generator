package proyect_maker

//func MakeRestAPI(P Proyect) {
func MakeRestAPI(P Proyect, Objs []Object) {
	switch P.Lang {
	case "PHP":
		//PHPMakeCode(P, Objs)
		JWTPHPMakeCode(P, Objs)
		break
	case "NODEJS":
		JWTNodeMakeCode(P, Objs)
		break
	case "GO":
		JWTGoMakeCode(P, Objs)
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
		switch P.Db {
		case "MYSQL":
			MySQLMakeCode(P, Objs)
			break
		case "PSQL":
			//PSQLMakeCode(P, Objs)
			break
		case "SQLSERVER":
			//SQLSERVERMakeCode(P, Objs)
			break
		default:
			//Some Stuff
			break
		}

	}

	if P.GenerateAndroid {
		AndroidMakeCode(P, Objs)
	}
}

func getObjects(Id int32) []Object {

	return nil
}
