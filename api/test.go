package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := "exemple.json" // имя файла с текстом в формате JSON
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data interface{}
	err = json.Unmarshal(bytes, &data) // разбор JSON
	if err != nil {
		fmt.Println(err)
		return
	}

	// вызов функции для обработки каждого уровня вложенности
	printLevel(data, 0)
}

// функция обработки уровня вложенности
func printLevel(data interface{}, level int) {
	switch value := data.(type) {
	case map[string]interface{}:
		for key, val := range value {
			fmt.Printf("%*s%s:\n", level*4, "", key)
			printLevel(val, level+1)
		}
	case []interface{}:
		for _, val := range value {
			printLevel(val, level+1)
		}
	default:
		fmt.Printf("%*s%v\n", level*4, "", value)
	}
}
