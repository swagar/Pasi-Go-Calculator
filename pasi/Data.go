package pasi

import (
	"log"
)

type Data struct {
	Name	string
	Kopf  Line
	Arme  Line
	Rumpf Line
	Beine Line
	Pasi  float64
}

type Line struct {
	Erythem      float64
	Schuppung    float64
	Infiltration float64
	Ausdehnung   float64
	Index        float64
}

func (data *Data) Calc() {
	data.calcKopf()
	data.calcArme()
	data.calcRumpf()
	data.calcBeine()

	data.Pasi = data.Kopf.Index + data.Arme.Index + data.Rumpf.Index + data.Beine.Index

	log.Print(data)
}

func (data *Data ) calcKopf() {
	var kopf = &data.Kopf

	kopf.Index = (kopf.Erythem + kopf.Schuppung + kopf.Infiltration) * kopf.Ausdehnung / 10
}

func (data *Data ) calcArme() {
	var arme = &data.Arme

	arme.Index = (arme.Erythem + arme.Schuppung + arme.Infiltration) * arme.Ausdehnung / 5
}

func (data *Data ) calcRumpf() {
	var rumpf = &data.Rumpf

	rumpf.Index = (rumpf.Erythem + rumpf.Schuppung + rumpf.Infiltration) * rumpf.Ausdehnung * 0.3
}

func (data *Data ) calcBeine() {
	var beine = &data.Beine

	beine.Index = (beine.Erythem + beine.Schuppung + beine.Infiltration) * beine.Ausdehnung * 0.4
}