package pasi

import (
	"log"
	"time"
	"strconv"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(pasiData *Data) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 20)

	pdf.Cell(40, 10, tr("P A S I - Berechnung für: "+ pasiData.Name))
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
	pdf.Cell(40, 10, tr("P A S I: "+d(pasiData.Pasi)))

	log.Print(pdf.OutputFileAndClose(pasiData.Name + "_" + time.Now().Format("02.01.2006") + ".pdf"))
}

func d(double float64) string{
	return strconv.FormatFloat(double, 'f', 1, 64)
}
