package main

import (
	"github.com/lxn/walk"
	//. "github.com/lxn/walk/declarative"
	"strconv"
	. "wt/pasi/pasi"
)

func main() {
	
	pasiData := new(Data)

	var updateFunc = func() {
		pasiData.Calc()
	}	

	var mW *walk.MainWindow

	var mainWindow = GenerateMainWindow(&mW, pasiData, updateFunc)

	var _, error = mainWindow.Run()

	if nil != error{
		walk.MsgBox(mW, "Fehler beim Starten der Anwendung", error.Error(), walk.MsgBoxIconError)
	}
}

func d(double float64) string{
	return strconv.FormatFloat(double, 'f', 1, 64)
}