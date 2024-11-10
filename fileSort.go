package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var fileGO = "" // сюда записываем файл для копирования и сравнения
var fileTime = time.Now()
var tmpName string

// ScanPathA - сканируем папку и ищем по шаблону файл
func ScanPathA(pathA string, pathB string, nameBase string) { // Сканируем
	files, err := ioutil.ReadDir(pathA)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileTmp := range files {
		// Если находим файл с началом имени базы то откладываем
		if strings.Contains(fileTmp.Name(), nameBase) && fileGO == "" {
			// Если за именем файла начинается _ЦИФРА - то наша база, если нет не наша
			tmpName = fileTmp.Name()
			_, err := strconv.Atoi(string(tmpName[len(nameBase)+3]))
			if err != nil {
				continue
			}
			fileGO = fileTmp.Name() // записываем имя файла и идем дальше
			fileTime = fileTmp.ModTime()
			continue
		}
		// Если находим подобный файл то сравниваем время
		if strings.Contains(fileTmp.Name(), nameBase) && fileGO != fileTmp.Name() {
			tmpName = fileTmp.Name()
			_, err := strconv.Atoi(string(tmpName[len(nameBase)+3]))
			if err != nil {
				continue
			}
			check := fileTmp.ModTime().Equal(fileTime) // Если новый файл fileTem малдше прошлого (fileGo(filetime))
			if !check {
				fileGO = fileTmp.Name() // записываем имя файла и идем дальше
				fileTime = fileTmp.ModTime()
			}
		}
	}
	//Копирование файлов
	//var daBLD = "\\"
	var srcFile = pathA + "\\" + fileGO
	var dstFile = pathB + "\\" + fileGO
	fmt.Println(srcFile)
	fmt.Println(dstFile)
	// check srcFile stats
	fileStat, err := os.Stat(srcFile)
	if err != nil {
		//fmt.Print("Failed to check stats for ", srcFile)
		panic(err)
	}

	// print srcFile stats
	perm := fileStat.Mode().Perm()
	perm = perm
	//fmt.Printf("File permission before copying %v \n", perm)

	// read srcFile
	srcContent, err := ioutil.ReadFile(srcFile)
	if err != nil {
		//fmt.Print("Failed to read file ", srcFile)
		panic(err)
	}

	// create dstFile and copy the content
	err = ioutil.WriteFile(dstFile, srcContent, fileStat.Mode())
	if err != nil {
		//fmt.Print("Failed to copy content into ", dstFile)
		panic(err)
	}

	// check dstFile stats
	newFileStats, err := os.Stat(dstFile)
	if err != nil {
		//fmt.Print("Failed to check stats for ", dstFile)
		panic(err)
	}

	// print dstFile stats
	perm2 := newFileStats.Mode().Perm()
	perm2 = perm2
	//fmt.Printf("File permission After copying %v \n", perm2)
}
