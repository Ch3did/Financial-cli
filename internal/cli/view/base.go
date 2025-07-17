package view

import (
	"encoding/json"
	"fmt"
	"reflect"

	"os"
	"os/exec"
)

func JsonOutput(i interface{}) {
	jsonData, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		fmt.Printf("Erro ao converter para JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func BaseOutput(i interface{}) {
	val := reflect.ValueOf(i)

	for idx := 0; idx < val.Len(); idx++ {
		item := val.Index(idx).Interface()
		jsonData, err := json.MarshalIndent(item, "", "    ")
		if err != nil {
			fmt.Printf("Erro ao converter item %d: %v\n", idx, err)
			continue
		}
		fmt.Println(string(jsonData))
	}

}

func RunIfNotDebug(fn func()) {
	if os.Getenv("DEBUG") == "true" {
		return
	}
	fn()
}
