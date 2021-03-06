package pasi

import (
	"time"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

/**
jo
*/
func GenerateMainWindow(mW **walk.MainWindow, pasiData *Data, onUpdate func()) MainWindow {
	var dataBinder *walk.DataBinder
	var windowSize = Size{Width: 800, Height: 400}
	var fontHeader = Font{PointSize: 20}
	var fontNormal = Font{PointSize: 16}

	var updateFunc = func() {
		dataBinder.Submit()
		onUpdate()
		dataBinder.Reset()
	}

	return MainWindow{
		Title:    "Pasi Rechner - V0.3",
		MaxSize:  windowSize,
		MinSize:  windowSize,
		Size:     windowSize,
		AssignTo: mW,
		Layout: VBox{
			Alignment:   AlignHNearVNear,
			MarginsZero: true,
			SpacingZero: true,
		},
		DataBinder: DataBinder{
			AssignTo:       &dataBinder,
			Name:           "pasiData",
			DataSource:     pasiData,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		Children: []Widget{
			Composite{
				Layout:     Grid{Columns: 2},
				Background: SolidColorBrush{Color: walk.RGB(245, 12, 12)},
				Border:     false,
				Children: []Widget{
					Label{Text: "Pat.-Name: ", Font: fontHeader},
					TextEdit{Text: Bind("Name", SelRequired{}), Font: fontHeader, OnTextChanged: updateFunc},
				},
			},
			Composite{
				Layout: Grid{Columns: 12},
				Children: []Widget{
					// Header
					Label{Text: "Körper-\nOberfläche", ColumnSpan: 2, Font: fontNormal},
					Label{Text: "X", Font: fontNormal},
					Label{Text: "(Erythem", Font: fontNormal},
					Label{Text: "+", Font: fontNormal},
					Label{Text: "Schuppung", Font: fontNormal},
					Label{Text: "+", Font: fontNormal},
					Label{Text: "Infiltration)", Font: fontNormal},
					Label{Text: "  X  ", Font: fontNormal},
					Label{Text: "Ausdehnung\nder Läsionen", Font: fontNormal},
					Label{Text: " = ", Font: fontNormal},
					Label{Text: "Index", Font: fontNormal},

					// Kopf
					Label{Text: "Kopf", Font: fontNormal},
					Label{Text: "0,1", Font: fontNormal},
					Label{Text: "X (", Font: fontNormal},

					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Kopf.Erythem", Range{0, 4}),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: "+", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Kopf.Schuppung"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: "+", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Kopf.Infiltration"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: ") X", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 6,
						Value:          Bind("Kopf.Ausdehnung"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: " = ", Font: fontNormal},
					NumberEdit{Decimals: 1, ReadOnly: true, Value: Bind("Kopf.Index"), Font: fontNormal},

					// Arme
					Label{Text: "Arme", Font: fontNormal},
					Label{Text: "0,2", Font: fontNormal},
					Label{Text: "X (", Font: fontNormal},

					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Arme.Erythem", Range{0, 4}),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: "+", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Arme.Schuppung"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: "+", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Arme.Infiltration"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: ") X", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 6,
						Value:          Bind("Arme.Ausdehnung"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: " = ", Font: fontNormal},
					NumberEdit{Decimals: 1, ReadOnly: true, Value: Bind("Arme.Index"), Font: fontNormal},

					// Rumpf
					Label{Text: "Rumpf", Font: fontNormal},
					Label{Text: "0,3", Font: fontNormal},
					Label{Text: "X (", Font: fontNormal},

					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Rumpf.Erythem", Range{0, 4}),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: "+", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Rumpf.Schuppung"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: "+", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Rumpf.Infiltration"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: ") X", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 6,
						Value:          Bind("Rumpf.Ausdehnung"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: " = ", Font: fontNormal},
					NumberEdit{Decimals: 1, ReadOnly: true, Value: Bind("Rumpf.Index"), Font: fontNormal},

					// Beine
					Label{Text: "Beine", Font: fontNormal},
					Label{Text: "0,4", Font: fontNormal},
					Label{Text: "X (", Font: fontNormal},

					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Beine.Erythem", Range{0, 4}),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: "+", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Beine.Schuppung"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: "+", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 4,
						Value:          Bind("Beine.Infiltration"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: ") X", Font: fontNormal},
					NumberEdit{Decimals: 0, MinValue: 0, MaxValue: 6,
						Value:          Bind("Beine.Ausdehnung"),
						OnValueChanged: updateFunc, Font: fontNormal,
					},
					Label{Text: " = ", Font: fontNormal},
					NumberEdit{Decimals: 1, ReadOnly: true, Value: Bind("Beine.Index"), Alignment: AlignHFarVFar, Font: fontNormal},
				},
			},
			Composite{
				Layout:     Grid{Columns: 4},
				Background: SolidColorBrush{Color: walk.RGB(245, 12, 12)},
				Border:     false,
				Children: []Widget{
					Label{Text: "P A S I: ", Font: fontHeader},
					NumberEdit{Decimals: 1, Font: fontHeader, ReadOnly: true, Value: Bind("Pasi"),
						MaxSize: Size{Width: 100, Height: 40},
					},
					Label{Text: " ", Font: fontNormal},
					PushButton{
						Text:    "Export PDF",
						Font:    fontNormal,
						MaxSize: Size{Width: 130, Height: 40},
						OnClicked: func() {
							var dialog = generateSaveDialog(pasiData)
							dialog.Run(*mW)
						},
					},
				},
			},
		},
	}
}

func generateSaveDialog(pasiData *Data) Dialog {
	var treeView *walk.TreeView
	dialogSize := Size{Width: 400, Height: 400}
	treeModel, _ := NewDirectoryTreeModel()
	var dialog *walk.Dialog

	var dataBinder *walk.DataBinder
	var fileNameData FileNameData = FileNameData{FileName: pasiData.Name + "_" + time.Now().Format("02.01.2006") + ".pdf"}

	var updateFunc = func() {
		dataBinder.Submit()
		dataBinder.Reset()
	}

	return Dialog{
		Title:    "Speichern unter",
		MinSize:  dialogSize,
		Size:     dialogSize,
		Layout:   VBox{},
		AssignTo: &dialog,
		DataBinder: DataBinder{
			AssignTo:       &dataBinder,
			Name:           "fileNameData",
			DataSource:     &fileNameData,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		Children: []Widget{
			TreeView{
				AssignTo: &treeView,
				Model:    treeModel, OnCurrentItemChanged: func() {
					dir := treeView.CurrentItem().(*Directory)
					fileNameData.FolderPath = dir.Path()
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{Text: "Dateiname: "},
					LineEdit{Text: Bind("FileName"), OnTextChanged: updateFunc},
					PushButton{
						Text: "Speichern",
						OnClicked: func() {
							GeneratePDF(pasiData, fileNameData.FolderPath+"\\"+fileNameData.FileName)
							dialog.Close(0)
						},
					},
				},
			},
		},
	}
}
