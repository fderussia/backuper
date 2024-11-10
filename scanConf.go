package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

var ListBase []string = make([]string, 0) // складываем имена файлов
var filePath = scanFilePath()
var checkPoint = false // если ключ найден то пишем в переменную

func scanFilePath() string {
	filePath := "config_backuper.txt"    // По условию файл называется - config_backuper.txt находится в папке программы
	filePath, _ = filepath.Abs(filePath) // Находим текущий путь и делаем абсолбтный путь
	return filePath
}

// ScanString - Сканирует строки
func ScanString(key string, val string) string {
	file, _ := os.Open(filePath)      //открываем файл
	scanner := bufio.NewScanner(file) // начинаем сканировать построчно
	for scanner.Scan() {
		tmpStr := scanner.Text()                                  // переменная получающая всю строку
		if strings.Contains(tmpStr, key) && checkPoint == false { // если ключ найден то дальше пишем в переменную val
			checkPoint = true
			continue
		} else if checkPoint && strings.Contains(tmpStr, "#") {
			checkPoint = false
			continue
		} else if checkPoint {
			val = strings.TrimSpace(tmpStr)
			continue
		}
	}

	return val
}

// ScanSlice - Сканирует в список
func ScanSlice() {
	file, _ := os.Open(filePath)      //открываем файл
	scanner := bufio.NewScanner(file) // начинаем сканировать построчно
	for scanner.Scan() {
		tmpStr := scanner.Text()                                         // переменная получающая всю строку
		if strings.Contains(tmpStr, "listBase") && checkPoint == false { // если ключ найден то дальше пишем в переменную val
			checkPoint = true
			continue
		} else if checkPoint && strings.Contains(tmpStr, "#") {
			checkPoint = false
			continue
		} else if checkPoint {
			tmpStr = strings.TrimSpace(tmpStr)
			ListBase = append(ListBase, tmpStr)
			continue
		}
	}

}
