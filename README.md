# 👾 GO_REST_API_Generator 👾

✨ This is a REST API generator based in GO language that can generate code in NodeJS, PHP and GO.

✨ Can connect to MySQL and SQLServer Databases and generate a json file with the basic information of the DB and tables Schema.

✨ Supported Languages with DB Connection:

| LANGUAGES | MySQL   |  MSSQL  |
| --------- | :-----: | :-----: |
| PHP       |   ✔    |  ❌    |
| NODEJS    |   ✔    |  ✔    |
| GO        |   ✔    |  ❌    |
        
## ✨ Flags Examples:

        ✔ go run main.go -mode=gen -folder=public -lang=nodejs
        
        ✔ go run main.go -mode=db_read -db=mysql -server=localhost -user=root -pass=Password -db_name=DB -folder=project_folder   
        
        ✔ go run main.go -mode=db_read_gen -db=mysql -server=localhost -user=root -pass=Password -db_name=DB -folder=project_folder

## ✨Flags Options

Use ```go run main.go -h``` or ```go run main.go --help```

  🚩 -db: string DB Type **(mysql, mssql)** (default "MYSQL")
         
  🚩 -db_name: string DB Name (default " ")
         
  🚩 -folder: string Folder Name of the Project to generate (default "project")
         
  🚩 -lang: string
        Server Language **(go, nodejs, php)**
         (default "php")
         
  🚩 -mode: string Generator mode: **(gen, db_read, db_read_gen)**
  
        * gen : Generate the REST API from the project.json file in the given folder
        * db_read : Read the given DB and generate the project.json file in the given folder
        * db_read_gen : Read the given DB and generate the project.json file and the REST API in the given folder
         (default "info")
         
  🚩 -pass: string DB Password (default " ")
         
  🚩 -port: string DB Port (default " ")
         
  🚩 -server: string DB Server URL (default "localhost")
         
         
 ## ✨Tasks for Future Versions

📌 Migrate all the generator logic to text/template go package

📌 Generate the REST APIS with ORMs in the code

📌 Support PostgresSQL and SQLite
