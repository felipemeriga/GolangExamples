package example

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Form struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
	Occupation  string `json:"occupation"`
	Documents   struct {
		Rg struct {
			Name           string `json:"name"`
			Identification string `json:"identification"`
		} `json:"rg"`
		Cpf struct {
			Name           string `json:"name"`
			Identification string `json:"identification"`
		} `json:"cpf"`
	} `json:"documents"`
}

func FormUnmarshall() {
	workingDirectory, _ := os.Getwd()
	jsonFile, _ := ioutil.ReadFile(workingDirectory + "/test.json")
	var form Form

	_ = json.Unmarshal(jsonFile, &form)

	fmt.Println(form)

}

func Reflection() {
	workingDirectory, _ := os.Getwd()
	jsonFile, _ := ioutil.ReadFile(workingDirectory + "/test.json")
	var form *Form

	_ = json.Unmarshal(jsonFile, &form)
	parametersToChange := map[string]interface{}{
		"Name": "Liscro",
		"Age":  12,
		"Documents": map[string]interface{}{
			"Rg": map[string]interface{}{
				"Name":           "Registro Geral",
				"Identification": "47.935.669-5",
			},
			"Cpf": map[string]interface{}{
				"Name":           "CPF",
				"Identification": "419.242.888-18",
			},
		},
	}

	for key, param := range parametersToChange {
		field := reflect.ValueOf(form).Elem().FieldByName(key)
		field.Set(reflect.ValueOf(param))
	}
}
