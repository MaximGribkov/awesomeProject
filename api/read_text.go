package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type asteroid struct {
	Name              string `json:"name"`
	ID                string `json:"id"`
	CloseApproachDate string `json:"close_approach_date"`
}

func main() {
	// Открытие файла с данными в формате JSON
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// Создание JSON декодера
	decoder := json.NewDecoder(file)

	var asteroids []asteroid

	// Декодирование JSON данных в структуру asteroid
	err = decoder.Decode(&asteroids)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Вывод информации о каждом астероиде по ключам "name", "id" и "close_approach_date"
	for _, a := range asteroids {
		fmt.Printf("Asteroid name: %s\n", a.Name)
		fmt.Printf("Asteroid ID: %s\n", a.ID)
		fmt.Printf("Close approach date: %s\n", a.CloseApproachDate)
		fmt.Println()
	}
}
