package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"time"
	"strconv"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	var dataBinder *walk.DataBinder
	pasiData := new(PasiData)

	var updateFunc = func() {
		dataBinder.Submit()
		pasiData.calc()
		dataBinder.Reset()
	}

	var windowSize = Size{Width: 800, Height: 400}
	var fontHeader = Font{PointSize: 20}
	var fontNormal = Font{PointSize: 16}

	var mW *walk.MainWindow

	var mainWindow = MainWindow{
		Title:   "Pasi Rechner - V0.2",
		MaxSize: windowSize,
		MinSize: windowSize,
		Size:    windowSize,
		AssignTo: &mW,
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
			ToolBar{
				ButtonStyle: ToolBarButtonImageBeforeText,
				Items: []MenuItem{
					Action{
						Text:        "Export PDF",
						Image:       "pdf.png",
						Enabled:     true,
						OnTriggered: func (){
							pdf := gofpdf.New("P", "mm", "A4", "")
							tr := pdf.UnicodeTranslatorFromDescriptor("")

							pdf.AddPage()

							pdf.SetFont("Arial", "B", 20)			

							pdf.Cell(40, 10, tr("P A S I - Berechnung für: " + pasiData.Name))
							pdf.Ln(15)

							pdf.SetFont("Arial", "", 10)				
							pdf.SetFillColor(200, 200, 200)

							pdf.CellFormat(32, 10, tr("Körper-Oberfläche"), "TL", 0, "", true, 0, "")
							pdf.CellFormat(8, 10, tr("X ("), "T", 0, "", true, 0, "")
							pdf.CellFormat(18, 10, tr("Erythem"), "T", 0, "", true, 0, "")
							pdf.CellFormat(5, 10, tr("+"), "T", 0, "", true, 0, "")
							pdf.CellFormat(20, 10, tr("Schuppung"), "T", 0, "", true, 0, "")
							pdf.CellFormat(5, 10, tr("+"), "T", 0, "", true, 0, "")
							pdf.CellFormat(18, 10, tr("Infiltration"), "T", 0, "", true, 0, "")
							pdf.CellFormat(7, 10, tr(") X"), "T", 0, "", true, 0, "")
							pdf.CellFormat(45, 10, tr("Ausdehnung der Läsionen"), "T", 0, "", true, 0, "")
							pdf.CellFormat(5, 10, tr("= "), "T", 0, "", true, 0, "")
							pdf.CellFormat(15, 10, tr("Index"), "TR", 0, "", true, 0, "")					

							pdf.Ln(-1)

							pdf.SetFillColor(240, 240, 240)

							pdf.CellFormat(17, 10, tr("Kopf"), "TL", 0, "", true, 0, "")
							pdf.CellFormat(15, 10, tr("0,1"), "T", 0, "C", true, 0, "")
							pdf.CellFormat(8, 10, tr("X ("), "T", 0, "", true, 0, "")
							pdf.CellFormat(18, 10, d(pasiData.Kopf.Erythem), "T", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("+"), "T", 0, "", true, 0, "")
							pdf.CellFormat(20, 10, d(pasiData.Kopf.Schuppung), "T", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("+"), "T", 0, "", true, 0, "")
							pdf.CellFormat(18, 10, d(pasiData.Kopf.Infiltration), "T", 0, "C", true, 0, "")
							pdf.CellFormat(7, 10, tr(") X"), "T", 0, "", true, 0, "")
							pdf.CellFormat(45, 10, d(pasiData.Kopf.Ausdehnung), "T", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("= "), "T", 0, "", true, 0, "")
							pdf.CellFormat(15, 10, d(pasiData.Kopf.Index), "TR", 0, "C", true, 0, "")

							pdf.Ln(-1)

							pdf.CellFormat(17, 10, tr("Arme"), "L", 0, "", true, 0, "")
							pdf.CellFormat(15, 10, tr("0,2"), "", 0, "C", true, 0, "")
							pdf.CellFormat(8, 10, tr("X ("), "", 0, "", true, 0, "")
							pdf.CellFormat(18, 10, d(pasiData.Arme.Erythem), "", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("+"), "", 0, "", true, 0, "")
							pdf.CellFormat(20, 10, d(pasiData.Arme.Schuppung), "", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("+"), "", 0, "", true, 0, "")
							pdf.CellFormat(18, 10, d(pasiData.Arme.Infiltration), "", 0, "C", true, 0, "")
							pdf.CellFormat(7, 10, tr(") X"), "", 0, "", true, 0, "")
							pdf.CellFormat(45, 10, d(pasiData.Arme.Ausdehnung), "", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("= "), "", 0, "", true, 0, "")
							pdf.CellFormat(15, 10, d(pasiData.Arme.Index), "R", 0, "C", true, 0, "")

							pdf.Ln(-1)

							pdf.CellFormat(17, 10, tr("Rumpf"), "L", 0, "", true, 0, "")
							pdf.CellFormat(15, 10, tr("0,3"), "", 0, "C", true, 0, "")
							pdf.CellFormat(8, 10, tr("X ("), "", 0, "", true, 0, "")
							pdf.CellFormat(18, 10, d(pasiData.Rumpf.Erythem), "", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("+"), "", 0, "", true, 0, "")
							pdf.CellFormat(20, 10, d(pasiData.Rumpf.Schuppung), "", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("+"), "", 0, "", true, 0, "")
							pdf.CellFormat(18, 10, d(pasiData.Rumpf.Infiltration), "", 0, "C", true, 0, "")
							pdf.CellFormat(7, 10, tr(") X"), "", 0, "", true, 0, "")
							pdf.CellFormat(45, 10, d(pasiData.Rumpf.Ausdehnung), "", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("= "), "", 0, "", true, 0, "")
							pdf.CellFormat(15, 10, d(pasiData.Rumpf.Index), "R", 0, "C", true, 0, "")

							pdf.Ln(-1)

							pdf.CellFormat(17, 10, tr("Beine"), "LB", 0, "", true, 0, "")
							pdf.CellFormat(15, 10, tr("0,4"), "B", 0, "C", true, 0, "")
							pdf.CellFormat(8, 10, tr("X ("), "B", 0, "", true, 0, "")
							pdf.CellFormat(18, 10, d(pasiData.Beine.Erythem), "B", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("+"), "B", 0, "", true, 0, "")
							pdf.CellFormat(20, 10, d(pasiData.Beine.Schuppung), "B", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("+"), "B", 0, "", true, 0, "")
							pdf.CellFormat(18, 10, d(pasiData.Beine.Infiltration), "B", 0, "C", true, 0, "")
							pdf.CellFormat(7, 10, tr(") X"), "B", 0, "", true, 0, "")
							pdf.CellFormat(45, 10, d(pasiData.Beine.Ausdehnung), "B", 0, "C", true, 0, "")
							pdf.CellFormat(5, 10, tr("= "), "B", 0, "", true, 0, "")
							pdf.CellFormat(15, 10, d(pasiData.Beine.Index), "RB", 0, "C", true, 0, "")

							pdf.Ln(-1)

							pdf.SetFont("Arial", "B", 14)	
							pdf.Cell(40, 10, tr("P A S I: " + d(pasiData.Pasi)))

							log.Print(pdf.OutputFileAndClose(pasiData.Name + "_" + time.Now().Format("02.01.2006") + ".pdf"))
							//walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
						},
					},
				},
			},
			Composite{
				Layout:     Grid{Columns: 2},
				Background: SolidColorBrush{Color: walk.RGB(245, 12, 12)},
				Border:     false,
				Children: []Widget{
					Label{Text: "Pat.-Name: ", Font: fontHeader},
					TextEdit{Text: Bind("Name",SelRequired{}), Font: fontHeader, OnTextChanged: updateFunc},
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
					NumberEdit{Decimals: 1, ReadOnly: true, Value: Bind("Beine.Index"),Alignment: AlignHFarVFar, Font: fontNormal},
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
					/*PushButton{
						Text: "Export PDF",
						Image: "pdf.png", 
						Font: fontHeader,
					},*/
				},
			},
		},
	}

	var _, error = mainWindow.Run()

	if nil != error{
		walk.MsgBox(mW, "Fehler beim Starten der Anwendung", error.Error(), walk.MsgBoxIconError)
	}
}

type PasiData struct {
	Name	string
	Kopf  PasiLine
	Arme  PasiLine
	Rumpf PasiLine
	Beine PasiLine
	Pasi  float64
}

type PasiLine struct {
	Erythem      float64
	Schuppung    float64
	Infiltration float64
	Ausdehnung   float64
	Index        float64
}

func (p *PasiData) calc() {
	p.calcKopf()
	p.calcArme()
	p.calcRumpf()
	p.calcBeine()

	p.Pasi = p.Kopf.Index + p.Arme.Index + p.Rumpf.Index + p.Beine.Index

	log.Print(p)
}

func (p *PasiData) calcKopf() {
	var kopf = &p.Kopf

	kopf.Index = (kopf.Erythem + kopf.Schuppung + kopf.Infiltration) * kopf.Ausdehnung / 10
}

func (p *PasiData) calcArme() {
	var arme = &p.Arme

	arme.Index = (arme.Erythem + arme.Schuppung + arme.Infiltration) * arme.Ausdehnung / 5
}

func (p *PasiData) calcRumpf() {
	var rumpf = &p.Rumpf

	rumpf.Index = (rumpf.Erythem + rumpf.Schuppung + rumpf.Infiltration) * rumpf.Ausdehnung * 0.3
}

func (p *PasiData) calcBeine() {
	var beine = &p.Beine

	beine.Index = (beine.Erythem + beine.Schuppung + beine.Infiltration) * beine.Ausdehnung * 0.4
}

func d(double float64) string{
	return strconv.FormatFloat(double, 'f', 1, 64)
}
