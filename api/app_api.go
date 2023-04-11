package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// код создания дб
// CREATE TABLE neo_count (
//                           date DATE NOT NULL,
//                           count INTEGER NOT NULL
//);

func main() {
	// Введите свой ключ API  ключ DC8v9CmHBaoelbzMGPGwncI1GfDRN2JuPqL6Jb1X
	apiKey := "API"

	// Установите даты начала и конца в нужном формате: "YYYY-MM-DD" по правилам 7 дней
	startDate := "2021-01-01"
	endDate := "2021-01-08"

	// Создайте URL для запроса
	url := fmt.Sprintf("https://api.nasa.gov/neo/rest/v1/feed?start_date=%s&end_date=%s&api_key=%s", startDate, endDate, apiKey)

	// Отправьте GET-запрос к API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка при выполнении запроса: %s", err.Error())
		return
	}

	// Прочитайте ответ от API
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении ответа API: %s", err.Error())
		return
	}

	// Напечатайте ответ от API
	fmt.Printf("Ответ от API: %s", string(body))
}
