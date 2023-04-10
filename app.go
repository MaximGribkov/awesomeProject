package main

//func main () {
//	planet[10460] = "Датомир"
//	planet[10465] = "Татуин"
//	planet[12120] = "Набу"
//	planet[12240] = "Корусант"
//	planet[12500] = "Альдеран"
//}
//
//На основе этой функции сделать рабочий код вывода map данных в порядке возрастания ключа
//Тестовое задание

import (
	"fmt"
	"sort"
)

func main() {
	planet := make(map[int]string) // начальная мапа
	planet[9460] = "Датомир"
	planet[80465] = "Татуин"
	planet[72120] = "Набу"
	planet[2240] = "Корусант"
	planet[100] = "Альдеран"

	keys := make([]int, 0, len(planet)) // отделили ключи
	for k := range planet {
		keys = append(keys, k)
	}
	sort.Ints(keys) // сортировка

	for _, k := range keys { // вывод
		fmt.Printf("Код: %d, Планета: %s\n", k, planet[k])
	}
}
