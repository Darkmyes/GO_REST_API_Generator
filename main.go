package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"./proyect_maker"
)

func main() {
	/*
		fs := http.FileServer(http.Dir("public"))
		http.Handle("/", fs)
		log.Println("Escuchando en el puerto 3000")
		http.ListenAndServe(":3000", nil)
	*/

	data, err := ioutil.ReadFile("./public/proyects.json")
	if err != nil {
		fmt.Print(err)
	}

	data2, err2 := ioutil.ReadFile("./public/objects.json")
	if err2 != nil {
		fmt.Print(err2)
	}

	// json data
	var obj proyect_maker.Proyect
	var objs []proyect_maker.Object

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	err2 = json.Unmarshal(data2, &objs)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

	proyect_maker.MakeRestAPI(obj, objs)
	//proyect_maker.Make (obj, objs)
}
