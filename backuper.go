package main

var cfgMap = map[string]string{
	"pathA": "", // Из какой папки будет копирование
	"pathB": "", // В какую папку будет копирование
}

func main() {
	for key, val := range cfgMap { // заполняю параметры конфига
		cfgMap[key] = ScanString(key, val)
	}
	ScanSlice() // Заполняем ListBase (список баз)
	for i := 0; i < len(ListBase); i++ {
		ScanPathA(cfgMap["pathA"], cfgMap["pathB"], ListBase[i])
	}
}
