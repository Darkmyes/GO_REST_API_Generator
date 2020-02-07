# GOBackendGenerator

Esta es un generador de REST APIs escrito en GO y puede generar código en NodeJS, PHPy GO.

This is a REST API generator based in GO language that can generate code in NodeJS, PHP and GO.

## Estructura del Proyecto / Project Structure

El directorio *public* contiene los archivos .json necesarios para la generación de la API REST.
El directorio *proyect_maker* contiene los archivos de generación del código de la API REST.

The directory *public* contain the .json needed files to generate the main API REST.
The directory *proyect_make* contain the needed code files to generate the main API REST code.

## Estructura de un Proyecto JSON / JSON Project Structue

lang puede ser "NODEJS","PHP" o "GO" y auth_mode "jwt" o "none".

lang can be "NODEJS","PHP" or "GO" and auth_mode can be "jwt" or "none".

```json
{
    "id" : 1,
    "name" : "prueba",
    "desc" : "description",
    "db" : "MYSQL",
    "db_data" : {
        "db_url" :  "127.0.0.1",
        "db_port" : "3306",
        "db_name" : "prueba",
        "db_user" : "root",
        "db_pass" :  "password",
        "db_charset" : "utf8",
        "db_collate" : "utf8_general_ci"
    },
    "lang" : "PHP", 
    "path" : "",
    "gen_sql" : true, 
    "gen_android" : false,    
    "auth_mode" : "jwt",     
    "init_date" : "2019-09-05",
    "upd_date" : "2019-09-05"
}
```

## Estructura de un Objeto de la Base de Datos JSON / Database Object JSON Structure

```
{
    "proyect_id" : "1",
    "name" : "table1",
    "tbl_name" : "table1",
    "is_login" : false,
    "login_id" : "codigo",
    "login_pass" : "codigo2",
    "columns" : [
          {
              "name" : "id",
              "col_name" : "id",
              "col_type" : "int",
              "col_lenght" : 20,
              "is_null" : false,
              "is_index" : true,
              "is_unique" : false,
              "primary_key" : true,
              "foreign_key" : false,
              "tbl_ref" : "",
              "col_refs" : [],
              "on_delete" : "",
              "on_update" : ""
          },
          {
              "name" : "cedula",
              "col_name" : "cedula",
              "col_type" : "varchar",
              "col_lenght" : 40,
              "is_null" : false,
              "is_index" : false,
              "is_unique" : false,
              "primary_key" : false,
              "foreign_key" : false,
              "tbl_ref" : "",
              "col_refs" : [],
              "on_delete" : "",
              "on_update" : ""
        }
    ]
}
```

## Como usar / How to use
Una vez descargado y descomprimido puedes modificar los archivos objects.json y proyects.json y luego ejecutar en una terminal dentro del directorio "go run main.go".

Once Downloaded and decompress, you can modify the objects.json and proyects.json files and run in the terminal while standing in the directory of the proyect "go run main.go".
