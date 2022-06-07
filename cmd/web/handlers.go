package main

import (
    "net/http"
	"html/template"
	"log"
	"strconv"
	"os"
	"image"
	_"image/png" //invoking only init method of this package, thats why this underscore
	"image/color"
	"strings"
)


func home(w http.ResponseWriter, r *http.Request) {

	// Handling 404 error
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

	//Parsing url query
	query := r.URL.Query()
	var str []string
	for k, _:= range query {
		str = strings.Split(k,",")
	}

	//Find district if coordinates received as url query
	var district string
	if len(str) == 2 {
		x, _ := strconv.Atoi(str[0])
		y, _ := strconv.Atoi(str[1])
		f, err := os.Open("./cmd/web/district.png")  //Reference image
		if err!=nil{
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return	
		}
		defer f.Close()
		img, _, err := image.Decode(f) // _ ==> format (not needed)
		if err!=nil	{
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		pixel := img.At(x,y)
		originalColor,_ := color.RGBAModel.Convert(pixel).(color.RGBA)
		red := int(originalColor.R)
		n := map[int]string { 42:"KASARGOD", 
							40: "KANNUR", 
							57: "KOZHIKODE", 
							127: "WAYANAD", 
							203: "MALAPPURAM", 
							190: "PALAKKAD", 
							142: "THRISSUR", 
							178: "ERNAKULAM", 
							76: "ALAPPUZHA", 
							150: "KOTTAYAM", 
							105: "IDUKKI", 
							224: "PATHANAMTHITTA", 
							91: "KOLLAM", 
							83: "THIRUVANANTHAPURAM"}
		log.Println("(" + strconv.Itoa(x) + ","+ strconv.Itoa(y)+"):" + n[red])
		district = n[red]
	}


	// Load templates
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	} 

	// Render template
	//err = ts.Execute(w, district)
	err = ts.ExecuteTemplate(w, "base", district)
	if err != nil	{
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}


func chumma(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Phew! Chumma another page"));
}